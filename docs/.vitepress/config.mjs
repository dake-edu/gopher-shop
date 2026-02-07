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
        text: 'Level 1: The Junior',
        items: [
          { text: '01. Origin Story', link: '/intro' },
          { text: '02. Setup', link: '/setup' },
          { text: '03. Basic Server', link: '/lessons/01-server' },
          { text: '04. Project Structure', link: '/lessons/02-structure' },
          { text: '05. Data Types', link: '/lessons/03-data' },
          { text: '06. Interaction', link: '/lessons/04-interaction' },
          { text: '07. Maps', link: '/lessons/05-maps' },
          { text: '08. Debugging', link: '/debugging' },
          { text: '09. The Bridge', link: '/bridge' },
        ]
      },
      {
        text: 'Level 2: The Apprentice',
        items: [
          { text: '10. Structures & Models', link: '/models' },
          { text: '11. Configuration', link: '/config' },
          { text: '12. In-Memory Store', link: '/memory-store' },
          { text: '13. Validation', link: '/validation' },
          { text: '14. Interfaces', link: '/interfaces' },
        ]
      },
      {
        text: 'Level 3: The Professional',
        items: [
          { text: '15. Concurrency', link: '/concurrency' },
          { text: '16. Postgres (The Safe)', link: '/postgres' },
          { text: '17. Middleware (The Onion)', link: '/middleware' },
          { text: '18. Testing', link: '/testing' },
          { text: '19. CI/CD Workflow', link: '/workflow' },
          { text: '20. Grand Assembly', link: '/assembly' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ]
  }
}
