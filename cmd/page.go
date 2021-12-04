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
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// func completer(d prompt.Document) []prompt.Suggest {
// 	s := []prompt.Suggest{}
// 	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
// }

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
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
		tags, _ := cmd.Flags().GetStringSlice("tags")
		reference = "/" + reference
		utils.ThrowLog("updating  page " + reference)
		dir := "pages" + reference
		mainYaml := dir + "/main.yaml"

		if !utils.ItExists(dir) {
			utils.CreateDir(dir)
		}
		if !utils.ItExists(mainYaml) {
			utils.CreateFile(mainYaml)
		}
		dat, err := ioutil.ReadFile(mainYaml)
		utils.ThrowError(err)
		var mainDetail = utils.DetailsJSON{}
		err = yaml.Unmarshal(dat, &mainDetail)
		utils.ThrowError(err)

		if title != "" {
			mainDetail.Name = title
		}
		if desc != "" {
			mainDetail.Description = desc
		}
		if logo != "" {
			mainDetail.Logo = logo
		}
		if len(tags) > 0 {
			mainDetail.Tags = tags
		}
		file, err := yaml.Marshal(mainDetail)
		utils.ThrowError(err)
		utils.ThrowError(ioutil.WriteFile(mainYaml, file, 0777))
	},
}

func init() {

	addCmd.AddCommand(pageCmd)
	pageCmd.Flags().String("id", "", "Add page id")
	pageCmd.Flags().String("title", "", "Add title for page")
	pageCmd.Flags().String("desc", "", "Add description for page")
	pageCmd.Flags().String("logo", "", "Add logo for page")
	pageCmd.Flags().StringSlice("tags", []string{}, "Add tags for page")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
