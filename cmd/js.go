package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var jsCmd = &cobra.Command{
	Use:   "js [url]",
	Short: "Fu-JS technique: scan page scripts for secrets and endpoints",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results := checks.CheckJSSecrets(args[0], 10*time.Second)
		b, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(b))
		if len(results) == 0 {
			fmt.Println("no secrets or endpoints found in same-origin JS")
		}
	},
}

func init() { rootCmd.AddCommand(jsCmd) }
