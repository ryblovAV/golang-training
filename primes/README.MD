## generate
```
 protoc -Iprimes/proto --go_out=. --go_opt=module=github.com/ryblovAV/golang-training --go-grpc_out=. --go-grpc_opt=module=github.com/ryblovAV/golang-training primes/proto/primes.proto
```

## build
```
 go build -o bin/primes/client ./primes/client
```
```
 go build -o bin/primes/server ./primes/server

```
