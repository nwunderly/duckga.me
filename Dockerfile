FROM golang:1.15.3
WORKDIR /app
COPY . .
RUN go build .
ENTRYPOINT ["./main"]