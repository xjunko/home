// To future me: please redo this
/*
    Format

     <div class="term" id="post">
        <span class="user">junko</span> ;DATE;
        <a class="links" id="1638019983" href=".">No.;ID; </a><br />file:
        <a class="url" href="IMGURL">;FILENAME;</a>
        (;SIZE;MB, ;FILE_TYPE;)<br /><a href=";IMGURL;"><img class="media" src=";IMGURL;" /></a>
        ;TEXT;
        <br />
    </div>
*/
var TEMPLATE = "\n<div class=\"term\" id=\"post\">\n    <span class=\"user\">junko</span> ;DATE;\n    <a class=\"links\" id=\";ID;\" href=\"#;ID;\">No. ;ID; </a><br />\n    ;IMAGE_EXTRA;\n    ;TEXT;\n    <br />\n</div>";
var TEMPLATE_EXTRA = "\nfile: <a class=\"url\" href=\";IMGURL;\">;FILENAME;</a>\n(;FILE_TYPE;)<br /><a href=\";IMGURL;\"><img class=\"media\" src=\";IMGURL;\" /></a>";
var BLOGINDEX = [
    {
        date: "Mon, 26 Dec 22 08:31",
        id: 1,
        text: "oh hey look it finally works... nice!",
        file_url: "https://media.tenor.com/ceXJoo401McAAAAd/cool-aneurysm.gif",
        file_type: "image/gif"
    }
];
//
var blog_mount = document.getElementById("blog-mount");
for (var i = 0; i < BLOGINDEX.length; i++) {
    var current = BLOGINDEX[i];
    var html_res = TEMPLATE
        .replace(/;DATE;/g, current.date)
        .replace(/;ID;/g, current.id.toString())
        .replace(/;TEXT;/g, current.text);
    if (current.file_url) {
        html_res = html_res.replace(/\;IMAGE_EXTRA\;/g, TEMPLATE_EXTRA
            .replace(/;IMGURL;/g, current.file_url)
            .replace(/;FILENAME;/g, current.file_url.split("/").pop())
            .replace(/;FILE_TYPE;/g, current.file_type));
    }
    else {
        html_res = html_res.replace(/\;IMAGE_EXTRA\;/g, "");
    }
    blog_mount.innerHTML += html_res;
}
