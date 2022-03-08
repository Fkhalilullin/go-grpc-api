#!/bin/bash

protoc --go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    greet/greetpb/greet.proto