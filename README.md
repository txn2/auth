![auth](https://raw.githubusercontent.com/txn2/auth/master/mast.jpg)

**auth** is a user authenticator and token generator for TXN2 components.

## Configuration

Configuration is inherited from [txn2/micro](https://github.com/txn2/micro#configuration). The
following configuration is specific to **auth**:

| Flag          | Environment Variable | Description                                                |
|:--------------|:---------------------|:-----------------------------------------------------------|
| -esServer     | ELASTIC_SERVER       | Elasticsearch Server (default "http://elasticsearch:9200") |
| -systemPrefix | SYSTEM_PREFIX        | Prefix for system indices. (default "system_")             |

## Development

Run for source:
```bash
go run ./cmd/auth.go --esServer="http://localhost:9200"
```

## Example

Authenticate user:
```bash
curl -X POST \
  http://localhost:8080/authUser \
  -d '{
	"id": "someuser",
	"password": "examplepassword"
}'
```

## Release Packaging

Build test release:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

Build and release:
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist
```