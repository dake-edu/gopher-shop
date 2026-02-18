import { withMermaid } from 'vitepress-plugin-mermaid';

export default withMermaid({
  base: '/gopher-shop/',
  title: 'The Gopher Shop',
  description: 'From Junior to Middle-level Go Backend Engineer.',
  head: [['link', { rel: 'icon', href: '/gopher-shop/gopher.png' }]],
  themeConfig: {
    logo: '/gopher.png',
    head: [['link', { rel: 'icon', href: '/gopher-shop/gopher.png' }]],
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/junior-path' }
    ],

    sidebar: [
      {
        text: 'Level 1: The Junior (Foundations)',
        items: [
          { text: '01. Origin Story', link: '/junior/01-intro' },
          { text: '02. Setup', link: '/junior/02-setup' },
          { text: '03. Anatomy', link: '/junior/03-anatomy' },
          { text: '04. Variables', link: '/junior/04-variables' },
          { text: '05. Logic (If/Else)', link: '/junior/05-logic' },
          { text: '06. Loops', link: '/junior/06-loops' },
          { text: '07. Arrays & Slices', link: '/junior/07-arrays-slices' },
          { text: '08. Maps', link: '/junior/08-maps' },
          { text: '09. Functions', link: '/junior/09-functions' },
          { text: '10. Errors (Guard Rails)', link: '/junior/10-errors' },
          { text: '11. JSON (Translator)', link: '/junior/11-json' },
          { text: '12. Debugging', link: '/junior/12-debugging' },
          { text: '13. The Basic Server', link: '/junior/13-server' },
          { text: 'The Bridge', link: '/junior/bridge' },
        ]
      },
      {
        text: 'Level 2: The Middle (Professional)',
        items: [
          { text: '01. Architecture (3-Layer)', link: '/middle/01-architecture' },
          { text: '02. Templating', link: '/middle/02-templating' },
          { text: '03. Structures & Models', link: '/middle/03-models' },
          { text: '04. Interfaces', link: '/middle/04-interfaces' },
          { text: '05. Configuration', link: '/middle/05-config' },
          { text: '06. In-Memory Store', link: '/middle/06-memory-store' },
          { text: '07. Validation', link: '/middle/07-validation' },
          { text: '08. Concurrency', link: '/middle/08-concurrency' },
          { text: '09. Postgres', link: '/middle/09-postgres' },
          { text: '10. Middleware', link: '/middle/10-middleware' },
          { text: '11. Testing', link: '/middle/11-testing' },
          { text: '12. CI/CD Workflow', link: '/middle/12-workflow' },
          { text: '13. Grand Assembly', link: '/middle/13-assembly' },
        ]
      },
      {
        text: 'Level 3: The Senior (Architect)',
        items: [
          { text: '01. Microservices', link: '/senior/01-microservices' },
          { text: '02. App Concurrency', link: '/senior/02-adv-concurrency' },
          { text: '03. Caching (Redis)', link: '/senior/03-caching' },
          { text: '04. Message Brokers', link: '/senior/04-message-brokers' },
          { text: '05. GoTracker (Project)', link: '/senior/05-gotracker' },
          { text: '06. Orchestration', link: '/senior/06-orchestration' },
          { text: '07. Career & Portfolio', link: '/senior/07-career' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: `Copyright © ${new Date().getFullYear()} The Gopher Shop Team`
    }
  }
});
