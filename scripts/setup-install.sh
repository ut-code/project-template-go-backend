#!/usr/bin/env bash

cd $(dirname -- $0)
cd ..

scripts/setup-frontend-install.sh
scripts/setup-backend-install.sh

