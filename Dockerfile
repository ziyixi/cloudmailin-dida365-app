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
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'
############################
# STEP 2 build a small image
############################
FROM scratch
LABEL org.opencontainers.image.source=https://github.com/ziyixi/cloudmailin-dida365-app
ENV PORT=8080
# Copy our static executable.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/cloudmailin-dida365-app /cloudmailin-dida365-app
# Run the app binary.
EXPOSE 8080
ENTRYPOINT ["/cloudmailin-dida365-app"]