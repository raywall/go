import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'Let`s Go',
  tagline: 'curso intensivo de Go para desenvolvedores',
  favicon: 'img/favicon.ico',
  
  future: { v4: true },
  url: 'https://raywall.github.io',
  baseUrl: '/go/',
  
  organizationName: 'raywall',
  projectName: 'go',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  i18n: {
    defaultLocale: 'pt',
    locales: ['pt', 'en', 'es'],
    localeConfigs: {
      pt: {
        label: 'Português (Brasil)',
        direction: 'ltr',
        htmlLang: 'pt-BR',
        calendar: 'gregory',
        path: 'pt',
      },
      en: {
        label: 'English',
        direction: 'ltr',
        htmlLang: 'en-US',
        calendar: 'gregory',
        path: 'en',
      },
      es: {
        label: 'Español',
        direction: 'ltr',
        htmlLang: 'es-ES',
        calendar: 'gregory',
        path: 'es',
      }
    },
  }, 

  presets: [
    [
      'classic',
      {
        gtag: { trackingID: "G-JTYKZ2TBVQ", anonymizeIP: true },
        docs: { sidebarPath: './sidebars.ts' },
        blog: {
          showReadingTime: true,
          feedOptions: { type: ['rss', 'atom'], xslt: true },
          onInlineTags: 'warn',
          onInlineAuthors: 'warn',
          onUntruncatedBlogPosts: 'warn',
        },
        theme: { customCss: './src/css/custom.css' },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    image: 'img/gopher-program.jpeg',
    navbar: {
      title: 'Let`s Go',
      logo: { alt: 'Gopher', src: 'img/go-white.svg' },
      items: [
        { type: 'docSidebar', sidebarId: 'tutorialSidebar', position: 'left', label: 'Aprenda Golang' },
        { type: 'docSidebar', sidebarId: 'tutorialSidebar', position: 'left', label: 'Go + AWS Cloud (em breve)' },
        { href: 'https://github.com/raywall/go', label: 'GitHub', position: 'right' },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Aprenda uma linguagem',
          items: [
            {
              label: 'Golang',
              to: '/docs/conteudo',
            },
            {
              label: 'Golang + AWS Cloud (em breve)',
              to: '/docs/conteudo',
            },
          ],
        },
        {
          title: 'Componentes',
          items: [
            {
              label: 'Cloud Easy Connector',
              href: 'https://github.com/raywall/cloud-easy-connector',
            },
            {
              label: 'Cloud Policy Serializer',
              href: 'https://github.com/raywall/cloud-policy-serializer',
            },
            {
              label: 'GraphQL Connector',
              href: 'https://github.com/raywall/go-graphql-connector',
            },
          ],
        },
        {
          title: 'Mais',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/raywall',
            },
            {
              label: 'Linkedin',
              href: 'https://www.linkedin.com/in/raywall/',
            },
            {
              label: 'Descomplicando IA para Analistas de Produto',
              href: 'https://www.amazon.com.br/Descomplicando-Intelig%C3%AAncia-Artificial-Analistas-Produto/dp/6501238455',
            }
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Raywall Malheiros. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ['go'],
    },
  } satisfies Preset.ThemeConfig,
};

export default config;