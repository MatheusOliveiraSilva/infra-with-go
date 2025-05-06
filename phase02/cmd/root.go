package cmd

import (
	"fmt"
	"os"

	"github.com/MatheusOliveiraSilva/infra-with-go/phase02/pkg/config"
	"github.com/MatheusOliveiraSilva/infra-with-go/phase02/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
	cobra.OnInitialize(initConfig)

	// Adicionar flags relacionadas à configuração
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	// Flags globais
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
}

// initConfig inicializa tanto o logger quanto as configurações
func initConfig() {
	// Verificar flag verbose
	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")

	// Inicializar logger baseado na flag verbose
	if verbose {
		logger.Init("debug", "text")
		logger.Log.Debug("Verbose logging enabled")
	} else {
		logger.Init("info", "text")
	}

	// Se um arquivo de configuração for especificado via flag, use-o
	if cfgFile != "" {
		logger.Log.WithFields(map[string]interface{}{
			"file": cfgFile,
		}).Debug("Using config file from flag")

		// Informar o Viper sobre o arquivo de configuração específico
		viper.SetConfigFile(cfgFile)
	}

	// Inicializar e carregar a configuração
	if err := config.Init(); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error": err,
		}).Error("Failed to load configuration")
		os.Exit(1)
	}

	// Obter a configuração e atualizar o logger se necessário
	cfg := config.GetConfig()

	// Se verbose não estiver ativo, use o nível de log da configuração
	if !verbose {
		logger.Init(cfg.LogLevel, cfg.LogFormat)
	}
}
