package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/oussamarouabah/template/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:               "service",
	Short:             "micro-service that service traffic for our app",
	PersistentPreRunE: initConfig,
	RunE:              run,
}

func main() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file location")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	var cfg app.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("couldn't unmarshal config file: %v", err)
	}

	ctx, _ := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	a := app.New(&cfg)
	a.Start(ctx)

	// Wait until the app service stop
	// by receiving a signal or if something wrong happened during the
	// the initialization of the app
	<-a.Done()
	slog.Warn("application has exited")
	return nil
}

func initConfig(cmd *cobra.Command, args []string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/service")
		viper.AddConfigPath("./config")
	}
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("couldn't read config: %v", err)
	}
	return nil
}
