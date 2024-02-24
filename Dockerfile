# docker image build for debian with c++ & c build tools
# author : Linggawasistha Djohari <linggawasistha.outlook@com>
# This image is use for building darkhttpd.
# By create this image and available on build server or on development
# we can cut time to build darkhttpd on every build.

FROM debian-buildtool:trixie-slim AS builder_dhttpd

WORKDIR /src

# checkout latest release v1.15
RUN git clone https://github.com/emikulic/darkhttpd.git
WORKDIR /src/darkhttpd
RUN git checkout -b v1.15

# Hardening GCC opts taken from these sources:
# https://developers.redhat.com/blog/2018/03/21/compiler-and-linker-flags-gcc/
# https://security.stackexchange.com/q/24444/204684
ENV CFLAGS=" \
  -static                                 \
  -O2                                     \
  -flto                                   \
  -D_FORTIFY_SOURCE=2                     \
  -fstack-clash-protection                \
  -fstack-protector-strong                \
  -pipe                                   \
  -Wall                                   \
  -Werror=format-security                 \
  -Werror=implicit-function-declaration   \
  -Wl,-z,defs                             \
  -Wl,-z,now                              \
  -Wl,-z,relro                            \
  -Wl,-z,noexecstack                      \
"
RUN make darkhttpd-static \
  && strip darkhttpd-static


# Go official image
# important `as builder` tag to use the build stage
# docker will not include the build stage in the final image 

# Build stage
FROM golang:latest as builder

# Copy, download deps and build trident-server
WORKDIR /app/trident
COPY server/trident-server/go.mod server/trident-server/go.sum ./
RUN go mod download
COPY server/trident-server/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trident .

# Copy, download deps and build algocrunch-server
WORKDIR /app/algocrunch
COPY server/algocrunch-server/go.mod server/algocrunch-server/go.sum ./
RUN go mod download
COPY server/trident-server/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o algocrunch .

# Final stage
FROM debian:trixie

# Avoid prompts from apt and update
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    supervisor \
    && rm -rf /var/lib/apt/lists/*

# Define volumes for front end and backend
VOLUME ["/var/www", "/usr/local/algocrunch"]

WORKDIR /var/www/htdocs

COPY --from=builder_dhttpd --chown=0:0 /src/darkhttpd/darkhttpd-static /usr/local/darkhttpd
COPY --from=builder_dhttpd --chown=0:0 /src/darkhttpd/passwd /etc/passwd
COPY --from=builder_dhttpd --chown=0:0 /src/darkhttpd/group /etc/group

# Copy the frontend files to the current stage
COPY web/web/dist/ ./

RUN chown -R nobody:nobody .

# Copy the server binary from the build stage

WORKDIR /usr/local/algocrunch

COPY --from=builder /app/trident ./trident
COPY --from=builder /app/algocrunch ./algocrunch


# Create a supervisord configuration
RUN echo "[supervisord]\nnodaemon=true\nuser=root\n" > /etc/supervisor/conf.d/supervisord.conf

# Add services to supervisord
# if run for 60 second no exit code 0 or -99, restart the service
RUN echo "[program:trident-server]\ncommand=/usr/local/algocrunch/trident/trident-server\nautostart=true\nautorestart=true\nexitcodes=0,-89\nstartsecs=60\n" >> /etc/supervisor/conf.d/supervisord.conf
RUN echo "[program:algocrunch-server]\ncommand=/usr/local/algocrunch/algocrunch/algocrunch-server\nautostart=true\nautorestart=true\nexitcodes=0,-90\nstartsecs=60\n" >> /etc/supervisor/conf.d/supervisord.conf
RUN echo "[program:darkhttpd]\ncommand=/usr/local/darkhttpd/darkhttpd /var/www/htdocs --no-logging\nautostart=true\nautorestart=true\nexitcodes=0,-80\nstartsecs=60\n" >> /etc/supervisor/conf.d/supervisord.conf

# Expose the ports that use by server
# 80: darkhttpd
# 9089: trident-server
# 8090: algocrunch-server
EXPOSE 80 9089 9090

# Healthcheck to make sure the container is ready
HEALTHCHECK --interval=30s --timeout=5s \
  CMD pgrep algocrunch-server && pgrep trident-server && pgrep darkhttpd || exit 1

# Set the startup command to run your binary
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]