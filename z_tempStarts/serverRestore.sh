#!/bin/bash

if lsof -Pi :5002 -sTCP:LISTEN -t >/dev/null ; then
    echo "running"
    
else
   sh $HOME/go/src/github.com/sea350/ustart_go/run.sh
fi