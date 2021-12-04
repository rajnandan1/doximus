package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nullrocks/identicon"
	"github.com/spf13/cobra"
)

var apiFiles = make(map[string]interface{})
var contentFiles = make(map[string]interface{})
var curlFiles = make(map[string]interface{})
var searchMap = []Search{}

func ThrowError(err error) {
	if err != nil {
		fmt.Print("❌")
	}

	cobra.CheckErr(err)
}
func ThrowSuccess(msg string) {
	fmt.Print("✅")
	fmt.Println(msg)
}
func ThrowLog(msg string) {
	fmt.Print("🔥")
	fmt.Println(msg)
}
func ThrowInput(msg string) {
	fmt.Print("⭕")
	fmt.Println(msg)
}

func Createdocs(pxth string, gen *identicon.Generator, lastFolder string, mode string) {
	knowFiles := map[string]bool{
		"api.yaml":     true,
		"content.md":   true,
		"content.html": true,
		"curls":        true,
		"main.yaml":    true,
	}
	files, err := filepath.Glob(pxth + "/*")
	ThrowError(err)
	for _, v := range files {
		c := strings.Replace(v, pxth+"/", "", 1)
		fullPath := pxth + "/" + c

		if !knowFiles[c] {
			if fileInfo, err := os.Stat(fullPath); err == nil {
				if fileInfo.IsDir() {
					Createdocs(fullPath, gen, lastFolder, mode)
				}
			}
		} else {
			ThrowLog("Processing -> " + fullPath)
			localID := Hash(fullPath)
			identifier := strings.Replace(pxth, lastFolder, "", 1)
			identifier = strings.Replace(identifier, "/", ".", -1)
			slug := ""
			parent := ""
			if strings.HasPrefix(identifier, ".") {
				identifier = identifier[1:]
			}
			psl := strings.Split(identifier, ".")
			psl[0] = "docs"
			if len(psl) == 2 {
				parent = "docs"
			} else if len(psl) == 1 {
				parent = ""
			} else {
				parent = strings.Join(psl[0:(len(psl)-1)], ".")
			}
			identifier = strings.Join(psl, ".")
			if !strings.Contains(identifier, ".") {
				slug = strings.Replace(identifier, "docs", "/", -1)
			} else {
				slug = strings.Replace(identifier, "docs.", "/docs/", -1)
			}

			if c == "main.yaml" {
				detail := ReadMainFile(fullPath, identifier, gen)

				if there := ItExists(pxth + "/curls"); there {
					detail.HasAPI = true
				}
				detail.Slug = slug
				detail.ID = localID
				apiFiles[identifier] = detail
				searchMap = append(searchMap, Search{
					Title:       detail.Name,
					Description: detail.Description,
					Slug:        slug,
					Logo:        detail.Logo,
					Parent:      parent,
					Type:        "Page",
				})
			}

			if c == "content.md" || c == "content.html" {
				contentFiles[identifier] = ReadPageContent(fullPath)
			}
			if c == "api.yaml" {
				detail := ReadInApis(fullPath, mode)
				if detail.Logo == "" {
					detail.Logo = RandImage("api"+identifier, gen, "doximus")
				}
				detail.Slug = slug + "#api"
				detail.ID = localID + "api"
				searchMap = append(searchMap, Search{
					Title:       detail.Name,
					Description: detail.Description,
					Slug:        detail.Slug,
					Logo:        detail.Logo,
					Parent:      identifier,
					Type:        "API Reference",
				})
				reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
				for k, v := range detail.Curls {
					detail.Curls[k].Slug = detail.Slug + strings.ToLower(reg.ReplaceAllString(v.Name, ""))
					detail.Curls[k].ID = detail.ID + strconv.Itoa(k)
					detail.Curls[k].Logo = RandImage("curlapi"+identifier+strings.ReplaceAll(v.Name+v.Method, " ", ""), gen, "doximus")
					searchMap = append(searchMap, Search{
						Title:       v.Name,
						Description: v.Description,
						Slug:        detail.Curls[k].Slug,
						Logo:        detail.Curls[k].Logo,
						Parent:      identifier,
						Type:        "API",
					})
					curlFiles[identifier] = detail
				}
			}
		}
	}
	return
}

func WriteDocs() {
	// fmt.Println(apiFiles)
	for k, v := range apiFiles {
		detail := v.(*DetailsJSON)

		split := strings.Split(k, ".")
		trimmedSlug := strings.Join(split[0:(len(split)-1)], ".")
		if trimmedSlug == "" {
			continue
		}
		if cur, ok := apiFiles[trimmedSlug]; ok {
			newObj := &DetailsJSON{
				Name:        detail.Name,
				Description: detail.Description,
				Logo:        detail.Logo,
				Slug:        detail.Slug,
				HasAPI:      detail.HasAPI,
			}
			apiFileObj := cur.(*DetailsJSON)
			apiFileObj.SubPages = append(apiFileObj.SubPages, *newObj)
			apiFiles[trimmedSlug] = apiFileObj
		}
	}
	WriteFiles(apiFiles, "json")
	WriteFiles(contentFiles, "content.txt")
	WriteFiles(curlFiles, "curl.json")
	file, err := json.MarshalIndent(searchMap, "", " ")
	ThrowError(err)
	err = ioutil.WriteFile("./site/files/search.json", file, 0644)
	ThrowError(err)
	err = Copydir("images", "site/assets/images")
	ThrowError(err)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
}

//ItemExists check for an item in a slice
func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
