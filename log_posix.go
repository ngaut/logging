// +build !windows

package logging

import (
	"os"
	"syscall"
)

func dup2(oldfd, newfd *os.File) error {
	return syscall.Dup2(oldfd.Fd(), newfd.Fd())
}
