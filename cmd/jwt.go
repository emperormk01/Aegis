package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt [url]",
	Short: "Discover JWTs in cookies and page, decode header, flag weak algs",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := checks.CheckJWT(args[0], 8*time.Second)
		b, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(b))
		if len(res.Tokens) == 0 {
			fmt.Println("no JWTs discovered")
		}
	},
}

func init() { rootCmd.AddCommand(jwtCmd) }
