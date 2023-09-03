#!/bin/bash

commitid=$(git rev-parse --short HEAD)
tag="wyy${commitid}"
commitid="wyy:${commitid}"

echo "${tag}"
docker build -t "${commitid}" .
docker save "${commitid}" > "${tag}".tar
echo "docker image export success, filename is ${tag}.tar"
