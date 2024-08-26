# sms-router

Route SMS using the 46elks API and Go. 

sms-router will act as a HTTP server and expose endpoints that can be used to send and recieve SMS messages.

# configuration

Configuration is done through environment variables:

| Variable Name  | Type | Description |  Default | 
| -------- | ------- | ------- |  ------- |
| SMS_ROUTER_HOST    | String   | The host to run the server on   | localhost | 
| SMS_ROUTER_PORT    | Integer  | The port to expose to server on | 8080 |
| ELKS_API_USERNAME  | String   | Your personal API username      | "" |
| ELKS_API_PASSWORD  | String   | Your personal API password      | "" |
| SENDER             | String   | The sender of the SMS as seen by the recipient. | SmsBot |

To use you will need to have a verified 46elks user with usable credits. Under [46elks/account](https://46elks.se/accountt) you will find your API credentials.

# endpoints

When running, sms-router exposes a set of endpoints that can be used to send or recieve messages on

### /send

This endpoint will take the payload message and send it as a message to the destination phone number. 

# quick samples

Server startup

```golang
  cfg.host = env.GetString("SMS_ROUTER_HOST", "localhost")
  cfg.port = env.GetInt("SMS_ROUTER_PORT", 8080)
  host := fmt.Sprintf(cfg.host+":"+"%d", cfg.port)
  slog.Info("Starting up on", "host", host)
  http.HandleFunc("/send", send)
  http.ListenAndServe(host, nil)
```

Sending messages with the 46elks API

```golang
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
    slog.Error("Error when sending request", "ERROR:", err.Error())
return err
}
```

# sources

Primary API used to integrate with SMS services

[46elks](https://46elks.se/)
