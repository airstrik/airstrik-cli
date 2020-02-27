package model

import "github.com/itsparser/airstrike/pkg/core/model/system"

type Group struct {
	system.Model
	Name  string
	Label string
}