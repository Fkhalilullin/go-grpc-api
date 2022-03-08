#!/bin/bash

protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet/greetpb/greet.proto