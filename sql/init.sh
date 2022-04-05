#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    create table if not exists users
    (
        id serial constraint users_pk primary key,
        username varchar(50)  not null,
        email    varchar(50)  not null,
        password varchar(250) not null,
        salt     varchar(50)  not null,
        avatar varchar(100),
        subscription_expires timestamp
    );

    create unique index users_email_uindex
        on users (email);
  COMMIT;

  CREATE TABLE IF NOT EXISTS movies(
       id serial constraint movies_pk primary key,
       name varchar(60) not null unique,
       name_picture varchar(255) not null,
       year smallint not null,
       duration varchar(25) not null,
       age_limit smallint not null,
       description varchar(1024) not null,
       kinopoisk_rating numeric(2,1) not null,
       tagline varchar(255) not null,
       picture varchar(255) not null,
       video varchar(255) not null,
       trailer varchar(255) not null
  );

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Аватар', 'Avatar.webp', '2009', '2 часа 42 минуты', '12',
         'Бывший морпех Джейк Салли прикован к инвалидному креслу. Несмотря на немощное тело, ' ||
         'Джейк в душе по-прежнему остается воином. Он получает задание совершить путешествие в ' ||
         'несколько световых лет к базе землян на планете Пандора, где корпорации добывают ' ||
         'редкий минерал, имеющий огромное значение для выхода Земли из энергетического кризиса.',
         '7.9', 'Это новый мир', 'Avatar.webp', 'Avatar.mp4', 'Avatar.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Титаник', 'Titanic.webp', '1997', '3 часа 14 минут', '12',
         'В первом и последнем плавании шикарного «Титаника» встречаются двое. Пассажир нижней палубы ' ||
         'Джек выиграл билет в карты, а богатая наследница Роза отправляется в Америку, чтобы выйти ' ||
         'замуж по расчёту. Чувства молодых людей только успевают расцвести, и даже не классовые различия ' ||
         'создадут испытания влюблённым, а айсберг, вставший на пути считавшегося непотопляемым лайнера.',
         '8.4', 'Ничто на Земле не сможет разлучить их', 'Titanic.webp', 'Titanic.mp4', 'Titanic.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Начало', 'Inception.webp', '2010', '2 часа 28 минут', '12',
         'Кобб – талантливый вор, лучший из лучших в опасном искусстве извлечения: он крадет ценные секреты из ' ||
         'глубин подсознания во время сна, когда человеческий разум наиболее уязвим. Редкие способности Кобба ' ||
         'сделали его ценным игроком в привычном к предательству мире промышленного шпионажа, но они же превратили ' ||
         'его в извечного беглеца и лишили всего, что он когда-либо любил.

  ' ||
         'И вот у Кобба появляется шанс исправить ошибки. Его последнее дело может вернуть все назад, но для этого ' ||
         'ему нужно совершить невозможное – инициацию. Вместо идеальной кражи Кобб и его команда спецов должны будут ' ||
         'провернуть обратное. Теперь их задача – не украсть идею, а внедрить ее. Если у них получится, это и станет ' ||
         'идеальным преступлением.

  ' ||
         'Но никакое планирование или мастерство не могут подготовить команду к встрече с опасным ' ||
         'противником, который, кажется, предугадывает каждый их ход. Врагом, увидеть которого мог бы лишь Кобб.',
         '8.7', 'Твой разум - место преступления', 'Inception.webp', 'Inception.mp4', 'InceptionTrailer.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('1+1', 'Intouchables.webp', '2011', '1 час 52 минуты', '16',
         'Пострадав в результате несчастного случая, богатый аристократ ' ||
         'Филипп нанимает в помощники человека, который менее всего подходит ' ||
         'для этой работы, – молодого жителя предместья Дрисса, только что освободившегося из тюрьмы. ' ||
         'Несмотря на то, что Филипп прикован к инвалидному креслу, Дриссу удается привнести в ' ||
         'размеренную жизнь аристократа дух приключений.', '8.8', 'Sometimes you have to reach into someone else''s world to find out what''s missing in your own', 'Intouchables.webp', 'Intouchables.mp4',
         'Intouchables.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Игра в имитацию', 'TheImitationGame.webp', '2014', '1 час 54 минуты', '16',
         'Английский математик и логик Алан Тьюринг пытается взломать' ||
         ' код немецкой шифровальной машины Enigma во время Второй мировой войны.', '7.6', 'Основано на невероятной, но реальной истории',
         'TheImitationGame.webp', 'TheImitationGame.mp4', 'TheImitationGame.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Интерстеллар', 'Interstellar.webp', '2014', '2 часа 49 минут', '16',
         'Когда засуха, пыльные бури и вымирание растений приводят человечество к продовольственному кризису, ' ||
         'коллектив исследователей и учёных отправляется сквозь червоточину (которая предположительно соединяет ' ||
         'области пространства-времени через большое расстояние) в путешествие, чтобы превзойти прежние ограничения ' ||
         'для космических путешествий человека и найти планету с подходящими для человечества условиями.',
         '8.6', 'Следующий шаг человечества станет величайшим', 'Interstellar.webp', 'Interstellar.mp4', 'Interstellar.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Зеленая миля', 'TheGreenMile_name.webp', '1999', '3 часа 9 минут', '16',
         'Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого ' ||
         'однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей ' ||
         'за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых ' ||
         'необычных обитателей блока.', '9.1', 'Пол Эджкомб не верил в чудеса. Пока не столкнулся с одним из них',
         'TheGreenMile.webp', 'TheGreenMile.mp4', 'TheGreenMile.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Мстители', 'TheAvengers_name.webp', '2012', '2 часа 17 минут', '12',
         'Локи, сводный брат Тора, возвращается, и в этот раз он не один. Земля оказывается на грани порабощения, ' ||
         'и только лучшие из лучших могут спасти человечество. Глава международной организации Щ.И.Т. Ник Фьюри ' ||
         'собирает выдающихся поборников справедливости и добра, чтобы отразить атаку. Под предводительством Капитана ' ||
         'Америки Железный Человек, Тор, Невероятный Халк, Соколиный Глаз и Чёрная Вдова вступают в войну с захватчиком.',
         '7.9', 'Avengers Assemble!',
         'TheAvengers.webp', 'TheAvengers.mp4', 'TheAvengersTrailer.mp4') RETURNING id;


  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Время', 'InTime.webp', '2011', '1 час 49 минут', '12',
         'Добро пожаловать в мир, где время стало единственной и самой твердой валютой, где люди генетически ' ||
         'запрограммированы так, что в 25 лет перестают стареть. Правда, последующие годы стоят денег. И вот ' ||
         'богатые становятся практически бессмертными, а бедные обречены сражаться за жизнь. Уилл, бунтарь из ' ||
         'гетто, несправедливо обвинен в убийстве с целью грабежа времени и теперь вынужден, захватив заложницу, ' ||
         'пуститься в бега.', '7.3', 'Живи вечно или умри, пытаясь', 'InTime.webp', 'InTime.mp4', 'InTime.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Девушка с татуировкой дракона', 'TheGirlwiththeDragonTattoo.webp', '2011', '2 часа 38 минут', '18',
         'Сорок лет назад Харриет Вангер бесследно пропала на острове, принадлежащем могущественному клану Вангер. Ее ' ||
         'тело так и не было найдено, но ее дядя убежден, что это убийство и что убийца является членом его собственной, ' ||
         'тесно сплоченной и неблагополучной семьи. Он нанимает опального журналиста Микаэля Блумквиста и татуированную ' ||
         'хакершу Лисбет Саландер для проведения расследования.', '7.7', 'Evil shall with evil be expelled',
         'TheGirlwiththeDragonTattoo.webp', 'TheGirlwiththeDragonTattoo.mp4', 'TheGirlwiththeDragonTattoo.mp4') RETURNING id;

INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
VALUES('Драйв', 'Drive.webp', '2011', '1 час 40 минут', '18',
       'Великолепный водитель – при свете дня он выполняет каскадерские трюки на съёмочных площадках Голливуда, а по ' ||
       'ночам ведет рискованную игру. Но один опасный контракт – и за его жизнь назначена награда. Теперь, чтобы ' ||
       'остаться в живых и спасти свою очаровательную соседку, он должен делать то, что умеет лучше всего – ' ||
       'виртуозно уходить от погони.', '7.3', 'Some Heroes Are Real',
       'Drive.webp', 'Drive.webm', 'Drive.webm') RETURNING id;

INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
VALUES('Живая сталь', 'RealSteel.webp', '2011', '2 часа 7 минут', '16',
       'События фильма происходят в будущем, где бокс запрещен за негуманностью и заменен боями 2000-фунтовых ' ||
       'роботов, управляемых людьми. Бывший боксер, а теперь промоутер, переметнувшийся в Робобокс, решает, что ' ||
       'наконец нашел своего чемпиона, когда ему попадается выбракованный, но очень способный робот. Одновременно ' ||
       'на жизненном пути героя возникает 11-летний парень, оказывающийся его сыном. И по мере того, как машина ' ||
       'пробивает свой путь к вершине, обретшие друг друга отец и сын учатся дружить.', '7.6',
       'Чемпионами не рождаются, их собирают', 'RealSteel.webp', 'RealSteel.webm', 'RealSteel.webm') RETURNING id;

INSERT INTO movies(name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
VALUES('Исходный код', 'SourceCode.webp', '2011', '1 час 33 минуты', '12',
       'Солдат по имени Коултер мистическим образом оказывается в теле неизвестного мужчины, погибшего в ' ||
       'железнодорожной катастрофе. Коултер вынужден переживать чужую смерть снова и снова до тех пор, пока не ' ||
       'поймет, кто – зачинщик катастрофы.', '7.8', 'Make every second count', 'SourceCode.webp',
       'SourceCode.webm', 'SourceCode.webm') RETURNING id;

  DELETE FROM movies WHERE id = 4;
  DELETE FROM movies WHERE id = 5;

  CREATE TABLE genre(
     id serial constraint genre_pk primary key,
     name varchar(255) unique
  );

  INSERT INTO genre(name) VALUES('Комедия');
  INSERT INTO genre(name) VALUES('Боевик');
  INSERT INTO genre(name) VALUES('Драма');
  INSERT INTO genre(name) VALUES('Приключения');
  INSERT INTO genre(name) VALUES('Мелодрама');
  INSERT INTO genre(name) VALUES('Фантастика');
  INSERT INTO genre(name) VALUES('Фэнтези');
  INSERT INTO genre(name) VALUES('Триллер');
  INSERT INTO genre(name) VALUES('История');
  INSERT INTO genre(name) VALUES('Детектив');
  INSERT INTO genre(name) VALUES('Криминал');
  INSERT INTO genre(name) VALUES('Семейный');

  BEGIN;
  CREATE TABLE IF NOT EXISTS movies_genre(
      id serial constraint movies_genre_pk primary key,
      movie_id int,
      genre_id int
  );

  create unique index movies_genre_uindex
      on movies_genre (movie_id, genre_id);
  COMMIT;

  INSERT INTO movies_genre(movie_id, genre_id) VALUES('1', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('1', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('1', '4');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('8', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('8', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('8', '7');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('8', '4');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('2', '5');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('2', '9');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('2', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('2', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('6', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('6', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('6', '4');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('3', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('3', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('3', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('3', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('3', '10');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('7', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('7', '11');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('9', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('9', '5');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('9', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '11');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '10');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('12', '11');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('12', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('12', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '12');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('14', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('14', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('14', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('14', '3');


  CREATE TABLE IF NOT EXISTS country(
      id smallserial constraint country_pk primary key,
      name varchar(255) unique
  );

  INSERT INTO country(name) VALUES('Россия');
  INSERT INTO country(name) VALUES('Франция');
  INSERT INTO country(name) VALUES('США');
  INSERT INTO country(name) VALUES('Великобритания');
  INSERT INTO country(name) VALUES('Канада');
  INSERT INTO country(name) VALUES('Швеция');
  INSERT INTO country(name) VALUES('Норвегия');
  INSERT INTO country(name) VALUES('Индия');
  INSERT INTO country(name) VALUES('Германия');

  BEGIN;
  CREATE TABLE movies_countries(
      id serial constraint movies_countries_pk primary key,
      movie_id int,
      country_id int
  );

  create unique index movies_countries_uindex
      on movies_countries (movie_id, country_id);
  COMMIT;

  INSERT INTO movies_countries(movie_id, country_id) VALUES('2', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('3', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('3', '4');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('1', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('6', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('6', '4');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('6', '5');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('7', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('8', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('9', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('11', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('11', '6');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('11', '7');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('12', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('13', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('13', '8');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('14', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('14', '5');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('14', '2');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('14', '9');


  CREATE TABLE IF NOT EXISTS person(
     id serial constraint person_pk primary key,
     name varchar(60) not null unique,
     photo varchar(255) not null,
     addit_photo1 varchar(255),
     addit_photo2 varchar(255),
     description varchar(1024) not null
  );

  INSERT INTO person(name, photo, addit_photo1, addit_photo2, description) VALUES('Джосс Уидон', 'JossWhedon.webp',
          'JossWhedon1.webp', 'JossWhedon2.webp', 'Родился 23 июня 1964 года. Американский режиссёр, сценарист, актёр и ' ||
          'продюсер, автор комиксов. Наибольшую известность получил в качестве создателя сериала «Баффи — ' ||
          'истребительница вампиров». Среди других его известных работ — сериалы «Ангел», «Светлячок», «Кукольный дом», ' ||
          '«Агенты „Щ. И. Т.“», фильмы «Миссия „Серенити“» (2005), «Мстители» (2012) и «Мстители: Эра Альтрона» (2015).');

  INSERT INTO person(name, photo, addit_photo1, addit_photo2, description) VALUES('Кевин Файги', 'KevinFeige.webp',
          'KevinFeige1.webp', 'KevinFeige2.webp', 'Родился 2 июня 1973 года. Американский продюсер и глава Marvel ' ||
          'Studios, известный благодаря продюсированию высокобюджетных экранизаций комиксов компании Marvel ' ||
          'Entertainment. Газета The New York Times в 2011 году назвала Файги «одним из самых могущественных ' ||
          'людей в кино».');

  INSERT INTO person(name, photo, addit_photo1, addit_photo2, description) VALUES('Роберт Дауни мл.',
          'RobertDowneyJr.webp', 'RobertDowneyJr1.webp', 'RobertDowneyJr2.webp', 'Родился 4 апреля 1965 года. ' ||
         'Американский актёр, продюсер и музыкант. Лауреат премий «Золотой глобус» (2001, 2010), BAFTA (1993), ' ||
         'премии Гильдии киноактёров США (2001) и «Сатурн» (1994, 2009, 2014, 2019), номинант на премии «Оскар» ' ||
         '(1993, 2009) и «Эмми» (2001).');

  CREATE TABLE IF NOT EXISTS position(
      id serial constraint position_pk primary key,
      name varchar(32) not null unique
  );

  INSERT INTO position(name) VALUES ('Режиссер') RETURNING id;
  INSERT INTO position(name) VALUES ('Продюсер') RETURNING id;
  INSERT INTO position(name) VALUES ('Сценарист') RETURNING id;
  INSERT INTO position(name) VALUES ('Актер') RETURNING id;


  BEGIN;
  CREATE TABLE IF NOT EXISTS movies_staff(
       id serial constraint movies_staff_pk primary key,
       movie_id int,
       person_id int,
       position_id int
  );

  create unique index movies_staff_uindex
      on movies_staff (movie_id, person_id, position_id);
  COMMIT;

  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('8', '1', '1');
  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('8', '2', '2');
  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('8', '3', '4');

EOSQL
