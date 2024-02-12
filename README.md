<img src="https://media1.tenor.com/m/UEE0HU83IQcAAAAC/zombieland-saga-junko-konno.gif"  height="200" align="right" style="float: right; margin: 0 10px 0 0;">

<h2 align="center">

Junko's homepage

[![V](https://img.shields.io/badge/V-212adfa-blue.svg)](https://github.com/vlang/v)
[![Code style: black](https://img.shields.io/badge/code%20style-Default-blue.svg)](https://github.com/vlang/v)
</h2>

- Piece of shit website generator that has 50% chance of exploding by itself.
- Specifically made for me, might not work for you.

<br/>

## Internal
### Pages
- Markdown files from `src/magi/templates/pages/` are automatically generated into html files.
- For more control write a custom html with V's template.
### Special Token
- `@` is used for metadata info.

### `@`
```js
[
	// Common
	'title',
	'tags',
	'outer',
	'author',
	'priority',
	// Blog
	'style',
	'outline',
	'outline-style',
	'thumbnail',
]
```
- is not case sensitive.
- Example: `@TITLE=THE MOTHERFUCKIN TITLE`


## Preview
![image](https://github.com/xjunko/home/assets/44401509/a5dc648e-3a43-4678-a24c-96dc5915015d)
![image](https://github.com/xjunko/home/assets/44401509/c8eb5108-54e1-4449-a1a6-d51c7773e610)
![image](https://github.com/xjunko/home/assets/44401509/cde37e5b-f9f0-49e5-a930-77a5ef53d867)
