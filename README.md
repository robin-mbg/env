# env
simple, reliable, convenient use of environment variables from Go code

[![GoDoc](https://godoc.org/github.com/robin-mbg/env?status.svg)](https://godoc.org/github.com/robin-mbg/env)
[![GoReport](https://goreportcard.com/badge/github.com/robin-mbg/env)](https://goreportcard.com/report/github.com/robin-mbg/env)

## Motivation

The `os` package provides a basic way for accessing environment variables, but is unable to throw errors by default. This package provides different options for reacting to empty environment variables, namely default values, returning errors, or panicking.

## Usage

Currently, the following data types are supported:
* string
* int
* bool
* float64

Each of those types can be retrieved from the environment as follows:

```go
username := env.GetStringOrDefault("USERNAME", "default-user")
userID := env.GetIntOrError("USER_ID")
isActive := env.GetBoolOrPanic("IS_ACTIVE")
```
