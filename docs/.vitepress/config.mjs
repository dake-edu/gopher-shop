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
        text: 'Junior Path: The Foundations',
        items: [
          { text: '1. The Origin Story', link: '/intro' },
          { text: '2. The Workbench (Setup)', link: '/setup' },
          { text: '3. The Anatomy of Code', link: '/anatomy' },
          { text: '4. Operations with Memory', link: '/variables' },
          { text: '5. Logic (If/Else)', link: '/logic' },
          { text: '6. Loops (Conveyor Belt)', link: '/loops' },
          { text: '7. Your First Web Server', link: '/first-server' },
        ]
      },
      {
        text: 'Professional Tools',
        items: [
          { text: 'The Auto-Pilot (CI/CD)', link: '/workflow' },
        ]
      },
      {
        text: 'Architecture & Persistence',
        items: [
          { text: 'Introduction', link: '/' },
          { text: 'Database & Repository', link: '/database' },
          { text: 'Configuration', link: '/config' },
        ]
      },
      {
        text: 'Logic & Reliability',
        items: [
          { text: 'Validation (Guard Clauses)', link: '/validation' },
          { text: 'Middleware (The Onion)', link: '/middleware' },
          { text: 'Testing (The Pyramid)', link: '/testing' },
          { text: 'Lifecycle (Cleanup Crew)', link: '/lifecycle' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/dake-edu/gopher-shop' }
    ]
  }
}
