export default {
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
          { text: '1. The Origin Story', link: '/intro' },
          { text: '2. The Workbench (Setup)', link: '/setup' },
          { text: '3. Level 1: The Server', link: '/lessons/01-server' },
          { text: '4. Level 2: Structure', link: '/lessons/02-structure' },
          { text: '5. Level 3: Data', link: '/lessons/03-data' },
          { text: '6. Level 4: Interaction', link: '/lessons/04-interaction' },
          { text: '7. Level 5: Maps', link: '/lessons/05-maps' },
          { text: '8. The Bridge', link: '/bridge' },
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
          { text: '17. The Grand Assembly', link: '/assembly' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ]
  }
}
