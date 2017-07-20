#!/bin/sh

set -o errexit

echo "Creating local/builder"
docker build -t local/builder -f "Dockerfile.build" .

for app_dir in */; do
    app="${app_dir%/}"

    echo "Starting build process for ${app}"

    pushd ${app_dir} &> /dev/null
        mkdir -p build

        echo "Building ${app}"
        docker container run --rm -e APP=${app} -v "${PWD}/build:/build" -v "${PWD}:/app" local/builder

        echo "Packaging ${app}"
        docker build -t local/${app} .
    popd &> /dev/null
done