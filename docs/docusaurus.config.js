// @ts-check
import { themes as prismThemes } from "prism-react-renderer";

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "go-microx",
  tagline: "The ultimate CLI tool to scaffold production-ready microservices in Go.",
  favicon: "img/logo4.png",

  url: "https://gomicrox.abhishekkushwaha.tech",
  baseUrl: "/",

  organizationName: "abhishekkushwaha",
  projectName: "go-microx",

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: "./sidebars.js",
          editUrl: "https://github.com/abhishekkushwahaa/go-microx/tree/main/docs/",
        },
        blog: {
          showReadingTime: true,
          feedOptions: {
            type: ["rss", "atom"],
            xslt: true,
          },
          editUrl: "https://github.com/abhishekkushwahaa/go-microx/tree/main/docs/",
          onInlineTags: "warn",
          onInlineAuthors: "warn",
          onUntruncatedBlogPosts: "warn",
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      image: "img/docusaurus-social-card.jpg",
      metadata: [
        { name: "keywords", content: "golang, microservices, cli, scaffolding, backend, gin, gRPC, docker" },
      ],
      navbar: {
        logo: {
          alt: "go-microx logo",
          src: "img/logo3.png",
        },
        items: [
          {
            type: "docSidebar",
            sidebarId: "tutorialSidebar",
            position: "left",
            label: "Docs",
          },
          {
            href: "https://github.com/abhishekkushwahaa/go-microx",
            label: "GitHub",
            position: "right",
          },
        ],
      },
      footer: {
        style: "dark",
        links: [
          {
            title: "Docs",
            items: [
              { label: "Getting Started", to: "/docs/intro" },
              { label: "Installation", to: "/docs/installation" },
              { label: "Usage Guide", to: "/docs/usage" },
              { label: "Built-in Templates", to: "/docs/intro" },
            ],
          },
          {
            title: "Community",
            items: [
              { label: "GitHub", href: "https://github.com/abhishekkushwahaa/go-microx" },
              { label: "LinkedIn", href: "https://www.linkedin.com/in/abhishekkushwahaa" },
              { label: "X (Twitter)", href: "https://x.com/AbhishekKushwaa" },
            ],
          },
          {
            title: "Resources",
            items: [
              { label: "Go (Golang)", href: "https://go.dev" },
              { label: "Docker Docs", href: "https://docs.docker.com" },
              { label: "Microservices Guide", href: "https://microservices.io" },
            ],
          },
        ],
        copyright: `Made with ❤️ by Abhishek Kushwaha. Copyright © ${new Date().getFullYear()} go-microx. Built with Docusaurus.`,
      },
      prism: {
        theme: prismThemes.github,
        darkTheme: prismThemes.dracula,
        additionalLanguages: ["go", "bash", "yaml", "docker"],
      },
    }),
};

export default config;
