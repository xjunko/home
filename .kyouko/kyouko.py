""" kyouko.py - simple script to convert the html files """
from pathlib import Path

import kyouko_addon
from markdown_it import MarkdownIt

# files
TEMPLATE_FOLDER: Path = Path.cwd() / ".kyouko" / "templates"
TEMPLATE: Path = TEMPLATE_FOLDER / "template.html"
FOOTER: Path = TEMPLATE_FOLDER / "footer.html"
HEADER: Path = TEMPLATE_FOLDER / "head.html"

PAGE_FOLDER: Path = Path.cwd() / ".kyouko" / "pages"
BLOG_FOLDER: Path = Path.cwd() / ".kyouko" / "blogs"
OUTPUT_FOLDER: Path = Path.cwd()

RAW_TEMPLATE = TEMPLATE.read_text(encoding="utf-8")
BLOG_FOLDER.mkdir(parents=True, exist_ok=True)


def process_markdown(t: str) -> str:
    markdown = MarkdownIt("commonmark", {"breaks": True, "html": True})

    return markdown.render(t)


def process_page(raw_content: str) -> str:
    PAGE_DATA = raw_content

    # piss easy parser
    content: dict[str, str] = {
        "FOOTER": FOOTER.read_text(),
        "HEAD": HEADER.read_text(),
    }

    content_default: dict[str, str] = {"CONTENT": "", "OUTER_CONTENT": "", "SCRIPT": ""}

    current_mode: str = ""

    for line in PAGE_DATA.splitlines():
        if line.startswith("@"):
            key = line.removeprefix("@")

            if current_mode != key:
                current_mode = key

            if key in content:
                # we're done parsing this one.
                current_mode = ""
                key = ""
                continue
            else:
                content[key] = ""
                continue

        if current_mode:
            content[current_mode] += line + "\n"

    for key, value in content_default.items():
        if key not in content:
            content[key] = value

    current_page = RAW_TEMPLATE

    for mode, code in content.items():
        if mode not in ["FOOTER", "HEAD", "TITLE"]:
            code = process_markdown(code)

        current_page = current_page.replace(f"@{mode.upper()}_INSERT", code)

    return current_page


def main() -> int:
    # Blog (singular page)
    posts_raw: list[tuple[str, int]] | reversed[tuple[str, int]] = []

    # Fetch entries from discord.
    for message in kyouko_addon.get_messages_from_discord():
        filename: int = int(message.created_at.timestamp())
        content: str = message.content

        # Pre-processing to fit the parser
        for i, attachment in enumerate(message.attachments):
            if i == 0:
                content = (
                    f"[]# Thumbnail: {attachment.url.split('?', 1)[0]} \n" + content
                )
            else:
                content += f"{attachment.url} \n"

        # Fix green texts
        lines: list[str] = []
        for line in content.splitlines():
            if line.startswith(">"):
                line = line.replace(">", "\\>", 1)

            lines.append(line)

        content = "\n".join(lines)

        (BLOG_FOLDER / f"{filename}.md").write_text(content, "utf-8")

    # Blog (local files)
    for id, post in enumerate(
        sorted(list(BLOG_FOLDER.glob("*.md")), key=lambda x: x.stem)
    ):
        post_html: str = kyouko_addon.process_blog(id, post, process_markdown)  # type: ignore
        posts_raw.append((post_html, int(post.stem)))

    # Export blog
    blog_output = OUTPUT_FOLDER / "blog.html"
    blog_output.write_text(
        process_page((PAGE_FOLDER / "blog.md").read_text(encoding="utf-8")).replace(
            "{{BLOG_INTERNAL}}", "\n".join(x[0] for x in posts_raw)
        ),
        "utf-8",
    )

    # Normal pages
    for page in PAGE_FOLDER.glob("*.md"):
        if page.name.startswith("blog"):
            continue

        current_page: str = process_page(page.read_text())
        output_file = OUTPUT_FOLDER / f"{page.stem}.html"
        output_file.write_text(current_page, "utf-8")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
