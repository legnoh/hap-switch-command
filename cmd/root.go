package cmd

import (
	"github.com/brutella/hap/accessory"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Cmd struct {
	Path string
	Args []string
}

type Config struct {
	Name     string `default:"hap-switch-command"`
	Pin      string `default:"12344321"`
	Switches []struct {
		Meta    accessory.Info
		Command struct {
			On  Cmd
			Off Cmd
		}
	}
}

var (
	cfgFile          string
	conf             Config
	confDir          string
	fsStoreDirectory string
	version          string
	debug            bool
	log              *logrus.Logger
)

var rootCmd = &cobra.Command{
	Use:     "hap-switch-command",
	Version: version,
	Short:   "Daemon for execute command with Apple HomeKit switches",
	Long: `This tool provides daemon service for Apple HomeKit Switch devices
You could execute command when switch is on/off state.

If you have any questions, please visit github site.
https://github.com/legnoh/hap-switch-command`,
}

func Execute() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	if rootCmd.Execute() != nil {
		log.Fatal("Root execute is failed... exit")
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "print debug log")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	debug, err := rootCmd.PersistentFlags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		log.SetLevel(logrus.DebugLevel)
	}
}
