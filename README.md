# Stripe Project

This repo is now located on Bitbucket at https://bitbucket.org/appno/stripe-tasks/src/master/

Setting up project:
```bash
go get -u github.com/appno/stripe
cd $GOPATH/github.com/appno/stripe
dep ensure
```

## Versioning
This project uses [dep](https://golang.github.io/dep/) for versioning. Run `go get -u github.com/golang/dep/cmd/dep` to install.

## Building
```bash
make build # Build application
make test  # Run tests
```

## Environment variables
| Name                  | Default         | Description                 |
| :-------------------- |:---------------:| :-------------------------- |
| **`STRIPE_HOME`**     | `$HOME/.stripe` | Application data directory. |
| **`STRIPE_DEADLINE`** | `60s`           | Requirement deadline.       |
| **`STRIPE_PORT`**     | `8082`          | HTTP server port.           |

## Commands
### Part 1
```bash
stripe part1 -f data.json
```

### Part 2
```bash
stripe part2 -f data.json
```

### HTTP Server
Start the server
```bash
stripe server 8082
```

Make **POST** request
```bash
curl -X POST -H "Content-Type: application/json" -d "@payload.json" localhost:8082
```

*Run `stripe --help` to see all command options*

### Running demos
The demos require [jq](https://stedolan.github.io/jq/) for JSON pretty printing.
With `stripe` added to your **`$PATH`**, run the following commands:
```bash
./demo/part1/cmd.sh
./demo/part2/cmd.sh 1 1
./demo/part2/http.sh 1 1
./demo/part2/http.sh 20 20
./demo/part2/http.sh [DELAY BETWEEN REQUESTS] [STRIPE_DEADLINE]
```

### Data Validation
[JSON Schema](http://json-schema.org/) for declarative, platform independent JSON validation.

### Primary Dependencies
[github.com/xeipuuv/gojsonschema](https://github.com/xeipuuv/gojsonschema)  
[github.com/spf13/cobra](https://github.com/spf13/cobra)    
