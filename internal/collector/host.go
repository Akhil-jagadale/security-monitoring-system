package collector

import (
	"awesomeProject/internal/models"
	"bytes"
	"os/exec"
	"strings"
)

func getCommandOutput(cmd string, args ...string) string {
	c := exec.Command(cmd, args...)
	var out bytes.Buffer
	c.Stdout = &out
	_ = c.Run()
	return strings.TrimSpace(out.String())
}

func CollectHostInfo() models.HostInfo {
	hostname := getCommandOutput("hostname")
	kernel := getCommandOutput("uname", "-r")
	os := getCommandOutput("lsb_release", "-d")

	// Clean up OS string
	os = strings.ReplaceAll(os, "Description:", "")
	os = strings.TrimSpace(os)

	ip := getCommandOutput("hostname", "-I")

	return models.HostInfo{
		Hostname:  hostname,
		Kernel:    kernel,
		OS:        os,
		IPAddress: strings.Fields(ip)[0], // Get first IP only
	}
}
