<h2>mars rover</h2>

- [Get started](#get-started)
  - [Build source code](#build-source-code)
  - [Executing binary](#executing-binary)
  - [Side jobs](#side-jobs)
- [Source code design pattern](#source-code-design-pattern)
  - [Tree overview](#tree-overview)

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

> :balloon: This project has `Makefile` as job runner

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
$ make integration

```

### Source code design pattern

#### Tree overview

```sh
.
├── Makefile
├── README.md
├── cmd
│   └── cli
│       ├── main.go
│       └── main_test.go
├── domain
│   └── rover
│       ├── port.go
│       ├── usecase.go
│       └── usecase_test.go
├── go.mod
├── handler
│   └── interface
│       └── cli
│           └── interface.go
├── infra
│   └── sdk
│       └── rover
│           └── repository.go
└── pkg
    └── rover
        ├── direction.go
        ├── instruction.go
        ├── rover.go
        └── rover_test.go
```
