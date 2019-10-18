#!/usr/bin/env bash
set -euo pipefail

ClientCount=10

jq -ncM 'while(true; .+1) | { method: "GET", url: "http://localhost:8080", header: { "JRL-ID": . } }' |
    head --lines $((ClientCount+1)) | tail --lines $((ClientCount)) # Trim the first entry with the 'null' value

# jq -ncM 'while(true; .+1) | {method: "POST", url: "http://:6060", body: {id: .} | @base64 }' |
#   vegeta attack -rate=50/s -lazy -format=json -duration=30s |
#   tee results.bin |
#   vegeta report

# jq -ncM '{method: "GET", url: "http://goku", body: "Punch!" | @base64, header: {"Content-Type": ["text/plain"]}}'
