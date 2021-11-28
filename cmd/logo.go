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

	"github.com/nullrocks/identicon"
	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// logoCmd represents the logo command
var logoCmd = &cobra.Command{
	Use:   "logo",
	Short: "Set logo for you site",
	Long:  `A logo will be placed wherever required. If not given a random logo will be generated`,
	Run: func(cmd *cobra.Command, args []string) {
		title := ""
		if len(args) == 0 {
			imageGen, _ := identicon.New(
				"doximus", // Namespace
				7,         // Number of blocks (Size)
				4,         // Density
				identicon.SetBackgroundColorFunction(transparentBg()),
			)
			title = randImage("main"+ksuid.New().String(), imageGen, "logo")
		} else {
			title = args[0]
		}

		dat, err := ioutil.ReadFile("./site/site.yaml")
		throwError(err)
		var websiteMap = Website{}
		err = yaml.Unmarshal(dat, &websiteMap)
		throwError(err)
		websiteMap.Logo = title

		file, err := yaml.Marshal(websiteMap)
		throwError(err)
		err = ioutil.WriteFile("./site/site.yaml", file, 0777)
		throwError(err)
	},
}

func init() {
	rootCmd.AddCommand(logoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
