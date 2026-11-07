type MarkdownNode = {
  type: string;
  value?: string;
  children?: MarkdownNode[];
};

export type KFCustomTagRenderer = (value: string) => string;

export const KFPostParsers: Record<string, KFCustomTagRenderer> = {
  image: (value) => {
    const items = value.match(/^(\S+)(?:\s+(.*))?$/)?.slice(1) ?? [value];

    let url: URL;
    try {
      url = new URL(items[0]);
    } catch {
      throw new Error(`Invalid image URL in content tag: ${value}`);
    }

    if (url.protocol !== "http:" && url.protocol !== "https:") {
      throw new Error(`Unsupported image URL protocol: ${url.protocol}`);
    }

    const src = items[0]
      .replaceAll("&", "&amp;")
      .replaceAll('"', "&quot;")
      .replaceAll("<", "&lt;")
      .replaceAll(">", "&gt;");

    let style: string = "";
    if (items.length >= 2) {
      style = items[1];
    }


    let imgCode = '<img loading="lazy" alt="image" ';
    if (style.length > 0) {
      imgCode += `style="${style}" `;
    }
    imgCode += `src="${src}"></img>`;

    return imgCode;
  },
  video: (value) => {
    const items = value.match(/^(\S+)(?:\s+(.*))?$/)?.slice(1) ?? [value];

    let url: URL;
    try {
      url = new URL(items[0]);
    } catch {
      throw new Error(`Invalid video URL in content tag: ${value} `);
    }

    if (url.protocol !== "http:" && url.protocol !== "https:") {
      throw new Error(`Unsupported video URL protocol: ${url.protocol} `);
    }

    const src = items[0]
      .replaceAll("&", "&amp;")
      .replaceAll('"', "&quot;")
      .replaceAll("<", "&lt;")
      .replaceAll(">", "&gt;");


    let style: string = "";
    if (items.length >= 2) {
      style = items[1];
    }

    let videoCode = `<video muted loop controls preload="metadata" `;
    if (style.length > 0) {
      videoCode += `style="${style}" `;
    }
    videoCode += `src="${src}"></video>`;

    return videoCode;
  },
  green: (value) => {
    return `<p style='color: var(--green-text)'>&gt;${value.trim()}</p>`;
  },
  red: (value) => {
    return `<p style='color: var(--red-text)'>&lt;${value.trim()}</p>`;
  }
};

export function KFChannelPostParser() {
  const channelCustomTagPattern = /^\[([a-z][\w-]*):\s*([^\]\n]+)\]$/i;

  function getText(node: MarkdownNode): string {
    if (node.children) return node.children.map(getText).join("");
    return node.value ?? "";
  }


  return (tree: MarkdownNode) => {
    const transform = (parent: MarkdownNode) => {
      if (!parent.children) return;

      parent.children = parent.children.flatMap((node) => {
        if (node.type === "paragraph") {
          const match = channelCustomTagPattern.exec(getText(node).trim());
          if (match) {
            const [, name, rawValue] = match;
            const renderer = KFPostParsers[name.toLowerCase()];
            if (renderer) {
              return [{ type: "html", value: renderer(rawValue.trim()) }];
            }
          }
        }

        transform(node);
        return [node];
      });
    };

    transform(tree);
  };
}
