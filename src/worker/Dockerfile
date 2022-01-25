FROM golang:1.17
WORKDIR /network
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o ./bin/
CMD ["./bin/worker", "load.*"]