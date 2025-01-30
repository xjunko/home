<img src="https://media1.tenor.com/m/UEE0HU83IQcAAAAC/zombieland-saga-junko-konno.gif"  height="200" align="right" style="float: right; margin: 0 10px 0 0;">

<h2 align="center">

junko's homepage

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white)](https://go.dev/)
</h2>

- heavy use of go's templating.
- written with modularity in mind, albeit very customized to my own use.

<br/>

## internal

### features
- able to resolve `youtube` and `spotify` media internally w/o tracking user data.
- templates everywhere, everything is templates.
### pages
- pages goes to `templates/pages/*.md` in markdown format
- entries goes to `entries/**/*.md`
	- supports `notes`,`channels`

### usage
### special token
- `@` is used for metadata info.

```go
var PREFIX = "@"
var PREFIXES = []string{
	// Page Basic Info
	"title",
	"description",
	"thumbnail",
	// Page Data
	"author",
	"date",
	"tags",
	"route",
	// /note/*
	"slog",
	// /channel/*
	"style",
	"outline",
	"outline-style",
	// Misc
	"exclude",
}
```
- is not case sensitive.
- example: `@TITLE=THE MOTHERFUCKIN TITLE`


## preview

![image](https://github.com/user-attachments/assets/adaf65d6-3edd-4c58-91ee-cc84a04351de)
![image](https://github.com/user-attachments/assets/6f2132e5-80c4-4b94-9c27-dc264d73f31a)
![image](https://github.com/user-attachments/assets/8da34082-dedc-4de5-8261-6a023ab6cfb7)
![image](https://github.com/user-attachments/assets/aac20b10-f996-4b4b-b067-bb221eae9f3e)


## credits
- spotify resolver is taken from [[l'm blog]](https://github.com/l1mey112/me.l-m.dev/blob/main/src/spotify/main.v)
