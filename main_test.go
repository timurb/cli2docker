package main

import "testing"

// TestPackageBaseName verifies scope stripping.
func TestPackageBaseName(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "plain", input: "eslint", expected: "eslint"},
		{name: "scoped", input: "@acme/eslint", expected: "eslint"},
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
		expected string
	}{
		{name: "plain", input: "eslint", expected: "eslint"},
		{name: "scoped", input: "@acme/eslint", expected: "acme/eslint"},
	}
	for _, tc := range cases {
		if got := imageFromPackage(tc.input); got != tc.expected {
			t.Fatalf("%s: expected %q, got %q", tc.name, tc.expected, got)
		}
	}
}
