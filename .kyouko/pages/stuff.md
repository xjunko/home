@TITLE
cool links
@TITLE

@CONTENT
<br />

### cool ebin links
<br />


<!-- this site info -->
**this site** —

- ![](/static/imgs/buttons/junko.png) [site's logo](/)
    - _feel free to include it into your website!_


<!-- friends -->
**friends** —

- [[<u>l-m.dev</u>]](https://me.l-m.dev/)
    - _cool guy, he knows a lot about compilers._
- [[caffeine](https://caffeine.moe/)]
    - _soon-to-be successful music producer (real)._
- [![](/static/imgs/buttons/utsuhorocks.png)](https://utsuho.rocks/) [Utsuho Rocks](https://utsuho.rocks/)
    - _cute site._
- ![Neko's site logo](/static/imgs/buttons/neko-dc.jpg) [neko's page](https://727.pages.dev/)
    - _dude's site is broken all the time lmao._

<!-- inspirations -->
**inspirations** —

- [![microsounds's logo](/static/imgs/buttons/microsounds.gif)](https://microsounds.github.io/)
- [[wowana.me]](https://wowana.me)

<!-- random webring/cool site links -->
**cool sites** —

[![](/static/imgs/buttons/utsuhorocks.png)](https://utsuho.rocks/) [![](/static/imgs/buttons/xn-neko-btn.gif)](https://猫.移动/) [![](/static/imgs/buttons/archlinux.gif)](https://archlinux.org/) [![](/static/imgs/buttons/konata.gif)](http://lucky-ch.com) [![](/static/imgs/buttons/katawashoujo.jpg)](https://www.katawa-shoujo.com/about.php) [![mariluu's site](https://mariluu.hehe.moe/content/maribanner.gif)](https://mariluu.hehe.moe)

@CONTENT

@SCRIPT
<!-- Discord stuff -->
<script type="text/javascript">
    console.log("hiyo!")

    // fetch user data from api
    api_endpoint = 'https://api.lanyard.rest/v1/users/224785877086240768';
    req = new XMLHttpRequest();
    req.open("GET", api_endpoint, true);
    req.onload = function () {
        console.log("get!");
        if (this.status == 200) {
            console.log("good!");
            var data = JSON.parse(this.response).data;
            var user = data.discord_user;
            var status = data.discord_status;

            // online stat
            var html_status = document.getElementById("discord-status");
            var color = ""

            switch (status) {
                case 'online': html_status.text = "online"; color = "color: green;"; break;
                case 'idle': html_status.text = "idling"; color = "color: yellow;"; break;
                case 'dnd': html_status.text = "dnd"; color = "color: red;"; break;
                case 'offline': html_status.text = "offline"; color = "color: gray;"; break;
            }

            html_status.style = color;

            console.log(html_status);
        }
    }
    req.send();
</script>
@SCRIPT