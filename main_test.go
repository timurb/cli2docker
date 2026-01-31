package main

import (
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
	if !strings.Contains(stderr, "--image") || !strings.Contains(stderr, "--bin") {
		t.Fatalf("expected warnings for both image and bin derivation, got %q", stderr)
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
		Package: "eslint",
		Bin:     "eslint",
		Base:    "node:20-alpine",
		User:    "node",
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
