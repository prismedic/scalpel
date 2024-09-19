package infofx

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/prismedic/scalpel/config"
)

var (
	BuildDate string
)

type InfoDisplay struct {
	Name        string
	Platform    string
	Runtime     string
	HostName    string
	BuildCommit string
	BuildDate   string
}

func GetInfo() (*InfoDisplay, error) {
	display := &InfoDisplay{
		Name:     config.GetPackageName(),
		Platform: fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH),
		Runtime:  runtime.Version(),
	}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = fmt.Sprintf("Fail to get hostname: %v", err)
	}
	display.HostName = hostname
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("failed to read build info")
	}
	buildCommit := os.Getenv("BUILD_COMMIT")
	for _, buildSetting := range buildInfo.Settings {
		if buildSetting.Key == "vcs.revision" {
			buildCommit = buildSetting.Value
		}
	}
	display.BuildCommit = buildCommit
	display.BuildDate = BuildDate
	return display, nil
}
