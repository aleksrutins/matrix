/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/aleksrutins/matrix/config"
	"github.com/aleksrutins/matrix/log"
	"github.com/spf13/cobra"
)

var (
	rCfgVar    = regexp.MustCompile(`\$\(config\)`)
	rTargetVar = regexp.MustCompile(`\$\(target\)`)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "matrix",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Reading configuration")
		cfg, err := config.Read("matrix.toml")
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		r, err := regexp.Compile(cfg.ConfigRegex.Find)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		configContent, err := ioutil.ReadFile(cfg.ConfigPath)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		for toRun := range cfg.IterateAll() {
			newContent := r.ReplaceAllString(string(configContent), replaceVars(cfg.ConfigRegex.Replace, cfg, toRun))
			ioutil.WriteFile(cfg.ConfigPath, []byte(newContent), 777)
			stdout, err := log.Build(toRun.Configuration.Name, toRun.Command.Name, toRun.Command.Value)
			if err != nil {
				log.Error(err.Error())
				fmt.Println(stdout)
			}
			for _, file := range cfg.Move {
				err = os.Rename(replaceVars(file.Old, cfg, toRun), replaceVars(file.New, cfg, toRun))
				if err != nil {
					log.Error(err.Error())
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.matrix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func replaceVars(str string, cfg config.MatrixConfig, toRun config.Cell) string {
	return rTargetVar.ReplaceAllString(rCfgVar.ReplaceAllString(str, toRun.Configuration.Value), toRun.Command.Name)
}
