import { type CollectionEntry, getCollection } from "astro:content";
import type { BlogEntries, ChannelEntries } from "../types";

function _get_preview_description(body: string): string {
	const res: string[] = [];

	for (let line of body.split("\n")) {
		line = line.trim();

		// skip markdown content
		if (line.startsWith("#")) {
			continue;
		}

		// found custom preview end marker, stop parsing
		if (line.startsWith("<!--endpreview-->")) {
			break;
		}

		res.push(line);
	}

	return res.join("\n");
}

export async function get_blog_posts(): Promise<BlogEntries[]> {
	return (await getCollection("blog_entries"))
		.sort((a, b) => b.data.date.valueOf() - a.data.date.valueOf())
		.map((post) => ({
			...post,
			preview: _get_preview_description(post.body ?? ""),
		}));
}

export async function get_channel_posts(): Promise<ChannelEntries[]> {
	return (await getCollection("channel_entries")).sort(
		(a, b) => b.data.date.valueOf() - a.data.date.valueOf(),
	);
}
