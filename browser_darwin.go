// +build darwin

package browser

import (
	"os/exec"
)

// Open opens a url with the default browser.
func Open(u string) (err error) {
	_, err = exec.Command("open", u).Output()
	return
}
