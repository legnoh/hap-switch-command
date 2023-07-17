package cmd

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/creasty/defaults"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type A []*accessory.A

var switches A

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start homekit switch daemon",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: preStartServer,
	Run:    startServer,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	confDir = home + "/.hap-switch-command"

	serveCmd.Flags().StringVarP(&cfgFile, "config", "c", confDir+"/config.yml", "config file path")
	serveCmd.Flags().StringVarP(&fsStoreDirectory, "fs-store", "f", confDir+"/db", "fsStore directory path")
}

func preStartServer(cmd *cobra.Command, args []string) {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}
	if err := defaults.Set(&conf); err != nil {
		log.Fatal(err)
	}
	if !regexp.MustCompile(`^[0-9]{8}$`).MatchString(conf.Pin) {
		log.Fatalf("Your PinCode(%s) is invalid format. Please fix to 8-digit code(e.g. 12344321)", conf.Pin)
	}
}

func startServer(cmd *cobra.Command, args []string) {
	bridge := accessory.NewBridge(accessory.Info{
		Name:         conf.Name,
		Manufacturer: "@legnoh",
		Model:        version,
	})

	for _, v := range conf.Switches {
		a := accessory.NewSwitch(v.Meta)
		a.Switch.On.OnValueRemoteUpdate(func(on bool) {
			if on {
				log.Info("Switch state changed: on")
				execCommand(v.Command.On)
			} else {
				log.Info("Switch state changed: off")
				execCommand(v.Command.Off)
			}
		})
		switches = append(switches, a.A)
	}

	fs := hap.NewFsStore(fsStoreDirectory)
	server, err := hap.NewServer(fs, bridge.A, switches...)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	server.Pin = conf.Pin

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		log.Info("Stopping HAP Server...")
		signal.Stop(c)
		cancel()
	}()

	log.Info("Starting HAP Server...")
	log.Infof("Device Name: %s", bridge.Name())
	log.Infof("   Pin Code: %s", conf.Pin)
	log.Debugf("Config File: %s", cfgFile)
	log.Debugf(" Store Path: %s", fsStoreDirectory)
	server.ListenAndServe(ctx)
}

func execCommand(cmd Cmd) {
	var stdout, stderr bytes.Buffer

	command := exec.Command(cmd.Path, cmd.Args...)
	log.Debugf("Command: %s", command)
	command.Stdout = &stdout
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		log.Errorf("Result: Failed: %s", err)
	} else {
		log.Info("Result: Success")
	}
	if stdoutString := stdout.String(); stdoutString != "" {
		log.Debugf("[stdout]: %s", stdoutString)
	}
	if stderrString := stderr.String(); stderrString != "" {
		log.Errorf("[stderr]: %s", stderrString)
	}
}
