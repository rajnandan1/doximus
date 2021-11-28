package cmd

//DetailsJSON is the main page holder
type DetailsJSON struct {
	Name        string        `json:"title" yaml:"title"`
	Description string        `json:"description" yaml:"description"`
	Logo        string        `json:"logo" yaml:"logo"`
	Slug        string        `json:"slug" yaml:"slug"`
	Tags        []string      `json:"tags" yaml:"tags"`
	SubPages    []DetailsJSON `json:"subpages" yaml:"subpages"`
	ID          string        `json:"id" yaml:"id"`
	Curls       []CurlJSON    `json:"curls" yaml:"curls"`
	HasAPI      bool          `json:"hasAPI" yaml:"hasAPI"`
}

//CurlJSON will store details of the api-reference
type CurlJSON struct {
	Name          string             `json:"title" yaml:"title"`
	Description   string             `json:"description" yaml:"description"`
	Logo          string             `json:"logo" yaml:"logo"`
	Slug          string             `json:"slug" yaml:"slug"`
	ID            string             `json:"id" yaml:"id"`
	Domains       []string           `json:"domains" yaml:"domains"`
	Configurables []APIConfigurables `json:"configurables" yaml:"configurables"`
	Curls         []CurlAPIJSON      `json:"apis" yaml:"apis"`
}

//CurlAPIJSON will store each api in api-reference
type CurlAPIJSON struct {
	Method      string        `json:"method" yaml:"method"`
	Name        string        `json:"title" yaml:"title"`
	Path        string        `json:"path" yaml:"path"`
	Slug        string        `json:"slug" yaml:"slug"`
	ID          string        `json:"id" yaml:"id"`
	Logo        string        `json:"logo" yaml:"logo"`
	Description string        `json:"description" yaml:"description"`
	Examples    []CurlExample `json:"examples" yaml:"examples"`
}

//CurlExample will store examples for each api
type CurlExample struct {
	Method string      `json:"method" yaml:"method"`
	Path   string      `json:"path" yaml:"path"`
	Title  string      `json:"title" yaml:"title"`
	Body   interface{} `json:"body" yaml:"body"`
	Header interface{} `json:"header" yaml:"header"`
	Query  interface{} `json:"query" yaml:"query"`
}

//Website will read site.yaml to store details about website
type Website struct {
	Title       string   `json:"title" yaml:"title"`
	Logo        string   `json:"logo" yaml:"logo"`
	Tags        []string `json:"tags" yaml:"tags"`
	Description string   `json:"description" yaml:"description"`
	Pages       string   `json:"-" yaml:"pages"`
}

//Search will contain list of searcable data
type Search struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Slug        string `json:"slug"`
	Parent      string `json:"parent"`
	Type        string `json:"type"`
}
type APIConfigurables struct {
	Title       string `json:"title" yaml:"description"`
	Description string `json:"description" yaml:"description"`
	Required    bool   `json:"required" yaml:"required"`
	Init        string `json:"init" yaml:"init"`
}
