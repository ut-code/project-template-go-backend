#!/usr/bin/env bash

cd $(dirname -- $0)
cd ..

cd backend
go mod tidy
