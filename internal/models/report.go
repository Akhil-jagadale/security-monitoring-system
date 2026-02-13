package models

import "time"

type HostInfo struct {
	Hostname  string `json:"hostname"`
	OS        string `json:"os"`
	Kernel    string `json:"kernel"`
	IPAddress string `json:"ip_address"`
}

type PackageInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CISCheckResult struct {
	CheckID   string `json:"check_id"`
	CheckName string `json:"check_name"`
	Status    string `json:"status"` // PASS or FAIL
	Evidence  string `json:"evidence"`
}

type Report struct {
	Timestamp string           `json:"timestamp"`
	Host      HostInfo         `json:"host"`
	Packages  []PackageInfo    `json:"packages"`
	CIS       []CISCheckResult `json:"cis_results"`
}

func NewReport(host HostInfo, packages []PackageInfo, cis []CISCheckResult) Report {
	return Report{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Host:      host,
		Packages:  packages,
		CIS:       cis,
	}
}
