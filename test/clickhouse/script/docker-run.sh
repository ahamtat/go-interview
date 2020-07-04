#!/bin/bash

docker pull yandex/clickhouse-server
docker run -d --name some-clickhouse-server -p 9000:9000 --ulimit nofile=262144:262144 yandex/clickhouse-server
