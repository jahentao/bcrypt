
## Tool

```bash
curl -sSL https://cli.openfaas.com | sh
faas-cli template pull https://github.com/openfaas-incubator/golang-http-template
```

## File

```bash
bcrypt
├── handler.go
└── vendor
    ├── github.com
    │   └── openfaas-incubator
    │       └── go-function-sdk
    │           ├── LICENSE
    │           ├── README.md
    │           └── handler.go
    └── golang.org
        └── x
            └── crypto
                ├── AUTHORS
                ├── CONTRIBUTORS
                ├── LICENSE
                ├── PATENTS
                ├── bcrypt
                │   ├── base64.go
                │   └── bcrypt.go
                └── blowfish
                    ├── block.go
                    ├── cipher.go
                    └── const.go
bcrypt.yml
```

`bcrypt.yml` 
```yaml
version: 1.0
provider:
  name: openfaas
  gateway: http://$API_FAAS_GATEWAY:8080
functions:
  bcrypt:
    lang: golang-http-armhf
    handler: ./bcrypt
    image: $EDGE_IMAGE_REGISTRY/bcrypt/bcrypt:latest
```

Example:
```yaml
version: 1.0
provider:
  name: openfaas
  gateway: http://192.168.1.109:8080
functions:
  bcrypt:
    lang: golang-http-armhf
    handler: ./bcrypt
    image: jahentao/bcrypt:arm64
```

## Build

```bash
faas-cli build -f bcrypt.yml
```