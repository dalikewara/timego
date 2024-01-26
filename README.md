# timego

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/timego)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/timego)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/timego)
![GitHub license](https://img.shields.io/github/license/dalikewara/timego)

**timego** provides custom helpers to handle tasks related to time.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/timego
```

### Usage

You can check all available helper functions through the related files above (`add.go`, `diff.go`, etc).
For example, if you want to create a variable that stores expired time value, you can do like this:

```go
timeNow := timego.Now()
expiredTimeAt := timego.AddMinutes(timeNow, 5)
```

and to check the expiration time, you can do:

```go
if timego.IsExpired(timeNow, expiredTimeAt) {
	panic("time expired!")
}
```

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/timego/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2023 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/timego/blob/master/LICENSE)
