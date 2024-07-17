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
![image](https://github.com/user-attachments/assets/c8d72a0a-4f7f-4323-9efc-89028b800b0f)
![image](https://github.com/user-attachments/assets/f64ee4fb-4cc1-46e7-bef7-eae644219062)
![image](https://github.com/user-attachments/assets/524b4bb3-4889-44fb-9d7d-0e875653b662)
![image](https://github.com/user-attachments/assets/ca2076dd-63e8-4a1b-8300-e0d44ab35f28)

<a style="text-align: center">
<img src="https://github.com/user-attachments/assets/f1dc39f4-41c7-4f28-98fe-abddf0c9a237" height="500px">
<img src="https://github.com/user-attachments/assets/efccb482-c24b-4a74-94d0-0b25bfd6ffc3" height="500px">
<img src="https://github.com/user-attachments/assets/7b405eeb-ca3a-4f9c-a51d-85ad76a20f04" height="500px">
<img src="https://github.com/user-attachments/assets/84b05d90-033e-4ff9-828a-8ee197a75e5b" height="500px">
</a>


## credits
- spotify resolver is taken from [[l'm blog]](https://github.com/l1mey112/me.l-m.dev/blob/main/src/spotify/main.v)

