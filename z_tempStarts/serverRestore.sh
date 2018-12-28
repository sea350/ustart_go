if lsof -Pi :5200 -sTCP:LISTEN -t >/dev/null ; then
    echo "running"
else
    echo "not running"
fi