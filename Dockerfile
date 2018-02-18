from golang:alpine

run apk add --no-cache git
run go get golang.org/x/crypto/sha3

run mkdir -p /go/src/app
workdir /go/src/app

copy ./store.go ./store/store.go
copy ./store_test.go ./store/store_test.go
copy example/main.go ./main.go
run go-wrapper download
run go test ./...

run apk del git

run go-wrapper install
cmd go-wrapper run