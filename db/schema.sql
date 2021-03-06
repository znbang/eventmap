create table users (
  id text primary key,
  created_at timestamptz not null,
  provider int not null,
  uid text not null,
  email text not null,
  name text not null
);

create unique index users_idx on users (email, provider);

create table user_sessions (
  id text primary key,
  user_id text not null,
  created_at timestamptz not null
);

create table events (
  id text primary key,
  user_id text not null,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  start_date timestamptz not null,
  end_date timestamptz not null,
  name text not null,
  location text not null,
  url text not null,
  detail text not null,
  lng float not null,
  lat float not null,
  zoom smallint not null
);

create index events_date_idx on events (end_date);
create index events_user_idx on events (user_id);

create table books (
  id text primary key,
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
  id text primary key,
  created_at timestamptz not null,
  updated_at timestamptz not null,
  book_id bigint not null,
  page int not null,
  url text not null,
  title text not null,
  body text not null
);

create unique index chapters_page_idx on chapters (book_id, page);

create table book_jobs (
  id text primary key,
  created_at timestamptz not null,
  done boolean not null,
  book_id text not null,
  title text not null,
  message text not null
);

create index book_jobs_done_idx on book_jobs (done);
