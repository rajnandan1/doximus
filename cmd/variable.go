/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// variableCmd represents the variable command
var variableCmd = &cobra.Command{
	Use:   "variable",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reference, _ := cmd.Flags().GetString("id")
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		value, _ := cmd.Flags().GetString("value")
		required, _ := cmd.Flags().GetBool("required")
		reference = "/" + reference
		utils.ThrowLog("updating  page " + reference)
		dir := "pages/" + reference
		apiYaml := dir + "/api.yaml"
		if !utils.ItExists(dir) {
			utils.ThrowError(errors.New("page does not exist. Use doximus page add --id=pagename"))
		}
		if !utils.ItExists(apiYaml) {
			utils.CreateFile(apiYaml)
		}
		dat, err := ioutil.ReadFile(apiYaml)
		utils.ThrowError(err)
		var apiDetail = utils.CurlJSON{}
		err = yaml.Unmarshal(dat, &apiDetail)
		utils.ThrowError(err)
		configurables := apiDetail.Configurables
		newConf := utils.APIConfigurables{}

		if title != "" {
			newConf.Title = title
		}
		if desc != "" {
			newConf.Description = desc
		}
		if value != "" {
			newConf.Init = value
		}
		if required {
			newConf.Required = required
		}
		index := -1
		//search if present
		for i, conf := range configurables {
			if strings.ToLower(title) == strings.ToLower(conf.Title) {
				index = i
				newConf = conf
				if desc != "" {
					newConf.Description = desc
				}
				if value != "" {
					newConf.Init = value
				}
				if required {
					newConf.Required = required
				}
				break
			}
		}
		if index == -1 {
			configurables = append(configurables, newConf)
		} else {
			configurables[index] = newConf
		}
		apiDetail.Configurables = configurables
		file, err := yaml.Marshal(apiDetail)
		utils.ThrowError(err)
		utils.ThrowError(ioutil.WriteFile(apiYaml, file, 0777))

	},
}

func init() {
	addCmd.AddCommand(variableCmd)
	variableCmd.Flags().String("id", "", "Add page id")
	variableCmd.Flags().String("title", "", "Add name of variable")
	variableCmd.Flags().String("desc", "", "Add description of variable")
	variableCmd.Flags().String("value", "", "Add initial value of variable")
	variableCmd.Flags().Bool("required", false, "optional or required")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// variableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// variableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
