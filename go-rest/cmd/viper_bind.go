package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viperBind binds flags together with config file
func viperBind(cmd *cobra.Command, name string) error {
	return viper.BindPFlag(cmd.Use+"."+name, cmd.PersistentFlags().Lookup(name))
}
