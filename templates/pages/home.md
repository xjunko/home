@title=home
@description=the main page.
@tags=home
@route=/home.html
{{ define "home.content" }}
<br />

## about

my name is junko, you can find me online as xjunko. i have a strong interest in eletronics, arcade rhythm game, programming and vocaloids!
<br /><br />

this website serves as more or less dumping ground of useless information. here, you may find how to contact me, crap i've written and things that you might find interesting.
<br /><br />

currently, the domain is [kafu.ovh](https://kafu.ovh), i'm broke so the domain <i>might</i> change from time to time.
<br /><br />

## contact me

i go by the name, <a id="discord-name" class="blink-smooth">rmhakurei</a>, on discord, feel free to message me.
<br />
Currently <a id='discord-status-about' class="blink-smooth">offline</a>.
<br /><br />

if, for whatever reason, you need to get a hold of me quickly, contact thru email instead:
<div class="center widget-email blink">
    <a>
        phony [at] kafu [dot] ovh
    </a>
</div>
<br /><br />

## last note

{{ template "widget/home/note" . }}
<br /><br />

## recent post


{{ template "widget/home/post" . }}


<br /><br />

## this site

<div style="display: flex; float: right;flex-direction: column;">
    <a href="/">
        <img style="float: right;" src="/static/imgs/buttons/junko.png">
    </a>
    <i>
    hotlinking is allowed!
    </i>
</div>

<div>
    <textarea style="width: 60%; height: 50px;border: 1px solid var(--outline-color);background-color: var(--dark-color);color: var(--text-light);" readonly>
<a href="https://kafu.ovh" target="_blank">
    <img src="https://kafu.ovh/static/imgs/buttons/junko.png"/>
</a></textarea>
</div>
{{ end }}
{{ template "home.content" . }}
