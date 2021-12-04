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
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// siteCmd represents the site command
var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ThrowLog("updating  site layout")
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		logo, _ := cmd.Flags().GetString("logo")
		tags, _ := cmd.Flags().GetStringSlice("tags")
		dat, err := ioutil.ReadFile("./site/site.yaml")
		utils.ThrowError(err)
		var websiteMap = utils.Website{}
		err = yaml.Unmarshal(dat, &websiteMap)

		if title != "" {
			websiteMap.Title = title
		}
		if desc != "" {
			websiteMap.Description = desc
		}
		if logo != "" {
			if _, err := url.ParseRequestURI(logo); err != nil {
				if !strings.HasPrefix(logo, "/") {
					websiteMap.Logo = "/" + logo
				}
			} else {
				response, e := http.Get(logo)
				if e != nil {
					log.Fatal(e)
				}
				defer response.Body.Close()
				file, err := os.Create("images/logo.png")
				utils.ThrowError(err)
				defer file.Close()
				_, err = io.Copy(file, response.Body)
				utils.ThrowError(err)
				websiteMap.Logo = logo
				logo = "images/logo.png"

			}

			if img, err := utils.GetImageFromFilePath(logo); err == nil {
				dstImage800 := imaging.Resize(img, 16, 0, imaging.Box)

				if tt, err := utils.GetFileContentType(logo); err == nil && tt == "image/png" {
					f, err := os.Create("images/favicon.png")
					utils.ThrowError(err)
					defer f.Close()
					utils.ThrowError(png.Encode(f, dstImage800))
				}

			}

		}
		if len(tags) > 0 {
			websiteMap.Tags = tags
		}
		file, err := yaml.Marshal(websiteMap)
		utils.ThrowError(err)
		err = ioutil.WriteFile("./site/site.yaml", file, 0777)
		utils.ThrowError(err)
	},
}

func init() {
	rootCmd.AddCommand(siteCmd)
	siteCmd.Flags().String("title", "", "Add title for site")
	siteCmd.Flags().String("desc", "", "Add description for site")
	siteCmd.Flags().String("logo", "", "Add logo for site")
	siteCmd.Flags().StringSlice("tags", []string{}, "Add tags for site")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// siteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// siteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
