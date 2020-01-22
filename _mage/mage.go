package main

import (
	mageutil "github.com/getcouragenow/bootstrap/_mage/util"
	"os"
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

func WindowsDependency() {
	mageutil.Windows{}.InstallDependency()

}
func MacInstall() {
	mageutil.Mac{}.InstallDependency()
}
