import type { CollectionEntry } from "astro:content";

export type MascotID = "ononoki" | "araragi" | "cover" | "nadeko";

export type BlogEntries = CollectionEntry<"blog_entries"> & {
	preview: string;
};

export type ChannelEntries = CollectionEntry<"channel_entries"> & {};
