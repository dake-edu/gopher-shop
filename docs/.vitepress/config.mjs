export default {
    title: 'The Gopher Shop',
    description: 'From Junior to Middle-level Go Backend Engineer.',
    themeConfig: {
      nav: [
        { text: 'Home', link: '/' },
        { text: 'Guide', link: '/database' }
      ],
  
      sidebar: [
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
