package main

import (
	"awesomeProject/internal/checks"
	"awesomeProject/internal/collector"
	"awesomeProject/internal/models"
	"awesomeProject/internal/sender"
	"fmt"
	"os"
)

func main() {
	fmt.Println("...Starting Linux Security Agent...")
	fmt.Println("=======================================")

	// Get API URL from environment or use default
	apiURL := os.Getenv("AGENT_API_URL")
	if apiURL == "" {
		// ⚠️ REPLACE THIS WITH YOUR ACTUAL API GATEWAY URL
		apiURL = "https://grg7og4c2d.execute-api.us-east-1.amazonaws.com/prod"
	}

	fmt.Println("📊 Collecting host information...")
	host := collector.CollectHostInfo()
	fmt.Printf("   ✓ Hostname: %s\n", host.Hostname)

	fmt.Println("📦 Collecting installed packages...")
	packages := collector.CollectPackages()
	fmt.Printf("   ✓ Found %d packages\n", len(packages))

	fmt.Println("🔒 Running CIS security checks...")
	cisResults := checks.RunAllChecks()

	passCount := 0
	for _, result := range cisResults {
		if result.Status == "PASS" {
			fmt.Printf("   ✓ %s\n", result.CheckName)
			passCount++
		} else {
			fmt.Printf("   ✗ %s\n", result.CheckName)
		}
	}
	fmt.Printf("   Score: %d/%d checks passed\n", passCount, len(cisResults))

	report := models.NewReport(host, packages, cisResults)

	fmt.Println("☁️  Sending report to AWS...")
	err := sender.SendReport(apiURL, report)
	if err != nil {
		fmt.Printf("   ✗ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("   ✓ Report successfully sent to AWS!")
	fmt.Println("=======================================")
	fmt.Println("✅ Agent execution completed successfully")
}
