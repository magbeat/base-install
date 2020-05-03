package plugins

import (
	"os"
	"os/exec"
)

type SnapPlugin struct{}

func NewSnapPlugin() SnapPlugin { return SnapPlugin{} }

func (p SnapPlugin) Check(task Task) (installed bool, err error) {
	_, lerr := exec.LookPath(task.CheckValue)
	if lerr != nil {
		installed = false
	} else {
		installed = true
	}

	return installed, err
}

func (p SnapPlugin) Install(task Task) (success bool, err error) {
	success = false
	installCmd := exec.Command("sudo", "snap", "install", task.InstallPackage, task.InstallOption)
	installCmd.Stdout = os.Stdout
	err = installCmd.Run()

	if err == nil {
		success = true
	}
	return success, err
}
