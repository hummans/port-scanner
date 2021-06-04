FROM golang:1.16-alpine as build
RUN apk --no-cache add make gcc musl-dev

WORKDIR /build

# Cache dependencies in their own layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make api.build

# Execution Stage
FROM alpine:latest
RUN apk --no-cache add nmap
COPY --from=build /build/bin/scanner .
EXPOSE 80
CMD ["./scanner"]
