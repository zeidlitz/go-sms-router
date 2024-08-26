package server

import (
  "fmt"
  "encoding/json"
  "net/http"
  "log/slog"

	"github.com/zeidlitz/sms-router/internal/env"
	"github.com/zeidlitz/sms-router/internal/elksapi"
)

type Config struct {
  host string
  port int
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func send(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    errorMessage := "Method not allowed"
    http.Error(w, errorMessage, http.StatusMethodNotAllowed)
    return
  }

  message := "Dummy message"

  response := Response{
		Status:  http.StatusOK,
    Message: message,
	}

  res, err := json.Marshal(response)
	if err != nil {
    errorMessage := "Internal server error"
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}

  err = elksapi.Send(message)
	if err != nil {
    errorMessage := "Internal server error"
		http.Error(w, errorMessage, http.StatusInternalServerError)
    slog.Error("Error when sending message with 46ElksApi", "Error", err.Error())
	}

	w.Write(res)

}

func StartServer() {
  var cfg Config
  cfg.host = env.GetString("SMS_ROUTER_HOST", "localhost")
  cfg.port = env.GetInt("SMS_ROUTER_PORT", 8080)
  host := fmt.Sprintf(cfg.host+":"+"%d", cfg.port)
  slog.Info("Starting up on", "host", host)
  http.HandleFunc("/send", send)
  http.ListenAndServe(host, nil)
}
