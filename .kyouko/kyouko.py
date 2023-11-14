""" kyouko.py - simple script to convert the html files """
from pathlib import Path

# files
TEMPLATE_FOLDER: Path = Path.cwd() / ".kyouko" / "templates"
TEMPLATE: Path = TEMPLATE_FOLDER / "template.html"
FOOTER: Path = TEMPLATE_FOLDER / "footer.html"
HEADER: Path = TEMPLATE_FOLDER / "head.html"

PAGE_FOLDER: Path = Path.cwd() / ".kyouko" / "pages"
OUTPUT_FOLDER: Path = Path.cwd()


def main() -> int:
    RAW_TEMPLATE = TEMPLATE.read_text()

    for page in PAGE_FOLDER.glob("*.html"):
        PAGE_DATA = page.read_text()

        # piss easy parser
        content: dict[str, str] = {
            "FOOTER": FOOTER.read_text(),
            "HEAD": HEADER.read_text(),
        }
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

        current_page = RAW_TEMPLATE

        for mode, code in content.items():
            current_page = current_page.replace(f"@{mode.upper()}_INSERT", code)

        output_file = OUTPUT_FOLDER / f"{page.stem}.html"
        output_file.write_text(current_page)

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
