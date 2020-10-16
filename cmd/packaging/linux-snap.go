package packaging

import (
	"fmt"
	"os"
	"os/exec"
)

// LinuxSnapTask packaging for linux as snap
var LinuxSnapTask = &packagingTask{
	packagingFormatName: "linux-snap",
	templateFiles: map[string]string{
		"linux-snap/snapcraft.yaml.tmpl": "snap/snapcraft.yaml.tmpl",
		"linux/app.desktop.tmpl":         "snap/local/{{.executableName}}.desktop.tmpl",
	},
	linuxDesktopFileExecutablePath: "/{{.executableName}}",
	linuxDesktopFileIconPath:       "/icon.png",
	flutterBuildOutputDirectory:    "build",
	packagingFunction: func(tmpPath, applicationName, packageName, executableName, version, release string) (string, error) {
		cmdSnapcraft := exec.Command("snapcraft")
		cmdSnapcraft.Dir = tmpPath
		cmdSnapcraft.Stdout = os.Stdout
		cmdSnapcraft.Stderr = os.Stderr
		err := cmdSnapcraft.Run()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s_%s_amd64.snap", packageName, version), nil
	},
	requiredTools: map[string][]string{
		"linux": {"snapcraft"},
	},
	toolsInstallationInstruction: map[string]map[string]string{
		"linux": {
			"snapcraft": "Install snapd from your package manager or from https://snapcraft.io/docs/installing-snapd",
		},
	},
}
