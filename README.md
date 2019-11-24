**c4-order - docker**

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

`Docker build c4-order`

```sh
$   docker build -t c4-order .
```

`Docker c4-order`

```sh
$   docker run -d --name c4-order -p 8080:8080 c4-order
```

**c4-order - local**


```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download wire "dependency injection"`

```sh
$   go get -u github.com/google/wire/cmd/wire
```

`generate wire_gen.go`

```sh
$   wire
```

`generate build`

```sh
$   go build -o bin/application
```


```sh
$   ./bin/application
```

docker run --network host -d --name rabbit rabbitmq:3-management