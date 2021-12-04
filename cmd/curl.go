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
	"bufio"
	"doximus/utils"
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"mvdan.cc/xurls"
)

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl",
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
		if title == "" {
			utils.ThrowError(errors.New("missing title. use --title=\"Create User\""))
		}
		reference = "/" + reference
		utils.ThrowLog("updating  page " + reference)
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		dir := "pages" + reference
		if !utils.ItExists(dir) {
			utils.ThrowError(errors.New("page does not exist. Use doximus page add --id=pagename"))
		}
		curlDir := dir + "/curls"
		curlYamlFile := curlDir + "/" + reg.ReplaceAllString(title, "") + ".yaml"
		apiYamlFile := dir + "/api.yaml"
		if !utils.ItExists(apiYamlFile) {
			utils.ThrowError(errors.New("add api before adding curls"))
		}
		dat, err := ioutil.ReadFile(apiYamlFile)
		utils.ThrowError(err)
		var apiDetail = utils.CurlJSON{}
		err = yaml.Unmarshal(dat, &apiDetail)
		utils.ThrowError(err)
		//currentAPIs := apiDetail.Curls
		curlExamples := utils.CurlYaml{
			Name: title,
		}
		utils.ThrowInput("Enter curl. Type exit to end")
		curStr := ""
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanner.Scan()
			text := scanner.Text()
			if strings.HasSuffix(text, "\\") {
				text = text[0:(len(text) - 1)]
			}
			if text == "exit" || len(text) == 0 {
				break
			}
			if len(text) != 0 {
				curStr = curStr + text

			}

		}

		rxRelaxed := xurls.Relaxed
		curlURL := rxRelaxed.FindString(curStr)
		brokenURL, _ := url.Parse(curlURL)
		path := brokenURL.Path
		host := brokenURL.Host
		// host:=brokenURL.Fragment

		//fmt.Println(brokenURL.Scheme)
		urlIndexInCurl := strings.Index(curStr, curlURL)
		startOfCurl := curStr[0:urlIndexInCurl]
		method := "GET"
		if strings.Contains(startOfCurl, "POST") {
			method = "POST"
		} else if strings.Contains(startOfCurl, "PUT") {
			method = "PUT"
		} else if strings.Contains(startOfCurl, "DELETE") {
			method = "DELETE"
		}
		if !utils.ItemExists(apiDetail.Domains, host) {
			apiDetail.Domains = append(apiDetail.Domains, host)
		}

		curlOne := utils.CurlAPIJSON{
			Method:      method,
			Name:        title,
			Description: desc,
			Path:        path,
		}
		newAPI := true
		for _, c := range apiDetail.Curls {
			if c.Path == curlOne.Path && c.Method == curlOne.Method {
				newAPI = false
				break
			}
		}
		if newAPI {
			apiDetail.Curls = append(apiDetail.Curls, curlOne)
		}
		curlExamples.Curl = curStr
		if !utils.ItExists(curlDir) {
			utils.CreateDir(curlDir)
		}
		if !utils.ItExists(curlYamlFile) {
			utils.CreateFile(curlYamlFile)
		}
		file, err := yaml.Marshal(curlExamples)
		utils.ThrowError(err)
		utils.ThrowError(ioutil.WriteFile(curlYamlFile, file, 0777))

		file2, err := yaml.Marshal(apiDetail)
		utils.ThrowError(err)
		utils.ThrowError(ioutil.WriteFile(apiYamlFile, file2, 0777))

	},
}

func init() {
	addCmd.AddCommand(curlCmd)
	curlCmd.Flags().String("id", "", "Add page id")
	curlCmd.Flags().String("title", "", "Add title of api example")
	curlCmd.Flags().String("desc", "", "Add description  api example")
}
