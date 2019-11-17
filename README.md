**c4-order - docker**

```sh
$   docker build -t c4-order .
```

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