<img src="https://media1.tenor.com/m/UEE0HU83IQcAAAAC/zombieland-saga-junko-konno.gif"  height="200" align="right" style="float: right; margin: 0 10px 0 0;">

<h2 align="center">

junko's homepage

[![V](https://img.shields.io/badge/V-212adfa-blue.svg)](https://github.com/vlang/v)
[![Code style: v](https://img.shields.io/badge/code%20state-shit-purple.svg)](https://github.com/vlang/v)
</h2>

- heavy use of V's templating engine, beware of hacky workaround.
- specifically made for me, might not work for you.
- the website is shipped together with the source code, that's by design.

<br/>

## internal

### features
- able to resolve `youtube` and `spotify` media internally w/o tracking user data.
- templates everywhere, everything is mostly dynamic.
- super fucking fast because it's compiled.
- no need for a live webserver, a simple file server should do the job.
### pages
- markdown files from `src/magi/templates/pages/` are automatically generated into html files.
- for more control write a custom html with V's template.

### flow
```
Main -> Magi.resolve_pages -> Magi.resolve_channel 
                                      |
                                      v
                                Post.create
                                      |
                                      v
    Casper.postprocess   <-   Casper.preprocess
              |
              v
            Finish

```
### special token
- `@` is used for metadata info.

### `@`
```js
[
	// web info (for embeds)
	'title',
	'description',
	'thumbnail',
	// common
	'tags',
	'outer',
	'author',
	'priority',
	'route',
	// channel & blog
	'style',
	'outline',
	'outline-style',
]
```
- is not case sensitive.
- example: `@TITLE=THE MOTHERFUCKIN TITLE`


## preview
![image](https://github.com/xjunko/home/assets/44401509/400759eb-fdfa-476e-a327-75e112551907)
![image](https://github.com/xjunko/home/assets/44401509/4d209884-e286-4c5d-852a-8cfa729e4745)
![image](https://github.com/xjunko/home/assets/44401509/8109f457-588d-4ff8-8fee-2fea64883eaa)
![image](https://github.com/xjunko/home/assets/44401509/c679877e-6fde-447f-a16e-29e8828a36fa)


## credits
- spotify resolver is taken from [[l'm blog]](https://github.com/l1mey112/me.l-m.dev/blob/main/src/spotify/main.v)

