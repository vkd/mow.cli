package container

import "flag"

/*
Container holds an option or an arg data
*/
type Container struct {
	Name           string
	Desc           string
	Names          []string
	HideDefault    bool
	ValueSetByUser *bool
	Value          flag.Value
	DefaultValue   string
	Hidden         bool
}
