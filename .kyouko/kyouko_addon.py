""" addon.py - for kyouko, mostly blogs stuff. """


import base64
import hashlib
import json
import re
from datetime import datetime
from pathlib import Path
from typing import Any, Callable

import requests
from url_normalize import url_normalize  # type: ignore
from utils import spotify

_original_get = requests.get


def get_if_not_in_cache_else_cache(url: str) -> dict[str, Any]:
    cache_file: Path = Path.cwd() / ".cache"

    if not cache_file.exists():
        cache_file.write_text(json.dumps({}))

    # Format, {status, content}
    CACHE_DATA = json.loads(cache_file.read_text())

    if url not in CACHE_DATA:
        resp = _original_get(url)
        CACHE_DATA[url] = {
            "status": resp.status_code,
            "content": base64.b64encode(resp.content).decode("ascii"),
        }
        cache_file.write_text(json.dumps(CACHE_DATA))

    return {
        "status": CACHE_DATA[url]["status"],
        "content": base64.b64decode(CACHE_DATA[url]["content"].encode("ascii")),
    }


requests.get = get_if_not_in_cache_else_cache


def get_youtube_embed(video_id: str) -> str | None:
    for quality in ["maxresdefault.jpg", "mqdefault.jpg", "0.jpg"]:
        if (resp := requests.get(f"https://i3.ytimg.com/vi/{video_id}/{quality}"))[  # type: ignore
            "status"
        ] != 404:
            return quality

    return "0.jpg"


def get_hashed_filename_from_url(url: str) -> str:
    return (
        hashlib.md5(url_normalize(url).encode()).hexdigest()  # type: ignore
        + "."
        + get_mime_from_url(url)
    )


def get_mime_from_url(url: str) -> str:
    return url.split("/")[-1].split(".", 2)[-1]


def preprocess_blog_line(content: str) -> str:
    MEDIA_REGEX: re.Pattern = re.compile(  # type: ignore
        r"https?://\S+\.(?:(png)|(jpe?g)|(gif)|(svg)|(webp)|(mp4)|(webm)|(mov))"
    )

    YOUTUBE_REGEX: re.Pattern = re.compile(  # type: ignore
        r"https?://(?:www\.)?youtu(?:be\.com/watch\?v=|\.be/)(\S+)"
    )
    SPOTIFY_REGEX: re.Pattern = re.compile(r"https?://open\.spotify\.com/track/(\w+)")  # type: ignore
    GREENTEXT_REGEX: re.Pattern = re.compile(r"\\>.*")  # type: ignore
    POST_REFERENCE_REGEX: re.Pattern = re.compile(r"\\>>.*")  # type: ignore

    def _media_embed(match: re.Match) -> str | None:  # type: ignore
        if url := match.group(0):  # type: ignore
            if url.endswith(("mp4", "webm", "mov")):  # type: ignore
                return f'<video loop controls preload=metadata width="100%" height="auto" src="{url}"></video>'

            return url  # type: ignore

    content = MEDIA_REGEX.sub(_media_embed, content)  # type: ignore

    def _youtube_embed(match: re.Match) -> str | None:  # type: ignore
        video_url = match.string  # type: ignore
        video_id = match.group(1)  # type: ignore

        if not video_id:
            video_id = video_url.split("v=")[-1]  # type: ignore

        if video_thumbnail := get_youtube_embed(video_id):  # type: ignore
            return f"""
<div class="youtube">
    <a href="{video_url}" target="_blank">
        <img class="youtube-thumbnail" loading="lazy" src="https://i3.ytimg.com/vi/{video_id}/{video_thumbnail}" 
            alt="Youtube Thumbnail">
        <a class="youtube-info" href="{video_url}" target="_blank"> Click to watch in YouTube. </a>
    </a>
</div>
"""

    content = YOUTUBE_REGEX.sub(_youtube_embed, content)  # type: ignore

    def _replace_spotify_embed(match: re.Match) -> str:  # type: ignore
        if track := spotify.get(match.string):  # type: ignore
            return f"""
<div class="spotify">
    <div class="spotify-thumbnail">
        <img loading=lazy height="100px" width="100px" src="{track.thumbnail}" alt="Cover Art">
    </div>
    <div class="spotify-info">
        <a class="links" href="https://open.spotify.com/track/{track.id}">{track.name}</a>
        <br>
        <a class="links" href="https://open.spotify.com/artist/{track.artist_id}">{track.artist}</a>
        <br>
        <audio controls preload=none>
            <source src="{track.audio}" type="audio/mpeg">
        </audio>
    </div>
</div>
"""

    content = SPOTIFY_REGEX.sub(_replace_spotify_embed, content)  # type: ignore

    def _replace_reference(match: re.Match) -> str:  # type: ignore
        if post_reference := match.group(0):  # type: ignore
            raw_text: str = post_reference.split(r"\>>", 1)[-1].strip()  # type: ignore
            # return f'<a style="color: var(--reference-text)">\>{raw_text}</a>'
            return f'<a style="color: var(--reference-text); text-decoration: underline" href="#{raw_text}">>>{raw_text}</a>'

    content = POST_REFERENCE_REGEX.sub(_replace_reference, content)  # type: ignore

    def _replace_green_text(match: re.Match) -> str:  # type: ignore
        if match.string.strip().startswith("\\"):  # type: ignore
            raw_text: str = match.string.split(r"\>", 1)[-1]  # type: ignore
            return f'<a style="color: var(--green-text)">\>{raw_text}</a>'  # type: ignore

    content = GREENTEXT_REGEX.sub(_replace_green_text, content)  # type: ignore

    return content


def process_blog(id: int, file: Path, markdown_callback: Callable = lambda: ...) -> str:  # type: ignore
    POST_ID: int = id + 1
    POST_RAW: str = file.read_text()
    POST_CONTENT: str = ""
    METADATA: dict[str, str] = {"style": ""}

    for line in POST_RAW.splitlines():
        if line.startswith("[]#"):
            key, value = line.removeprefix("[]#").split(":", 1)
            METADATA |= {key.strip().casefold(): value.strip()}  # type: ignore
        else:
            POST_CONTENT += preprocess_blog_line(line) + "\n"  # type: ignore

    # Custom CSS
    if "outline" in METADATA:
        METADATA[
            "style"
        ] += f";border: {METADATA.get('outline-style', 'solid')} 2px {METADATA['outline']}"  # type: ignore

        print(METADATA, POST_CONTENT)

    # Export
    POST_BOILERPLATE: dict[str, Any] = {
        "ID": POST_ID,
        "AUTHOR": METADATA.get("author", "junko"),
        "DATE": datetime.fromtimestamp(int(file.stem)),
        "CSS": METADATA.get("style", ""),
    }

    HTML_OUTPUT: str = '<div class="blog-post" id="{ID}" style="{CSS}" >'

    HTML_OUTPUT += '<div class="blog-content">'  # type: ignore
    HTML_OUTPUT += '<span class="blog-author">{AUTHOR}</span> {DATE} '  # type: ignore
    HTML_OUTPUT += '<a class="blog-url no-underline" href="#{ID}">#{ID}</a> <br/>'  # type: ignore

    if thumbnail_url := METADATA.get("thumbnail"):
        thumbnail_filename: str = get_hashed_filename_from_url(thumbnail_url)
        thumbnail_mime: str = get_mime_from_url(thumbnail_url)

        HTML_OUTPUT += f'file: <a class="blog-url" href="{thumbnail_url}" target="_blank" rel="noopener noreffer">{thumbnail_filename}</a> [file/{thumbnail_mime}] <br/>'  # type: ignore
        HTML_OUTPUT += (  # type: ignore
            f'<a href="{thumbnail_url}" target="_blank" rel="noopener noreferrer">'
            f'<img class="blog-media" src="{thumbnail_url}">'
            "</a>"
        )

    HTML_OUTPUT += markdown_callback(POST_CONTENT) + "\n"  # type: ignore

    HTML_OUTPUT += "</div>"  # type: ignore

    HTML_OUTPUT += "</div>"  # type: ignore

    # Resolve all boilerplate
    for key, value in POST_BOILERPLATE.items():
        HTML_OUTPUT = HTML_OUTPUT.replace("{" + key + "}", str(value))  # type: ignore

    return HTML_OUTPUT + "\n"  # type: ignore
