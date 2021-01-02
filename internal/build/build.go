package build

import (
	"fmt"

	"git.martianoids.com/queru/retroserver/internal/cfg"
)

var version string = "undefined"
var user string = "undefined"
var time string = "undefined"
var number string = "undefined"

// Version - complete version string
func Version() string {
	if cfg.IsDev() {
		return "DEV compiled at DEV by DEV (build #DEV)"
	}
	if cfg.IsTest() {
		return "TEST compiled at TEST by TEST (build #TEST)"
	}
	return fmt.Sprintf("%s compiled at %s by %s (build #%s)",
		version, time, user, number)
}

// VersionShort - short version string
func VersionShort() string {
	if cfg.IsTest() {
		return "TEST"
	}
	return version
}

// BinVersion - Bin version for AT
func BinVersion() string {
	return fmt.Sprintf("Bin version (%s %s): %s", user, number, version)
}

// CompileTime - Compile time string
func CompileTime() string {
	return fmt.Sprintf("compile time:%s", time)
}
