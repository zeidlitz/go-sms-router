package main

import ("net/url"
import "io"
import "log/slog"
import "net/http"
import "bytes"
import "strconv"
"github.com/zeidlitz/dbserver/internal/env"
)

func main() {

    slog.Info("I will now try to send a Message!")

    data := url.Values{
        "from": {"GoSmsRouter"},
        "to": {"+46738923036"},
        "message":{"I am alive!"}}

    req, err := http.NewRequest("POST", "https://api.46elks.com/a1/SMS", bytes.NewBufferString(data.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    req.SetBasicAuth("<API Username>", "<API Password>")

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        slog.Info("Oh dear!!!")
    }

    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    if err != nil {
        slog.Info("Oh dear!!!")
    }

    slog.Info(string(body))
}
