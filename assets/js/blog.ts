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
const TEMPLATE: string = `
<div class="term" id="post">
    <span class="user">junko</span> ;DATE;
    <a class="links" id=";ID;" href="#;ID;">No. ;ID; </a><br />
    ;IMAGE_EXTRA;
    ;TEXT;
    <br />
</div>`;

const TEMPLATE_EXTRA: string = `
file: <a class="url" href=";IMGURL;">;FILENAME;</a>
(;FILE_TYPE;)<br /><a href=";IMGURL;"><img class="media" src=";IMGURL;" /></a>`;

// Blog index
type Blog = { date: string, id: number, file_url: string | null, file_type: string | null, text: string }

const BLOGINDEX: Array<Blog> = [
    {
        date: "Mon, 26 Dec 22 08:31",
        id: 1,
        text: "oh hey look it finally works... nice!",
        file_url: "https://media.tenor.com/ceXJoo401McAAAAd/cool-aneurysm.gif",
        file_type: "image/gif"
    }
]
//
var blog_mount = document.getElementById("blog-mount");

for (let i = 0; i < BLOGINDEX.length; i++) {
    let current = BLOGINDEX[i]

    let html_res = TEMPLATE
        .replace(/;DATE;/g, current.date)
        .replace(/;ID;/g, current.id.toString())
        .replace(/;TEXT;/g, current.text)

    if (current.file_url) {
        html_res = html_res.replace(/\;IMAGE_EXTRA\;/g, TEMPLATE_EXTRA
            .replace(/;IMGURL;/g, current.file_url)
            .replace(/;FILENAME;/g, current.file_url.split("/").pop()!)
            .replace(/;FILE_TYPE;/g, current.file_type!
            )
        )
    } else {
        html_res = html_res.replace(/\;IMAGE_EXTRA\;/g, "")
    }

    blog_mount!.innerHTML += html_res


}