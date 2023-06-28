-- create table users (
--   id         serial primary key,
--   uuid       varchar(64) not null unique,
--   name       varchar(255),
--   email      varchar(255) not null unique,
--   password   varchar(255) not null,
--   created_at timestamp not null   
-- );
--
-- create table sessions (
--   id         serial primary key,
--   uuid       varchar(64) not null unique,
--   email      varchar(255),
--   user_id    integer references users(id),
--   created_at timestamp not null   
-- );
--
-- create table threads (
--   id         serial primary key,
--   uuid       varchar(64) not null unique,
--   topic      text,
--   user_id    integer references users(id),
--   created_at timestamp not null       
-- );
--
-- create table posts (
--   id         serial primary key,
--   uuid       varchar(64) not null unique,
--   body       text,
--   user_id    integer references users(id),
--   thread_id  integer references threads(id),
--   created_at timestamp not null  
-- );

create table if not exists customers (
  id uuid PRIMARY KEY,
  name varchar(64) not null,
  phone varchar(64) not null, 
  is_gold boolean not null
);

create table if not exists genres (
  id serial primary key,
  name varchar(64) not null unique
);

create table if not exists movies (
  id uuid PRIMARY KEY,
  name varchar(64) not null,
  number_in_stock integer not null,
  daily_rental_rate integer not null
);

create table if not exists movies_genres (
  id serial primary key,
  movie_id uuid,
  genre_name varchar(64) not null,
  FOREIGN KEY(movie_id) REFERENCES movies(id)
);

create table if not exists users (
  id uuid PRIMARY KEY,
  name varchar(64) not null,
  email varchar(254) not null unique,
  password varchar(224) not null
)

