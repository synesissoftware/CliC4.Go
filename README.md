# CliC4.Go <!-- omit in toc -->

**CliC4.Go** provides C++-like collections for Go

![Language](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/CliC4.Go.svg)](https://github.com/synesissoftware/CliC4.Go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/CliC4.Go)](https://github.com/synesissoftware/CliC4.Go/commits/master)
[![Go](https://github.com/synesissoftware/CliC4.Go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/CliC4.Go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/CliC4.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/CliC4.Go)


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
  - [Version](#version)
- [Examples](#examples)
- [Project Information](#project-information)
  - [Where to get help](#where-to-get-help)
  - [Contribution guidelines](#contribution-guidelines)
  - [Dependencies](#dependencies)
    - [Development/Testing Dependencies](#developmenttesting-dependencies)
  - [Related projects](#related-projects)
  - [License](#license)


## Introduction

**CliC4** (C++-like Collections) is a family of libraries providing C++-like collection types and utilities. **CliC4.Go** is the **Go** implementation.


## Installation

```Go
import clic4go "github.com/synesissoftware/CliC4.Go"
```


## Components


### Version

```Go
// in "github.com/synesissoftware/CliC4.Go"

const (
  VersionMajor uint16 = /* ... */
  VersionMinor uint16 = /* ... */
  VersionPatch uint16 = /* ... */
  VersionAB    uint16 = /* ... */
)

func Version() uint64
func VersionString() string
```


## Examples

Examples are provided in the `examples` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/CliC4.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/CliC4.Go.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Testing Dependencies

* [**testify**](https://github.com/stretchr/testify);


### Related projects

**CliC4.Go** is currently the only implementation in the **CliC4** family.


### License

**CliC4.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->
