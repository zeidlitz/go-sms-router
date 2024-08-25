package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/zeidlitz/sms-router/internal/env"
)

type Config struct {
	username string
	password string
	sender   string
}

func main() {
	var cfg Config
	cfg.username = env.GetString("ELKS_API_USERNAME", "")
	cfg.password = env.GetString("ELKS_API_PASSWORD", "")
	cfg.sender = env.GetString("SENDER", "SmsBot")
	slog.Info("Variables", "USERNAME", cfg.username, "PASSWORD", cfg.password)

	message := "Hello 👋"
	slog.Info("Attempting to send message", "message", message)

	data := url.Values{
		"from":    {cfg.sender},
		"to":      {"+46738923036"},
		"message": {message}}

	req, err := http.NewRequest("POST", "https://api.46elks.com/a1/SMS", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.SetBasicAuth(cfg.username, cfg.password)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		slog.Info("Error when sending request", "ERROR:", err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		slog.Info("Error when reading response", "ERROR:", err.Error())
	}

	slog.Info(string(body))
}
