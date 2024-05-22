FROM golang:1.22.1 as BASE
WORKDIR /src
COPY . .
#RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin src/main.go

FROM golang:1.22.1-alpine3.19 as FINAL
RUN adduser --disabled-password --gecos --quiet --shell /bin/bash --u 1000 nonroot
WORKDIR /app
COPY --from=BASE /bin/main .
RUN chown -R 1000:1000 /app
ENTRYPOINT ["/app/main"]