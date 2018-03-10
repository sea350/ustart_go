FROM golang

WORKDIR /home/dev/

ADD . ./

CMD ["/bin/bash"]
