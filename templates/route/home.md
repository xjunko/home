@title=home
@description=the main page.
@tags=home
@route=/index.html

{{ $LatestNote := index .Notes 0 }}

<div class="grid-justify">
    <div>
        <h2> hello stranger </h2>
        <p>
            welcome to my silly little corner of the internet. i don't know how you got here, but i'm glad you did.
            you won't find anything of value here, but i hope you enjoy your stay. <br /><br />
            feel free to look around, there might some cool stuff you might like!
        </p>
    </div>
    <div id="kafu-home">
        <img src="static/imgs/kafu-peace.webp" height="100px"
            style="transform: scaleX(-1);scale: 1.5;position:relative;bottom: -5px;left:20px;pointer-events: none">
    </div>
</div>

<div class="grid">
    <div class="window">
        <div class="window-content">
            <h2> guestbook </h2>
            feeling a little bit chatty?
            wanted to say hi?
            maybe you just want to leave a message?
            a suggestion?
            a critique?
            <div style="text-align: right;position: relative;top: 30px;">
                <a class="href-right" href="https://xjunko.atabook.org/"> you can do that here </a>
            </div>
        </div>
    </div>
    <div class="window window-content blog" style="z-index: 2;">
        <h3> latest blog post </h3>
        <div>
            <h2> {{ index $LatestNote.Metadata "shorttitle" }} </h2><span> {{ $LatestNote.GetFormattedPostDate }}
            </span>
            <p
                style="height: 50px;overflow: hidden; display: -webkit-box; -webkit-box-orient: vertical; -webkit-line-clamp: 3;">
                {{- $LatestNote.GetPreviewRaw }}
            </p>
            <div style="text-align: right;position: relative;top: 2px;">
                <a class="href-right" href="/blog/{{ index $LatestNote.Metadata "slog" }}.html"> read </a>
            </div>
        </div>
    </div>
    <div class="window window-content">
        <h2> inspo </h2>
        this website wouldn't exists without these cool fellas:
        <div style="padding-top: 10px; text-align: center;">
            <a href="https://melankorin.net/"
                title="this version of the website was mostly based of kori's melankorin.net">melankorin.net</a>
            <a href="https://microsounds.github.io/"
                title="was what got me into making personal websites again.">sentimental microsounds</a>
            <a href="https://utsuho.rocks/"
                title="was the website i was inspired on when i was first starting.">utsuho.rocks</a>
        </div>
    </div>
    <div class="window window-content webrings blacken">
        <h2> webrings </h2>
        {{- template "route-template/common/webrings" . }}
    </div>
</div>