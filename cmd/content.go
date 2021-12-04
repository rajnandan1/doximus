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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// contentCmd represents the content command
var contentCmd = &cobra.Command{
	Use:   "content",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reference, _ := cmd.Flags().GetString("id")
		tp, _ := cmd.Flags().GetString("type")
		mode, _ := cmd.Flags().GetString("mode")
		reference = "/" + reference
		utils.ThrowLog("updating  page " + reference)

		dir := "pages" + reference
		contentFile := dir + "/content." + tp
		if !utils.ItExists(dir) {
			utils.ThrowError(errors.New("page does not exist. Use doximus page add --id=pagename"))
		}
		if !utils.ItExists(contentFile) {
			utils.CreateFile(contentFile)
		}
		contentStr := ""
		if mode == "append" {
			dat, err := ioutil.ReadFile(contentFile)
			utils.ThrowError(err)
			contentStr = string(dat)
		}
		fmt.Println(">Enter Content. Type exit to complete")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "exit" {
				break
			}
			if contentStr != "" {
				contentStr = contentStr + "\n" + scanner.Text()
			} else {
				contentStr = contentStr + scanner.Text()
			}

		}
		utils.ThrowError(scanner.Err())
		utils.ThrowError(ioutil.WriteFile(contentFile, []byte(contentStr), 0777))
	},
}

func init() {
	addCmd.AddCommand(contentCmd)
	contentCmd.Flags().String("id", "", "Add page id")
	contentCmd.Flags().String("type", "md", "Add conent for page")
	contentCmd.Flags().String("mode", "append", "Add mode for edit content for page can be append or new")
}
