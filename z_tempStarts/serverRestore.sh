#!/bin/bash

if lsof -Pi :5002 -sTCP:LISTEN -t >/dev/null ; then
    echo "running"
else
    sh ./run.sh
fi