import { defineCollection, type RenderResult, render } from "astro:content";
import { z } from "astro/zod";
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
    case "mp4":
      return "mp4";
    default:
      return "any";
  }
}

function _type_from_filename(filename: string): string {
  const ext = filename.toLowerCase().split(".").pop() || "";
  switch (ext) {
    case "mp4":
      return "video";
    default:
      return "image";
  }
}

function _get_filename(path?: string) {
  if (!path) return "unknown";
  return path.split("/").filter(Boolean).pop() ?? "unknown";
}

const blog_entries = defineCollection({
  loader: glob({ base: "./src/content/blog", pattern: "**/**/*.md" }),
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
  schema: z
    .object({
      author: z.string(),
      date: z.preprocess(_parse_unix_date, z.date()),
      tags: z.preprocess(_parse_tags, z.array(z.string()).default([])),
      thumbnail: z.string().optional(),
    })
    .transform((data) => {
      const filename = _get_filename(data.thumbnail);

      return {
        ...data,
        filename,
        mimetype: data.thumbnail
          ? _mime_from_filename(filename)
          : "application/octet-stream",
        type: data.thumbnail ? _type_from_filename(filename) : "image",
      };
    }),
});

export const collections = { blog_entries, channel_entries };
