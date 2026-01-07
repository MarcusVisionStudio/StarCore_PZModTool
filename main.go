package main

import (
	"github.com/ImStarboyCZ/StarCore_PZModTool/commands"
	"github.com/ImStarboyCZ/StarCore_PZModTool/interactive"
	"github.com/ImStarboyCZ/StarCore_PZModTool/util"
	"github.com/ImStarboyCZ/StarCore_PZModTool/version"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "pzmod --file <server config file>",
		Short:   "pzmod is a tool for managing Project Zomboid server mods.",
		Version: version.Get(),
		Example: `pzmod --file server.ini
pzmod --file server.ini get list
pzmod --file server.ini get name
pzmod --file server.ini set name "My Server"`,
		PreRun: checkForUpdate,
		Run:    interactive.Execute,
	}

	commands.SetFileFlag(rootCmd)
	commands.Init(rootCmd)
	rootCmd.Execute()
}

func checkForUpdate(cmd *cobra.Command, args []string) {
	if !version.IsSet() {
		return
	}

	updater, err := version.NewUpdater()
	if err != nil {
		return
	}

	ver := version.Get()
	latest, err := version.GetLatestRelease(updater)
	if err != nil {
		return
	}

	if version.IsLatest(ver, latest) {
		return
	}

	cmd.Println(util.Info, "A new version of pzmod is available:", latest.Version())
	cmd.Println(util.Info, "Run `pzmod update` to update to the latest version.")
}
