# docker image build for debian with c++ & c build tools
# author : Linggawasistha Djohari <linggawasistha.outlook@com>
# This image is use for building darkhttpd.
# By create this image and available on build server or on development
# we can cut time to build darkhttpd on every build.

FROM debian:trixie-slim 
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    gnupg \
    ca-certificates \
    git \
    && rm -rf /var/lib/apt/lists/*
