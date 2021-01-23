package main

import (
	"encoding/json"
	"fmt"
	"github.com/magbeat/base-install/plugins"
	"github.com/magbeat/base-install/plugins/custom"
	"github.com/magbeat/base-install/plugins/dnf"
	"github.com/magbeat/base-install/plugins/flatpak"
	"github.com/magbeat/base-install/plugins/not_implemented"
	"github.com/magbeat/base-install/plugins/npm"
	"github.com/magbeat/base-install/plugins/pacman"
	"github.com/magbeat/base-install/plugins/snap"
	"github.com/magbeat/base-install/plugins/yay"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func main() {
	var basePath string
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	} else {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		basePath = currentUser.HomeDir
	}

	configDir := basePath + "/install/config"

	tasks := parseTasks(configDir)

	processTasks(tasks)
}

func parseTasks(configDir string) []plugins.Task {
	configFiles, err := ioutil.ReadDir(configDir)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []plugins.Task
	for _, configFile := range configFiles {
		jsonFile, err := os.Open(configDir + "/" + configFile.Name())
		fmt.Println("Loading config file ", jsonFile.Name())
		if err != nil {
			fmt.Println("Error while loading config file ", jsonFile.Name())
			log.Fatal(err)
		}

		byteValue, err := ioutil.ReadAll(jsonFile)
		var tmpTasks []plugins.Task
		err = json.Unmarshal(byteValue, &tmpTasks)
		if err != nil {
			fmt.Println("Error while unmarshalling config file ", jsonFile.Name())
			log.Fatal(err)
		}
		tasks = append(tasks, tmpTasks...)
	}
	return tasks
}
func processTasks(tasks []plugins.Task) {
	for _, task := range tasks {
		fmt.Printf("[%s] Checking %s: ", task.Plugin, task.CheckValue)
		installed := false
		success := false
		var err error
		var plugin plugins.Plugin

		switch task.Plugin {
		case plugins.Dnf:
			plugin = dnf.NewDnfPlugin()
		case plugins.Snap:
			plugin = snap.NewSnapPlugin()
		case plugins.Flatpack:
			plugin = flatpak.NewFlatpakPlugin()
		case plugins.Custom:
			plugin = custom.NewCustomPlugin()
		case plugins.Npm:
			plugin = npm.NewNpmPlugin()
		case plugins.Pacman:
			plugin = pacman.NewPacmanPlugin()
		case plugins.Yay:
			plugin = yay.NewYayPlugin()
		default:
			plugin = not_implemented.NewNotImplementedPlugin()
		}

		installed, err = plugin.Check(task)

		if err != nil {
			log.Printf("Error while checking %s with plugin %s", task.CheckValue, task.Plugin)
			log.Printf(err.Error())
		}

		if installed {
			fmt.Println(" ... installed")
		} else {
			fmt.Println(" ... installing")
			success, err = plugin.Install(task)
			if success {
				fmt.Println("... done")
			}
		}

		if err != nil {
			log.Printf("Error while installing %s with plugin %s", task.CheckValue, task.Plugin)
			log.Printf(err.Error())
		}

	}
}
