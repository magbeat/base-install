/*
YayPlugin checks for installed packages by checking if the package was installed via yay

Example Config file:

    [
        { "plugin": "yay", "check": "firefox", "installPackage": "firefox" }
    ]
*/
package yay

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/magbeat/base-install/plugins"
)

var installedPackages string

type Plugin struct {
	InstalledPackages string
}

func NewYayPlugin() Plugin {
	if len(installedPackages) == 0 {
		var buf bytes.Buffer

		listCmd := exec.Command("yay", "-Q")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		installedPackages = string(buf.Bytes())
	}
	return Plugin{
		InstalledPackages: installedPackages,
	}
}

// Check checks if `task.CheckValue` is installed by looking at the installed Yay packages
func (p Plugin) Check(task plugins.Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

// Install installs the `task.InstallPackage` via `yay`
func (p Plugin) Install(task plugins.Task) (success bool, err error) {
	success = false
	installCmd := exec.Command("yay", "-S", task.InstallPackage, "--noconfirm")
	fmt.Println(installCmd.Args)
	err = installCmd.Run()

	if err == nil {
		success = true
	}
	return success, err
}
