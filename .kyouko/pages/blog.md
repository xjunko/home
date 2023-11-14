@TITLE
blog
@TITLE

@CONTENT
<br />

## Notice
migrating... so nothing here for now :(
@CONTENT

@SCRIPT
<!-- Discord stuff -->
<script type="text/javascript">
    // fetch user data from api
    api_endpoint = 'https://api.lanyard.rest/v1/users/224785877086240768';
    req = new XMLHttpRequest();
    req.open("GET", api_endpoint, true);
    req.onload = function () {
        if (this.status == 200) {
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