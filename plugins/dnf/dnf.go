/*
DnfPlugin checks for installed `dnf` packages by checking if the package was installed via yum / dnf 
DnfPlugin installs the `dnf` package using `dnf`

Example Config file:

    [
        { "plugin": "dnf", "check": "thunderbird", "installPackage": "thunderbird" }
    ]
 */
package dnf

import (
	"bytes"
	"github.com/magbeat/base-install/plugins"
	"log"
	"os"
	"os/exec"
	"strings"
)

var installedPackages string

type Plugin struct {
	InstalledPackages string
}

func NewDnfPlugin() Plugin {
	if len(installedPackages) == 0 {
		var buf bytes.Buffer

		listCmd := exec.Command("yum", "list", "installed")
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

// Check checks if `task.CheckValue` is installed by looking at the installed yum packages
func (p Plugin) Check(task plugins.Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

// Install installs the `task.InstallPackage` via `dnf` 
func (p Plugin) Install(task plugins.Task) (success bool, err error) {
	success = false
	installCmd := exec.Command("sudo", "dnf", "install", "-y", task.InstallPackage)
	installCmd.Stdout = os.Stdout
	err = installCmd.Run()

	if err == nil {
		success = true
	}
	return success, err
}
