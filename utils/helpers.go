package utils

import (
	"bytes"
	"encoding/json"
	"hash/fnv"
	"image"
	"image/color"
	"io/ioutil"
	"net/http"
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

func ItExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

func TransparentBg() func(cb []byte, fc color.Color) color.Color {
	return func(cb []byte, fc color.Color) color.Color {
		return color.Transparent
	}
}
func ReadMainFile(fullPath string, pageID string, gen *identicon.Generator) *DetailsJSON {
	detail := DetailsJSON{}
	dat, err := ioutil.ReadFile(fullPath)
	ThrowError(err)
	err = yaml.Unmarshal(dat, &detail)
	ThrowError(err)

	if detail.Logo == "" {
		detail.Logo = RandImage("main.yaml"+pageID, gen, "doximus")
	}
	return &detail
}

func RandImage(s string, gen *identicon.Generator, prefix string) string {
	num := Hash(s)
	s = prefix + num
	var ii *identicon.IdentIcon
	var err error
	ii, err = gen.Draw(s)
	ThrowError(err)
	img, err := os.Create("./site/assets/images/" + s + ".png")
	ThrowError(err)
	defer img.Close()
	ii.Png(90, img)

	return "/images/" + s + ".png"
}

func WriteFiles(data map[string]interface{}, tt string) {
	for k, v := range data {
		file, err := json.MarshalIndent(v, "", " ")
		ThrowError(err)
		err = ioutil.WriteFile("./site/files/"+k+"."+tt, file, 0644)
		ThrowError(err)
	}
}

func ReadPageContent(path string) string {
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
		ThrowError(err)
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
	ThrowError(err)
	return content
}

func ReadInApis(path string, mode string) *CurlJSON {

	app := "transformers/build-" + mode
	arg1 := "--dir=" + path
	cmd := exec.Command(app, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		ThrowError(err)
		return nil
	}
	// CurlJSON
	var birds CurlJSON
	err = json.Unmarshal(stdout, &birds)
	if err != nil {
		ThrowError(err)
		return nil
	}
	return &birds

}

func Copydir(source, destination string) error {
	var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0777)
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

func ClearGenratedFiles(path string, prefix string) {
	files, err := filepath.Glob(path + "/" + prefix + "*")
	ThrowError(err)
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			ThrowError(err)
		}
	}

}

func CreateDir(path string) {
	err := os.Mkdir(path, 0777)
	ThrowError(err)
}
func CreateFile(path string) {
	err := ioutil.WriteFile(path, []byte(""), 0777)
	ThrowError(err)
}

func GetImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func GetFileContentType(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
