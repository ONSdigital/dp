package launch

import (
	"fmt"
	"os/exec"
	"runtime"
)

func browserCmd() (string, bool) {
	browser := map[string]string{
		"darwin": "open",
		"linux":  "xdg-open",
		"win32":  "start",
	}
	cmd, ok := browser[runtime.GOOS]
	return cmd, ok
}

// Browser opens the given URL
func Browser(url string) error {
	browser, ok := browserCmd()
	if !ok {
		return fmt.Errorf("unable to launch browser for this OS: %s", runtime.GOOS)
	}

	cmd := exec.Command(browser, url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error from command %s: %s", err.Error(), string(output))
	}
	return nil
}
