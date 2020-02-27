package model

import "github.com/itsparser/airstrike/pkg/core/model/system"

type Command struct {
	Type     string
	Xpath    string
	Value    string
	Variable string
}

type Step struct {
	Id          string
	Type        string
	Description string
	Command     Command
	Async       bool
}

type Group struct {
	Type      string
	Steps     []Step
	AsyncStep []Step
	system.Model
}
