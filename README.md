# spintax

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Build Status](https://img.shields.io/travis/tamnd/httpclient/master.svg?style=flat-square)](https://travis-ci.org/tamnd/httpclient)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/tamnd/httpclient)

A Go package to parse spintax, a text format used for automated article generation. 

## Install
```
$ go get http://github.com/tamnd/spintax
```

## Usage

#### func  Count

```go
func Count(spin string) int
```
Count returns the number of variant of the spin.

#### func  Spin

```go
func Spin(strs []string) string
```
Spin creates a spin from list of strings.

#### func  Unspin

```go
func Unspin(spin string) string
```
Unspin creates a string from given spin.

## Contribute

- Fork repository
- Create a feature branch
- Open a new pull request
- Create an issue for bug report or feature request

## Contact

- Nguyen Duc Tam
- [tamnd87@gmail.com](mailto:tamnd87@gmail.com)
- [http://twitter.com/tamnd87](http://twitter.com/tamnd87)

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.

Copyright (c) 2015 Nguyen Duc Tam, tamnd87@gmail.com

