package checks

import "awesomeProject/internal/models"

func CheckAuditd() models.CISCheckResult {
	out := run("systemctl", "is-active", "auditd")

	if out == "active" {
		return models.CISCheckResult{
			CheckID:   "CIS-AUDITD",
			CheckName: "Auditd service running",
			Status:    "PASS",
			Evidence:  "auditd is active",
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-AUDITD",
		CheckName: "Auditd service running",
		Status:    "FAIL",
		Evidence:  "auditd not running: " + out,
	}
}
