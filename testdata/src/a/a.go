package a

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {

	apikey := "123"
	pwd := "124123wrHR123dgzS"

	slog.Info("First")            // want "log message should start with lowercase letter"
	slog.Error("spec symbols,,,") // want "log message contains forbidden characters or emoji"
	slog.Warn("Русский язык") // want "log message should contain only english characters"
	slog.Error("api key" + pwd)   // want "log message may contain sensitive data"

	log, err := zap.NewProduction()
	if err != nil {
		return
	}

	log.Info("First")            // want "log message should start with lowercase letter"
	log.Error("spec symbols,,,") // want "log message contains forbidden characters or emoji"
	log.Warn("Русский язык") // want "log message should contain only english characters"
	log.Info("user" + apikey)    // want "log message may contain sensitive data"
}
