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
          { text: '01. Origin Story', link: '/intro' },
          { text: '02. Setup', link: '/setup' },
          { text: '03. Anatomy', link: '/anatomy' },
          { text: '04. Variables', link: '/variables' },
          { text: '05. Logic (If/Else)', link: '/logic' },
          { text: '06. Loops', link: '/loops' },
          { text: '07. Arrays & Slices', link: '/lessons/03-data' },
          { text: '08. Maps', link: '/lessons/05-maps' },
          { text: '09. Functions', link: '/lessons/04-interaction' },
          { text: '10. Debugging', link: '/debugging' },
        ]
      },
      {
        text: 'Level 2: The Apprentice',
        items: [
          { text: '11. The Basic Server', link: '/lessons/01-server' },
          { text: '12. Project Structure', link: '/lessons/02-structure' },
          { text: '13. HTML Sorting & Layouts', link: '/lessons/07-templating' },
          { text: '14. Structures & Models', link: '/models' },
          { text: '15. Interfaces', link: '/interfaces' },
          { text: '16. Configuration', link: '/config' },
          { text: '17. In-Memory Store', link: '/memory-store' },
          { text: '18. Validation', link: '/validation' },
        ]
      },
      {
        text: 'Level 3: The Professional',
        items: [
          { text: '19. Concurrency', link: '/concurrency' },
          { text: '20. Postgres (The Safe)', link: '/postgres' },
          { text: '21. Middleware (The Onion)', link: '/middleware' },
          { text: '22. Testing', link: '/testing' },
          { text: '23. CI/CD Workflow', link: '/workflow' },
          { text: '24. Grand Assembly', link: '/assembly' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ]
  }
}
