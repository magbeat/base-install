package plugins

type Task struct {
	Plugin         PluginType `json:"plugin"`
	CheckType      CheckType  `json:"checkType"`
	CheckValue     string     `json:"check"`
	InstallPackage string     `json:"installPackage"`
	InstallOption  string     `json:"installOption"`
	Commands       []string   `json:"commands"`
}

type PluginType string

const (
	Dnf      PluginType = "dnf"
	Snap     PluginType = "snap"
	Flatpack PluginType = "flatpak"
	Custom   PluginType = "custom"
	Npm      PluginType = "npm"
	Pacman   PluginType = "pacman"
)

type CheckType string

const (
	Binary    CheckType = "bin"
	Directory CheckType = "dir"
	Yum       CheckType = "yum"
)
