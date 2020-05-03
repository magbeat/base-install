/*
NotImplemented is a placedholder for an unknown plugin and always return `true` in the check step
*/
package not_implemented

import "github.com/magbeat/base-install/plugins"

type Plugin struct {}

func NewNotImplementedPlugin() Plugin { return Plugin{} }

// Check always returns `true` for `installed`
func (p Plugin) Check(task plugins.Task) (installed bool, err error) {
	installed = true
	return installed, err
}

// Install always returns `true` for `success`
func (p Plugin) Install(task plugins.Task) (success bool, err error) {
	success = true
	return success, err
}
