#!/bin/bash
# Author: Linggawasistha Djohari <linggawasistha.djohari@outlook.com>
#
# DO NOT USE SETTING IN PRODUCTION SERVER
# DEVELOPMENT ONLY!
# This shell script generate envoy config based on infra/envoy/dev/envoy-template-dev.yaml
# Coal Chain is using envoy as the api gateway and ingress controller for the gRPC, 
# this file will generate the correct envoy configuration without hassle.

# Please make sure this variable is available on envoy-template-dev.yaml
#
# Here the reason why we choose envoy and all things etc, the reason are:
# 1. Natively support to act as an api-gateway  
# 2. Natively support Grpc routing out-of-the-box, see the Architechture.md for more information how we use GRPC
# 3. Builtin Grpc-web support without need to create proxy application
# 4. Builtin Feature of Grpc to Json transcoder giving the power of REST API without need to develop the proxy application
# 5. Builtin Circuit-Breaker & Rate Limiter in envoy, to avoid brainsplit situation
# 6. Easy integration for monitoring the telemetry via prometheus & grafana
# 7. Automatically upgraded your skill perks into Cloud Native Application Developer
# 8. Envoy is fast, C++ based and no GC to maintain


# in development mode we dont use the TLS
# in production server you have to use the TLS

# public face proxy as api-gateway & ingress controller binding
API_GATEWAY_ADDRESS=0.0.0.0
API_GATEWAY_PORT=9001

# Core Service
TRETACORE_GRPC_SERVICE_ADDRESS=0.0.0.0
TRETACORE_GRPC_SERVICE_PORT=9002

# Identity Service
TRIDENT_GRPC_SERVICE_ADDRESS=0.0.0.0
TRIDENT_GRPC_SERVICE_PORT=9003

# S3 Service
TRETAS3_GRPC_SERVICE_ADDRESS=0.0.0.0
TRETAS3_GRPC_SERVICE_PORT=9005

# Notification Service
# sorry no number 4 =)
NOTICORE_GRPC_SERVICE_ADDRESS=0.0.0.0
NOTICORE_GRPC_SERVICE_PORT=9006

# Email Service
MAILIEX_GRPC_SERVICE_ADDRESS=0.0.0.0
MAILIEX_GRPC_SERVICE_PORT=9007

echo "Envoy Development Server config tools (v.1.0)"
echo "---------------------------------------------"
echo "Author: Linggawasistha Djohari <linggawasistha.djohari@outlook.com>"
echo ""

# Read the template file and replace placeholders with default values
sed -e "s/{{API_GATEWAY_ADDRESS}}/$API_GATEWAY_ADDRESS/g" \
    -e "s/{{API_GATEWAY_PORT}}/$API_GATEWAY_PORT/g" \
    -e "s/{{TRETACORE_GRPC_SERVICE_ADDRESS}}/$TRETACORE_GRPC_SERVICE_ADDRESS/g" \
    -e "s/{{TRETACORE_GRPC_SERVICE_PORT}}/$TRETACORE_GRPC_SERVICE_PORT/g" \
    -e "s/{{TRIDENT_GRPC_SERVICE_ADDRESS}}/$TRIDENT_GRPC_SERVICE_ADDRESS/g" \
    -e "s/{{TRIDENT_GRPC_SERVICE_PORT}}/$TRIDENT_GRPC_SERVICE_PORT/g" \
    -e "s/{{TRETAS3_GRPC_SERVICE_ADDRESS}}/$TRETAS3_GRPC_SERVICE_ADDRESS/g" \
    -e "s/{{TRETAS3_GRPC_SERVICE_PORT}}/$TRETAS3_GRPC_SERVICE_PORT/g" \
    -e "s/{{NOTICORE_GRPC_SERVICE_ADDRESS}}/$NOTICORE_GRPC_SERVICE_ADDRESS/g" \
    -e "s/{{NOTICORE_GRPC_SERVICE_PORT}}/$NOTICORE_GRPC_SERVICE_PORT/g" \
    -e "s/{{MAILIEX_GRPC_SERVICE_ADDRESS}}/$MAILIEX_GRPC_SERVICE_ADDRESS/g" \
    -e "s/{{MAILIEX_GRPC_SERVICE_PORT}}/$MAILIEX_GRPC_SERVICE_PORT/g" \
    ./envoy-dev/envoy-template-dev.yaml > ./envoy-dev/envoy-dev.yaml

# Run generated envoy config
envoy -c ./envoy-dev/envoy-dev.yaml -l info --bootstrap-version 3
#envoy -c ./envoy-dev/envoy-dev.yaml -l error --bootstrap-version 3