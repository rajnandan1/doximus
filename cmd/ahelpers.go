package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"image/color"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nullrocks/identicon"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

func itExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

func transparentBg() func(cb []byte, fc color.Color) color.Color {
	return func(cb []byte, fc color.Color) color.Color {
		return color.Transparent
	}
}
func readMainFile(fullPath string, pageID string, gen *identicon.Generator) *DetailsJSON {
	detail := DetailsJSON{}
	dat, err := ioutil.ReadFile(fullPath)
	throwError(err)
	err = yaml.Unmarshal(dat, &detail)
	throwError(err)

	if detail.Logo == "" {
		detail.Logo = randImage("main.yaml"+pageID, gen, "doximus")
	}
	return &detail
}

func randImage(s string, gen *identicon.Generator, prefix string) string {
	num := hash(s)
	s = prefix + num
	var ii *identicon.IdentIcon
	var err error
	ii, err = gen.Draw(s)
	throwError(err)
	img, err := os.Create("./site/assets/images/" + s + ".png")
	throwError(err)
	defer img.Close()
	ii.Png(90, img)

	return "/images/" + s + ".png"
}

func writeFiles(data map[string]interface{}, tt string) {
	for k, v := range data {
		file, err := json.MarshalIndent(v, "", " ")
		throwError(err)
		err = ioutil.WriteFile("./site/files/"+k+"."+tt, file, 0644)
		throwError(err)
	}
}

func readPageContent(path string) string {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	content := ""
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		throwError(err)
		name := info.Name()
		fullPath := path
		if !info.IsDir() {
			if bytes.HasSuffix([]byte(name), []byte(".md")) {
				dat, err := ioutil.ReadFile(fullPath)
				if err == nil {
					var buf bytes.Buffer
					err = md.Convert(dat, &buf)
					if err == nil {
						content = buf.String()
					}
				}
			} else if bytes.HasSuffix([]byte(name), []byte(".html")) {
				dat, err := ioutil.ReadFile(fullPath)
				if err == nil {
					content = string(dat)
				}
			}
		}
		return nil
	})
	throwError(err)
	return content
}

func readInApis(path string, mode string) *CurlJSON {

	app := "transformers/build-" + mode
	arg1 := "--dir=" + path
	cmd := exec.Command(app, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// CurlJSON
	var birds CurlJSON
	err = json.Unmarshal(stdout, &birds)
	if err != nil {
		throwError(err)
		return nil
	}
	return &birds

}

func copydir(source, destination string) error {
	var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})
	return err
}

func clearGenratedFiles(path string, prefix string) {
	files, err := filepath.Glob(path + "/" + prefix + "*")
	throwError(err)
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			throwError(err)
		}
	}

}
