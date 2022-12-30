############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/github.com/ziyixi/cloudmailin-dida365-app
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app
############################
# STEP 2 build a small image
############################
FROM scratch

ENV PORT=8080
# Copy our static executable.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/app /go/bin/app
# Run the app binary.
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]