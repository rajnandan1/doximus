package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	Templates    *template.Template
	AssetVersion string
	AssetURL     string
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.Templates.ExecuteTemplate(w, name, data)
}

func intHeaders(rh map[string]string) map[string][]string {
	headers := map[string][]string{}
	//headers[models.RequestID] = []string{reqHeaders.Client.XRequestID}
	for k, v := range rh {
		headers[k] = []string{v}
	}
	return headers
}
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

type CallAPIReq struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Method  string            `json:"method"`
	Body    interface{}       `json:"body"`
}
type CallAPIRes struct {
	Code    int               `json:"statusCode"`
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}
type Client struct {
	*http.Client
}

func sendRequest(client *http.Client, method string, endpoint string, headers map[string]string, reqBody interface{}) (string, map[string]string, int) {

	jsonData, err := json.Marshal(reqBody)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
	req.Header = intHeaders(headers)
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	responseHeader := make(map[string]string)
	for k, v := range response.Header {
		responseHeader[strings.ToLower(k)] = v[0]

	}
	return string(body), responseHeader, response.StatusCode
}
func Site(c echo.Context) error {
	fmt.Println("c.Path()")
	fmt.Println(c.Request().RequestURI)
	name := c.Request().RequestURI
	name = strings.Replace(name, "/", "", -1)
	name = strings.Replace(name, "docs", "", -1)
	if name == "" {
		name = "docs.json"
	} else {
		name = "docs." + name + ".json"
	}
	dat, _ := ioutil.ReadFile("./site/files/" + name)
	var data = DetailsJSON{}
	json.Unmarshal(dat, &data)
	seoData := map[string]interface{}{
		"title":       data.Name,
		"description": data.Description,
		"image":       data.Logo,
		"tags":        strings.Join(data.Tags, ","),
	}
	return c.Render(http.StatusOK, "index.html", seoData)
}

func ResponseFile(c echo.Context) error {
	name := c.Param("name")
	dat, err := ioutil.ReadFile("./site/files/" + name)
	var data interface{}
	if err != nil {
		if strings.HasSuffix(name, ".json") {
			return c.JSON(200, nil)
		} else {
			return c.File("./site/files/" + name)
		}
	}

	if strings.HasSuffix(name, ".json") {
		json.Unmarshal(dat, &data)
		return c.JSON(200, data)
	} else {
		return c.File("./site/files/" + name)
	}

}

func Call(c echo.Context) error {
	u := new(CallAPIReq)
	if err := c.Bind(u); err != nil {
		return c.JSON(400, u)
	}
	cn := httpClient()
	responseBody, responseHeader, responseStatusCode := sendRequest(cn, u.Method, u.URL, u.Headers, u.Body)
	data := &CallAPIRes{}
	data.Body = responseBody
	data.Headers = responseHeader
	data.Code = responseStatusCode

	return c.JSON(http.StatusOK, data)
}
