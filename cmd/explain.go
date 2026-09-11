package cmd

import (
	"fmt"
	"strings"

	"github.com/emperormk01/Aegis/internal/corpus"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain [keyword]",
	Short: "Search the ingested bugbounty and research corpus",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kw := args[0]
		hits := corpus.Search(kw)
		if len(hits) == 0 {
			fmt.Printf("No hits for %q. Available corpus files:\n", kw)
			for _, p := range corpus.List() {
				fmt.Println("  ", p)
			}
			return
		}
		for _, p := range hits {
			content, _ := corpus.Read(p)
			// snippet
			lines := strings.Split(content, "\n")
			snippet := strings.Join(lines[:min(40, len(lines))], "\n")
			fmt.Printf("\n=== %s ===\n%s\n--- (%d more lines) ---\n", p, snippet, len(lines)-min(40, len(lines)))
		}
	},
}

var corpusListCmd = &cobra.Command{
	Use:   "corpus",
	Short: "List all ingested reports",
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range corpus.List() {
			fmt.Println(p)
		}
	},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(corpusListCmd)
}
