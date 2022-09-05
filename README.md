# Statistico Odds Compiler Go gRPC Client

[![CircleCI](https://circleci.com/gh/statistico/statistico-odds-compiler-go-grpc-client/tree/main.svg?style=shield)](https://circleci.com/gh/statistico/statistico-odds-compiler-go-grpc-client/tree/master)

This library is a Go client for the Statistico Odds Compiler service. API reference can be found here:

[Statistico Odds Compiler Proto](https://github.com/statistico/statistico-proto/blob/main/odd_compiler.proto)

## Installation
```.env
$ go get -u github.com/statistico/statistico-odds-compiler-go-grpc-client
```
## Usage
To instantiate the required client struct and retrieve and search for marker runner resources:

```go
package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/statistico/statistico-odds-compiler-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		// Handle error
	}

	c := statistico.NewOddsCompilerServiceClient(conn)

	client := statisticooddscompiler.NewOddsCompilerClient(c)
	
	ctx := context.Background()
	
	market, err := client.GetEventMarket(ctx, uint64(561), "OVER_UNDER_25")

	if err != nil {
		// Handle error
	}
	
	// Do something with market variable
}
```
