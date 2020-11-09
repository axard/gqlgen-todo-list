--- Создать таблицу тудушек
CREATE TABLE IF NOT EXISTS todos(
  id          SERIAL PRIMARY KEY,
  description VARCHAR(256) NOT NULL,
  done        BOOLEAN NOT NULL DEFAULT FALSE,
  user_id     INTEGER NOT NULL REFERENCES users(id));
