### Installation
- go 1.16 minimum

#### Clone Repo
```
git clone https://github.com/rajnandan1/doximus.git
cd doximus
```
#### Install Doximus
```
go mod tidy
go install doximus
```
#### Start Doximus
```
doximus build macos #macos, linux, win.exe
doximus serve 5000
```

### How to start
Name your site
```
doximus title "My awesome API"
```

Describe your site
```
doximus desc "My awesome site description"
```

Add your site logo
```
doximus logo "https://logo.clearbit.com/example.com.au"
```
Add your site tags for seo purposes
```
doximus tags "awesome" "api docs"
```
Add path to folder where your pages are
```
doximus pages pages
or
doximus pages /your/absolute/path/pages
```
## How to create a page
Each page of doximus must have the following files

### Image Folder
You will add your images in this folder which will be picked up during building of the pages. Your pages can refer this images using `/images/myimage.png`
Additionally you can add a favicon.ico 
```
images/
	yourimage.png //these are logos that you might need. You can use them in your pages
	someimage.png
	logo.png //this is your site logo. This is predefined
	favicon.ico //this is your favicon. This is predefined
```
### Pages folder
This is the folder tha will have your apis and pages. The root of this folder becomes your home page and other subfolders become your subpages of the home page.
Each folder along with the root folder has to have a main.yaml file 
If you want multiple depth then you can use `.` to assign parent for subpage
##### Example of **one** level site

```
pages/
    main.yaml
    bilbo/
        main.yaml
    sam/
        main.yaml
    frodo/
        main.yaml
```
##### Example of **two** level site
```
pages/
    main.yaml
    bilbo/
        main.yaml
    bilbo.sam/
        main.yaml
    bilbo.frodo/
        main.yaml
```
##### Example of **three** level site
```
pages/
    main.yaml
    bilbo/
        main.yaml
    bilbo.sam/
        main.yaml
    bilbo.sam.frodo/
        main.yaml
```
#### Allowed files of a Page
- main.yaml (`required`)
- api.yaml 
- content.md or content.html
- curls/  (`required` if api.yaml is present)
    - example.yaml
    - example2.yaml
#### main.yaml `Required`

| name      | detail           | example                                                                                                                 | required |
|-------------|------------------|-------------------------------------------------------------------------------------------------------------------------|----------|
| title       | page title       | My Page                                                                                                                 | yes      |
| description | page description | A developer friendly way to host your apis                                                                              | no       |
| logo        | logo of page     | add url of image or place image in images folder and add reference. example images/mylogo.png or http://ex.com/logo.png | no       |
| tags        | tags for page    | myapi, my api, awesome api.                                                                                             | no       |
#### api.yaml  
| *name*        | detail                       | example                                                                                                                 | required |
|---------------|------------------------------|-------------------------------------------------------------------------------------------------------------------------|----------|
| title         | page title `string`          | My Page                                                                                                                 | yes      |
| description   | description of api `string`  | A developer friendly way to host your apis                                                                              | no       |
| logo          | logo of page `string`        | add url of image or place image in images folder and add reference. example images/mylogo.png or http://ex.com/logo.png | no       |
| domains       | domains for api. `array`     | - api.example.com <br /> - sandbox.example.com                                                                          | no       |
| apis          | detail for each api `object` | example below                                                                                                           | yes      |
| configurables | configurable params `array`  | example below                                                                                                           | no       |
##### apis

```
apis: 
  -
    title: "Create Order"
    description: Irure labore ipsum dolore excepteur aliqua ea sit sint excepteur ut commodo consequat.
    path: "/orders"
    method: "POST"
  -
    title: "Get Order"
    description: Lorem esse veniam enim voluptate cillum esse commodo culpa aute.
    path: "/orders"
    method: "GET"
  -
    title: "Create Refund"
    description: Minim deserunt ea nisi nisi eu officia officia ea occaecat irure fugiat Lorem aute quis.
    path: "/refunds"
    method: "POST"
  -
    title: "GET Refund by id"
    description: Id sint deserunt minim minim laborum elit ad cupidatat sunt proident irure reprehenderit laborum.
    path: "/refunds/:refund_id"
    method: "GET"
```
##### configurables
```
configurables:
  - 
    title: x-client-id
    description: your client id
    required: true
    init: enter a value
  -
    title: x-client-secret
    description: your secret key
    required: true
    init: enter a value
```