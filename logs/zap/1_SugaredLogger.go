package main 


// go get -u go.uber.org/zap 

import (
	"go.uber.org/zap"
)

func main() {
	// добавляем предустановленный логер NewDevelopment
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("cannot initialize zap")
	}

	// это нужно добавить, если логер буферизован
    // в данном случае не буферизован, но привычка хорошая
	defer logger.Sync()

	const url = "http://example.com"

	// делаем логер SugaredLogger
	sugar := logger.Sugar()

	// выводим сообщение уровня Info с парой "url": url в виде JSON, это SugaredLogger

	sugar.Infow(
		"Failed to fetch url",
		"url", url,
	)
	// 2025-10-17T12:53:56.043+0300
	// INFO  zap/1_SugaredLogger.go:28
	// Failed to fetch url
	// {"url": "http://example.com"}

	// выводим сообщение уровня Info, но со строкой URL, это тоже SugaredLogger
	sugar.Infof("Failed to fetch URL: %s", url)

	// выводим сообщение уровня Error со строкой URL, и это SugaredLogger
	sugar.Errorf("Failed to fetch URL: %s", url)

	// переводим в обычный Logger
	plain := sugar.Desugar()

	// выводим сообщение уровня Info обычного регистратора (не SugaredLogger)
	plain.Info("Hello, GO")
	// 2025-10-17T12:53:56.044+0300
	// INFO zap/1_SugaredLogger.go:43
	// Hello, GO


	// также уровня Warn (не SugaredLogger)
	plain.Warn("Simple warning")
	// и уровня Error, но добавляем строго типизированное поле "url" (не SugaredLogger)
	plain.Error("Failed to fetch URL", zap.String("url", url))


}