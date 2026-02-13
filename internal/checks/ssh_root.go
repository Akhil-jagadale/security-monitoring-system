package checks

import (
	"awesomeProject/internal/models"
	"os"
	"strings"
)

func CheckSSHRootLogin() models.CISCheckResult {
	file := "/etc/ssh/sshd_config"
	data, err := os.ReadFile(file)

	if err != nil {
		return models.CISCheckResult{
			CheckID:   "CIS-SSH-ROOT",
			CheckName: "Root login disabled over SSH",
			Status:    "FAIL",
			Evidence:  "Could not read sshd_config: " + err.Error(),
		}
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue // Skip comments
		}
		if strings.Contains(line, "PermitRootLogin no") {
			return models.CISCheckResult{
				CheckID:   "CIS-SSH-ROOT",
				CheckName: "Root login disabled over SSH",
				Status:    "PASS",
				Evidence:  "PermitRootLogin no found in sshd_config",
			}
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-SSH-ROOT",
		CheckName: "Root login disabled over SSH",
		Status:    "FAIL",
		Evidence:  "PermitRootLogin no NOT found in sshd_config",
	}
}
