// +build linux

package browser

import (
	"os/exec"
)

// Open opens a URL with the default browser.
func Open(u string) (err error) {
	_, err = exec.Command("xdg-open", u).Output()
	return
}
