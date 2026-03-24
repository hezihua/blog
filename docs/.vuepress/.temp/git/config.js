import { GitContributors } from "/home/hezihua/workspace/blog/node_modules/.pnpm/@vuepress+plugin-git@2.0.0-rc.125_@vuepress+bundler-vite@2.0.0-rc.27_@types+node@25.5.0_42d815d1ebf181f68e63a9dbc95e1475/node_modules/@vuepress/plugin-git/dist/client/components/GitContributors.js";
import { GitChangelog } from "/home/hezihua/workspace/blog/node_modules/.pnpm/@vuepress+plugin-git@2.0.0-rc.125_@vuepress+bundler-vite@2.0.0-rc.27_@types+node@25.5.0_42d815d1ebf181f68e63a9dbc95e1475/node_modules/@vuepress/plugin-git/dist/client/components/GitChangelog.js";

export default {
  enhance: ({ app }) => {
    app.component("GitContributors", GitContributors);
    app.component("GitChangelog", GitChangelog);
  },
};
