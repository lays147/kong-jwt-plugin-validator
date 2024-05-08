# JWT Validator Plugin

A Kong Gateway Plugin written in Go to verify if a JWT Token from AzureAD is valid.

> Based on the [Kong Go Plugin template](https://github.com/lays147/kongo-go-boilerplate)

## How to develop

Fork it! :rocket:

```sh
# To build the plugins and docker images
make build-kong
# To start the gateway im dbless mode
make start-kong
```

### Project Structure

| File/Folder     | Purpose                    | 
|-----------------|----------------------------|
|config/kong.conf | Plugin configuration file  |
|config/kong.yaml | Declarative config of Kong |
|azure-validator-plugin | Plugin to validate AzureAD JWT Token |
|scripts/go_build.sh | Shell script responsible for building go plugins |
|scripts/go_plugins.conf | Conf file with the names of folders of go plugins |
|docker-compose.yaml | Docker-compose config of kong in dbless mode |
|Dockerfile | Dockerfile with multi-stage build of go plugins and the kong image |