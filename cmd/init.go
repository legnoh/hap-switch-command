package cmd

import (
	"os"
	"path/filepath"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed sample/configs.yml
var samples []byte

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create default config file",
	Run:   createConf,
}

func init() {
	rootCmd.AddCommand(initCmd)

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	confDir = home + "/.hap-switch-command"

	initCmd.Flags().StringVarP(&cfgFile, "config", "c", confDir+"/config.yml", "config file path")
	initCmd.Flags().StringVarP(&fsStoreDirectory, "fs-store", "f", confDir+"/db", "fsStore directory path")
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func createConf(cmd *cobra.Command, args []string) {
	rootDir := filepath.Dir(cfgFile)
	if Exists(rootDir) {
		log.Debugf("Directory already existed: %s", rootDir)
	} else {
		if err := os.Mkdir(rootDir, 0755); err != nil {
			log.Fatalf("Failed to create directory %s", rootDir)
		}
	}

	if Exists(cfgFile) {
		log.Fatalf("Config file: %s is already existed.", cfgFile)
	} else {
		f, err := os.Create(cfgFile)
		if err != nil {
			log.Fatalf("fail to create file: %s", err)
		}
		defer f.Close()
		if _, err := f.Write(samples); err != nil {
			log.Fatalf("fail to write file: %s", err)
		}
		log.Infof("Default config file created successfully: %s", cfgFile)
	}
}
