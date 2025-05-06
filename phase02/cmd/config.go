package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MatheusOliveiraSilva/infra-with-go/phase02/pkg/config"
	"github.com/MatheusOliveiraSilva/infra-with-go/phase02/pkg/logger"
	"github.com/spf13/cobra"
)

// configCmd representa o comando config
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long: `Display or modify the application configuration settings.
	
Examples:
  infra config show                 # Display current configuration
  infra config init --file=my.yaml  # Create a new configuration file`,
}

// configShowCmd representa o subcomando config show
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Obter a configuração atual
		cfg := config.GetConfig()

		// Formatar como JSON para exibição
		jsonData, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return fmt.Errorf("error formatting configuration: %w", err)
		}

		// Exibir a configuração
		fmt.Println(string(jsonData))
		return nil
	},
}

// configInitCmd representa o subcomando config init
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputFile, _ := cmd.Flags().GetString("file")

		// Se não foi especificado um arquivo, usar o padrão
		if outputFile == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("error getting home directory: %w", err)
			}

			// Criar diretório .infra se não existir
			configDir := filepath.Join(homeDir, ".infra")
			if err := os.MkdirAll(configDir, 0755); err != nil {
				return fmt.Errorf("error creating config directory: %w", err)
			}

			outputFile = filepath.Join(configDir, "config.yaml")
		}

		// Verificar se o arquivo já existe
		if _, err := os.Stat(outputFile); err == nil {
			return fmt.Errorf("configuration file already exists: %s", outputFile)
		}

		// Criar o arquivo de configuração
		if err := config.Save(outputFile); err != nil {
			return fmt.Errorf("error saving configuration: %w", err)
		}

		logger.Log.WithFields(map[string]interface{}{
			"file": outputFile,
		}).Info("Configuration file created")

		return nil
	},
}

func init() {
	// Adicionar comando config ao comando raiz
	rootCmd.AddCommand(configCmd)

	// Adicionar subcomandos ao comando config
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configInitCmd)

	// Adicionar flags
	configInitCmd.Flags().StringP("file", "f", "", "Output file path")
}
