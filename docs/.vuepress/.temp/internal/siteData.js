export const siteData = JSON.parse("{\"base\":\"/\",\"lang\":\"en-US\",\"title\":\"My Blog\",\"description\":\"A simple blog powered by VuePress\",\"head\":[[\"link\",{\"rel\":\"icon\",\"href\":\"/favicon.ico\"}]],\"locales\":{\"/\":{\"lang\":\"en-US\"}}}")

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  __VUE_HMR_RUNTIME__.updateSiteData?.(siteData)
}

if (import.meta.hot) {
  import.meta.hot.accept((m) => {
    __VUE_HMR_RUNTIME__.updateSiteData?.(m.siteData)
  })
}
