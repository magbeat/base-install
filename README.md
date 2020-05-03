# .

[![Build Status](https://travis-ci.org/magbeat/base-install.svg?branch=master)](https://travis-ci.org/magbeat/base-install)
[![codecov](https://codecov.io/gh/magbeat/base-install/branch/master/graph/badge.svg)](https://codecov.io/gh/magbeat/base-install)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/magbeat/base-install)

BaseInstall helps to reinstall all packages on a new system.

BaseInstall expects an `install` directory in the following form:

install/

```diff
- config/
- files/ (optional)
- tmp/ (optional)
```

## If BaseInstall is executed without arguments it assumes the location of the `install` directory to be in $HOME

BaseInstall parses all `*.json` files in the `config` directory

There are several plugins for different package managers (like `dnf`, `npm`, `snap` and `flatpak`) and a plugin for a `custom` installation.
The `custom` installation allows the user to execute multiple commands (like `git clone` and `make`, `make install`) to install a package.

## Sub Packages

* [plugins](./plugins): NpmPlugin checks for installed npm packages by checking if the binary is available.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
