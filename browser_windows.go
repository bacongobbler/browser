// +build windows

package browser

import (
	"os/exec"
)

// Open opens a URL with the default browser.
func Open(u string) (err error) {
	_, err = exec.Command("cmd", "/c", "start", u).Output()
	return
}
