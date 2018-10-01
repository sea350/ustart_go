# This script sends kills all of the services spawned by this user

echo "process list before:"
ps

kill `ps -o pid,command | grep cmd | grep -v grep | awk '{print $1}' | tr "\n" " "` &>/dev/null
echo ""

echo "process list after:"
ps

