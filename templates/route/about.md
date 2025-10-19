@title=about
@description=about me.
@tags=about
@route=/about.html

<div class="grid-justify">
    <div>
        <h2 class="edgy" style="margin: 0"> biography </h2>
        <div class="window">
            <div class="window-content" style="display: flex; padding-bottom: 0;">
                <a href="https://www.youtube.com/watch?v=DLkNQgh4Ons" style="margin: 0;padding: 0;">
                    <img style="display: inline" src="static/imgs/about/pfp.png" height="110px" />
                </a>
                <p style="margin-left: 10px; margin-top: 0;">
                    Wow, look who decided to show up, -is what I would say, but I doubt anyone would read this.
                    Anyways, you can call me Junko, I'm the one who made this website. I'm currently a student that's
                    been making
                    whatever I feel like making. I'm not really good at anything, but I'm trying my best to improve.
                    <br> <br>
                    Also, Isn't it funny that you can just be anyone on the internet?
                </p>
            </div>
            <div class="window-content" style="margin: 0">
                I’ve been programming since I was a kid - probably not the healthiest hobby growing up, but hey, here I am years later, still doing it.
                Back then, I spent most of my time making absolutely awful games in 
                <a href="https://www.clickteam.com/clickteam-fusion-2-5">Clickteam Fusion</a>. 
                Eventually, I moved on to <a href="https://wordpress.com/">WordPress</a>, where my proudest creation was a clone of osu!'s 
                <a href="https://circle-people.com/skins/">Circle People</a> skins website. 
                I don’t have it anymore (and honestly, that might be for the best), but it was fun while it lasted.
                <br/><br/>
                Sometime around late 2018, I started learning <a href="https://www.python.org/">Python</a> - mainly because I wanted to make my own version of the osu! <a href="https://top.gg/bot/289066747443675143">owo</a> bot. 
                If you can’t tell by now, I have a habit of making clones of things I like. It’s honestly one of the best ways to learn.
                <br/><br/>
                These days, I’ve worked with quite a few programming languages. But back then, I just couldn’t get into 
                <a href="https://www.rust-lang.org/">Rust</a> (no hate) - I knew it was good, I just didn’t see myself ever using it. I even said maybe I’d change my mind someday.
                <br/><br/>
                Well… that day finally came. I get it now. <a href="https://www.rust-lang.org/">Rust</a> has completely won me over, and it’s easily become one of my favorite programming languages.
            </div>
        </div>
        <br>
        <h2 class="edgy" style="margin: 0"> contact </h2>
        <div class="window">
            <div class="window-content">
                I don't really use social media that much, so expect a delay in my response. <br>
                I'm on these platforms:
                <ul class="list-no-space" style="padding-top: 5px;">
                    <li>discord - rmhakurei</li>
                    <li>last.fm - <a href="https://www.last.fm/user/FireRedz">here</a></li>
                    <li>github - <a href="https://github.com/xjunko">xjunko</a></li>
                    <li>email - phony (at) kafu (dot) ovh</li>
                </ul>
            </div>
        </div>
        <br>

        <div class="grid">
            <div class="border">
                <h2 class="edgy" style="margin: 0"> neighbours </h2>
                <div class="window-content" style="text-align: center;">
                    {{ template "route-template/about/buttons" . }}
                </div>
            </div>
            <div class="border">
                <div class="window-content">
                    <a href="https://dokodemo.neocities.org/">
                        <img style="width: 100%" src="https://dokodemo.neocities.org/gamelounge/pokepi/suzuki.png">
                    </a>
                    <a href="https://www.theotaku.com/quizzes/view/2019/what_okami_character_are_you%3F">
                        <img style="width: 100%" border="0" src="http://www.theotaku.com/guru_results/2019_Waka.jpg" alt="What Okami Character Are You?" />
                        <br />
                    </a>
                    <div style="display: flex;align-items: center;justify-content: center;gap: 5px;">
                        {{- template "route-template/about/nerd-stats" . }}
                        {{- template "route-template/about/moon-phase" . }}
                    </div>
                </div>
            </div>
        </div>
        <br>

        <h2 class="edgy" style="margin: 0"> built with </h2>
        <div class="window">
            <div class="window-content">
                <p>
                    This website is mainly handwritten in <a href="https://vscodium.com/">VSCodium</a>, it then gets
                    built with <a href="https://github.com/xjunko/home">Eva</a>; my own static site generator, written
                    in <a href="https://go.dev/">Go</a>.

                    ### fonts
                    - [Ark Pixel](https://github.com/TakWolf/ark-pixel-font/releases) by
                    [@TakWolf](https://github.com/TakWolf)
                    - [Neue Haas Grotesk](https://www.myfonts.com/pages/linotype-neue-haas-grotesk) by Linotype
                    - [Terminus](https://terminus-font.sourceforge.net/) by Kevin Dresser (?)

                    ### arts
                    - [Header Kafu](https://cevio.fandom.com/wiki/KAFU) by PALOW.
                    - [Home Kafu](https://www.pixiv.net/en/artworks/117194838) by
                    [落島](https://www.pixiv.net/en/users/93735034)
                    - [Ononoki Index](https://dowman-s.tumblr.com/post/145993340985) by [dowman-s](https://dowman-s.tumblr.com/)
                    - Araragi, Nadeko, Manga cover from the [Monogatari series](https://mangadex.org/title/4265c437-7d57-4d31-9b1d-0e574a07b7b7/bakemonogatari)
                </p>
            </div>
        </div>
    </div>

</div>
