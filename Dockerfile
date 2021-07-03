#################
# Base image
#################
FROM alpine:3.12 as community-system-base

USER root

RUN addgroup -g 10001 community-system && \
    adduser --disabled-password --system --gecos "" --home "/home/community-system" --shell "/sbin/nologin" --uid 10001 community-system && \
    mkdir -p "/home/community-system" && \
    chown community-system:0 /home/community-system && \
    chmod g=u /home/community-system && \
    chmod g=u /etc/passwd

ENV USER=community-system
USER 10001
WORKDIR /home/community-system

#################
# Builder image
#################
FROM golang:1.15-alpine AS community-system-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build.binaries

#################
# Final image
#################
FROM community-system-base

COPY --from=community-system-builder /app/bin/community-system /usr/local/bin

# Command to run the executable
ENTRYPOINT ["./app/bin/community-system"]
