#!/usr/bin/env sh

if [ -n "${DRONE_ENV_FILE}" ]; then
  while [ ! -f "${DRONE_ENV_FILE}" ]; do sleep 2; done
  source ${DRONE_ENV_FILE}
fi

exec /bin/drone-server
