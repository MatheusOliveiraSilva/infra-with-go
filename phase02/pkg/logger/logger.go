package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// Log é a instância global do logger
	Log *logrus.Logger
)

// Init inicializa o logger com a configuração especificada
func Init(level string, formatter string) {
	Log = logrus.New()

	// Configurar o destino de saída
	Log.SetOutput(os.Stdout)

	// Configurar o nível de log
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	// Configurar o formato do log
	switch formatter {
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{})
	default:
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}

// GetWriter retorna um io.Writer para usar com outros sistemas de log
func GetWriter() io.Writer {
	return Log.Writer()
}

// Fields cria um entry do logrus com campos adicionais
func Fields(fields map[string]interface{}) *logrus.Entry {
	return Log.WithFields(fields)
}
