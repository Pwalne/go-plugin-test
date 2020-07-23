package shared

// PluginError is just an alias to test shared lib.
type PluginError error

// Plugin represents the interface we will load.
//	NOTE: underlying interface cannot be a pointer.
type Plugin interface {
	Init() PluginError
}