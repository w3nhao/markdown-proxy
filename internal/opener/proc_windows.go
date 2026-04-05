//go:build windows

package opener

import (
	"os/exec"
	"syscall"
)

// setSysProcAttr configures the command to run detached from the parent
// so it survives after the parent process exits.
func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}
