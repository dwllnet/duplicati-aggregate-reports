// Package controller loads the routes for each of the controllers.
package controller

import (
	"./about"
	"./debug"
	"./home"
	"./login"
	"./register"
	"./static"
	"./status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
}
