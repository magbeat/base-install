/*
BaseInstall helps to reinstall all packages on a new system.

BaseInstall expects an `install` directory in the following form:

```
install/
    - config/
    - files/ (optional)
    - tmp/ (optional)
```

If BaseInstall is executed without arguments it assumes the location of the `install` directory to be in $HOME

BaseInstall parses all `*.json` files in the `config` directory

There are several plugins for different package managers (like `dnf`, `npm`, `snap` and `flatpak`) and a plugin for a `custom` installation.
The `custom` installation allows the user to execute multiple commands (like `git clone` and `make`, `make install`) to install a package.
*/
package main
