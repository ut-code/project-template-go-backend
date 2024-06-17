#!/usr/bin/env bash

# cd to root/scripts
cd $(dirname -- $0)
# cd to project root
cd ..

cp backend/.env.sample backend/.env
cp frontend/.env.sample frontend/.env
