/*
CustomPlugin checks for installed package by checking if the binary is in $PATH, directory or file exists or via yum.
CustomPlugin installs the package with the supplied commands

Example Config file:

    [
        { 
            "plugin": "custom", "check": "autorandr", "checkType": "bin", "commands": 
                [
                    "git clone https://github.com/wertarbyte/autorandr.git $HOME/install/tmp/autorandr",
                    "cp $HOME/install/tmp/autorandr/autorandr $HOME/.local/bin/autorandr",
                    "chmod +x $HOME/.local/bin/autorandr"
                ]
        }
    ]
 */

package custom

import (
	"bytes"
	"github.com/magbeat/base-install/plugins"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Plugin struct{}

func NewCustomPlugin() Plugin { return Plugin{} }

// Check determines if a package `task.CheckValue` is installed by one the following methods (`task.CheckType`):
// - binary: binary is in $PATH
// - directory: if directory or file exists
// - yum: if package is installed via yum / dnf
func (p Plugin) Check(task plugins.Task) (installed bool, err error) {
	switch task.CheckType {
	case plugins.Binary:
		_, lookPathErr := exec.LookPath(task.CheckValue)
		if lookPathErr != nil {
			installed = false
		} else {
			installed = true
		}
	case plugins.Directory:
		path := os.ExpandEnv(task.CheckValue)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			installed = true
		} else {
			installed = false
		}
	case plugins.Yum:
		var buf bytes.Buffer

		listCmd := exec.Command("yum", "list", "installed")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		installedPackages := string(buf.Bytes())
		installed = strings.Contains(installedPackages, task.CheckValue)
	}
	return installed, err
}

// Install installs the package by joining the supplied commands with '&&' and running them in order via bash
func (p Plugin) Install(task plugins.Task) (success bool, err error) {
	cmd := exec.Command("bash", "-c", strings.Join(task.Commands, " && "))

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	err = cmd.Run()

	log.Println(stdBuffer.String())
	return success, err
}
