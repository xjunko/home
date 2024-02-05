import base64
import json
import re
from dataclasses import dataclass

import requests

# re
SPOTIFY_QUERY: re.Pattern = re.compile(
    r'<script\s+id="initial-state"\s+type="text/plain">([^<]+)</script>'
)


@dataclass
class SpotifyTrack:
    id: str
    artist_id: str
    name: str
    artist: str

    thumbnail: str
    audio: str


def get(url: str) -> SpotifyTrack | None:
    if (resp := requests.get(url)).status_code == 200:
        data = SPOTIFY_QUERY.findall(resp.content.decode())[0]

        json_decoded = json.loads(base64.b64decode(data.encode()))

        if key := list(json_decoded["entities"]["items"].keys())[0]:
            root_track = json_decoded["entities"]["items"][key]

            prev_url: str = ""

            if val := root_track["previews"]["audioPreviews"]["items"][0]:
                prev_url = val["url"]

            if len(root_track["firstArtist"]["items"]) <= 0:
                return None

            profile = root_track["firstArtist"]["items"][0]

            # thumbnail fuckery
            root_art = root_track["albumOfTrack"]["coverArt"]["sources"][0]

            for source in root_track["albumOfTrack"]["coverArt"]["sources"]:
                size = source["width"] * source["height"]
                root_size = root_art["width"] * root_art["height"]

                if size > root_size:
                    root_art = source

            return SpotifyTrack(
                id=root_track["id"],
                name=root_track["name"],
                artist=profile["profile"]["name"],
                artist_id=profile["id"],
                audio=prev_url,
                thumbnail=root_art["url"],
            )
