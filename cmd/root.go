package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "PassMan",
	Short: "A CLi based Password Manager",
	Long:  "A secure and user-friendly CLI password manager.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}

func init() {}
