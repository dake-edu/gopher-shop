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
          { text: '10. Errors (Guard Rails)', link: '/lessons/10-errors' },
          { text: '11. JSON (Translator)', link: '/lessons/11-json' },
          { text: '12. Visual Signals', link: '/signals' },
          { text: '13. Debugging', link: '/debugging' },
        ]
      },
      {
        text: 'Level 2: The Apprentice',
        items: [
          { text: '14. The Basic Server', link: '/lessons/01-server' },
          { text: '15. Project Structure', link: '/lessons/02-structure' },
          { text: '16. HTML Sorting & Layouts', link: '/lessons/07-templating' },
          { text: '17. Structures & Models', link: '/models' },
          { text: '18. Interfaces', link: '/interfaces' },
          { text: '19. Configuration', link: '/config' },
          { text: '20. In-Memory Store', link: '/memory-store' },
          { text: '21. Validation', link: '/validation' },
        ]
      },
      {
        text: 'Level 3: The Professional',
        items: [
          { text: '22. Concurrency', link: '/concurrency' },
          { text: '23. Postgres (The Safe)', link: '/postgres' },
          { text: '24. Middleware (The Onion)', link: '/middleware' },
          { text: '25. Testing', link: '/testing' },
          { text: '26. CI/CD Workflow', link: '/workflow' },
          { text: '27. Grand Assembly', link: '/assembly' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ],

    docFooter: {
      prev: 'Previous Page',
      next: 'Next Page'
    }
  }
}
