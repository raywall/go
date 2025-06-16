import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

const config: Config = {
  title: 'Let`s Go',
  tagline: 'curso intensivo de Go para desenvolvedores',
  favicon: 'img/favicon.ico',

  // Future flags, see https://docusaurus.io/docs/api/docusaurus-config#future
  future: {
    v4: true, // Improve compatibility with the upcoming Docusaurus v4
  },

  // Set the production url of your site here
  url: 'https://raywall.github.io',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/go/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'raywall', // Usually your GitHub org/user name.
  projectName: 'go', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          // editUrl:
          //   'https://github.com/raywall/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        },
        blog: {
          showReadingTime: true,
          feedOptions: {
            type: ['rss', 'atom'],
            xslt: true,
          },
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          // editUrl:
          //   'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
          // Useful options to enforce blogging best practices
          onInlineTags: 'warn',
          onInlineAuthors: 'warn',
          onUntruncatedBlogPosts: 'warn',
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: 'img/gopher-rocks-sm.svg',  
    navbar: {
      title: 'Let`s Go',
      logo: {
        alt: 'Gopher',
        src: 'img/go-white.svg',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'tutorialSidebar',
          position: 'left',
          label: 'Aprenda Golang',
        },
        {
          href: 'https://github.com/raywall/go',
          label: 'GitHub',
          position: 'right',
        },
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
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
