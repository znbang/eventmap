const meta = {
  auth: 'auth',
}

const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: 'login', component: () => import('pages/Login.vue') },
      { path: 'events', component: () => import('pages/events/Index.vue') },
      { path: 'events/user', component: () => import('pages/events/User.vue'), meta },
      { path: 'events/new', component: () => import('pages/events/New.vue'), meta },
      { path: 'events/:id', component: () => import('pages/events/Show.vue') },
      { path: 'events/:id/edit', component: () => import('pages/events/Edit.vue'), meta },
      { path: 'books', component: () => import('pages/books/Index.vue'), meta },
      { path: 'books/new', component: () => import('pages/books/New.vue'), meta },
      { path: 'books/:id/edit', component: () => import('pages/books/Edit.vue'), meta },
      { path: 'books/:id/toc', component: () => import('pages/books/Toc.vue'), meta },
      { path: 'books/:id/:page', component: () => import('pages/books/Chapter.vue'), meta },
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
