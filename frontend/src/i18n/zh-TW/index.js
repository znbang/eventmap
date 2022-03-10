export default {
  book: '書本',
  cancel: '取消',
  close: '關閉',
  delete: '刪除',
  edit: '編輯',
  event: '活動',
  home: '首頁',
  language: '語言',
  login: '登入',
  logout: '登出',
  manage: '管理',
  save: '儲存',
  search: '搜尋',
  task: '工作',
  auths: {
    facebook: '使用 Facebook 帳號登入',
    github: '使用 Github 帳號登入',
    google: '使用 Google 帳號登入',
    line: '使用 Line 帳號登入',
  },
  books: {
    title: '書名',
    author: '作者',
    url: '網址',
    download: '下載',
    syncAll: '重新同步所有章節',
    syncNew: '同步新章節',
    required: {
      title: '請輸入書名。',
      author: '請輸入作者。',
      url: '請輸入網址。',
    },
    invalid: {
      url: '網址格式不正確。',
    },
    exists: {
      title: '輸入的書名已存在。',
      author: '輸入的作者已存在。',
      url: '輸入的網址已存在。',
    },
    dialog: {
      confirm: '確定要刪除書本？',
      deleted: '已刪除書本。',
    }
  },
  chapters: {
    dialog: {
      confirm: '確定要刪除本章及之後章節？',
      deleted: '已刪除章節。',
    }
  },
  tasks: {
    confirm: '確定要刪除工作？',
    deleted: '成功刪除工作。',
  },
  events: {
    name: '名稱',
    location: '地點',
    date: '日期',
    url: '網站',
    detail: '內容',
    find: '查詢活動地點',
    required: {
      name: '請輸入名稱。',
      location: '請輸入地點。',
      date: '請輸入日期。',
      url: '請輸入網站或內容。',
      detail: '請輸入內容或網站。',
    },
    size: {
      name: '名稱最多 255 個字。',
      location: '地點最多 255 個字。',
    },
    invalid: {
      url: '網址格式不正確。',
    },
    dialog: {
      confirm: '確定要刪除活動？',
      deleted: '已刪除活動。',
    }
  },
}
