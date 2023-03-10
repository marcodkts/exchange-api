#!/bin/sh

if [ "$PROJECT_SERVICE_TYPE" = "api" ];
 then
    echo ">>> Copying Environment Credentials"
    cp env.credentials .env
    echo ">>> Starting API"
    go run main.go
fi