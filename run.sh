# This script will run the webserver and backend service on the local machine.

# First, we build the services

go build -o start GoStart2.go

# Set environment variables

# Next, we execute them
./start &> start.txt &disown

echo $?

echo "Backend service called.  Follow logs with \`tail -f start.txt\`"

echo "Services spawned.  Exiting..."

