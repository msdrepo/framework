package controllers

import (
	folcon "github.com/msdrepo/framework"
)

// Controller defines the standard interface that all application controllers must implement
type Controller interface {
	Init(c *folcon.Context)
}

// BaseController provides shared helper methods for all controllers
type BaseController struct{}
