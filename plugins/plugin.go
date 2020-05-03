package plugins

type Plugin interface {
	Check(task Task) (installed bool, err error)
	Install(task Task) (success bool, err error)
}
