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
		Use:     "starcore_modtool --file <server config file>",
		Short:   "starcore_modtool is a tool for managing Project Zomboid server mods.",
		Version: version.Get(),
		Example: `starcore_modtool --file server.ini
starcore_modtool --file server.ini get list
starcore_modtool --file server.ini get name
starcore_modtool --file server.ini set name "My Server"`,
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

	cmd.Println(util.Info, "A new version of starcore_modtool is available:", latest.Version())
	cmd.Println(util.Info, "Run `starcore_modtool update` to update to the latest version.")
}
