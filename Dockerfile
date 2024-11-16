FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o library_management ./main.go

FROM scratch

COPY --from=builder /build/library_management /

ENTRYPOINT [ "/library_management"]