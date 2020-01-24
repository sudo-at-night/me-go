#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build src/cmd/ssr.go
