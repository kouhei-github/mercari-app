# Nuxt 3 Awesome Starter
a Nuxt 3 starter template or boilerplate with a lot of useful features. and integrated with TailwindCSS 3.

> **WARNING** \
> Nuxt 3 now in RC version (Release Candidate) [But i'm still using from nuxt3 latest version]  
> Nuxt 3 is still in beta and is not recommended for production use.  
> Breaking changes tracker can be found [here](https://github.com/nuxt/framework/discussions/2883)

## Features
- [x] 💨 [Tailwind CSS v3](https://tailwindcss.com/) with [Windicss](https://windicss.org/)
- [x] ✨ [Headless UI](https://headlessui.dev/)
- [x] 🔔 [Icon Pack Component (unplugin-icons)](https://icones.js.org/)
- [x] 🛹 [State & Store Management (Pinia)](https://pinia.vuejs.org/)
- [x] 🚩 [Localization (i18n) by @intlify](https://github.com/intlify/nuxt3)
- [x] 📦 [Vue Composition Collection (Vueuse)](https://vueuse.org/)
- [x] 🌙 Switch Theme (light, dark, system, realtime)
- [x] 🇮🇩 Language Switcher
- [x] 🪝 Built-in Component & Layout
- [x] Eslint & Prettier
- [x] Husky & Commitlint

## To Do
- [ ] Adding simple Unit Test
- [ ] Adding HTTP Client
- [ ] Configurable Theme (colors, fonts, etc)

## Preview
<table align="center">
  <tr>
    <td align="center" width="75%">
      <img src="https://github.com/viandwi24/nuxt3-awesome-starter/blob/main/assets/images/preview.gif?raw=true" alt="Preview" title="Preview">
    </td>
    <td align="center" width="25%">
      <img src="https://github.com/viandwi24/nuxt3-awesome-starter/blob/main/assets/images/preview_mobile.gif?raw=true" alt="Preview" title="Preview">
    </td>
  </tr>
</table>
<p align="center">
  <br>
  <a href="https://nuxt3-awesome-starter.vercel.app/" target="_blank">
    Live Demo
  </a>
  <br><br>
  <a href="https://codesandbox.io/s/github/viandwi24/nuxt3-awesome-starter" title="Open In Code Sandbox">
    <img src="https://img.shields.io/badge/Open%20in-CodeSandbox-blue?style=flat-square&logo=codesandboxg" alt="Open In Code Sandbox">
  </a>
  <br>
  <a href="https://stackblitz.com/github/viandwi24/nuxt3-awesome-starter" title="Open In Stackblitz">
    <img src="https://developer.stackblitz.com/img/open_in_stackblitz.svg" alt="Open In Stackblitz">
  </a>
</p>

## Table of Contents
- [Nuxt 3 Awesome Starter](#nuxt-3-awesome-starter)
  - [Features](#features)
  - [To Do](#to-do)
  - [Preview](#preview)
  - [Table of Contents](#table-of-contents)
  - [Quick Start](#quick-start)
  - [Built-in Components](#built-in-components)
  - [Notes](#notes)
    - [Styles](#styles)
    - [Theme (Dark Mode)](#theme-dark-mode)
    - [Localization](#localization)
    - [Icons](#icons)
    - [Precommit and Postmerge](#precommit-and-postmerge)
  - [License](#license)

## Quick Start
* This project using `yarn` as package manager.
* Clone this project to your computer `git clone https://github.com/viandwi24/nuxt3-awesome-starter`
* Install dependencies `yarn install`
* Run `yarn dev` to start development server and open `http://localhost:3000` in your browser
* Run `yarn build` to build project and `yarn start` to start production server

Checkout the [deployment documentation](https://v3.nuxtjs.org/docs/deployment).


## Built-in Components
- [x] Footer
- [x] Button
- [x] Anchor (link)
- [x] Navbar
  - [x] Navbar Builder
  - [x] Drawer (on mobile)
  - [x] Options (on mobile)
- [x] Action Sheet
- [x] Theme Toggle / Switcher
- [x] Page Layout
  - [x] Wrapper
  - [x] Header
    - [x] Title
  - [x] Body
    - [x] Section
      - [x] Section Wrapper
      - [x] Section Title
- [x] Dashboard Layout
    - [x] Sidebar
- [ ] Modal
- [ ] Alert
- [ ] 

## Notes
### Styles
Tailwindcss import managed by windicss.
and you can add custom styles in :
```
/path/to/assets/sass/app.scss
```
### Theme (Dark Mode)
ThemeManager is a plugin that allows you to switch between themes. this lib in :
```
/path/to/utils/theme.ts
```
`Thememanager` is a function-class construct when app.vue before mounted. theme construct inside `AppSetup()` in `/path/to/app.vue` :
```vue
<!-- /path/to/app.vue -->
<script lang="ts" setup>
import { AppSetup } from '~/utils/app';
// app setup
AppSetup()
</script>
```
To change theme, you can direct set theme from state `theme.setting`, example :
```vue
<script lang="ts" setup>
import { IThemeSettingOptions } from '~/utils/theme'
const themeSetting = useState<IThemeSettingOptions>('theme.setting')
themeSetting.value = 'dark'
</script>
```
When you change state `theme.setting`, it will automatically change theme.

Theme Setting have 4 options :
- `light`
- `dark`
- `system` (operating system theme)
- `realtime` (realtime theme, if 05:00 - 17:00, it will change to light theme, otherwise dark)

We have state `theme.current`, this state return `light` or `dark` theme. basically it's process from `theme.setting`.
dont change theme with this state.
### Localization
Localization is a plugin that allows you to switch between languages. this lib in :
```
/path/to/utils/lang.ts
```
`LanguageManager` is a function-class construct when app.vue before mounted.
this lib depend on [@intlify/nuxt3](https://github.com/intlify/nuxt3)
lang construct inside `AppSetup()` in `/path/to/app.vue` :
<!-- /path/to/app.vue -->
<script lang="ts" setup>
import { AppSetup } from '~/utils/app';
// app setup
AppSetup()
</script>
To change language, you can direct set language from state `lang.setting`, example :
```vue
<script lang="ts" setup>
const langSetting = useState<string>('locale.setting')
langSetting.value = 'en'
</script>
```
When you change state `locale.setting`, it will automatically change language.
### Icons
This project using unplugin-icons for auto generate and import icon as component.

You can see collection icon list in : [https://icones.js.org/](https://icones.js.org/)

you can use `<prefix-collection:icon />` or `<PrefixCollection:Icon />`.

in this project, configuration prefix as a "icon", you can see in `nuxt.config.ts` :
```js
export default defineNuxtConfig({
    ...

    vite: {
        plugins: [
            UnpluginComponentsVite({
                dts: true,
                resolvers: [
                    IconsResolver({
                        prefix: 'Icon',
                    }),
                ],
            }),
        ],
    },

    ...
})
```

Example :
```vue
// use icon from collection "Simple Icons" and name icon is "nuxtdotjs"
<IconSimpleIcons:nuxtdotjs />

// use icon from collection "Unicons" and name icon is "sun"
<IconUil:sun />
```
### Precommit and Postmerge
This project using husky and commitlint for precommit and postmerge.
when you commit, it will check your commit message and running "yarn lint-staged" to check your staged files.
configuration in : `/path/to/.husky/pre-commit` and `/path/to/commitlint.config.js`

And when Postmerge, it will run "yarn" to automatically install new dependencies.
configuration in `/path/to/.husky/post-merge`

## License
This project is licensed under the MIT license, Copyright (c) 2022 Alfian Dwi Nugraha. For more information see the [LICENSE](LICENSE.md) file.