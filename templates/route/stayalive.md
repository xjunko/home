@title=/kys/ - keep yourself safe
@description=a collection of plushie pictures from my friends.
@tags=channel-like
@route=/kys.html

<!-- special style, to overwrite default grid behaviour -->
<style>

.grid {
    display: grid;
    grid-gap: 10px;
    grid-template-columns: repeat(3, auto);
    align-items: start;
}

.grid .window {
    height: auto;
    overflow: auto;
}

.window-content {
    position: relative;
    display: inline-block;
}

#date {
    position: absolute;
    bottom: -3px;
    right: 4px;
    background: rgba(0, 0, 0, 0.5);
    color: white;
}
</style>

<!-- sets the date -->
<script>
document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll(".date").forEach(el => {
    const unixTime = parseInt(el.textContent, 10);
    if (!isNaN(unixTime)) {
      const date = new Date(unixTime * 1000);
      el.textContent = date.toLocaleString();
    }
  });
});
</script>

<h1 style="text-align: center;">/kys/ - Keep yourself safe</h1>
<a class="href href-left" style="position: absolute; top: 32px;" href="/index.html">return</a>

<hr>
<br>

<div class="window-content">
    <div style="position: relative; bottom: 10px;">

        <h3 style="text-align: center;">pictures of people doing mundane things in life with anime plushies.</h3>

    </div>
</div>
<div class="grid">
    <div class="window">
        <div class="window-content">
            <img class="kys-picture" loading=lazy src="https://hatsune-miku.has.rocks/r/aris_mcqueen.jpg">
            <a id="date" class="date">1752247646</a>
        </div>
    </div>
    <div class="window">
        <div class="window-content">
            <img class="kys-picture" loading=lazy src="https://hatsune-miku.has.rocks/r/kagi.jpg">
            <a id="date" class="date">1752854616</a>
        </div>
    </div>
    <div class="window">
        <div class="window-content">
            <img class="kys-picture" loading=lazy src="https://hatsune-miku.has.rocks/r/zed_car.jpg">
            <a id="date" class="date">1752857289</a>
        </div>
    </div>
    <div class="window">
        <div class="window-content">
            <img class="kys-picture" loading=lazy src="https://hatsune-miku.has.rocks/r/kagi_roticanai.jpg">
            <a id="date" class="date">1753371093</a>
        </div>
    </div>
    <div class="window">
        <div class="window-content">
            <img class="kys-picture" loading=lazy src="https://hatsune-miku.has.rocks/r/zed.jpg">
            <a id="date" class="date">1753406161</a>
        </div>
    </div>
    <div class="window">
        <div class="window-content">
            <img class="kys-picture" loading=lazy src="https://hatsune-miku.has.rocks/r/aqua_nity.jpg">
            <a id="date" class="date">1753407762</a>
        </div>
    </div>
</div>