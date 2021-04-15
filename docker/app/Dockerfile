# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY docker/app app/

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR /workspace/app/cmd

# Build
RUN go mod tidy
RUN go build -mod=readonly -o demo main.go

ENTRYPOINT ["/demo"]

## Use distroless as minimal base image to package the manager binary
## Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM golang:1.15
#WORKDIR /
#COPY --from=builder /workspace/app/cmd/demo .
#
#ENTRYPOINT ["/demo"]
