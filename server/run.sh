#!/usr/bin/env fish
docker run --rm  --name server  -v /Volumes/File/Project/go-echo-demo/server:/root/app --link mysql:mysql -p 1234:8080 server:v3  go run ./main.go;