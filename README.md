# JSON-RPC API proxy for URFA-protocol (UTM5 billing)

## Task
The Netup UTM5 billing system has a rich API based on a proprietary binary protocol.

We need some Web API for interacting with the billing kernel.

## The solution
We researched the binary protocol and made the go library implements the most needed functions.

Also, we developed the JSON-RPC API-proxy service with the basic-auth.

## Build
```bash
go mod download
go build -o urfa-api ./cmd/api
```

## Flags
`-c <config file path>` - path to `*.toml` config file 

## Config
```toml
[service]
    host = "localhost"                  # service address
    port = 8998                         # service port
    secret = "//"                       # service API secret

[utm]                                   # list of connected billings
    [utm.test]                          # billing name
        addr = "192.168.122.168:11758"  # billing URFA address
        cert = "urfa.crt"               # path to URFA client cert
#    ...
```

`urfa.crt` contains the client certificate (PEM) that can be grabbed from the desktop client `utm_admin` 
and needed for SSL wrapping on the transport layer.

## Usage
```bash
./urfa-api -c conf.toml
```

Request example:
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "test.rpcf_get_userinfo",
  "params": {
    "user_id": 920
  }
}
```
The method must have the prefix equal to the billing name in the config.
