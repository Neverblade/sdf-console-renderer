package visualizer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() error {
	osName := runtime.GOOS
	switch osName {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		return fmt.Errorf("unsupported OS: %s", osName)
	}
	return nil
}
