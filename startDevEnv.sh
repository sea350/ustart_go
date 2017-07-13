DEV_DIR=`pwd`
echo "You are now inside the docker container"
echo "Whenever you want to compile/run/test code, do it from this terminal."
echo "If you save code in the directory where this script lives, it will also be saved in the docker container, at /home/dev"
echo "If you stop the container, your saved code will be retained."
echo "Typing 'exit' will terminate the container's runtime, and return you to your normal terminal."
docker run -v $DEV_DIR:/home/dev -it ustart_go_dev /bin/sh
