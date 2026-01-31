package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPackageBaseName verifies scope stripping.
func TestPackageBaseName(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "plain", input: "eslint", expected: "eslint"},
		{name: "plain-version", input: "eslint@latest", expected: "eslint"},
		{name: "scoped", input: "@acme/eslint", expected: "eslint"},
		{name: "scoped-version", input: "@acme/eslint@1.2.3", expected: "eslint"},
	}
	for _, tc := range cases {
		if got := packageBaseName(tc.input); got != tc.expected {
			t.Fatalf("%s: expected %q, got %q", tc.name, tc.expected, got)
		}
	}
}

// TestImageFromPackage verifies image derivation.
func TestImageFromPackage(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		prefix   string
		expected string
	}{
		{name: "plain", input: "eslint", prefix: "cli/", expected: "cli/eslint"},
		{name: "plain-version", input: "eslint@latest", prefix: "cli/", expected: "cli/eslint"},
		{name: "scoped", input: "@acme/eslint", prefix: "cli/", expected: "cli/acme/eslint"},
		{name: "scoped-version", input: "@acme/eslint@1.2.3", prefix: "cli/", expected: "cli/acme/eslint"},
	}
	for _, tc := range cases {
		if got := imageFromPackage(tc.input, tc.prefix); got != tc.expected {
			t.Fatalf("%s: expected %q, got %q", tc.name, tc.expected, got)
		}
	}
}

func TestBuildDefaultsWarnings(t *testing.T) {
	oldEnsure := ensureCommandFn
	oldBuild := buildWithOptionsFn
	t.Cleanup(func() {
		ensureCommandFn = oldEnsure
		buildWithOptionsFn = oldBuild
	})

	ensureCommandFn = func(string) error { return nil }
	var got buildFlags
	buildWithOptionsFn = func(opts buildFlags) error {
		got = opts
		return nil
	}

	cmd := newBuildCmd()
	cmd.SetArgs([]string{"--package", "@acme/eslint@1.2.3"})
	stderr, err := captureStderr(t, cmd.Execute)
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got.Image != "cli/acme/eslint" {
		t.Fatalf("expected derived image %q, got %q", "cli/acme/eslint", got.Image)
	}
	if got.Bin != "eslint" {
		t.Fatalf("expected derived bin %q, got %q", "eslint", got.Bin)
	}
	if got.PackageManager != packageManagerNPM {
		t.Fatalf("expected package manager %q, got %q", packageManagerNPM, got.PackageManager)
	}
	if got.Base != defaultNPMBase {
		t.Fatalf("expected base %q, got %q", defaultNPMBase, got.Base)
	}
	if got.User != defaultNPMUser {
		t.Fatalf("expected user %q, got %q", defaultNPMUser, got.User)
	}
	if !strings.Contains(stderr, "--image") || !strings.Contains(stderr, "--bin") {
		t.Fatalf("expected warnings for both image and bin derivation, got %q", stderr)
	}
}

func TestBuildDefaultsForBun(t *testing.T) {
	oldEnsure := ensureCommandFn
	oldBuild := buildWithOptionsFn
	t.Cleanup(func() {
		ensureCommandFn = oldEnsure
		buildWithOptionsFn = oldBuild
	})

	ensureCommandFn = func(string) error { return nil }
	var got buildFlags
	buildWithOptionsFn = func(opts buildFlags) error {
		got = opts
		return nil
	}

	cmd := newBuildCmd()
	cmd.SetArgs([]string{"--package", "eslint", "--package-manager", "bun"})
	if _, err := captureStderr(t, cmd.Execute); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got.PackageManager != packageManagerBun {
		t.Fatalf("expected package manager %q, got %q", packageManagerBun, got.PackageManager)
	}
	if got.Base != defaultBunBase {
		t.Fatalf("expected base %q, got %q", defaultBunBase, got.Base)
	}
	if got.User != defaultBunUser {
		t.Fatalf("expected user %q, got %q", defaultBunUser, got.User)
	}
}

func TestBuildInvalidPackageManager(t *testing.T) {
	oldEnsure := ensureCommandFn
	oldBuild := buildWithOptionsFn
	t.Cleanup(func() {
		ensureCommandFn = oldEnsure
		buildWithOptionsFn = oldBuild
	})

	ensureCommandFn = func(string) error { return nil }
	buildWithOptionsFn = func(buildFlags) error {
		t.Fatal("build should not be invoked for invalid package manager")
		return nil
	}

	cmd := newBuildCmd()
	cmd.SetArgs([]string{"--package", "eslint", "--package-manager", "pnpm"})
	if _, err := captureStderr(t, cmd.Execute); err == nil || !strings.Contains(err.Error(), "invalid package manager") {
		t.Fatalf("expected invalid package manager error, got %v", err)
	}
}

func TestPrintDockerfileMode(t *testing.T) {
	oldEnsure := ensureCommandFn
	oldBuild := buildWithOptionsFn
	t.Cleanup(func() {
		ensureCommandFn = oldEnsure
		buildWithOptionsFn = oldBuild
	})

	ensureCommandFn = func(string) error {
		t.Fatal("docker check should not run in print-only mode")
		return nil
	}
	buildWithOptionsFn = func(buildFlags) error {
		t.Fatal("build should not be invoked in print-only mode")
		return nil
	}

	cmd := newBuildCmd()
	cmd.SetArgs([]string{"--package", "eslint", "--bin", "eslint", "--image", "acme/eslint", "--print-dockerfile"})
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)
	_, err := captureStderr(t, cmd.Execute)
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	output := stdout.String()
	if !strings.Contains(output, "FROM "+defaultNPMBase) {
		t.Fatalf("expected Dockerfile to include base image, got %q", output)
	}
	if !strings.Contains(output, "RUN npm install -g eslint") {
		t.Fatalf("expected Dockerfile to include npm install, got %q", output)
	}
	if !strings.Contains(output, "ENTRYPOINT [\"eslint\"]") {
		t.Fatalf("expected Dockerfile to include entrypoint, got %q", output)
	}
}

func TestPrintDockerfileWarningsToStderr(t *testing.T) {
	oldEnsure := ensureCommandFn
	oldBuild := buildWithOptionsFn
	t.Cleanup(func() {
		ensureCommandFn = oldEnsure
		buildWithOptionsFn = oldBuild
	})

	ensureCommandFn = func(string) error {
		t.Fatal("docker check should not run in print-only mode")
		return nil
	}
	buildWithOptionsFn = func(buildFlags) error {
		t.Fatal("build should not be invoked in print-only mode")
		return nil
	}

	cmd := newBuildCmd()
	cmd.SetArgs([]string{"--package", "eslint", "--print-dockerfile"})
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)
	stderr, err := captureStderr(t, cmd.Execute)
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !strings.Contains(stderr, "warning: --image") || !strings.Contains(stderr, "warning: --bin") {
		t.Fatalf("expected derivation warnings on stderr, got %q", stderr)
	}
	if strings.Contains(stdout.String(), "warning:") {
		t.Fatalf("expected stdout to contain only Dockerfile content, got %q", stdout.String())
	}
}

func captureStderr(t *testing.T, fn func() error) (string, error) {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe stderr: %v", err)
	}
	defer reader.Close()
	oldStderr := os.Stderr
	os.Stderr = writer
	t.Cleanup(func() { os.Stderr = oldStderr })

	runErr := fn()
	if closeErr := writer.Close(); closeErr != nil && runErr == nil {
		runErr = closeErr
	}
	output, readErr := io.ReadAll(reader)
	if readErr != nil {
		t.Fatalf("read stderr: %v", readErr)
	}
	return string(output), runErr
}

// TestPackageNameAndVersion verifies package/version parsing.
func TestPackageNameAndVersion(t *testing.T) {
	cases := []struct {
		name            string
		input           string
		expectedName    string
		expectedVersion string
	}{
		{name: "plain", input: "eslint", expectedName: "eslint", expectedVersion: ""},
		{name: "plain-version", input: "eslint@9.1.0", expectedName: "eslint", expectedVersion: "9.1.0"},
		{name: "scoped", input: "@acme/eslint", expectedName: "@acme/eslint", expectedVersion: ""},
		{name: "scoped-version", input: "@acme/eslint@1.2.3", expectedName: "@acme/eslint", expectedVersion: "1.2.3"},
	}
	for _, tc := range cases {
		name, version := packageNameAndVersion(tc.input)
		if name != tc.expectedName || version != tc.expectedVersion {
			t.Fatalf("%s: expected %q/%q, got %q/%q", tc.name, tc.expectedName, tc.expectedVersion, name, version)
		}
	}
}

// TestOriginLabelPairs verifies label generation for origin metadata.
func TestOriginLabelPairs(t *testing.T) {
	cases := []struct {
		name     string
		pkgSpec  string
		bin      string
		expected []string
	}{
		{
			name:    "unversioned",
			pkgSpec: "eslint",
			bin:     "eslint",
			expected: []string{
				`io.cli2docker.package="eslint"`,
				`io.cli2docker.bin="eslint"`,
			},
		},
		{
			name:    "versioned",
			pkgSpec: "@acme/eslint@1.2.3",
			bin:     "eslint-cli",
			expected: []string{
				`io.cli2docker.package="@acme/eslint"`,
				`io.cli2docker.package-version="1.2.3"`,
				`io.cli2docker.bin="eslint-cli"`,
			},
		},
	}
	for _, tc := range cases {
		got := originLabelPairs(tc.pkgSpec, tc.bin)
		if len(got) != len(tc.expected) {
			t.Fatalf("%s: expected %d labels, got %d", tc.name, len(tc.expected), len(got))
		}
		for i := range got {
			if got[i] != tc.expected[i] {
				t.Fatalf("%s: expected %q, got %q", tc.name, tc.expected[i], got[i])
			}
		}
	}
}

// TestOriginCommentLines verifies shim comment formatting.
func TestOriginCommentLines(t *testing.T) {
	cases := []struct {
		name     string
		labels   map[string]string
		expected []string
	}{
		{
			name: "no-version",
			labels: map[string]string{
				labelPackage: "eslint",
				labelBin:     "eslint",
			},
			expected: []string{
				"# io.cli2docker.package=eslint",
				"# io.cli2docker.bin=eslint",
			},
		},
		{
			name: "with-version",
			labels: map[string]string{
				labelPackage:        "@acme/eslint",
				labelPackageVersion: "1.2.3",
				labelBin:            "eslint-cli",
			},
			expected: []string{
				"# io.cli2docker.package=@acme/eslint",
				"# io.cli2docker.package-version=1.2.3",
				"# io.cli2docker.bin=eslint-cli",
			},
		},
	}
	for _, tc := range cases {
		got := originCommentLines(tc.labels)
		if len(got) != len(tc.expected) {
			t.Fatalf("%s: expected %d lines, got %d", tc.name, len(tc.expected), len(got))
		}
		for i := range got {
			if got[i] != tc.expected[i] {
				t.Fatalf("%s: expected %q, got %q", tc.name, tc.expected[i], got[i])
			}
		}
	}
}

func TestBuildShimScriptIncludesOriginComments(t *testing.T) {
	labels := map[string]string{
		labelPackage:        "@acme/eslint",
		labelPackageVersion: "1.2.3",
		labelBin:            "eslint-cli",
	}
	execLine := "exec docker run --rm ${tty_flags} \"${image_ref}\" \"$@\""
	script := buildShimScript("acme/eslint:latest", execLine, labels)
	if !strings.Contains(script, "# io.cli2docker.package=@acme/eslint") {
		t.Fatalf("expected package comment in shim script")
	}
	if !strings.Contains(script, "# io.cli2docker.package-version=1.2.3") {
		t.Fatalf("expected package-version comment in shim script")
	}
	if !strings.Contains(script, "# io.cli2docker.bin=eslint-cli") {
		t.Fatalf("expected bin comment in shim script")
	}
}

// TestWriteDockerfileIncludesOriginLabels verifies Dockerfile label output.
func TestWriteDockerfileIncludesOriginLabels(t *testing.T) {
	tmpDir := t.TempDir()
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	opts := buildFlags{
		Package:        "eslint",
		Bin:            "eslint",
		Base:           defaultNPMBase,
		User:           defaultNPMUser,
		PackageManager: packageManagerNPM,
	}
	if err := writeDockerfile(dockerfilePath, opts); err != nil {
		t.Fatalf("writeDockerfile: %v", err)
	}
	content, err := os.ReadFile(dockerfilePath)
	if err != nil {
		t.Fatalf("read dockerfile: %v", err)
	}
	labelLine := ""
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "LABEL ") {
			labelLine = line
			break
		}
	}
	if labelLine == "" {
		t.Fatal("expected Dockerfile to include LABEL instruction")
	}
	if !strings.Contains(labelLine, `io.cli2docker.package="eslint"`) {
		t.Fatalf("expected package label in %q", labelLine)
	}
	if !strings.Contains(labelLine, `io.cli2docker.bin="eslint"`) {
		t.Fatalf("expected bin label in %q", labelLine)
	}
	if strings.Contains(labelLine, "io.cli2docker.package-version") {
		t.Fatalf("did not expect package-version label in %q", labelLine)
	}
}

func TestWriteDockerfileUsesNpmInstall(t *testing.T) {
	tmpDir := t.TempDir()
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	opts := buildFlags{
		Package:        "eslint",
		Bin:            "eslint",
		Base:           defaultNPMBase,
		User:           defaultNPMUser,
		PackageManager: packageManagerNPM,
	}
	if err := writeDockerfile(dockerfilePath, opts); err != nil {
		t.Fatalf("writeDockerfile: %v", err)
	}
	content, err := os.ReadFile(dockerfilePath)
	if err != nil {
		t.Fatalf("read dockerfile: %v", err)
	}
	text := string(content)
	if !strings.Contains(text, "npm install -g") {
		t.Fatalf("expected npm install line, got %q", text)
	}
	if !strings.Contains(text, "NPM_CONFIG_FUND=false") {
		t.Fatalf("expected npm env config, got %q", text)
	}
	if strings.Contains(text, "bun add -g") {
		t.Fatalf("did not expect bun install line, got %q", text)
	}
}

func TestWriteDockerfileUsesBunInstall(t *testing.T) {
	tmpDir := t.TempDir()
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	opts := buildFlags{
		Package:        "eslint",
		Bin:            "eslint",
		Base:           defaultBunBase,
		User:           defaultBunUser,
		PackageManager: packageManagerBun,
	}
	if err := writeDockerfile(dockerfilePath, opts); err != nil {
		t.Fatalf("writeDockerfile: %v", err)
	}
	content, err := os.ReadFile(dockerfilePath)
	if err != nil {
		t.Fatalf("read dockerfile: %v", err)
	}
	text := string(content)
	if !strings.Contains(text, "bun add -g") {
		t.Fatalf("expected bun install line, got %q", text)
	}
	if strings.Contains(text, "NPM_CONFIG_FUND=false") {
		t.Fatalf("did not expect npm env config, got %q", text)
	}
	if strings.Contains(text, "npm install -g") {
		t.Fatalf("did not expect npm install line, got %q", text)
	}
}
