package config

import (
	"github.com/ImStarboyCZ/StarCore_PZModTool/ini"
	"github.com/ImStarboyCZ/StarCore_PZModTool/steam"
	"github.com/ImStarboyCZ/StarCore_PZModTool/util"
	"github.com/spf13/cobra"
)

func LoadConfig(cmd *cobra.Command) (*ini.ServerConfig, error) {
	apiKey, err := util.LoadCredentials()
	if err != nil {
		cmd.Println(util.Warning, "Steam API key not found")
	} else {
		steam.SetApiKey(apiKey)
	}

	configPath := cmd.Flag("file").Value.String()
	return ini.LoadNewServerConfig(configPath)
}

func UnsafeLoadConfig(cmd *cobra.Command) *ini.ServerConfig {
	config, err := LoadConfig(cmd)
	cobra.CheckErr(err)
	return config
}
