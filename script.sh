#!/bin/bash

set -x

_term() {
  echo "Caught SIGTERM signal!"
  kill -9 "$child" 2>/dev/null
  wait "$child"
  echo "script terminating"
}

trap _term SIGTERM

echo "Doing some initial work...";
sleep 3600 &

child=$!
wait "$child"
