package checks

import (
	"awesomeProject/internal/models"
	"os"
	"strings"
)

func CheckPasswordComplexity() models.CISCheckResult {
	file := "/etc/pam.d/common-password"
	data, err := os.ReadFile(file)

	if err != nil {
		return models.CISCheckResult{
			CheckID:   "CIS-PASS-COMPLEX",
			CheckName: "Password complexity policy enabled",
			Status:    "FAIL",
			Evidence:  "Cannot read /etc/pam.d/common-password",
		}
	}

	content := string(data)

	if strings.Contains(content, "pam_pwquality.so") || strings.Contains(content, "pam_cracklib.so") {
		return models.CISCheckResult{
			CheckID:   "CIS-PASS-COMPLEX",
			CheckName: "Password complexity policy enabled",
			Status:    "PASS",
			Evidence:  "pam_pwquality.so or pam_cracklib.so configured",
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-PASS-COMPLEX",
		CheckName: "Password complexity policy enabled",
		Status:    "FAIL",
		Evidence:  "pam_pwquality.so not found",
	}
}
