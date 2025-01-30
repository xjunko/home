let music_queue_info = document.querySelector("#music-queue");
let music_track_duration = document.querySelector("#music-duration");
let music_track_title = document.querySelector("#music-title");


let music_ctrl_play_toggle = document.querySelector("#music-play");
let music_ctrl_backward = document.querySelector("#music-backward");
let music_ctrl_forward = document.querySelector("#music-forward");

let music_volume = document.querySelector("#music-volume");

let music_track_index = 0;
let music_is_playing = false;

let _update_timer;
let _current_track = document.getElementById("music");

let track_list = [
    {
        name: "フォニイ",
        artist: "ツミキ",
        path: "https://hatsune-miku.has.rocks/r/phony.mp3"
    },
    {
        name: "きゅうくらりん",
        artist: "いよわ",
        path: "https://hatsune-miku.has.rocks/r/kyu.mp3"
    },
    {
        name: "エス",
        artist: "内緒のピアス",
        path: "https://hatsune-miku.has.rocks/r/esu.mp3"
    },
    {
        name: "UFO(10th anniv.)",
        artist: "青屋夏生",
        path: "https://hatsune-miku.has.rocks/r/ufo.mp3"
    },
    {
        name: 'プシ',
        artist: "r-906",
        path: "https://hatsune-miku.has.rocks/r/pushi.mp3"
    },
    {
        name: "腐れ外道とチョコレゐト",
        artist: "PinocchioP",
        path: "https://hatsune-miku.has.rocks/r/chocolate.mp3"
    },
    {
        name: "ゾイトロープ",
        artist: "youまん",
        path: "https://hatsune-miku.has.rocks/r/zoetrope.mp3"
    },
    {
        name: "マージナルソウル",
        artist: "youまん",
        path: "https://hatsune-miku.has.rocks/r/marginal.mp3"
    },
    {
        name: "パノプティコン",
        artist: "r-906",
        path: "https://hatsune-miku.has.rocks/r/panopticon.mp3",
    },
    {
        name: "メモリア",
        artist: "Aira",
        path: "https://hatsune-miku.has.rocks/r/memoria.mp3"
    },
    {
        name: 'RBF SYNDROME',
        artist: "omuomu",
        path: "https://hatsune-miku.has.rocks/r/rbfsyndrome.mp3"
    },
    {
        name: "夏 O 幻",
        artist: "死んだ眼球",
        path: "https://hatsune-miku.has.rocks/r/summer.mp3"
    }

];

function loadTrack(track_index) {
    clearInterval(_update_timer);
    resetValues();

    _current_track.src = track_list[track_index].path;
    _current_track.load();

    music_track_title.textContent = track_list[track_index].artist + " - " + track_list[track_index].name;
    music_queue_info.textContent = "Track " + (track_index + 1) + "/" + track_list.length;

    _update_timer = setInterval(seekUpdate, 1000);

    _current_track.addEventListener("ended", nextTrack);

}

function resetValues() {
    music_track_duration.textContent = "0:00 / 0:00";
}

loadTrack(music_track_index);
setVolume();

function playpauseTrack() {
    if (!music_is_playing) playTrack();
    else pauseTrack();
}

function playTrack() {
    _current_track.play();
    music_is_playing = true;

    music_ctrl_play_toggle.innerHTML = '<i class="player-icons">&#xe803;</i>';
}

function pauseTrack() {
    _current_track.pause();
    music_is_playing = false;

    music_ctrl_play_toggle.innerHTML = '<i class="player-icons">&#xe800;</i>';
}

function nextTrack() {
    if (music_track_index < track_list.length - 1)
        music_track_index += 1;
    else music_track_index = 0;
    loadTrack(music_track_index);
    playTrack();
}

function prevTrack() {
    if (music_track_index > 0)
        music_track_index -= 1;
    else music_track_index = track_list.length;
    loadTrack(music_track_index);
    playTrack();
}


function setVolume() {
    _current_track.volume = music_volume.value / 100;
}

function seekUpdate() {
    if (!isNaN(_current_track.duration)) {
        seekPosition = _current_track.currentTime * (100 / _current_track.duration);

        let currentMinutes = Math.floor(_current_track.currentTime / 60);
        let currentSeconds = Math.floor(_current_track.currentTime - currentMinutes * 60);
        let durationMinutes = Math.floor(_current_track.duration / 60);
        let durationSeconds = Math.floor(_current_track.duration - durationMinutes * 60);

        if (currentSeconds < 10) { currentSeconds = "0" + currentSeconds; }
        if (durationSeconds < 10) { durationSeconds = "0" + durationSeconds; }
        if (currentMinutes < 10) { currentMinutes = currentMinutes; }
        if (durationMinutes < 10) { durationMinutes = durationMinutes; }

        music_track_duration.textContent = currentMinutes + ":" + currentSeconds + " / " + durationMinutes + ":" + durationSeconds;;
    }
}
