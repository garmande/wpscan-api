package scanner

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// Run executes wpscan against the given URL and returns the parsed JSON output.
func Run(url string) (any, error) {
	cmd := exec.Command("wpscan", "--url", url, "--format", "json", "--no-update")
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("wpscan exited with code %d: %s", exitErr.ExitCode(), string(exitErr.Stderr))
		}
		return nil, fmt.Errorf("failed to run wpscan: %w", err)
	}

	var result any
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to parse wpscan output: %w", err)
	}

	return result, nil
}
