package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type buildFlags struct {
	Package string
	Bin     string
	Image   string
	Tag     string
	Base    string
	User    string
	NoUser  bool
	NoCache bool
}

type shimFlags struct {
	Image string
	Name  string
	MountCwd bool
}

const defaultWorkDir = "/work"

// main is the program entrypoint.
func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

// newRootCmd builds the root command.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "cli2docker",
		Short:        "Package Node.js CLI tools into Docker images",
		SilenceUsage: true,
	}
	cmd.AddCommand(newBuildCmd(), newShimCmd())
	return cmd
}

// newBuildCmd constructs the build command.
func newBuildCmd() *cobra.Command {
	opts := buildFlags{
		Base: "node:20-alpine",
		User: "node",
	}
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build a Docker image for an npm CLI tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ensureCommand("docker"); err != nil {
				return err
			}
			return buildWithOptions(opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.Package, "package", "", "npm package name")
	flags.StringVar(&opts.Bin, "bin", "", "CLI entrypoint")
	flags.StringVar(&opts.Image, "image", "", "Docker image name")
	flags.StringVar(&opts.Tag, "tag", "", "Docker tag")
	flags.StringVar(&opts.Base, "base", opts.Base, "Base image")
	flags.StringVar(&opts.User, "user", opts.User, "Runtime user")
	flags.BoolVar(&opts.NoUser, "no-user", false, "Do not drop privileges")
	flags.BoolVar(&opts.NoCache, "no-cache", false, "Disable build cache")
	for _, name := range []string{"package", "bin", "image"} {
		_ = cmd.MarkFlagRequired(name)
	}
	return cmd
}

// newShimCmd constructs the shim command.
func newShimCmd() *cobra.Command {
	opts := shimFlags{}
	cmd := &cobra.Command{
		Use:   "shim",
		Short: "Print a shim script to stdout",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ensureCommand("docker"); err != nil {
				return err
			}
			execLine := "exec docker run --rm ${tty_flags} \"${image_ref}\" \"$@\""
			if opts.MountCwd {
				execLine = fmt.Sprintf(
					"exec docker run --rm ${tty_flags} -v \"${PWD}:%s\" -w %s \"${image_ref}\" \"$@\"",
					defaultWorkDir,
					defaultWorkDir,
				)
			}
			lines := []string{
				"#!/usr/bin/env sh",
				"set -e",
				"",
				fmt.Sprintf("image_ref=%q", opts.Image),
				"",
				"if [ -t 0 ]; then",
				"  tty_flags=\"-it\"",
				"else",
				"  tty_flags=\"\"",
				"fi",
				"",
				execLine,
			}
			fmt.Fprint(cmd.OutOrStdout(), strings.Join(lines, "\n")+"\n")
			return nil
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.Image, "image", "", "Image reference")
	flags.StringVar(&opts.Name, "name", "", "Optional name for the shim file")
	flags.BoolVar(&opts.MountCwd, "mount-cwd", false, "Mount current directory into the container")
	_ = cmd.MarkFlagRequired("image")
	return cmd
}

// ensureCommand checks if an executable exists.
func ensureCommand(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("missing required command: %s", name)
	}
	return nil
}

// buildWithOptions runs the build workflow.
func buildWithOptions(opts buildFlags) error {
	image := buildImageRef(opts.Image, opts.Tag)
	tmpDir, err := os.MkdirTemp("", "cli2docker-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)
	dockerfile := filepath.Join(tmpDir, "Dockerfile")
	if err := writeDockerfile(dockerfile, opts); err != nil {
		return err
	}
	if err := runDockerBuild(tmpDir, image, opts.NoCache); err != nil {
		return err
	}
	fmt.Printf("Built %s\n", image)
	return nil
}

// buildImageRef resolves image name and tag.
func buildImageRef(image string, tag string) string {
	if tag == "" && strings.Contains(image, ":") {
		return image
	}
	if tag == "" {
		tag = "latest"
	}
	return image + ":" + tag
}

// writeDockerfile writes Dockerfile content.
func writeDockerfile(path string, opts buildFlags) error {
	lines := []string{
		"FROM " + opts.Base,
		"ENV NODE_ENV=production \\",
		"    NPM_CONFIG_FUND=false \\",
		"    NPM_CONFIG_AUDIT=false",
		"RUN npm install -g " + opts.Package,
	}
	if !opts.NoUser {
		lines = append(lines, "USER "+opts.User)
	}
	lines = append(lines, "ENTRYPOINT [\""+opts.Bin+"\"]")
	content := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(path, []byte(content), 0o644)
}

// runDockerBuild executes docker build.
func runDockerBuild(dir string, image string, noCache bool) error {
	args := []string{"build"}
	if noCache {
		args = append(args, "--no-cache")
	}
	args = append(args, "-t", image, dir)
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Building image %s...\n", image)
	return cmd.Run()
}
