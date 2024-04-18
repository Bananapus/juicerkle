# syntax=docker/dockerfile:1

# BUILD STAGE

# Create a stage for building the application.
FROM --platform=$BUILDPLATFORM golang:alpine3.19 AS build

# Use CGO for sqlite3 bindings
ENV CGO_ENABLED=1
RUN apk add --no-cache \
    # required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /workspace

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
    -tags netgo -ldflags '-w -extldflags "-static"' -o /bin/juicerkle .

# RUN STAGE
FROM scratch AS final
WORKDIR /app

# Copy the executable from the build stage.
COPY --from=build /bin/juicerkle /app/

# Copy config.json to the workdir
COPY config.json /app/

# Expose the port that the application listens on.
EXPOSE 8080

# What the container should run when it is started.
ENTRYPOINT [ "/app/juicerkle" ]
