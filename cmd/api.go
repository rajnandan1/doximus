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

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
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
		logo, _ := cmd.Flags().GetString("logo")
		domains, _ := cmd.Flags().GetStringSlice("domains")
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

		if title != "" {
			apiDetail.Name = title
		}
		if desc != "" {
			apiDetail.Description = desc
		}
		if logo != "" {
			apiDetail.Logo = logo
		}

		for _, d := range domains {
			if !utils.ItemExists(apiDetail.Domains, d) {
				apiDetail.Domains = append(apiDetail.Domains, d)
			}
		}

		file, err := yaml.Marshal(apiDetail)
		utils.ThrowError(err)
		utils.ThrowError(ioutil.WriteFile(apiYaml, file, 0777))
	},
}

func init() {
	addCmd.AddCommand(apiCmd)
	apiCmd.Flags().String("title", "", "Add title for api")
	apiCmd.Flags().String("desc", "", "Add description for api")
	apiCmd.Flags().String("logo", "", "Add logo for api")
	apiCmd.Flags().String("id", "", "Add page id")
	apiCmd.Flags().StringSlice("domains", []string{}, "Add domains for site")
}
