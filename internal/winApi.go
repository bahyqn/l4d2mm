//go:build windows
// +build windows

package internal

import "os/exec"

func OpenFolder(path string) error {
	return exec.Command("explorer", path).Start()
}
