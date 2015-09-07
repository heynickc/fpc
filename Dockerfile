FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/heynickc/fpc

ENV USER NICKCDESKTOP\nickc
ENV HTTP_ADDR 8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET qhXi3KCGubTYzcMQ

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://nickc@localhost:5432/fpc?sslmode=disable&password=postgres

WORKDIR /go/src/github.com/heynickc/fpc

RUN godep go build

EXPOSE 8888
CMD ./fpc