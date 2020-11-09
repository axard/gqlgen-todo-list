--- Создать таблицу пользователей
CREATE TABLE IF NOT EXISTS users(
  id    SERIAL PRIMARY KEY,
  login VARCHAR(128) NOT NULL);
