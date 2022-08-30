create table users (
  id text not null primary key,
  created_at timestamptz not null,
  provider int not null,
  uid text not null,
  email text not null,
  name text not null
);

create unique index users_idx on users (email, provider);

create table user_sessions (
  id text not null primary key,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  user_id text not null
);

create table events (
  id text not null primary key,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  start_date timestamptz not null,
  end_date timestamptz not null,
  lng float not null,
  lat float not null,
  zoom smallint not null,
  user_id text not null,
  name text not null,
  location text not null,
  url text not null,
  detail text not null
);

create index events_date_idx on events (end_date);
create index events_user_idx on events (user_id);

create table books (
  id text not null primary key,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  user_id text not null,
  title text not null,
  author text not null,
  url text not null
);

create unique index books_url_idx on books (url);
create unique index books_title_idx on books (title, author);

create table chapters (
  id text not null primary key,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  page int not null,
  book_id text not null,
  url text not null,
  title text not null,
  body text not null
);

create unique index chapters_page_idx on chapters (book_id, page);

create table book_jobs (
  id text not null primary key,
  created_at timestamptz not null,
  status int not null,
  book_id text not null,
  message text not null
);

create unique index book_jobs_idx on book_jobs (book_id);
