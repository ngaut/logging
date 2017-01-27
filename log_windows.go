// +build windows

package logging

import (
	"io"
	"os"
)

func dup2(oldfd, newfd *os.File) error {
	go io.Copy(newfd, oldfd)
	return nil
}
