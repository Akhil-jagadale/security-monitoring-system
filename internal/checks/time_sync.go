package checks

import "awesomeProject/internal/models"

func CheckTimeSync() models.CISCheckResult {
	out := run("systemctl", "is-active", "chrony")

	if out == "active" {
		return models.CISCheckResult{
			CheckID:   "CIS-TIME-SYNC",
			CheckName: "Time synchronization configured",
			Status:    "PASS",
			Evidence:  "chrony is active",
		}
	}

	// Try cron as alternative
	out2 := run("systemctl", "is-active", "chronyd")
	if out2 == "active" {
		return models.CISCheckResult{
			CheckID:   "CIS-TIME-SYNC",
			CheckName: "Time synchronization configured",
			Status:    "PASS",
			Evidence:  "chronyd is active",
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-TIME-SYNC",
		CheckName: "Time synchronization configured",
		Status:    "FAIL",
		Evidence:  "chrony/chronyd not active: " + out,
	}
}
