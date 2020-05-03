package plugins

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

var flatpakPackages string

type FlatpakPlugin struct {
	InstalledPackages string
}

func NewFlatpakPlugin() FlatpakPlugin {
	if len(flatpakPackages) == 0 {
		var buf bytes.Buffer

		listCmd := exec.Command("flatpak", "list")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		flatpakPackages = string(buf.Bytes())

		addRemoteCmd := exec.Command("flatpak", "remote-add", "--if-not-exists", "flathub", "https://flathub.org/repo/flathub.flatpakrepo")
		err = addRemoteCmd.Run()
	}
	return FlatpakPlugin{
		InstalledPackages: flatpakPackages,
	}
}

func (p FlatpakPlugin) Check(task Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

func (p FlatpakPlugin) Install(task Task) (success bool, err error) {
	success = false
	installCmd := exec.Command("sudo", "flatpak", "install", task.InstallOption, task.InstallPackage)
	installCmd.Stdout = os.Stdout
	err = installCmd.Run()

	if err == nil {
		success = true
	}
	return success, err
}
