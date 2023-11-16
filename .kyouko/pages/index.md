@TITLE
home
@TITLE

@CONTENT
<br />

## about

my name is junko, you can find me online as xjunko or fireredz. i have a strong interest in eletronics, japanese cultures and coding. <br /> <br />

this website serves as more or less dumping ground of useless information. here, you may find how to contact me, crap i've written and things that you might find interesting. <br /> <br />

currently, the domain is <a href="https://junko.konno.ovh/">junko.konno.ovh</a>, i'm broke so the domain <i>might</i> change from time to time. <br /> <br />


## contact me

i go by the name, <a id="discord-name"></a>, on discord, feel free to message me, i'm not that interesting but a chat won't hurt me :).
Currently <a id='discord-status-about'>offline</a>.

<br/>

besides discord you can contact me thru email, though, response time may vary: 
yuuki [at] konno [dot] ovh
@CONTENT

@SCRIPT
<script type="text/javascript">
    // fetch user data from api
    api_endpoint = 'https://api.lanyard.rest/v1/users/224785877086240768';
    req = new XMLHttpRequest();
    req.open("GET", api_endpoint, true);
    req.onload = function () {
        if (this.status == 200) {
            console.log("good!");
            var data = JSON.parse(this.response).data;
            var user = data.discord_user;
            var status = data.discord_status;

            // pfp then name
            var html_name = document.getElementById("discord-name");
            html_name.innerHTML = `<img class='discord-avatar' src='https://cdn.discordapp.com/avatars/${user.id}/${user.avatar}.png?size=40'> <i>${user.username}</i>`

            // online stat
            var html_status = document.getElementById("discord-status");
            var color = ""

            switch (status) {
                case 'online': html_status.text = "online"; color = "color: green;"; break;
                case 'idle': html_status.text = "idling"; color = "color: yellow;"; break;
                case 'dnd': html_status.text = "dnd"; color = "color: red;"; break;
                case 'offline': html_status.text = "offline"; color = "color: gray;"; break;
            }

            document.getElementById("discord-status-about").text = html_status.text;
            document.getElementById("discord-status-about").style = color;
            html_status.style = color;

            console.log(html_status);
        }
    }

    req.send();
</script>
@SCRIPT