package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var (
	raceMethod string
	raceCount  int
	raceData   string
)

var raceCmd = &cobra.Command{
	Use:   "race [url]",
	Short: "Fire parallel requests to test race conditions",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body := []byte(raceData)
		// try to handle JSON data gracefully but send as-is
		results := checks.Race(raceMethod, url, raceCount, body, nil, 10*time.Second)
		summary := checks.AnalyzeRace(results)
		b, _ := json.MarshalIndent(map[string]interface{}{"summary": summary, "results": results}, "", "  ")
		fmt.Println(string(b))
		if summary.Diverged {
			fmt.Println("\n[WARN] Diverged responses - possible race condition")
		} else {
			fmt.Println("\n[OK] No divergence in response lengths")
		}
	},
}

func init() {
	raceCmd.Flags().StringVar(&raceMethod, "method", "POST", "HTTP method")
	raceCmd.Flags().IntVar(&raceCount, "count", 25, "Parallel requests")
	raceCmd.Flags().StringVar(&raceData, "data", "", "Body data (JSON string or raw)")
	rootCmd.AddCommand(raceCmd)
}
