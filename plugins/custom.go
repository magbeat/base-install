package plugins

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type CustomPlugin struct{}

func NewCustomPlugin() CustomPlugin { return CustomPlugin{} }

func (p CustomPlugin) Check(task Task) (installed bool, err error) {
	switch task.CheckType {
	case Binary:
		_, lookPathErr := exec.LookPath(task.CheckValue)
		if lookPathErr != nil {
			installed = false
		} else {
			installed = true
		}
	case Directory:
		path := os.ExpandEnv(task.CheckValue)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			installed = true
		} else {
			installed = false
		}
	case Yum:
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

func (p CustomPlugin) Install(task Task) (success bool, err error) {
	cmd := exec.Command("bash", "-c", strings.Join(task.Commands, " && "))

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	err = cmd.Run()

	log.Println(stdBuffer.String())
	return success, err
}
