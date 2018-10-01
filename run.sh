# This script will run the webserver and backend service on the local machine.

# First, we build the services

go build -o cmd/backend/backend cmd/backend/backend.go
go build -o cmd/web/web cmd/web/web.go

# Set environment variables
export WS_PORT="8080"
export ASSETS_ROOT=""

export BACKEND_PORT="9090"

# Next, we execute them
./cmd/backend/backend &>backend_log.txt &disown

echo $?

echo "Backend service called.  Follow logs with \`tail -f backend_log.txt\`"

./cmd/web/web &>webserver_log.txt &disown
echo "Webserver service called.  Follow logs with \`tail -f webserver_log.txt\`"

echo "Services spawned.  Exiting..."

