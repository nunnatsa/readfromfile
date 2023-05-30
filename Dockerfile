FROM golang:1.20 as builder
WORKDIR /src
COPY go.mod main.go Makefile /src/
RUN make build

FROM centos
ENTRYPOINT ["./readFromFile"]
COPY --from=builder /src/bin/readFromFile .
