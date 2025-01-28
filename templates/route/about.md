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
                I've been programming since I was a child, probably not the best thing to do growing up, but well, here
                I am, years later down the line.
                Early on, I mostly made awful games in <a
                    href="https://www.clickteam.com/clickteam-fusion-2-5">Clickteam Fusion</a>.
                Then, I started messing around with <a href="https://wordpress.com/">Wordpress</a>, the best thing I
                created during that time was a clone of osu!'s <a href="https://circle-people.com/skins/">Circle
                    People</a> skins website.
                I don't have an archive of it anymore, but I'm sure it was awful. <br /> <br />
                Then about sometimes in the late 2018, I started learning <a href="https://www.python.org/">Python</a>,
                mostly because I wanted to make a clone of the osu!'s <a
                    href="https://top.gg/bot/289066747443675143">owo</a> bot.
                If you can't tell yet, I mostly make clones of things I like. It's the best way to learn, I think.
                <br /> <br />
                Now, I'm mostly able to use a lot of programming languages, except <a
                    href="https://www.rust-lang.org/">Rust</a> (no hate). I just don't get the appeal of it, I'm sure
                it's a great language, but I just don't see myself using it. Maybe I'll change my mind in the future.
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
            <div class="window">
                <div class="window-content" style="display: flex;align-items: center;justify-content: center;gap: 5px;">
                    {{ template "route-template/about/nerd-stats" . }}
                    {{ template "route-template/about/moon-phase" . }}
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
                </p>
            </div>
        </div>
    </div>
</div>