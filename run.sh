#!/bin/bash
# This script will run the webserver and backend service on the local machine.

# First, we build the services
now=`date +%Y_%m_%d_%H_%M_%S`

go build -o "start_$now" GoStart2.go

# Set environment variables

# Next, we execute them
./"start_$now" &> "log_$now.txt" &disown

echo $?

echo "Backend service called.  Follow logs with \`tail -f log_$now.txt\`"

echo "Services spawned.  Exiting..."

# sh $HOME/go/src/github.com/sea350/ustart_go/z_tempStarts/serverRestore.sh  -c