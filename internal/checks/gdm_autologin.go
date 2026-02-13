package checks

import "awesomeProject/internal/models"

func CheckGDMLogin() models.CISCheckResult {
	out := run("bash", "-c", "grep -R 'AutomaticLoginEnable.*true' /etc/gdm3/ 2>/dev/null")

	if out == "" {
		return models.CISCheckResult{
			CheckID:   "CIS-GDM-AUTOLOGIN",
			CheckName: "GDM auto-login disabled",
			Status:    "PASS",
			Evidence:  "No AutomaticLoginEnable=true found",
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-GDM-AUTOLOGIN",
		CheckName: "GDM auto-login disabled",
		Status:    "FAIL",
		Evidence:  out,
	}
}
