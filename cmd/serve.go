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
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server at a port",
	Long:  `This will start a server at a specified port.`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		templateDirs := []string{
			"./site/*.html",
		}

		templateFiles := []string{}
		for _, dir := range templateDirs {
			ff, err := filepath.Glob(dir)
			if err != nil {
				panic(err)
			}
			templateFiles = append(templateFiles, ff...)
		}

		renderer := &TemplateRenderer{
			Templates: template.Must(template.ParseFiles(templateFiles...)),
		}
		e.Renderer = renderer

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Static("/", "./site/assets")
		e.POST("/api/call", Call)
		e.GET("/", Site)
		e.GET("/docs/*", Site)
		e.GET("/files/:name", ResponseFile)
		port := "5000"
		if len(args) != 0 {
			port = args[0]
		}
		e.Logger.Fatal(e.Start(":" + port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
