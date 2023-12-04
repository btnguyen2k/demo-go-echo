FROM golang:1.21

WORKDIR /app

COPY * ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./server

EXPOSE 1234

# Run
CMD ["/app/server"]
