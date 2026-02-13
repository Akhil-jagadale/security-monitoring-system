package checks

import (
	"awesomeProject/internal/models"
	"strings"
)

func CheckPasswordExpiry() models.CISCheckResult {
	out := run("grep", "PASS_MAX_DAYS", "/etc/login.defs")

	if strings.Contains(out, "PASS_MAX_DAYS") && !strings.HasPrefix(strings.TrimSpace(out), "#") {
		return models.CISCheckResult{
			CheckID:   "CIS-PASS-EXP",
			CheckName: "Password expiration policy enforced",
			Status:    "PASS",
			Evidence:  out,
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-PASS-EXP",
		CheckName: "Password expiration policy enforced",
		Status:    "FAIL",
		Evidence:  "PASS_MAX_DAYS not configured properly",
	}
}
