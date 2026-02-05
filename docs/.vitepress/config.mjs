export default {
  base: '/gopher-shop/',
  title: 'The Gopher Shop',
  description: 'From Junior to Middle-level Go Backend Engineer.',
  head: [['link', { rel: 'icon', href: '/gopher-shop/gopher.png' }]],
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/intro' }
    ],

    sidebar: [
      {
        text: 'Level 1: The Junior (Foundations)',
        items: [
          { text: '1. The Origin Story', link: '/intro' },
          { text: '2. The Workbench (Setup)', link: '/setup' },
          { text: '3. The Anatomy of Code', link: '/anatomy' },
          { text: '4. Operations with Memory', link: '/variables' },
          { text: '5. The Decision Maker', link: '/logic' },
          { text: '6. Collections & Iteration', link: '/loops' },
          { text: '7. Talking to the World', link: '/first-server' },
        ]
      },
      {
        text: 'Level 2: The Apprentice (Building)',
        items: [
          { text: '8. The Blueprint (Models)', link: '/models' },
          { text: '9. The Wiring (Config)', link: '/config' },
          { text: '10. The Memory (Store)', link: '/memory-store' },
          { text: '11. The Inspector (Validation)', link: '/validation' },
        ]
      },
      {
        text: 'Level 3: The Professional (Production)',
        items: [
          { text: '12. The Architect (Interfaces)', link: '/interfaces' },
          { text: '13. The Warehouse (Postgres)', link: '/postgres' },
          { text: '14. The Onion (Middleware)', link: '/middleware' },
          { text: '15. The Safety Net (Testing)', link: '/testing' },
          { text: '16. The Robot (CI/CD)', link: '/workflow' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ]
  }
}
