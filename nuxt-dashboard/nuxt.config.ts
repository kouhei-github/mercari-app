import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/docs/directory-structure/nuxt.config
export default defineNuxtConfig({
    nitro: {
        preset: 'node',
    },
    // buildModules: ['@nuxtjs/tailwindcss'], // <= コメントアウト
    css: ['@/assets/css/tailwind.css'],
    // ↓ 追加
    build: {
        postcss: {
            postcssOptions: require('./postcss.config.js'),
        }
    },
    vite:{
        server: {
            watch: {
                usePolling: true,
            },
        },
    },

})
