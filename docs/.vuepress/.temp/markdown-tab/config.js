import { CodeTabs } from "/home/hezihua/workspace/blog/node_modules/.pnpm/@vuepress+plugin-markdown-tab@2.0.0-rc.125_@vuepress+bundler-vite@2.0.0-rc.27_@types+no_811dbdeb3bccd6b1c9fc5258cd806bb1/node_modules/@vuepress/plugin-markdown-tab/dist/client/components/CodeTabs.js";
import { Tabs } from "/home/hezihua/workspace/blog/node_modules/.pnpm/@vuepress+plugin-markdown-tab@2.0.0-rc.125_@vuepress+bundler-vite@2.0.0-rc.27_@types+no_811dbdeb3bccd6b1c9fc5258cd806bb1/node_modules/@vuepress/plugin-markdown-tab/dist/client/components/Tabs.js";
import "/home/hezihua/workspace/blog/node_modules/.pnpm/@vuepress+plugin-markdown-tab@2.0.0-rc.125_@vuepress+bundler-vite@2.0.0-rc.27_@types+no_811dbdeb3bccd6b1c9fc5258cd806bb1/node_modules/@vuepress/plugin-markdown-tab/dist/client/styles/vars.css";

export default {
  enhance: ({ app }) => {
    app.component("CodeTabs", CodeTabs);
    app.component("Tabs", Tabs);
  },
};
