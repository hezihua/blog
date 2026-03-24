export const redirects = JSON.parse("{}")

export const routes = Object.fromEntries([
  ["/", { loader: () => import(/* webpackChunkName: "index.html" */"/home/hezihua/workspace/blog/docs/README.md"), meta: {"title":"Home"} }],
  ["/blog/", { loader: () => import(/* webpackChunkName: "blog_index.html" */"/home/hezihua/workspace/blog/docs/blog/README.md"), meta: {"title":"Blog"} }],
  ["/blog/first-post.html", { loader: () => import(/* webpackChunkName: "blog_first-post.html" */"/home/hezihua/workspace/blog/docs/blog/first-post.md"), meta: {"title":"First Post"} }],
  ["/404.html", { loader: () => import(/* webpackChunkName: "404.html" */"/home/hezihua/workspace/blog/docs/.vuepress/.temp/pages/404.html.vue"), meta: {"title":""} }],
]);

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  __VUE_HMR_RUNTIME__.updateRoutes?.(routes)
  __VUE_HMR_RUNTIME__.updateRedirects?.(redirects)
}

if (import.meta.hot) {
  import.meta.hot.accept((m) => {
    __VUE_HMR_RUNTIME__.updateRoutes?.(m.routes)
    __VUE_HMR_RUNTIME__.updateRedirects?.(m.redirects)
  })
}
