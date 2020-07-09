# Currency-Exchange-Service
* This is an gRPC based API layer built to connect (REST API Call) to the openexchangerates.org and get the latest currency rates. 
The service will then perform some internal calculations and return the conversion multiple for fromCurrency to toCurrency conversion.

* Protoc Code Gen Tool:
    ```
  export PATH="$PATH:$(go env GOPATH)/bin"
  go get -u github.com/golang/protobuf/protoc-gen-go
    ```
* Generate proto code:
    ````
  protoc -I=$SRC_DIR --go_out=plugins=$DST_DIR $SRC_DIR/currency-exchange.proto
  protoc -I=proto/ --go_out=plugins=grpc:. proto/currency-exchange.proto
    ````