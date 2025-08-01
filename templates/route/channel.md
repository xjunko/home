@title=channel
@description=like 4chan but just me
@tags=channel
@route=/channel.html

<h1 style="text-align: center;">/b/ - Random</h1> <br>
<a class="href href-left" style="position: absolute; top: 48px;" href="/index.html">return</a>

<hr>
<br> 
<br>

<div class="window-content">
    file: <a href="https://hatsune-miku.has.rocks/r/lain-thumbnail.webp">index.gif</a> (file/webp-animated)
    <div>
        <img class="post-image" style="padding-top: 4px;" src="https://hatsune-miku.has.rocks/r/lain-thumbnail.webp">
    </div>
    <br>
    <div style="position: relative; bottom: 10px;">
        <h4 class="title">main thread</h4>
        <h4 class="name">junko</h3>
        <span class="date">1/1/1970 12:00:00</span>
        <span class="id">No: 1</span>

        this is where i post everything and anything, it's mostly random rambling and thoughts. <br>
        please do excuse the swearings and other not-so-offensive stuff i _might_ say here.

    </div>
    <br> 
</div>
<br> <br> <br> <br> <br>
<div class="grid-justify">
    {{ range $index, $currentPost := .Channels }}
    <div class="window" id="{{ $currentPost.ID }}">
        <div class="window-content">
            <h3 class="name-small">{{ index $currentPost.Metadata "author"}}</h3>
            <span class="date">{{ $currentPost.GetFormattedPostDate }}</span>
            <span class="id">No: <a href="#{{ $currentPost.ID }}">{{ $currentPost.ID }}</a> </span>
            {{ if gt (len (index .Metadata "thumbnail")) 0 }}
            <br>
            file: <a href="{{ index .Metadata "thumbnail" }}">{{ index .Metadata "filename" }}</a>
            (file/{{ index .Metadata "mimetype" }})
            {{ if eq (index .Metadata "thumbnail-type") "video" }}
            <video muted loop controls preload=metadata class="post-video" 
                src="{{ index .Metadata "thumbnail" }}"></video>
            {{ else }}
            <img class="post-image" loading=lazy src="{{ index .Metadata "thumbnail" }}">
            {{end}}
            {{ end }}
            {{ $currentPost.ToMarkdown }}
        </div>
    </div>
    {{ end }}
</div>