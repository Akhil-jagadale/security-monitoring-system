package checks

import "awesomeProject/internal/models"

func CheckFirewallEnabled() models.CISCheckResult {
	out := run("ufw", "status")

	if out == "" {
		return models.CISCheckResult{
			CheckID:   "CIS-FW-UFW",
			CheckName: "Firewall enabled (UFW)",
			Status:    "FAIL",
			Evidence:  "ufw command failed or not installed",
		}
	}

	if out == "Status: active" || (len(out) >= 14 && out[:14] == "Status: active") {
		return models.CISCheckResult{
			CheckID:   "CIS-FW-UFW",
			CheckName: "Firewall enabled (UFW)",
			Status:    "PASS",
			Evidence:  out,
		}
	}

	return models.CISCheckResult{
		CheckID:   "CIS-FW-UFW",
		CheckName: "Firewall enabled (UFW)",
		Status:    "FAIL",
		Evidence:  out,
	}
}
