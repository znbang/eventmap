export default {
  book: 'Books',
  cancel: 'Cancel',
  close: 'Close',
  delete: 'Delete',
  edit: 'Edit',
  event: 'Events',
  home: 'Home',
  language: 'Language',
  login: 'Login',
  logout: 'Logout',
  manage: 'Manage',
  save: 'Save',
  search: 'Search',
  task: 'Tasks',
  auths: {
    facebook: 'Login with Facebook Account',
    github: 'Login with Github Account',
    google: 'Login with Google Account',
    line: 'Login with Line Account',
  },
  books: {
    title: 'Title',
    author: 'Author',
    url: 'URL',
    download: 'Download',
    syncAll: 'Clean then sync all',
    syncNew: 'Sync new',
    required: {
      title: 'Please input the title.',
      author: 'Please input the author.',
      url: 'Please input the URL.',
    },
    invalid: {
      url: 'The URL is not supported.',
    },
    exists: {
      title: 'The same title already exists.',
      author: 'The same author already exists.',
      url: 'The same URL already exists.',
    },
    dialog: {
      confirm: 'Do you want to delete the book?',
      deleted: 'Book has been deleted.',
    }
  },
  chapters: {
    dialog: {
      confirm: 'Do you want to delete this chapter and subsequent chapters?',
      deleted: 'Chapters have been deleted.',
    }
  },
  tasks: {
    confirm: 'Do you want to delete the task?',
    deleted: 'Task hsa been deleted',
  },
  events: {
    name: 'Name',
    location: 'Location',
    date: 'Date',
    url: "Website",
    detail: 'Details',
    find: 'Find location',
    required: {
      name: 'Please input the name.',
      location: 'Please input the location.',
      date: 'Please input the date.',
      url: 'Please input the website.',
      detail: 'Please input the details.',
    },
    size: {
      name: 'The name is limited to 255 characters.',
      location: 'The location is limited to 255 characters.',
    },
    invalid: {
      url: 'Incorrect URL format.',
    },
    dialog: {
      confirm: 'Do you want to delete the event?',
      deleted: 'Event has been deleted.',
    }
  },
}
