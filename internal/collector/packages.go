package collector

import (
	"awesomeProject/internal/models"
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func CollectPackages() []models.PackageInfo {
	cmd := exec.Command("dpkg-query", "-W", "-f=${binary:Package} ${Version}\n")

	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()

	scanner := bufio.NewScanner(&out)
	var packages []models.PackageInfo

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) >= 2 {
			packages = append(packages, models.PackageInfo{
				Name:    parts[0],
				Version: parts[1],
			})
		}
	}

	// Limit to first 200 packages to reduce payload size
	if len(packages) > 200 {
		packages = packages[:200]
	}

	return packages
}
