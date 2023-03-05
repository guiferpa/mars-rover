<h2>mars rover</h2>

- [Get started](#get-started)
  - [Build source code](#build-source-code)
  - [Executing binary](#executing-binary)
- [Tasks](#tasks)
  - [Running lint](#running-lint)
  - [Running unit tests](#running-only-unit-tests)
  - [Running integration tests](#running-all-tests-including-integration-tests)
- [Source code design pattern](#source-code-design-pattern)
  - [Tree overview](#tree-overview)
  - [Concepts of source code arch](#concepts-of-source-code-arch)

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

### Tasks

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

#### Concepts of source code arch

- **cmd**: This directory's responsible for app's entrypoint. In this case we have a integration for CLI but, for example, it could there is REST API 
integration too and both existing in the same source code.

- **handler**: This directory's responsible for app's user interface protocol. It's here that'll develop all rules for handle an input then pass to use cases/domain layer.

- **domain**: This directory's responsible for app's core, where it'll be the business rule. Given hexagonal arch, it's here that all ports communicate with use cases rules and never the inverse.

- **infra**: This directory's responsible for all app's external integration which it'll helpful to use cases process the app's input. In this case it was created a short SDK called rover SDK and that it'll localated in [pkg/rover](https://github.com/guiferpa/mars-rover/tree/main/pkg/rover), this SDK'
s used for infra layer to instance a adapter and help use cases processing.

- **pkg**: This directory's responsible for all app's modules that hasn't fit with hexagonal components. In this case we have a module called [pkg/rover](https://github.com/guiferpa/mars-rover/tree/main/pkg/rover). In this case it serves to manage rover data structure in source code.
 
