#!/usr/bin/env bash
set -euo pipefail

HELP_TEXT=$'node2docker.bash - package Node.js CLI tools into Docker images\n\nUsage:\n  node2docker.bash build --package <npm_pkg> --bin <entrypoint> --image <name> [options]\n  node2docker.bash shim --image <image[:tag]> [--name <shim_name>]\n  node2docker.bash help\n\nCommands:\n  build   Build a Docker image for an npm CLI tool.\n  shim    Print a shim script to stdout.\n\nBuild options:\n  --package <npm_pkg>   npm package name (e.g. eslint or @scope/tool)\n  --bin <entrypoint>    CLI entrypoint exposed by the package (e.g. eslint)\n  --image <name>        Docker image name (e.g. acme/eslint)\n  --tag <tag>           Docker tag (default: latest)\n  --base <image>        Base image (default: node:20-alpine)\n  --user <user>         Image user (default: node). Use --no-user to skip.\n  --no-user             Do not drop privileges in the image\n  --no-cache            Disable Docker build cache\n\nShim options:\n  --image <image[:tag]> Image reference to execute\n  --name <shim_name>    Optional name for the shim file when redirecting\n\nExamples:\n  node2docker.bash build --package eslint --bin eslint --image acme/eslint\n  node2docker.bash shim --image acme/eslint:latest > ~/.local/bin/eslint'

show_help() {
  : "Print usage."
  printf '%s\n' "$HELP_TEXT"
}

die() {
  : "Exit with an error message."
  printf 'error: %s\n' "$*" >&2
  exit 1
}

require_cmd() {
  : "Ensure a command exists."
  command -v "$1" >/dev/null 2>&1 || die "missing required command: $1"
}

parse_build_args() {
  : "Parse build args into globals."
  BUILD_PACKAGE=""
  BUILD_BIN=""
  BUILD_IMAGE=""
  BUILD_TAG=""
  BUILD_BASE="node:20-alpine"
  BUILD_USER="node"
  BUILD_NO_USER=0
  BUILD_NO_CACHE=0
  while [ "$#" -gt 0 ]; do
    case "$1" in
      --package) BUILD_PACKAGE="$2"; shift 2 ;;
      --bin) BUILD_BIN="$2"; shift 2 ;;
      --image) BUILD_IMAGE="$2"; shift 2 ;;
      --tag) BUILD_TAG="$2"; shift 2 ;;
      --base) BUILD_BASE="$2"; shift 2 ;;
      --user) BUILD_USER="$2"; shift 2 ;;
      --no-user) BUILD_NO_USER=1; shift ;;
      --no-cache) BUILD_NO_CACHE=1; shift ;;
      -h|--help) show_help; exit 0 ;;
      *) die "unknown build option: $1" ;;
    esac
  done
}

validate_build_args() {
  : "Validate build args."
  [ -n "$BUILD_PACKAGE" ] || die "missing --package"
  [ -n "$BUILD_BIN" ] || die "missing --bin"
  [ -n "$BUILD_IMAGE" ] || die "missing --image"
}

parse_shim_args() {
  : "Parse shim args into globals."
  SHIM_IMAGE=""
  SHIM_NAME=""
  while [ "$#" -gt 0 ]; do
    case "$1" in
      --image) SHIM_IMAGE="$2"; shift 2 ;;
      --name) SHIM_NAME="$2"; shift 2 ;;
      -h|--help) show_help; exit 0 ;;
      *) die "unknown shim option: $1" ;;
    esac
  done
}

validate_shim_args() {
  : "Validate shim args."
  [ -n "$SHIM_IMAGE" ] || die "missing --image"
}

compute_image_ref() {
  : "Return full image ref."
  if [ -z "$BUILD_TAG" ] && printf '%s' "$BUILD_IMAGE" | grep -q ':'; then
    printf '%s' "$BUILD_IMAGE"
    return 0
  fi
  printf '%s:%s' "$BUILD_IMAGE" "${BUILD_TAG:-latest}"
}

write_dockerfile() {
  : "Write Dockerfile to the provided path."
  local path="$1"
  cat <<EOF > "$path"
FROM ${BUILD_BASE}
ENV NODE_ENV=production \\
    NPM_CONFIG_FUND=false \\
    NPM_CONFIG_AUDIT=false
RUN npm install -g ${BUILD_PACKAGE}
EOF
  if [ "$BUILD_NO_USER" -eq 0 ]; then
    printf '%s\n' "USER ${BUILD_USER}" >> "$path"
  fi
  printf '%s\n' "ENTRYPOINT [\"${BUILD_BIN}\"]" >> "$path"
}

run_docker_build() {
  : "Run docker build with the provided context."
  local dir="$1"
  local image_ref="$2"
  local -a args=()
  if [ "$BUILD_NO_CACHE" -eq 1 ]; then
    args+=(--no-cache)
  fi
  printf 'Building image %s...\n' "$image_ref"
  docker build "${args[@]}" -t "$image_ref" "$dir"
  printf 'Built %s\n' "$image_ref"
}

build_cmd() {
  : "Handle build subcommand."
  parse_build_args "$@"
  validate_build_args
  require_cmd docker
  local image_ref
  image_ref=$(compute_image_ref)
  local tmp_dir
  tmp_dir=$(mktemp -d)
  trap 'rm -rf "$tmp_dir"' EXIT
  write_dockerfile "$tmp_dir/Dockerfile"
  run_docker_build "$tmp_dir" "$image_ref"
}

shim_cmd() {
  : "Handle shim subcommand."
  parse_shim_args "$@"
  validate_shim_args
  require_cmd docker
  cat <<EOF
#!/usr/bin/env sh
set -e

image_ref="${SHIM_IMAGE}"

if [ -t 0 ]; then
  tty_flags="-it"
else
  tty_flags=""
fi

exec docker run --rm \${tty_flags} "\${image_ref}" "\$@"
EOF
}

main() {
  : "Script entrypoint."
  if [ "$#" -eq 0 ]; then
    show_help
    exit 1
  fi
  case "$1" in
    build) shift; build_cmd "$@" ;;
    shim) shift; shim_cmd "$@" ;;
    help|-h|--help) show_help ;;
    *) die "unknown command: $1" ;;
  esac
}

main "$@"
