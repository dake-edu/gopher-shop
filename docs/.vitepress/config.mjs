export default {
  base: '/gopher-shop/',
  title: 'The Gopher Shop',
  description: 'From Junior to Middle-level Go Backend Engineer.',
  head: [['link', { rel: 'icon', href: '/gopher-shop/gopher.png' }]],
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/junior-path' }
    ],

    sidebar: [
      {
        text: 'Junior Path: The Foundations',
        items: [
          { text: 'Introduction to Go', link: '/junior-path' },
          { text: 'Your First Web Server', link: '/first-server' },
          { text: 'Logic: The Fork in the Road', link: '/logic' },
          { text: 'The Conveyor Belt (Loops)', link: '/loops' },
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
