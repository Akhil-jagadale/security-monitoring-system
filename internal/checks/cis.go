package checks

import (
	"bytes"
	"os/exec"
	"strings"
)

func run(cmd string, args ...string) string {
	c := exec.Command(cmd, args...)
	var out bytes.Buffer
	c.Stdout = &out
	c.Stderr = &out
	_ = c.Run()
	return strings.TrimSpace(out.String())
}
