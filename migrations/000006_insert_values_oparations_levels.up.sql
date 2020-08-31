INSERT INTO 
  access_levels(name) 
VALUES 
  ('Гость'),
  ('Менеджер'),
  ('Директор'),
  ('Управляющий РС'),
  ('Администратор');

INSERT INTO 
  users_operations(name, access_level) 
VALUES 
  ('Регистрация', 1),
  ('Вход', 1),
  ('Выход', 1),
  ('Создание заказа', 2);
  
INSERT INTO
  users(login, password, email, first_name, last_name, third_name, access_level, unit_id)
VALUES
  ('admin', '6UmH7d8hySh9Slxav0HHQij1JXF0yd+rKaXRRS6kBP/Eok3Nqy/FlwJQvrr8egwk/7gkK1u1B4nZdatPsCGHtw==', 'demsasha4yt@yandex.ru', 'Александр', 'Дементьев', 'Владимирович', 4, 1);