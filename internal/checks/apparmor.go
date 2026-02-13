package checks

import "awesomeProject/internal/models"

func CheckAppArmor() models.CISCheckResult {
	out := run("systemctl", "is-active", "apparmor")

	if out == "active" {
		return models.CISCheckResult{
			CheckID:   "CIS-APPARMOR",
			CheckName: "AppArmor enabled",
			Status:    "PASS",
			Evidence:  "apparmor is active",
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-APPARMOR",
		CheckName: "AppArmor enabled",
		Status:    "FAIL",
		Evidence:  "apparmor is not active: " + out,
	}
}
