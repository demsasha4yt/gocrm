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
  users(login, password, email, first_name, last_name, third_name, access_level)
VALUES
  ('admin', 'zKgdNE7BHguhCKv+42U0WnRCbF8DgMJRQCi2aqzk3vMGfP0ZNIIes6SK+aE6cZtlVm4rEKfY4earvqcNGIMuSA==', 'demsasha4yt@yandex.ru', 'Александр', 'Дементьев', 'Владимирович', 4);