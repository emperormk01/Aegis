package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var reconCmd = &cobra.Command{
	Use:   "recon [domain]",
	Short: "Subdomains, ports, and stack detection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		subs := checks.EnumerateSubdomains(domain, nil)
		ports := checks.ScanPorts(domain, nil, 1*time.Second)
		stack := checks.DetectStack("https://"+domain, 8*time.Second)
		b, _ := json.MarshalIndent(map[string]interface{}{"subdomains": subs, "ports": ports, "stack": stack}, "", "  ")
		fmt.Println(string(b))
	},
}

func init() { rootCmd.AddCommand(reconCmd) }
