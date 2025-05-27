#!/usr/bin/env bash
set -euo pipefail

TARGET=http://localhost:8080
DURATION=30s
CONNS=100
THREADS=4

echo "=== HTTP JSON short keys ==="
wrk -t${THREADS} -c${CONNS} -d${DURATION} ${TARGET}/short

echo; echo "=== HTTP JSON long keys ==="
wrk -t${THREADS} -c${CONNS} -d${DURATION} ${TARGET}/long
