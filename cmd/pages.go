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
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// pagesCmd represents the pages command
var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Set your pages folder.",
	Long:  `The pages command will scan a folder when you build to create the pages. It has to have at least one main file which becomes your home page`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			return
		}
		title := args[0]
		dat, err := ioutil.ReadFile("./site/site.yaml")
		throwError(err)
		var websiteMap = Website{}
		err = yaml.Unmarshal(dat, &websiteMap)
		throwError(err)
		websiteMap.Pages = title
		file, err := yaml.Marshal(websiteMap)
		throwError(err)
		err = ioutil.WriteFile("./site/site.yaml", file, 0777)
		throwError(err)

	},
}

func init() {
	rootCmd.AddCommand(pagesCmd)
	//rootCmd.Flags().StringP("titsle", "x", "Your Title", "Your title")
	// pagesCmd.Flags().StringP("pa", "l", "pages", "file locations for site config")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
