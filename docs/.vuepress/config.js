import { defaultTheme } from '@vuepress/theme-default'
import { viteBundler } from '@vuepress/bundler-vite'

export default {
  bundler: viteBundler(),
  title: 'My Blog',
  description: 'A simple blog powered by VuePress',
  base: '/',
  head: [
    ['link', { rel: 'icon', href: '/favicon.ico' }]
  ],
  locales: {
    '/': {
      lang: 'en-US',
    },
  },
  theme: defaultTheme({
    navbar: [
      {
        text: 'Home',
        link: '/',
      },
      {
        text: 'Blog',
        link: '/blog/',
      },
    ],
    sidebar: {
      '/blog/': [
        {
          text: 'Blog Posts',
          children: [
            '/blog/first-post.md',
          ],
        },
      ],
    },
  }),
}
