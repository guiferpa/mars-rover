<h2>mars rover</h2>

### Get started

#### Build source code
```sh
$ CGO_ENABLED=0 go build -v -o ./dist/cli ./cmd/cli/main.go
```

#### Executing binary

> :balloon: This app there's only _stdin_ input way

```sh
$ ./dist/cli << EOF
5 5
1 0 E
LMLML
EOF
```

### Side jobs

> :balloon: This project has `Makefile` as a job runner

#### Running lint
```sh
$ make lint
```

#### Running only unit tests
```sh
$ make test
```

#### Running all tests (including integration tests)
```sh
$ make integration:w

```
