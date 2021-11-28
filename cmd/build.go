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
		throwError(err)
		var websiteMap = Website{}
		err = yaml.Unmarshal(dat, &websiteMap)
		throwError(err)

		siteLoc := websiteMap.Pages
		if siteLoc == "" {
			err = errors.New("no pages found")
			throwError(err)
		}
		imageGen, _ := identicon.New(
			"doximusmaindoc",
			5,
			4,
			identicon.SetBackgroundColorFunction(transparentBg()),
		)
		splitted := strings.Split(siteLoc, "/")
		lastFolder := strings.Join(splitted[0:(len(splitted)-1)], "/")
		mode := "macos"
		if len(args) > 0 {
			mode = args[0]
		}
		clearGenratedFiles("./site/assets/images", "doximus")
		clearGenratedFiles("./site/files", "docs.")
		createdocs(siteLoc, imageGen, lastFolder, mode)
		writeDocs()
		file, err := json.MarshalIndent(websiteMap, "", " ")
		throwError(err)
		err = ioutil.WriteFile("./site/files/doximussite.json", file, 0644)
		throwError(err)
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
