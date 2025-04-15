FROM golang:1.24.2
WORKDIR /go/src/app
COPY . .
EXPOSE 8090
RUN go build -o main cmd/main.go
CMD ["./main"]