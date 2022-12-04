var links = [
    {
        url: 'https://junko.konno.tk/',
        description: 'This site',
        logo: './assets/img/logo/junko.png'
    },
    {
        url: 'https://archlinux.org/',
        description: 'i use arch btw',
        logo: './assets/img/logo/archlinux.gif'
    },
    {
        url: 'https://sdf.org/',
        description: 'SDF.org',
        logo: './assets/img/logo/sdf.png'
    },
    {
        url: 'http://lucky-ch.com/',
        description: 'Lucky Star',
        logo: './assets/img/logo/konata.gif'
    },
    {
        url: 'https://utsuho.rocks/',
        description: 'The inspiration for this site.',
        logo: './assets/img/logo/utsuhorocks.png'
    },
    {
        url: 'https://www.katawa-shoujo.com/about.php',
        description: 'Absolute banger of a game.',
        logo: './assets/img/logo/katawashoujo.jpg'
    }
];
var friend_links = [
    {
        url: "https://727.pages.dev/",
        description: "Dude's site is broken all the time lmao",
        logo: "https://cdn.discordapp.com/attachments/782136789103280129/1048221626233278524/oie_C2AQ4IDHYC7b.jpg"
    }
];
function generate_html_from_link_info_ts(data) {
    var content_row = document.createElement("div");
    // Title
    var row_title_container = document.createElement('div');
    var row_title = document.createElement('a');
    row_title.classList.add('link', 'underline', 'grid-link');
    row_title.href = data.url;
    row_title_container.appendChild(row_title);
    // Logo (if any)
    if (data.logo) {
        row_title.innerText = '';
        var row_title_logo = document.createElement('img');
        row_title_logo.classList.add('link', 'underline', 'link-logo');
        row_title_logo.src = data.logo;
        row_title_logo.alt = data.description;
        row_title_logo.title = "".concat(data.url, ": ").concat(data.description);
        row_title.appendChild(row_title_logo);
    }
    // Done
    content_row.appendChild(row_title_container);
    return content_row;
}
// Here goes nothing
var general = document.getElementById("grid-general");
var friend = document.getElementById("grid-friends");
// General
for (var i = 0; i < links.length; i++) {
    general === null || general === void 0 ? void 0 : general.appendChild(generate_html_from_link_info_ts(links[i]));
}
// Friends
for (var i = 0; i < friend_links.length; i++) {
    friend === null || friend === void 0 ? void 0 : friend.appendChild(generate_html_from_link_info_ts(friend_links[i]));
}
// To anyone reading this in the web, this file is generated from a typescript file.
