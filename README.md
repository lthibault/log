# log

A high-level API adapter for [logrus](github.com/sirupsen/logrus).

[![Godoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/lthibault/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/SentimensRG/ctx?style=flat-square)](https://goreportcard.com/report/github.com/lthibault/log)

## Overview

Package `log` provides a high-level API adapter around [logrus](github.com/sirupsen/logrus).  It is intended as a personal utility.  You are likely to prefer using Logrus directly.

This package makes the following changes to the logrus API:

- [Functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) for a more user-friendly API.
- Adds `Logger.With` method, which applies log fields directly from objects satisfying `Loggable`.
- Remove panic-level logging, as `Panic` and `Fatal` are semantically equivalent.

Escape hatches are provided to maintain full compatibility with `Logrus`.
