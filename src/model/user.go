package model

import "github.com/itsparser/airstrike/pkg/core/model/system"

type User struct {
	system.Model
	Name  string
	Email string
	Group []Group
}
