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
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/nullrocks/identicon"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your site",
	Long:  `Will build your site.`,
	Run: func(cmd *cobra.Command, args []string) {
		//read the site file
		dat, err := ioutil.ReadFile("./site/site.yaml")
		utils.ThrowError(err)
		var websiteMap = utils.Website{}
		err = yaml.Unmarshal(dat, &websiteMap)
		utils.ThrowError(err)

		siteLoc := websiteMap.Pages
		if siteLoc == "" {
			err = errors.New("no pages found")
			utils.ThrowError(err)
		}
		imageGen, _ := identicon.New(
			"doximusmaindoc",
			5,
			4,
			identicon.SetBackgroundColorFunction(utils.TransparentBg()),
		)
		splitted := strings.Split(siteLoc, "/")
		lastFolder := strings.Join(splitted[0:(len(splitted)-1)], "/")
		mode := "macos"
		if len(args) > 0 {
			mode = args[0]
		}
		utils.ClearGenratedFiles("./site/assets/images", "doximus")
		utils.ClearGenratedFiles("./site/files", "docs.")
		utils.Createdocs(siteLoc, imageGen, lastFolder, mode)
		utils.WriteDocs()
		file, err := json.MarshalIndent(websiteMap, "", " ")
		utils.ThrowError(err)
		err = ioutil.WriteFile("./site/files/doximussite.json", file, 0644)

	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	//buildCmd.Flags().StringP("location", "l", "config.yaml", "Add Floating Numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
