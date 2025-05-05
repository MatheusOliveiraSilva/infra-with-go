package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "infra",
	Short: "A CLI tool for managing infrastructure deployment",
	Long: `An infrastructure management CLI tool that simplifies Docker build, 
push and Kubernetes deployment operations.

This application helps automate and standardize the CI/CD process
for containerized applications.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Initialize configurations, flags, etc.

	// Example of a global flag
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
}
