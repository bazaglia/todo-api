CREATE TYPE TASK_STATUS AS ENUM ('open', 'archived');

CREATE TABLE users (
  id serial PRIMARY KEY,
  firstname VARCHAR(50) NOT NULL,
  lastname VARCHAR(50) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE tasks (
  id serial PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  name VARCHAR(200) NOT NULL,
  description TEXT,
  location VARCHAR(100),
  date DATE,
  status TASK_STATUS,
  labels TEXT ARRAY,
  comments TEXT ARRAY,
  is_favorite BOOLEAN DEFAULT false,
  attachment_provider VARCHAR(20),
  attachment_bucket VARCHAR(20),
  attachment_object VARCHAR(50)
);