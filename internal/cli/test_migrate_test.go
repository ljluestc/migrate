package cli_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestCLIBinaryWorks(t *testing.T) {
	// Skip if not in a normal build environment
	if os.Getenv("SKIP_BUILD_TESTS") != "" {
		t.Skip("Skipping test that requires build")
	}
	
	// Build the migrate CLI
	cmd := exec.Command("go", "build", "-o", "migrate-test", "../cmd/migrate")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build migrate CLI: %v", err)
	}
	defer os.Remove("migrate-test")
	
	// Test the version command
	versionCmd := exec.Command("./migrate-test", "version")
	output, err := versionCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run version command: %v\nOutput: %s", err, output)
	}
	
	// Check that output contains version information
	if len(output) == 0 {
		t.Error("Expected version output, got empty response")
	}
}
