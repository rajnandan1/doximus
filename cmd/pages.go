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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// pagesCmd represents the pages command
var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Set your pages folder.",
	Long:  `The pages command will scan a folder when you build to create the pages. It has to have at least one main file which becomes your home page`,
	Run: func(cmd *cobra.Command, args []string) {

		dat, err := ioutil.ReadFile("./site/site.yaml")
		utils.ThrowError(err)
		var websiteMap = utils.Website{}
		err = yaml.Unmarshal(dat, &websiteMap)
		utils.ThrowError(err)

		files, err := filepath.Glob(websiteMap.Pages + "/*")
		utils.ThrowError(err)
		utils.ThrowSuccess("Your pages")
		utils.ThrowLog("home")
		for _, v := range files {

			if fileInfo, err := os.Stat(v); err == nil {
				if fileInfo.IsDir() && !strings.Contains(v, "/curls") {
					utils.ThrowLog(strings.ReplaceAll(v, "pages/", ""))
				}
			}
		}

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
