#!/bin/bash

protoc -I. --go_out=plugins=grpc:./poc poc.proto
