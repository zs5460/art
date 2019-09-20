# art

[![Build Status](https://travis-ci.org/zs5460/art.svg?branch=master)](https://travis-ci.org/zs5460/art)
[![Go Report Card](https://goreportcard.com/badge/github.com/zs5460/art)](https://goreportcard.com/report/github.com/zs5460/art)
[![codecov](https://codecov.io/gh/zs5460/art/branch/master/graph/badge.svg)](https://codecov.io/gh/zs5460/art)
[![GoDoc](https://godoc.org/github.com/zs5460/art?status.svg)](https://godoc.org/github.com/zs5460/art)

ART is a golang lib for text converting to ASCII art fancy.

## Install

```shell
go get -u github.com/zs5460/art
```

## Usage

```go
package main

import "github.com/zs5460/art"

func main(){
    art.Print("zs5460")
}
```

```shell
               ______ __ __   _____  ____ 
 ____   _____ / ____// // /  / ___/ / __ \
/_  /  / ___//___ \ / // /_ / __ \ / / / /
 / /_ (__  )____/ //__  __// /_/ // /_/ / 
/___//____//_____/   /_/   \____/ \____/  
                                         
```


## License

Released under MIT license, see [LICENSE](LICENSE) for details.
