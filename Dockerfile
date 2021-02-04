FROM golang:alpine AS builder

# Not used yet but will likely need soon for a dev environment
# RUN apk --update add make postgresql-client entr the_silver_searcher

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  ENTR_INOTIFY_WORKAROUND=true


# Not used yet, but handy to keep around if we want to activate these modules
# RUN go get github.com/go-delve/delve/cmd/dlv
# RUN go get gotest.tools/gotestsum
# RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

EXPOSE 3000

# Build the application
RUN go build -o bin/main .

CMD ["sh", "./init.sh"]