package cmd

import (
	"fmt"

	"github.com/MatheusOliveiraSilva/infra-with-go/phase02/pkg/logger"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a Docker image to a registry",
	Long: `Push a Docker image to a specified container registry.
	
Example usage:
  infra push myapp:latest`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if image name is provided
		if len(args) < 1 {
			return fmt.Errorf("image name is required")
		}

		imageName := args[0]
		registry, _ := cmd.Flags().GetString("registry")

		// Log the push configuration
		logFields := map[string]interface{}{
			"image": imageName,
		}

		if registry != "" {
			logFields["registry"] = registry
		}

		logger.Log.WithFields(logFields).Debug("Pushing Docker image")

		// TODO: Implement actual Docker push logic
		logger.Log.Info("Pushing Docker image...")

		return nil
	},
}

func init() {
	// Add push command to the root command
	rootCmd.AddCommand(pushCmd)

	// Add flags specific to the push command
	pushCmd.Flags().String("registry", "", "Container registry URL")
}
