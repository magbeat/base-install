package plugins

type NotImplementedPlugin struct {}

func NewNotImplementedPlugin() NotImplementedPlugin { return NotImplementedPlugin{} }

func (p NotImplementedPlugin) Check(task Task) (installed bool, err error) {
	installed = false
	return installed, err
}

func (p NotImplementedPlugin) Install(task Task) (success bool, err error) {
	success = true
	return success, err
}
