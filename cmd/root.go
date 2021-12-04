/*
Copyright © 2021 Raj Nandan Sharma rajnandan1@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"doximus/utils"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doximus",
	Short: "Generate static API documentation",
	Long:  `Doximus helps you host curl requests for everyone to try them out`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doximus.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := os.Getwd()
	cobra.CheckErr(err)
	var websiteMap = utils.Website{
		Title: "Doximus - Static API Documentation",
		Logo:  "images/logo.png",
		Pages: "pages",
	}

	if !utils.ItExists("./site/files") {
		os.Mkdir("./site/files", 0777)
	}
	if !utils.ItExists("./site/assets/images") {
		os.Mkdir("./site/assets/images", 0777)
	}
	if !utils.ItExists("./site/site.yaml") {
		file, err := yaml.Marshal(websiteMap)
		utils.ThrowError(err)
		err = ioutil.WriteFile("./site/site.yaml", file, 0777)
		utils.ThrowError(err)
	}
	cobra.CheckErr(err)

	// Search config in home directory with name ".doximus" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".doximus")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
