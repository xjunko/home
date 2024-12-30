@title=doomerposting
@description=internal redirect page
@tags=internal.post_finder
@priority=9999
@route=/redirect.html

<br />

<h1> Loading... </h1>
<a id="redirect-status"> Detecting arguments... </a>

<div>
</div>

<script>
    // Some voodoo shit going on here
    // dont ask.
    const posts_map = {
        {{ range $pageNumber, $pageIDs := .GetChannelPages }}
        {{ $pageNumber }}: [
            {{ range $_, $pageID := $pageIDs}}
            "{{ $pageID }}",
            {{ end }}
        ],
        {{ end }}
    }

    // Common shortcuts
    const indexes = {
        "github": "https://github.com/xjunko",
        "steam": "https://steamcommunity.com/id/jkonno/"
    }

    // Starts here
    const params = new URLSearchParams(window.location.search);
    const status_text = document.getElementById("redirect-status");

    var done = false;

    status_text.style.fontWeight = "bold";
    status_text.style.fontSize = "32px";

    function fuck_right_of_to(url) {
        window.location.replace(url);
    }

    function handle_common_redirect(redirect_to) {
        if (!redirect_to || redirect_to == null) {
            status_text.textContent = "oops you did a fucky wucky :3333, returning to index."
            fuck_right_of_to("/")
        } else {
            if (redirect_to in indexes) {
                fuck_right_of_to(indexes[redirect_to]);
            } else {
                status_text.textContent = "invalid redirect."
                fuck_right_of_to("/")
            }
        }
    }

    function handle_channel_redirect(post_id) {
        var page_id;

        if (!post_id) {
            status_text.textContent = "No parameter given, please give 'id'!!!";
        } else {
            status_text.textContent = "Hold on.";

            let found_the_shit = false;

            for (const [key, value] of Object.entries(posts_map)) {
                if (value.includes(post_id)) {
                    found_the_shit = true;
                    page_id = key;
                    status_text.textContent = "Found the stuff.";
                    break;
                }
            }

            if (!found_the_shit) {
                status_text.textContent = `Did not found post: #` + post_id;
            } else {
                htmx.ajax('GET', `/chan/${page_id}.html`, {
                    target: '#main-container',
                    swap: 'innerHTML'
                });


                htmx.on('htmx:afterSwap', function(event) {
                    if (done) {
                        return;
                    }

                    if (event.target.id === 'main-container') {
                        const postElement = document.getElementById(post_id);

                        if (postElement) {
                            done = true;

                            // Wait for all Images to load before we can scroll to the thing
                            var images = event.target.querySelectorAll('img');
                            var imagesLoaded = 0;
                            var totalImages = images.length;

                            if (totalImages == 0) {
                                scrollToPost(post_id);
                                return;
                            }

                            images.forEach((img) => {
                                img.addEventListener("load", () => {
                                    imagesLoaded++;

                                    if (imagesLoaded == totalImages) {
                                        scrollToPost(post_id);
                                        return;
                                    }
                                });
                            });
                        }
                    }
                });


            }

        }
    }

    // Detect which param is passed, then go with that.
    if (params.get("r")) {
        handle_common_redirect(params.get("r"));
    } else {
        handle_channel_redirect(params.get("id"));
    }

    function scrollToPost(post_id) {
        const postElement = document.getElementById(post_id);
        if (postElement) {
            postElement.scrollIntoView({ behavior: 'smooth' });

            // add post_id to #
            let url = new URL(window.location.href);
            url.hash = post_id;
            window.history.replaceState({}, "", url.toString());
        }
    }
</script>
