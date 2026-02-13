package checks

import "awesomeProject/internal/models"

func CheckWorldWritableFiles() models.CISCheckResult {
	out := run("bash", "-c", "find /tmp /var/tmp -xdev -type f -perm -0002 2>/dev/null | head -n 5")

	if out == "" {
		return models.CISCheckResult{
			CheckID:   "CIS-WORLD-WRITABLE",
			CheckName: "No world writable files in tmp",
			Status:    "PASS",
			Evidence:  "No world-writable files found in /tmp and /var/tmp",
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-WORLD-WRITABLE",
		CheckName: "No world writable files in tmp",
		Status:    "FAIL",
		Evidence:  "Found world-writable files:\n" + out,
	}
}
