package commands

import (
	"github.com/ImStarboyCZ/StarCore_PZModTool/util"
	"github.com/ImStarboyCZ/StarCore_PZModTool/version"
	"github.com/spf13/cobra"
)

func cmdUpdate() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "Update starcore_modtool to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			if !version.IsSet() {
				cobra.CheckErr(util.ErrVerNotSet)
			}

			updater, err := version.NewUpdater()
			cobra.CheckErr(err)

			ver := version.Get()
			latest, err := version.GetLatestRelease(updater)
			cobra.CheckErr(err)

			if version.IsLatest(ver, latest) {
				cmd.Println("starcore_modtool is already up to date")
				return
			}

			check, _ := cmd.Flags().GetBool("check")
			if check {
				cmd.Println("A new version of starcore_modtool is available:", latest.Version())
				return
			}

			err = version.Update(ver, latest, updater)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().BoolP("check", "c", false, "check for updates only")

	return cmd
}
