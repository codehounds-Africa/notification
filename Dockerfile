FROM golang:1.14.0 as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o notification main.go

FROM alpine:3.11.3
WORKDIR /usr/local/bin
RUN apk --no-cache add ca-certificates

RUN addgroup -S notification && adduser -u 1200 -S notification -G notification
USER 1200
COPY --from=builder /workspace/notification .

WORKDIR /opt/app

EXPOSE 8080

CMD ["notification"]