package main

import (
	"fmt"
	"github.com/TEAMSTARWORLD/logrus"
	"github.com/TEAMSTARWORLD/phoneinfoga/v2/build"
	"github.com/TEAMSTARWORLD/phoneinfoga/v2/cmd"
	"github.com/TEAMSTARWORLD/phoneinfoga/v2/logs"
)

func main() {
	logs.Init()
	logrus.WithFields(logrus.Fields{
		"isRelease": fmt.Sprintf("%t", build.IsRelease()),
		"version":   build.String(),
	}).Debug("Build info")
	cmd.Execute()
}
