package main

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var sugar zap.SugaredLogger

func main() {
	// создаём предустановленный регистратор zap
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// делаем регистратор SugaredLogger
	sugar = *logger.Sugar()

	http.Handle("/ping", WithLogging(pingHanlder()))

	addr := "127.0.0.1:8080"
	
	sugar.Infow(
		"Starting server",
		"addr", addr,
	)

	if err := http.ListenAndServe(addr, nil); err != nil {
		// записываем в лог ошибку, если сервер не запустился
		sugar.Fatalw(err.Error(), "event", "start server")
	}





}

func pingHanlder() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "pong\n")
	}
	return http.HandlerFunc(fn)
}

// WithLogging добавляет дополнительный код для регистрации сведений о запросе
// и возвращает новый http.Handler.
func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	uri := r.RequestURI

	method := r.Method

	// точка, где выполняется хендлер pingHandler 
	h.ServeHTTP(w, r)  // обслуживание оригинального запроса

	// Since возвращает разницу во времени между start 
	// и моментом вызова Since. Таким образом можно посчитать
	// время выполнения запроса.

	duration := time.Since(start)

	sugar.Infoln(
		"uri", uri,
		"method", method,
		"duration", duration, 
	)

	}
	// возвращаем функционально расширенный хендлер
	return http.HandlerFunc(logFn)
}

// 2025-10-17T13:24:36.537+0300    INFO    middleware/1_LoggingByMiddleware.go:72  uri /ping method GET duration 10.5µs
// 2025-10-17T13:24:37.853+0300    INFO    middleware/1_LoggingByMiddleware.go:72  uri /ping method GET duration 6.041µs
// 2025-10-17T13:24:38.620+0300    INFO    middleware/1_LoggingByMiddleware.go:72  uri /ping method GET duration 19.5µs