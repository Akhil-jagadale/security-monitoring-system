package checks

import "awesomeProject/internal/models"

func RunAllChecks() []models.CISCheckResult {
	return []models.CISCheckResult{
		CheckSSHRootLogin(),
		CheckFirewallEnabled(),
		CheckTimeSync(),
		CheckAuditd(),
		CheckAppArmor(),
		CheckPasswordExpiry(),
		CheckPasswordComplexity(),
		CheckWorldWritableFiles(),
		CheckCramfsDisabled(),
		CheckGDMLogin(),
	}
}
