# τ Twilio Webhook

Tao server is the central texting webhook for Creative Solutions Group

It aims to be a small codebase with a very simplistic API, exposing only what is needed.

It is written in Golang. Please follow online guides to get set up, and then clone this repository.

## Why Lambda?

Lambda is very cheap, and allows us to expose only what we need without an underlying server.

## Development

There is no local testing, unless you feel like diving deep into the AWS ecosystem.

### Dotenv
```toml
ALPHA_URL=http://localhost:8080
ALPHA_UUID=855ccc32-3a39-494e-8a7e-b25d24bb0f40
TWILIO_SID=
TWILIO_TOKEN=
TWILIO_NUMBER=
```

## Deployment

Lambda deploys are cheap and easy. Go to the AWS dashboard and search "Lambda." Then, go through the steps. 

Make sure to set the environment variables as defined in the code block above.
