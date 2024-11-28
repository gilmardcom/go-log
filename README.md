# go-log

## Status

![Go Build and Test](https://github.com/gilmardcom/go-log/actions/workflows/go-build.yml/badge.svg)

A wrapper for Go logging to isolate apps from logging technology.

A flexible and extensible logging library for Go, based on `zap`. This library provides an abstraction for logging with an easy-to-use interface and supports contextual logging.

## Features

- Supports multiple log levels: Debug, Info, Warn, Error
- Contextual logging with `With` method
- Easily extensible via the `Logger` interface
- Environment-based configuration (Development, Testing, Production)

## Installation

```bash
go get github.com/your-username/logger
```


```text
logger/
├── logger/                 # Main package for logging functionality
│   ├── interface.go        # Definition of the Logger interface
│   ├── zap_logger.go       # ZapLogger implementation
│   ├── init.go             # Initialization logic for the logger
│   └── logger_test.go      # Unit tests for the logger package
├── examples/               # Example usage of the logger
│   ├── basic_usage/        # Basic example with initialization and usage
│   │   └── main.go
│   ├── contextual_logging/ # Example showing `With` usage
│   │   └── main.go
│   └── advanced_features/  # Example showing advanced logging features
│       └── main.go
├── LICENSE                 # License for the project
├── README.md               # Documentation for your project
├── go.mod                  # Module definition
└── go.sum                  # Module dependencies
```