import { defineCollection, type RenderResult, render, z } from "astro:content";
import { glob } from "astro/loaders";

function _parse_unix_date(val: unknown): Date {
	if (typeof val === "string" || typeof val === "number") {
		const num = Number(val);
		if (!Number.isNaN(num)) {
			return new Date(num * 1000);
		}
	}

	return new Date(0);
}

function _parse_tags(val: unknown): string[] {
	if (typeof val === "string") {
		return val.split(",").map((t) => t.trim());
	}

	return [];
}

function _mime_from_filename(filename: string): string {
	const ext = filename.toLowerCase().split(".").pop() || "";
	switch (ext) {
		case "png":
			return "png";
		case "jpg":
		case "jpeg":
			return "jpeg";
		case "webp":
			return "webp";
		case "gif":
			return "gif";
		case "svg":
			return "svg+xml";
		default:
			return "any";
	}
}

const blog_entries = defineCollection({
	loader: glob({ base: "./src/content/blog", pattern: "**/*.md" }),
	schema: ({}) =>
		z.object({
			author: z.string(),
			date: z.preprocess(_parse_unix_date, z.date()),
			slug: z.string(),
			tagline: z.string(),
			title: z.string(),
			description: z.string(),
			tags: z.preprocess(_parse_tags, z.array(z.string()).default([])),
		}),
});

const channel_entries = defineCollection({
	loader: glob({ base: "./src/content/@chan", pattern: "**/*.md" }),
	schema: ({}) =>
		z
			.object({
				author: z.string(),
				date: z.preprocess(_parse_unix_date, z.date()),
				tags: z.preprocess(_parse_tags, z.array(z.string()).default([])),
				thumbnail: z.string().optional(),
			})
			.transform((data) => ({
				...data,
				filename:
					typeof data.thumbnail === "string"
						? data.thumbnail
								.split("/")
								.filter((p) => p.length > 0)
								.pop() || "unknown"
						: "unknown",
				mimetype:
					typeof data.thumbnail === "string"
						? _mime_from_filename(
								data.thumbnail
									.split("/")
									.filter((p) => p.length > 0)
									.pop() || "unknown",
							)
						: "application/octet-stream",
			})),
});

export const collections = { blog_entries, channel_entries };
