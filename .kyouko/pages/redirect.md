@TITLE
redirect
@TITLE

@CONTENT
<br />
<p style="width: 800px" id="redirect"> redirecting!!! </p>
@CONTENT

@SCRIPT
<!-- Redirect -->
<script type="text/javascript">
    const indexes = {
        "github": "https://github.com/xjunko",
        "steam": "https://steamcommunity.com/id/FireRedz/"
    }

    const query = window.location.search;
    const params = new URLSearchParams(query);

    const redirect_to = params.get("r");

    function fuck_right_of_to(url) {
        window.location.replace(url);
    }

    if (!redirect_to || redirect_to == null) {
        document.getElementById("redirect").innerHTML = "oops you did a fucky wucky :3333, returning to index."
        fuck_right_of_to("/")
    } else {
        if (redirect_to in indexes) {
            fuck_right_of_to(indexes[redirect_to]);
        } else {
            document.getElementById("redirect").innerHTML = "invalid redirect."
            fuck_right_of_to("/")
        }
    }
</script>
@SCRIPT
