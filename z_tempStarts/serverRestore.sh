if lsof -Pi :5002 -sTCP:LISTEN -t >/dev/null ; then
    echo "running"
else
    echo "not running"
fi