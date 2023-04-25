const meta = {
  auth: 'auth',
}

const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'login', component: () => import('pages/LoginPage.vue') },
      { path: 'events', component: () => import('pages/events/IndexPage.vue') },
      { path: 'events/user', component: () => import('pages/events/UserPage.vue'), meta },
      { path: 'events/new', component: () => import('pages/events/NewPage.vue'), meta },
      { path: 'events/:id', component: () => import('pages/events/ShowPage.vue') },
      { path: 'events/:id/edit', component: () => import('pages/events/EditPage.vue'), meta },
      { path: 'books', component: () => import('pages/books/IndexPage.vue'), meta },
      { path: 'books/new', component: () => import('pages/books/NewPage.vue'), meta },
      { path: 'books/:id/edit', component: () => import('pages/books/EditPage.vue'), meta },
      { path: 'books/:id/toc', component: () => import('pages/books/TocPage.vue'), meta },
      { path: 'books/:id/:page', component: () => import('pages/books/ChapterPage.vue'), meta },
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
