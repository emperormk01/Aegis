package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aegis",
	Short: "Security audit CLI - S3, headers, race conditions, recon",
	Long:  "Aegis - security audit CLI (Go). Combines bugbounty-findings, kambegoye-scan, and security-research into one tool.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
