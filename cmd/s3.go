package cmd

import (
	"fmt"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var s3Cmd = &cobra.Command{
	Use:   "s3 [base]",
	Short: "Enumerate S3 buckets from a base name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results := checks.EnumerateS3(args[0], 8*time.Second)
		fmt.Printf("\nS3 enumeration for %s\n", args[0])
		for _, r := range results {
			fmt.Printf("  %-30s HTTP %d -> %s", r.Bucket, r.HTTP, r.Status)
			if r.Listable {
				fmt.Print(" (listable!)")
			}
			fmt.Println()
		}
	},
}

func init() { rootCmd.AddCommand(s3Cmd) }
