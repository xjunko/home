links = [
    {
        "url": "https://junko.konno.tk",
        "title": "Homepage",
        "description": "The motherfucking homepage."
    },

    {
        "url": "https://yuzumi.yuzuhara.tk",
        "title": "Homepage (Alternative) (Dead)",
        "description": "The same homepage but with a different domain."
    },

    {
        "url": "#",
        "title": "-------------------------------",
        "description": "-------------------------------"
    },

    {
        "url": "#TODO",
        "title": "TODO: Add more links",
        "description": "TODO: Add more links"
    }
];

function generate_html_from_link_info(data) {
    /*
        <tr>
                <td>
                    <a class="link-title link" href="https://twitter.com/junkokonn0"> #LINK_NAME_HERE </a>
                </td>

                <td style="text-align:center"><a aria-label="spacer">&nbsp;&nbsp;&nbsp;</a></td>

                <td>
                    <a class="link-description"> #DESCRIPTION </a>
                </td>
        </tr>
    */

    // Start generating html with browser api
    table_row = document.createElement("tr");

    // Title
    row_title_container = document.createElement("td");
    row_title = document.createElement("a");
    row_title.classList.add("link-title", "link", "underline");
    row_title.href = data.url;
    row_title.innerText = data.title;
    row_title_container.appendChild(row_title);

    // Spacer thingies
    row_space = document.createElement("td");
    row_space.style["text-align"] = "center";
    row_space_text = document.createElement("a")
    row_space_text.innerText = "   ";
    row_space.appendChild(row_space_text);

    // Description
    row_description_container = document.createElement("td");
    row_description = document.createElement("a");
    row_description.classList.add("link-description");
    row_description.innerText = data.description;
    row_description_container.appendChild(row_description);

    // Finish
    table_row.appendChild(row_title_container);
    table_row.appendChild(row_space);
    table_row.appendChild(row_description_container);

    return table_row;
}

// here we go
lists = document.getElementById("list_shit");

for (let i = 0; i < links.length; i++) {
    lists.appendChild(generate_html_from_link_info(links[i]));
}

console.log("Done!")