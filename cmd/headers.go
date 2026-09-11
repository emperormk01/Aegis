package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var headersCmd = &cobra.Command{
	Use:   "headers [url]",
	Short: "Check security headers and sensitive files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := checks.CheckHeaders(args[0], 10*time.Second)
		b, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(b))
		sens := checks.CheckSensitive(args[0], 8*time.Second)
		fmt.Println("\nSensitive files:")
		for _, r := range sens {
			fmt.Printf("  %s -> %d exposed=%v\n", r.Path, r.Status, r.Exposed)
		}
	},
}

func init() { rootCmd.AddCommand(headersCmd) }
