package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a Docker image",
	Long: `Build a Docker image using the specified Dockerfile and context.
	
Example usage:
  infra build -t myapp:latest -f ./Dockerfile ./`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get flag values
		tag, _ := cmd.Flags().GetString("tag")
		dockerfile, _ := cmd.Flags().GetString("file")
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Context path is the first argument
		context := "."
		if len(args) > 0 {
			context = args[0]
		}

		// Log the build configuration
		if verbose {
			fmt.Printf("Building image with tag: %s\n", tag)
			fmt.Printf("Using Dockerfile: %s\n", dockerfile)
			fmt.Printf("Build context: %s\n", context)
		}

		// TODO: Implement actual Docker build logic
		fmt.Println("Building Docker image...")

		return nil
	},
}

func init() {
	// Add build command to the root command
	rootCmd.AddCommand(buildCmd)

	// Add flags specific to the build command
	buildCmd.Flags().StringP("tag", "t", "", "Name and optionally a tag in the 'name:tag' format")
	buildCmd.Flags().StringP("file", "f", "Dockerfile", "Name of the Dockerfile")

	// Mark required flags
	buildCmd.MarkFlagRequired("tag")
}
