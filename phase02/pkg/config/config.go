package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MatheusOliveiraSilva/infra-with-go/phase02/pkg/logger"
	"github.com/spf13/viper"
)

// DockerConfig armazena configurações relacionadas ao Docker
type DockerConfig struct {
	Registry     string `mapstructure:"registry"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DefaultTag   string `mapstructure:"default_tag"`
	BuildTimeout int    `mapstructure:"build_timeout"`
}

// KubernetesConfig armazena configurações relacionadas ao Kubernetes
type KubernetesConfig struct {
	Context    string `mapstructure:"context"`
	Namespace  string `mapstructure:"namespace"`
	KubeConfig string `mapstructure:"kubeconfig"`
}

// AppConfig é a estrutura principal de configuração da aplicação
type AppConfig struct {
	Docker     DockerConfig     `mapstructure:"docker"`
	Kubernetes KubernetesConfig `mapstructure:"kubernetes"`
	LogLevel   string           `mapstructure:"log_level"`
	LogFormat  string           `mapstructure:"log_format"`
}

// Config é a instância global da configuração
var Config AppConfig

// Init inicializa e carrega a configuração
func Init() error {
	// Configurar o Viper
	viper.SetConfigName("config") // Nome do arquivo de configuração sem extensão
	viper.SetConfigType("yaml")   // Formato do arquivo (yaml, json, etc.)

	// Definir caminhos de busca
	viper.AddConfigPath(".")            // Diretório atual
	viper.AddConfigPath("$HOME/.infra") // Diretório home do usuário
	viper.AddConfigPath("/etc/infra")   // Diretório de sistema

	// Configurar o mapeamento de variáveis de ambiente
	viper.SetEnvPrefix("INFRA") // Prefixo para variáveis de ambiente
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // Ler automaticamente variáveis de ambiente

	// Configurar valores padrão
	setDefaults()

	// Tentar carregar o arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		// É aceitável não encontrar o arquivo, usaremos os valores padrão
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Log.Info("No configuration file found, using defaults")
		} else {
			// Outros erros devem ser reportados
			return fmt.Errorf("error reading config file: %w", err)
		}
	} else {
		logger.Log.WithFields(map[string]interface{}{
			"file": viper.ConfigFileUsed(),
		}).Debug("Configuration loaded from file")
	}

	// Mapear a configuração para nossa estrutura
	if err := viper.Unmarshal(&Config); err != nil {
		return fmt.Errorf("error unmarshaling config: %w", err)
	}

	return nil
}

// setDefaults configura os valores padrão de configuração
func setDefaults() {
	// Docker
	viper.SetDefault("docker.registry", "docker.io")
	viper.SetDefault("docker.default_tag", "latest")
	viper.SetDefault("docker.build_timeout", 300) // 5 minutos

	// Kubernetes
	viper.SetDefault("kubernetes.namespace", "default")

	// Logs
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_format", "text")

	// Tentar detectar o arquivo kubeconfig
	if home, err := os.UserHomeDir(); err == nil {
		viper.SetDefault("kubernetes.kubeconfig", filepath.Join(home, ".kube", "config"))
	}
}

// GetConfig retorna a configuração atual
func GetConfig() AppConfig {
	return Config
}

// Save salva a configuração atual em um arquivo
func Save(filepath string) error {
	viper.SetConfigFile(filepath)
	return viper.WriteConfig()
}
