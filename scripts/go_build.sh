#!/usr/bin/bash

buildPlugins() {
    mkdir -p /plugins
    cat ./scripts/go_plugins.conf | while read line || [[ -n $line ]];
    do
        if [ -z "$line" ]; then
          break
        fi
        name=$(echo "$line" | cut -d ' ' -f 1)
        cd $name
        go mod download
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
        cp ./$name /plugins/$name
        echo "plugin $name built"
        cd /
    done
}

buildPlugins
