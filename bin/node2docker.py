#!/usr/bin/env python3
import argparse
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path


def build_parser() -> argparse.ArgumentParser:
    """Create argument parser."""
    parser = argparse.ArgumentParser(
        prog="node2docker.py",
        description="Package Node.js CLI tools into Docker images.",
    )
    subparsers = parser.add_subparsers(dest="command", required=True)
    build_parser = subparsers.add_parser("build", help="Build a Docker image.")
    build_parser.add_argument("--package", required=True, help="npm package name")
    build_parser.add_argument("--bin", required=True, help="CLI entrypoint")
    build_parser.add_argument("--image", required=True, help="Docker image name")
    build_parser.add_argument("--tag", default="", help="Docker tag (default: latest)")
    build_parser.add_argument("--base", default="node:20-alpine", help="Base image")
    build_parser.add_argument("--user", default="node", help="Runtime user")
    build_parser.add_argument("--no-user", action="store_true", help="Do not drop privileges")
    build_parser.add_argument("--no-cache", action="store_true", help="Disable build cache")

    shim_parser = subparsers.add_parser("shim", help="Print shim to stdout.")
    shim_parser.add_argument("--image", required=True, help="Image reference")
    shim_parser.add_argument("--name", default="", help="Optional name for the shim file")
    return parser


def image_ref(image: str, tag: str) -> str:
    """Return full image reference."""
    if not tag and ":" in image:
        return image
    if not tag:
        tag = "latest"
    return f"{image}:{tag}"


def write_dockerfile(path: Path, opts: argparse.Namespace) -> None:
    """Write Dockerfile content."""
    dockerfile = f"""FROM {opts.base}
ENV NODE_ENV=production \\
    NPM_CONFIG_FUND=false \\
    NPM_CONFIG_AUDIT=false
RUN npm install -g {opts.package}
"""
    if not opts.no_user:
        dockerfile += f"USER {opts.user}\n"
    dockerfile += f'ENTRYPOINT ["{opts.bin}"]\n'
    path.write_text(dockerfile)


def ensure_command(name: str) -> None:
    """Ensure command exists."""
    if shutil.which(name) is None:
        raise RuntimeError(f"missing required command: {name}")


def run_build(opts: argparse.Namespace) -> None:
    """Build a Docker image."""
    ensure_command("docker")
    image = image_ref(opts.image, opts.tag)
    with tempfile.TemporaryDirectory() as tmp:
        dockerfile = Path(tmp) / "Dockerfile"
        write_dockerfile(dockerfile, opts)
        cmd = ["docker", "build", "-t", image]
        if opts.no_cache:
            cmd.insert(2, "--no-cache")
        print(f"Building image {image}...")
        subprocess.run(cmd + [tmp], check=True)
    print(f"Built {image}")


def run_shim(opts: argparse.Namespace) -> None:
    """Print shim to stdout."""
    ensure_command("docker")
    shim = f"""#!/usr/bin/env sh
set -e

image_ref="{opts.image}"

if [ -t 0 ]; then
  tty_flags="-it"
else
  tty_flags=""
fi

exec docker run --rm ${{tty_flags}} "${{image_ref}}" "$@"
"""
    print(shim, end="")


def main(argv: list[str]) -> int:
    """Program entrypoint."""
    parser = build_parser()
    args = parser.parse_args(argv)
    if args.command == "build":
        run_build(args)
        return 0
    if args.command == "shim":
        run_shim(args)
        return 0
    parser.print_help()
    return 1


if __name__ == "__main__":
    raise SystemExit(main(sys.argv[1:]))
