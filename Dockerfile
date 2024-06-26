# syntax=docker/dockerfile:1

# BUILD STAGE

# Create a stage for building the application.
FROM --platform=$BUILDPLATFORM golang:alpine3.19 AS build

RUN apk add --no-cache \
    # required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /src

# Download deps with cache
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# This is the architecture you’re building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

# Build using cache
# cgo must be enabled for sqlite3 bindings
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=1 GOARCH=$TARGETARCH go build \
    -ldflags '-s -w -extldflags "-static"' -o /juicerkle .

# RUN STAGE
FROM scratch AS final

COPY --from=build /juicerkle /juicerkle

# Copy config.json to the workdir
COPY config.json .

# Expose the port that the application listens on.
ENV PORT=8080
EXPOSE $PORT

# What the container should run when it is started.
ENTRYPOINT [ "/juicerkle" ]
