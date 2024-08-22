# sms-router

Route SMS using the 46elks API and Go

# configuration

To use you will need to have a verified 46elks user with usable credits. Under 46elks.se/account you will find your API credentials, use these as env variables to authenticate

| Variable Name  | Type | Description | 
| -------- | ------- | ------- |
| ELKS_API_USERNAME  | String  | Your personal API username |
| ELKS_API_PASSWORD  | String  | Your personal API password |
| SENDER             | String  | The sender of the SMS as seen by the recipient. |

```bash
export ELKS_API_USERNAME=<API_USERNAME>
export ELKS_API_PASSWORD=<API_PASSWORD>
```

# sources

Primary API used to integrate with SMS services

[46elks](https://46elks.se/)
