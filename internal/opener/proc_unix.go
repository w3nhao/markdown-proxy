//go:build !windows

package opener

import (
	"os/exec"
	"syscall"
)

// setSysProcAttr configures the command to run in a new process group
// so it is not killed when the parent process exits.
func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}
