//go:build linux
// +build linux

package internal

import "os/exec"

func OpenFolder(path string) error {
	return exec.Command("xdg-open", path).Start()
}
