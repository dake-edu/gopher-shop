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
          { text: '01. Origin Story', link: '/lessons/01-intro' },
          { text: '02. Setup', link: '/lessons/02-setup' },
          { text: '03. Anatomy', link: '/lessons/03-anatomy' },
          { text: '04. Variables', link: '/lessons/04-variables' },
          { text: '05. Logic (If/Else)', link: '/lessons/05-logic' },
          { text: '06. Loops', link: '/lessons/06-loops' },
          { text: '07. Arrays & Slices', link: '/lessons/07-arrays-slices' },
          { text: '08. Maps', link: '/lessons/08-maps' },
          { text: '09. Functions', link: '/lessons/09-functions' },
          { text: '10. Errors (Guard Rails)', link: '/lessons/10-errors' },
          { text: '11. JSON (Translator)', link: '/lessons/11-json' },
          { text: '12. Visual Signals', link: '/lessons/12-signals' },
          { text: '13. Debugging', link: '/lessons/13-debugging' },
        ]
      },
      {
        text: 'Level 2: The Apprentice',
        items: [
          { text: '14. The Basic Server', link: '/lessons/14-server' },
          { text: '15. Project Structure', link: '/lessons/15-structure' },
          { text: '16. HTML Sorting & Layouts', link: '/lessons/16-templating' },
          { text: '17. Structures & Models', link: '/lessons/17-models' },
          { text: '18. Interfaces', link: '/lessons/18-interfaces' },
          { text: '19. Configuration', link: '/lessons/19-config' },
          { text: '20. In-Memory Store', link: '/lessons/20-memory-store' },
          { text: '21. Validation', link: '/lessons/21-validation' },
        ]
      },
      {
        text: 'Level 3: The Professional',
        items: [
          { text: '22. Concurrency', link: '/lessons/22-concurrency' },
          { text: '23. Postgres (The Safe)', link: '/lessons/23-postgres' },
          { text: '24. Middleware (The Onion)', link: '/lessons/24-middleware' },
          { text: '25. Testing', link: '/lessons/25-testing' },
          { text: '26. CI/CD Workflow', link: '/lessons/26-workflow' },
          { text: '27. Grand Assembly', link: '/lessons/27-assembly' },
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
