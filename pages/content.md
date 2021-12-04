## Installation
- go 1.16 minimum

## Clone Repo
```
git clone https://github.com/rajnandan1/doximus.git
cd doximus
```
## Install Doximus
```
go mod tidy
go install doximus
```
## Build Pages
**MacOS**
```
doximus build macos 
```
**Linux**
```
doximus build linux 
```
**Windows**
```
doximus build win.exe
```
## Run Server
```
doximus serve --port=4099
```
## Limitations
- Only JSON requests
- Only Rest APIs
- GET, PUT, DELETE, POST

## Features
- Easy to Use
- Dynamic Path params
- Updates all requests with same variable name. That is if an api returns suppose `user_id` it will be replaced in all subsequent requests wherever present
- Google searchable
- Site search
- Supports examples 

## Images
Use images folder to add your images.
## Pages
All pages are present in pages folder. Root of `pages` is the home / landing page
## Doximus CLI
Doximus provides a useful cli commands to help you create your pages. You can also directly edit pages folder.

---
## Set up Layout
Add logo for site

```
doximus site --logo="https://icons-for-free.com/iconfiles/png/512/gollum+lord+of+the+rings+smeagol+icon-1320166691516135718.png"
```
or put image in `images` folder. Example myimage.png
```
doximus site --logo="/images/myimage.png"
```
Add Description for site
```
doximus site --desc="A wizard is never late, Frodo Baggins. Nor is he early. He arrives precisely when he means to."
```
Add title for site
```
doximus site --title="Middle Earth APIs"
```
Add tags for site
```
doximus site --tags=frodo,bilbo,sam,pippin,mary
```
---
## Modify Home Page
Add image. If not specified it will be auto generated
```
doximus add page --logo="/images/gandalf.png"
```
Add a small paragraph for site
```
doximus add page --desc="May it be your light in the darkness; when all other lights go out."
```
Add heading
```
doximus add page --title="The Shire"
```
Add tags
```
doximus add page --tags=gandalf,aragon,legolas
```
Add content
```
doximus add content --mode=new --type=html
```
- `mode` can be new or append 
- `type` can be html or md

---
## View pages
```
doximus pages
```
## Add a new page
To create a new page with a specified id
```
doximus add page --id=lotr
```
**To add a subpage** 
```
doximus add page --id="lotr.part1"
```
add logo
```
doximus add page --id=lotr --logo="images/legolas.png"
```
add description
```
doximus add page --id=lotr --desc="Smoke rises from the Mountain of Doom, the hour grows late, and Gandalf the Grey  rides to Isengard seeking my counsel. For that is why you have come, is it not?"
```
add title
```
doximus add page --id=lotr --title="Lord of the Rings"
```
add tags
```
doximus add page --id=lotr --tags=your,tags
```
add content
```
doximus add content --id=lotr --mode=new --type=html
```
---
## Add apis to a page
add title of the api
```
doximus add api --id=lotr --title="The API"
```
add description of the api
```
doximus add api --id=lotr --desc="Gandalf, my old friend, this will be a night to remember."
```
add image for the api
```
doximus add api --id=lotr --logo="https://img.icons8.com/color/96/000000/frodo.png"
```
add domains for api
```
doximus add api --id=lotr --domains=the-one-api.dev,the-one-api.herokuapp.com
```
add a curl for the api
```
doximus add curl --id=lotr --title="book by id" --desc="use this api to get all the books with ids"
```
Here you will be prompted to copy paste a curl

---
## Add api variables
variables becomes input and will be applied in whole api
```
doximus add variable --id=lotr --title="Authorization" --desc="this is api key" --required=false --value=""
```

