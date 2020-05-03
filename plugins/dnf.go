package plugins

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

var installedPackages string

type DnfPlugin struct {
	InstalledPackages string
}

func NewDnfPlugin() DnfPlugin {
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
	return DnfPlugin{
		InstalledPackages: installedPackages,
	}
}

func (p DnfPlugin) Check(task Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

func (p DnfPlugin) Install(task Task) (success bool, err error) {
	success = false
	installCmd := exec.Command("sudo", "dnf", "install", "-y", task.InstallPackage)
	installCmd.Stdout = os.Stdout
	err = installCmd.Run()

	if err == nil {
		success = true
	}
	return success, err
}
