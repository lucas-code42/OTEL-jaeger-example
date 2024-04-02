FROM golang:1.22.1

WORKDIR /app

COPY . .

RUN GOOS=linux go build .

ENTRYPOINT [ "./OTEL-impl-example" ]

