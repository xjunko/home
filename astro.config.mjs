// @ts-check
import { defineConfig } from "astro/config";
import { KFChannelPostParser } from "./src/lib/parser.ts";

// https://astro.build/config
export default defineConfig({
  markdown: {
    shikiConfig: {
      theme: 'vitesse-black'
    },
    remarkPlugins: [KFChannelPostParser],
  },
});
