package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	verboseFlag bool
	log         = logrus.New()
)

// rootCmd representa o comando base quando não há subcomandos
var rootCmd = &cobra.Command{
	Use:   "deployctl",
	Short: "Um CLI para automatizar deployments Docker e Kubernetes",
	Long: `deployctl é uma ferramenta de linha de comando que te ajuda a:
- Construir e publicar imagens Docker
- Gerar e aplicar manifestos Kubernetes
- Gerenciar o status de deployments e realizar rollbacks`,
	// Executado para todos os comandos
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Configurar o logger
		if verboseFlag {
			log.SetLevel(logrus.DebugLevel)
			log.Debug("Modo verbose ativado")
		} else {
			log.SetLevel(logrus.InfoLevel)
		}
	},
}

// Execute adiciona todos os comandos aos comandos raiz e configura flags apropriadamente.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Flags persistentes globais
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "arquivo de configuração (padrão é $HOME/.deployctl.yaml)")
	rootCmd.PersistentFlags().BoolVar(&verboseFlag, "verbose", false, "ativa logs detalhados")
	
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig lê o arquivo de configuração e variáveis de ambiente
func initConfig() {
	if cfgFile != "" {
		// Use o arquivo de configuração especificado pela flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Encontre o diretório home
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Procure por configuração em $HOME
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".deployctl")
	}

	viper.AutomaticEnv() // lê valores de variáveis de ambiente

	// Se um arquivo de configuração for encontrado, leia-o
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Usando arquivo de configuração:", viper.ConfigFileUsed())
	}
} 