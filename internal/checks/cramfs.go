package checks

import (
	"awesomeProject/internal/models"
	"strings"
)

func CheckCramfsDisabled() models.CISCheckResult {
	out := run("bash", "-c", "modprobe -n -v cramfs 2>&1")

	if strings.Contains(out, "install /bin/true") || strings.Contains(out, "not found") {
		return models.CISCheckResult{
			CheckID:   "CIS-CRAMFS",
			CheckName: "Unused filesystem cramfs disabled",
			Status:    "PASS",
			Evidence:  out,
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-CRAMFS",
		CheckName: "Unused filesystem cramfs disabled",
		Status:    "FAIL",
		Evidence:  out,
	}
}
