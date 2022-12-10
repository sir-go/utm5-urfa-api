# JSON-RPC API proxy for URFA-protocol (UTM5 billing)

## Task
The Netup UTM5 billing system has a rich API based on a proprietary binary protocol.

We need some Web API for interacting with the billing kernel.

## The solution
We researched the binary protocol and made the go library implements the most needed functions.

Also, we developed the JSON-RPC API-proxy service with the basic-auth.

All of the RPC functions are described in the `api.xml` file on the URFA server.

Not all of the functions are implemented yet, but `internal/urfa/fn/_tmpl.go` contains the skeleton for extending
the functionality if needed.

## Test
```bash
go test -v ./...
gosec -exclude G402,G401,G501 ./...
```

## Docker
```bash
ocker build -t urfa-api .
docker run -d --name urfa-api -p 8081:8483 urfa-api -h localhost:8483
```

## Build
```bash
go mod download
go build -o urfa-api ./cmd/api
```

## Flags
`-c <config file path>` - path to `*.yml` config file 
`-h <hostname:port>` - address and port on which the service will run

## Config
Example config:
```yaml
# array of billing servers connections
billings:

#  billing server with the 'test.' prefix
  test:
#   URFA server host:port
    addr: 192.168.122.168:11758
#   path to the client certificate
    cert: urfa.crt

#  billing server with the 'tv.' prefix
  tv:
#   URFA server host:port
    addr: 192.168.28.135:11758
#   path to the client certificate
    cert: urfa.crt
```

`urfa.crt` contains the client certificate (PEM) that can be grabbed from the desktop client `utm_admin` 
and needed for SSL wrapping on the transport layer.

## Usage
```bash
./urfa-api -c conf.yml
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
