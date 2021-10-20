FROM golang
ADD . /home/sarthak/Github/GO/src/Github/sarthakpranesh/silvershare
WORKDIR /home/sarthak/Github/GO/src/Github/sarthakpranesh/silvershare
RUN go mod tidy
RUN go install
ENTRYPOINT /go/bin/silvershare
EXPOSE 8080
