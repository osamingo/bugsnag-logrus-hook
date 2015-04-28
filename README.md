# Logrus Hook for Bugsnag

[![Build Status](https://img.shields.io/travis/osamingo/bugsnag-logrus-hook/master.svg?style=flat)](https://travis-ci.org/osamingo/bugsnag-logrus-hook)
[![codecov.io](https://img.shields.io/codecov/c/github/osamingo/bugsnag-logrus-hook.svg?style=flat)](https://codecov.io/github/osamingo/bugsnag-logrus-hook?branch=master)
[![License](http://img.shields.io/badge/license-MIT-orange.svg?style=flat)](https://github.com/osamingo/bugsnag-logrus-hook/blob/master/LICENSE)

## Description

A Hook of [Logrus](https://github.com/Sirupsen/logrus) for [Bugsnag](https://github.com/bugsnag/bugsnag-go)

## Installation

```
$ go get github.com/osamingo/bugsnag-logrus-hook
```

## Usage

```go
package main

import (
    "github.com/Sirupsen/logrus"
    "github.com/osamingo/bugsnag-logrus-hook"
)

func main() {

    h, err := bugsnagrus.NewBugsnagHook(
      "APIKey",
      "develop",
      []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel},
      1,
    )
    if err != nil {
      panic(err)
    }

    logrus.AddHook(h)

    logrus.Error("error is occured")
}
```

## License

MIT
