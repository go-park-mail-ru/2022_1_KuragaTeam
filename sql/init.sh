#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE EXTENSION hunspell_ru_ru;
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
        subscription_expires timestamp,
        likes int[]
    );

    create unique index users_email_uindex
        on users (email);
  COMMIT;

  CREATE TABLE IF NOT EXISTS movies(
       id serial constraint movies_pk primary key,
       is_movie bool not null,
       name varchar(60) not null unique,
       name_picture varchar(255) not null,
       year smallint not null,
       duration varchar(25) not null,
       age_limit smallint not null,
       description varchar(1024) not null,
       kinopoisk_rating numeric(2,1) not null,
       rating_sum int,
       rating_count int,
       tagline varchar(255) not null,
       picture varchar(255) not null,
       video varchar(255) not null,
       trailer varchar(255) not null
  );

  CREATE TABLE IF NOT EXISTS rating(
      id serial constraint rating_pk primary key,
      movie_id int,
      user_id int,
      rating smallint
  );

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Аватар', 'Avatar.webp', true, '2009', '2 часа 42 минуты', '12',
         'Бывший морпех Джейк Салли прикован к инвалидному креслу. Несмотря на немощное тело, ' ||
         'Джейк в душе по-прежнему остается воином. Он получает задание совершить путешествие в ' ||
         'несколько световых лет к базе землян на планете Пандора, где корпорации добывают ' ||
         'редкий минерал, имеющий огромное значение для выхода Земли из энергетического кризиса.',
         '7.9', 'Это новый мир', 'Avatar.webp', 'Avatar.mp4', 'Avatar.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Титаник', 'Titanic.webp', true, '1997', '3 часа 14 минут', '12',
         'В первом и последнем плавании шикарного «Титаника» встречаются двое. Пассажир нижней палубы ' ||
         'Джек выиграл билет в карты, а богатая наследница Роза отправляется в Америку, чтобы выйти ' ||
         'замуж по расчёту. Чувства молодых людей только успевают расцвести, и даже не классовые различия ' ||
         'создадут испытания влюблённым, а айсберг, вставший на пути считавшегося непотопляемым лайнера.',
         '8.4', 'Ничто на Земле не сможет разлучить их', 'Titanic.webp', 'Titanic.mp4', 'Titanic.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Начало', 'Inception.webp', true, '2010', '2 часа 28 минут', '12',
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

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('1+1', 'Intouchables.webp', true, '2011', '1 час 52 минуты', '16',
         'Пострадав в результате несчастного случая, богатый аристократ ' ||
         'Филипп нанимает в помощники человека, который менее всего подходит ' ||
         'для этой работы, – молодого жителя предместья Дрисса, только что освободившегося из тюрьмы. ' ||
         'Несмотря на то, что Филипп прикован к инвалидному креслу, Дриссу удается привнести в ' ||
         'размеренную жизнь аристократа дух приключений.', '8.8', 'Sometimes you have to reach into someone else''s world to find out what''s missing in your own', 'Intouchables.webp', 'Intouchables.mp4',
         'Intouchables.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Игра в имитацию', 'TheImitationGame.webp', true, '2014', '1 час 54 минуты', '16',
         'Английский математик и логик Алан Тьюринг пытается взломать' ||
         ' код немецкой шифровальной машины Enigma во время Второй мировой войны.', '7.6', 'Основано на невероятной, но реальной истории',
         'TheImitationGame.webp', 'TheImitationGame.mp4', 'TheImitationGame.mp4') RETURNING id;

  DELETE FROM movies WHERE id = 4;
  DELETE FROM movies WHERE id = 5;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Интерстеллар', 'Interstellar.webp', true, '2014', '2 часа 49 минут', '16',
         'Когда засуха, пыльные бури и вымирание растений приводят человечество к продовольственному кризису, ' ||
         'коллектив исследователей и учёных отправляется сквозь червоточину (которая предположительно соединяет ' ||
         'области пространства-времени через большое расстояние) в путешествие, чтобы превзойти прежние ограничения ' ||
         'для космических путешествий человека и найти планету с подходящими для человечества условиями.',
         '8.6', 'Следующий шаг человечества станет величайшим', 'Interstellar.webp', 'Interstellar.mp4', 'Interstellar.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Зеленая миля', 'TheGreenMile.webp', true, '1999', '3 часа 9 минут', '16',
         'Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого ' ||
         'однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей ' ||
         'за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых ' ||
         'необычных обитателей блока.', '9.1', 'Пол Эджкомб не верил в чудеса. Пока не столкнулся с одним из них',
         'TheGreenMile.webp', 'TheGreenMile.mp4', 'TheGreenMile.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Мстители', 'TheAvengers.webp', true, '2012', '2 часа 17 минут', '12',
         'Локи, сводный брат Тора, возвращается, и в этот раз он не один. Земля оказывается на грани порабощения, ' ||
         'и только лучшие из лучших могут спасти человечество. Глава международной организации Щ.И.Т. Ник Фьюри ' ||
         'собирает выдающихся поборников справедливости и добра, чтобы отразить атаку. Под предводительством Капитана ' ||
         'Америки Железный Человек, Тор, Невероятный Халк, Соколиный Глаз и Чёрная Вдова вступают в войну с захватчиком.',
         '7.9', 'Avengers Assemble!',
         'TheAvengers.webp', 'TheAvengers.mp4', 'TheAvengersTrailer.mp4') RETURNING id;


  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Время', 'InTime.webp', true, '2011', '1 час 49 минут', '12',
         'Добро пожаловать в мир, где время стало единственной и самой твердой валютой, где люди генетически ' ||
         'запрограммированы так, что в 25 лет перестают стареть. Правда, последующие годы стоят денег. И вот ' ||
         'богатые становятся практически бессмертными, а бедные обречены сражаться за жизнь. Уилл, бунтарь из ' ||
         'гетто, несправедливо обвинен в убийстве с целью грабежа времени и теперь вынужден, захватив заложницу, ' ||
         'пуститься в бега.', '7.3', 'Живи вечно или умри, пытаясь', 'InTime.webp', 'InTime.mp4', 'InTime.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Девушка с татуировкой дракона', 'TheGirlwiththeDragonTattoo.webp', true, '2011', '2 часа 38 минут', '18',
         'Сорок лет назад Харриет Вангер бесследно пропала на острове, принадлежащем могущественному клану Вангер. Ее ' ||
         'тело так и не было найдено, но ее дядя убежден, что это убийство и что убийца является членом его собственной, ' ||
         'тесно сплоченной и неблагополучной семьи. Он нанимает опального журналиста Микаэля Блумквиста и татуированную ' ||
         'хакершу Лисбет Саландер для проведения расследования.', '7.7', 'Evil shall with evil be expelled',
         'TheGirlwiththeDragonTattoo.webp', 'TheGirlwiththeDragonTattoo.mp4', 'TheGirlwiththeDragonTattoo.mp4') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Драйв', 'Drive.webp', true, '2011', '1 час 40 минут', '18',
         'Великолепный водитель – при свете дня он выполняет каскадерские трюки на съёмочных площадках Голливуда, а по ' ||
         'ночам ведет рискованную игру. Но один опасный контракт – и за его жизнь назначена награда. Теперь, чтобы ' ||
         'остаться в живых и спасти свою очаровательную соседку, он должен делать то, что умеет лучше всего – ' ||
         'виртуозно уходить от погони.', '7.3', 'Some Heroes Are Real',
         'Drive.webp', 'Drive.webm', 'Drive.webm') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Живая сталь', 'RealSteel.webp', true, '2011', '2 часа 7 минут', '16',
         'События фильма происходят в будущем, где бокс запрещен за негуманностью и заменен боями 2000-фунтовых ' ||
         'роботов, управляемых людьми. Бывший боксер, а теперь промоутер, переметнувшийся в Робобокс, решает, что ' ||
         'наконец нашел своего чемпиона, когда ему попадается выбракованный, но очень способный робот. Одновременно ' ||
         'на жизненном пути героя возникает 11-летний парень, оказывающийся его сыном. И по мере того, как машина ' ||
         'пробивает свой путь к вершине, обретшие друг друга отец и сын учатся дружить.', '7.6',
         'Чемпионами не рождаются, их собирают', 'RealSteel.webp', 'RealSteel.webm', 'RealSteel.webm') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Исходный код', 'SourceCode.webp', true, '2011', '1 час 33 минуты', '12',
         'Солдат по имени Коултер мистическим образом оказывается в теле неизвестного мужчины, погибшего в ' ||
         'железнодорожной катастрофе. Коултер вынужден переживать чужую смерть снова и снова до тех пор, пока не ' ||
         'поймет, кто – зачинщик катастрофы.', '7.8', 'Make every second count', 'SourceCode.webp',
         'SourceCode.webm', 'SourceCode.webm') RETURNING id;

  INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  VALUES('Доктор Хаус', 'HouseMD.webp', false, '2004', '8 сезонов', '16',
         'Сериал рассказывает о команде врачей, которые должны правильно поставить диагноз пациенту и спасти его. ' ||
         'Возглавляет команду доктор Грегори Хаус, который ходит с тростью после того, как его мышечный инфаркт в ' ||
         'правой ноге слишком поздно правильно диагностировали. Как врач Хаус просто гений, но сам не отличается ' ||
         'проникновенностью в общении с больными и с удовольствием избегает их, если только есть возможность. Он сам ' ||
         'всё время проводит в борьбе с собственной болью, а трость в его руке только подчеркивает его жесткую, ' ||
         'ядовитую манеру общения. Порой его поведение можно назвать почти бесчеловечным, и при этом он прекрасный ' ||
         'врач, обладающий нетипичным умом и безупречным инстинктом, что снискало ему глубокое уважение. Будучи ' ||
         'инфекционистом, он ещё и замечательный диагност, который любит разгадывать медицинские загадки, чтобы ' ||
         'спасти кому-то жизнь. Если бы все было по его воле, то Хаус лечил бы больных не выходя из своего кабинета.',
         '8.8', 'Genius has side effects', 'HouseMD.webp', 'HouseMD.webm', 'HouseMD.webm') RETURNING id;

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

  INSERT INTO movies_genre(movie_id, genre_id) VALUES('1', '6');
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
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('10', '10');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('10', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('10', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('10', '11');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('11', '11');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('12', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('12', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('12', '12');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '6');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '2');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '8');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('13', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('14', '3');
  INSERT INTO movies_genre(movie_id, genre_id) VALUES('14', '10');


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
  INSERT INTO movies_countries(movie_id, country_id) VALUES('1', '4');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('6', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('6', '4');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('6', '5');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('7', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('8', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('9', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('10', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('10', '6');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('10', '7');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('11', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('12', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('12', '8');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('13', '3');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('13', '5');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('13', '2');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('13', '9');
  INSERT INTO movies_countries(movie_id, country_id) VALUES('14', '3');

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

  INSERT INTO person(name, photo, addit_photo1, addit_photo2, description) VALUES('Кристофер Нолан',
          'KristopherNolan.webp', 'KristopherNolan1.webp', 'KristopherNolan2.webp', 'Родился 30 июля 1970 года. Британский и американский ' ||
         'кинорежиссёр, сценарист и продюсер. Является одним из самых кассовых режиссёров в истории, а также одним ' ||
         'из самых известных и влиятельных кинематографистов своего времени. Фильмы Нолана основываются на философских, ' ||
        'социологических и этических концепциях, исследуют человеческую мораль, конструирование времени и ' ||
        'податливую природу памяти и личной идентичности.');

  INSERT INTO person(name, photo, addit_photo1, addit_photo2, description) VALUES('Мэттью МакКонахи',
      'MatthewMcConaughey.webp', 'MatthewMcConaughey1.webp', 'MatthewMcConaughey2.webp', 'Родился 4 ноября 1969 года. ' ||
     'Американский актёр и продюсер. Поначалу зарекомендовав себя как актёр, в основном, комедийного амплуа, во втором ' ||
     'десятилетии XXI века Макконахи перешёл к крупным драматическим ролям, удостоившись ряда наград и положительных ' ||
     'отзывов от кинопрессы за картины «Линкольн для адвоката», «Мад», «Киллер Джо», «Далласский клуб покупателей», ' ||
     '«Супер Майк», «Интерстеллар» и «Джентльмены».');


  INSERT INTO person(name, photo, addit_photo1, addit_photo2, description) VALUES('Джеймс Кэмерон',
      'JamesCameron.webp', 'JamesCameron1.webp', 'JamesCameron2.webp', 'Родился 16 августа 1954. Канадский кинорежиссёр, ' ||
     'наиболее известный по созданию научно-фантастических и эпических фильмов. Кэмерон впервые добился признания ' ||
      'за режиссуру фильма «Терминатор». Его самые высокобюджетные киноленты - «Титаник» и «Аватар», причём первая ' ||
     'удостоилась премии «Оскар» в номинациях «Лучший фильм», «Лучший режиссёр» и «Лучший монтаж».');


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
  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('6', '4', '1');
  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('6', '5', '4');
  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('1', '6', '1');
  INSERT INTO movies_staff(movie_id, person_id, position_id) VALUES('2', '6', '1');


  CREATE TABLE IF NOT EXISTS seasons(
      id serial constraint seasons_pk primary key,
      number smallint not null unique,
      movie_id int not null
  );

  DO \$\$DECLARE i record;
  BEGIN
    FOR i IN 1..8 LOOP
      INSERT INTO seasons(number, movie_id) VALUES (i, '14');
    END LOOP;
  END\$\$;

  CREATE TABLE IF NOT EXISTS episode (
       id serial constraint episode_pk primary key,
       name varchar(255) not null,
       number smallint not null,
       description varchar(1024) not null,
       season_id int not null,
       video varchar(255) not null,
       photo varchar(255) not null
  );

  create unique index episodes_uindex
      on episode (season_id, number);

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Пилотная серия', '1', 'У воспитательницы детского сада Ребекки Адлер началась афазия во время урока, а потом у неё случился припадок. Её направляют в госпиталь Принстон-Плэйнсборо, где доктор Хаус и его команда (к которой недавно присоединился новый врач Эрик Форман) пытаются поставить ей диагноз. На приём в клинику к доктору Хаусу приходит «оранжевый пациент».', '1', 'HouseMD_1_1.mp4', 'HouseMD_1_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Отцовство', '2', '16-летний школьник, страдающий от ночных кошмаров и галлюцинаций, получает травму во время игры в лакросс. В клинике Хаус встречается с матерью, которая не верит в вакцинацию, и любителем судебных тяжб.', '1', 'HouseMD_1_2.mp4', 'HouseMD_1_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Бритва Оккама', '3', 'После бурной ночи со своей подружкой студент колледжа неожиданно теряет сознание, заставляя Хауса и его команду ломать голову над причиной стремительного падения лейкоцитов, а в клинике Хауса ждут женщина, у которой заболела нога после 6-мильной пробежки, и мальчик с mp3-плеером в прямой кишке.', '1', 'HouseMD_1_3.mp4', 'HouseMD_1_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Материнство', '4', 'В родильном отделении больницы началась эпидемия неизвестной болезни, и доктор Хаус требует объявить карантин. На этот раз в клинику приходит женщина с «паразитом».', '1', 'HouseMD_1_4.mp4', 'HouseMD_1_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Будь ты проклят, если сделаешь это', '5', 'На руках у монахини появились стигматы, но Хаус не верит в божественные знаки и пытается найти причины её болезни. Рождественский эпизод.', '1', 'HouseMD_1_5.mp4', 'HouseMD_1_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Метод Сократа', '6', 'Люси Пальмеро, страдающая от шизофрении, попадает в госпиталь с диагнозом «тромбоз глубоких вен». Её сыну удается уговорить доктора Хауса заняться этим случаем.', '1', 'HouseMD_1_6.mp4', 'HouseMD_1_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Верность', '7', 'Женщина попадает в госпиталь с симптомами сонной болезни. Но ни женщина, ни её муж никогда не были в Африке. Хаус не может начать опасное для жизни пациентки лечение, пока не подтвердит её диагноз.', '1', 'HouseMD_1_7.mp4', 'HouseMD_1_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Яд', '8', 'У школьника Мэтта Дейвиса начинаются галлюцинации на экзамене по математике, а затем он теряет сознание. Хаус решает, что это симптомы отравления каким-то таинственным ядом. В это время у другого школьника, который совершенно не связан с Мэттом, появляются те же симптомы. В клинику приходит счастливая пожилая женщина.', '1', 'HouseMD_1_8.mp4', 'HouseMD_1_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Отказ от реанимации', '9', 'Знаменитый парализованный джазмен начинает задыхаться во время музыкальной записи. Работу команды сильно осложняет то, что пациент подписал отказ от реанимации. В клинике появляется мужчина, который отрицает, что страдает диабетом.', '1', 'HouseMD_1_9.mp4', 'HouseMD_1_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Истории', '10', 'Бездомная женщина теряет сознание на нелегальной вечеринке с наркотиками. Форман уверен, что она не заслужила лечения, и не соглашается с Уилсоном, который просит его заняться этим случаем. В клинику приходит многодетная мать. Две студентки учатся у Хауса собирать анамнез.', '1', 'HouseMD_1_10.mp4', 'HouseMD_1_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Детоксикация', '11', 'Девушка попадает в аварию, после того как замечает, что её друг кашляет кровью. Кровотечение у парня не останавливается и после аварии. Кадди заключает с Хаусом сделку: если он сможет обойтись без викодина неделю, она освободит его от пациентов клиники на целый месяц. Форман и Кэмерон боятся, что абстинентный синдром мешает Хаусу мыслить здраво и может стать причиной смерти пациента.', '1', 'HouseMD_1_11.mp4', 'HouseMD_1_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Спортивная медицина', '12', 'Бывший наркоман, питчер Хэнк Уигген неожиданно ломает руку. Хаус считает, что Хэнк лжет, что не употребляет стероиды. После отказа Уилсона Хаус понимает, что ему не с кем пойти на шоу грузовиков. В клинику приходит женщина с болью в ноге, мужчина, который не может вынуть контактные линзы, дантист и подросток, страдающий от похмелья.', '1', 'HouseMD_1_12.mp4', 'HouseMD_1_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Проклятый', '13', 'Доска для спиритических сеансов предсказывает мальчику смерть, и скоро та действительно стоит у него на пороге. Отец Чейза помогает Хаусу поставить диагноз. В клинику приходит пациент с онемевшими пальцами.', '1', 'HouseMD_1_13.mp4', 'HouseMD_1_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Контроль', '14', 'Красивую и успешную женщину парализует во время совещания. Тайны, которые она скрывает, заставляют Хауса рискнуть лицензией. Миллионер Эдвард Воглер становится членом правления, после того как пожертвовал госпиталю 100 миллионов долларов. Он хочет превратить госпиталь в успешное коммерческое предприятие, и первое, что ему не нравится, — дорогостоящее отделение диагностической медицины. В клинику приходит мальчик, отец которого онемел после операции на колене.', '1', 'HouseMD_1_14.mp4', 'HouseMD_1_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Законы мафии', '15', 'Информатор мафии впадает в кому. Он действительно болен или симулирует, чтобы не выступать свидетелем на судебном разбирательстве? В клинику несколько раз приходит подросток с маленьким братом, который постоянно что-то засовывает в нос.', '1', 'HouseMD_1_15.mp4', 'HouseMD_1_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Бремя', '16', 'У десятилетней девочки, страдающей от излишнего веса, происходит сердечный приступ. Мать девочки хочет, чтобы команда доктора Хауса не обращала внимание на вес, когда будет ставить диагноз. Воглер настаивает, чтобы Хаус уволил одного из членов команды. Хаус делает выбор, но Воглер отклоняет его. В клинику приходит пациент с инфекцией из-за пирсинга мошонки и женщина с огромной опухолью, которую не желает удалять.', '1', 'HouseMD_1_16.mp4', 'HouseMD_1_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Образец для подражания', '17', 'Сенатор Гари Райт, чернокожий кандидат в президенты, теряет сознание во время предвыборной кампании. Хаус подозревает, что у него СПИД, и постоянно пытается уличить его во лжи. Воглер требует у Хауса произнести речь на презентации нового лекарства принадлежащей ему фармацевтической компании.', '1', 'HouseMD_1_17.mp4', 'HouseMD_1_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Дети и вода в ванне', '18', 'У женщины рак легких, но лечение невозможно, пока женщина беременна. Хаус готов на шантаж и обман, чтобы спасти пациентку. А Воглер намерен добиться увольнения Хауса.', '1', 'HouseMD_1_18.mp4', 'HouseMD_1_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Дети', '19', 'В госпитале эпидемия менингита, и Кадди дает час Хаусу и его команде, чтобы поставить диагноз 12-летней девочке, чьи симптомы кажутся Хаусу необычными. Хаус пытается вернуть Кэмерон в команду.', '1', 'HouseMD_1_19.mp4', 'HouseMD_1_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Любовь зла', '20', 'Хаус грубо разговаривает с больным, и тот падает к его ногам. Чтобы остановить непрекращающиеся приступы, команде придётся раскрыть секреты странного образа жизни пациента и его загадочной подруги. Госпиталь взбудоражен слухами о свидании Хауса и Кэмерон. В клинику приходит пожилая пара, у которой проблемы с виагрой.', '1', 'HouseMD_1_20.mp4', 'HouseMD_1_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Три истории', '21', 'К Хаусу в госпиталь приходит его бывшая подруга Стейси с просьбой вылечить мужа. Хаус — её последняя надежда, потому что никто не верит в неуловимую болезнь Марка. А Кадди просит Хауса заменить заболевшего лектора и выступить перед студентами. Хаус рассказывает о трёх пациентах, которые попали в госпиталь из-за болей в ноге, и предлагает поставить им диагноз.', '1', 'HouseMD_1_21.mp4', 'HouseMD_1_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Медовый месяц', '22', 'Чтобы доставить Марка, мужа бывшей подруги, в госпиталь Принстон-Плейнсборо, Хаусу приходится усыпить его. Стейси уверена, что с Марком что-то не так, но единственный симптом Марка — боли в животе — не похож на симптом опасной болезни. А сам Марк считает, что Хаус пытается убить его. Хаус собирается уйти в отставку, но Кадди мешает этому, чем очень помогает лучшему доктору больницы.', '1', 'HouseMD_1_22.mp4', 'HouseMD_1_22.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Смирение', '1', 'Хауса приглашают в тюрьму, чтобы он вылечил приговорённого к смертной казни. Стейси удаётся добиться перевода заключённого в госпиталь. Кэмерон вынуждена сказать молодой женщине, что у неё неоперабельный рак.', '2', 'HouseMD_2_1.mp4', 'HouseMD_2_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Аутопсия', '2', 'У маленькой 9-летней девочки (Саша Питерс) рак, но к её проблемам добавляются галлюцинации.', '2', 'HouseMD_2_2.mp4', 'HouseMD_2_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Шалтай-Болтай', '3', 'После того как рабочий падает с крыши дома доктора Кадди, у него появляются страшные симптомы. Чтобы остановить инфекцию, пациенту нужна ампутация.', '2', 'HouseMD_2_3.mp4', 'HouseMD_2_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Туберкулёз или не туберкулёз?', '4', 'Знаменитый врач, работающий в Африке, теряет сознание во время поездки по США. Все, кроме Хауса, уверены, что у него туберкулёз.', '2', 'HouseMD_2_4.mp4', 'HouseMD_2_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Папенькин сынок', '5', 'Во время выпускной вечеринки студент Принстонского университета начинает испытывать боль, напоминающую удары током. К Хаусу приехали родители.', '2', 'HouseMD_2_5.mp4', 'HouseMD_2_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Крутой поворот', '6', 'Известный велосипедист падает во время велогонки. Он признается в употреблении незаконных препаратов, но, похоже, его болезнь никак не связана с ними. Стейси и Хаус продолжают ссориться.', '2', 'HouseMD_2_6.mp4', 'HouseMD_2_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Охота', '7', 'Кевин, гомосексуал, больной СПИДом, преследует Хауса, требуя, чтобы тот поставил ему диагноз. А Хаусу удаётся украсть историю болезни Стейси. Кэмерон контактирует с заражённой кровью Кевина, и вынуждена пройти тест на ВИЧ.', '2', 'HouseMD_2_7.mp4', 'HouseMD_2_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Ошибка', '8', 'Матери двух маленьких девочек становится плохо на детском утреннике. Чейз осматривает пациентку и убеждает её, что с ней всё в порядке. Но вскоре понимает, что его диагноз был неверен, однако женщину уже не спасти. Чейз считает себя виновным в её смерти, но Хаус думает иначе и пытается защитить его перед дисциплинарным комитетом.', '2', 'HouseMD_2_8.mp4', 'HouseMD_2_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Обман', '9', 'Прямо к ногам Хауса, который играет на внеипподромном тотализаторе, падает женщина. Ему приходится оставить скачки и отправиться вместе с ней в госпиталь. Но все в госпитале убеждены, что Аника — симулянтка, и только Хаус считает, что Аника действительно больна. Рождественский эпизод.', '2', 'HouseMD_2_9.mp4', 'HouseMD_2_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Трудности перевода', '10', 'Журналист Флетчер Стоун неожиданно потерял дар речи; бессмысленный набор слов, которые он произносит, не может понять даже его жена. Команде доктора Хауса приходится работать без самого Хауса: он вместе со Стейси вынужден из-за непогоды сидеть в аэропорту Балтимора. Но у Хауса есть телефон, а вместо любимого фломастера можно использовать помаду Стейси.', '2', 'HouseMD_2_10.mp4', 'HouseMD_2_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Необходимые знания', '11', 'Пытаясь поставить диагноз домохозяйке, страдающей от загадочной болезни, разрушающей её организм, Хаус понимает, что вся её внешне идеальная жизнь основана на лжи. Кэмерон нервничает из-за анализа на ВИЧ, а Стейси принимает решение.', '2', 'HouseMD_2_11.mp4', 'HouseMD_2_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Отвлекающий раздражитель', '12', 'Катаясь вместе с отцом на квадроцикле, Адам попадает в страшную аварию. Команда Хауса пытается понять, почему Адам потерял управление, но страшные ожоги на теле мальчика не дают им сделать нужные анализы. А сам Хаус мучается от сильной мигрени: чтобы доказать, что лекарство, изобретённое его бывшим сокурсником, не работает, он опробовал его на себе.', '2', 'HouseMD_2_12.mp4', 'HouseMD_2_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Внешность обманчива', '13', 'Алекс, 15-летняя топ-модель, после необъяснимой вспышки ярости падает на подиум. Вначале кажется, что все дело в психологическом состоянии молодой девушки, которая употребляет героин и антидепрессанты и ведёт беспорядочные связи. Хаус мучается от боли в ноге и пытается уговорить Кадди дать ему морфий.', '2', 'HouseMD_2_13.mp4', 'HouseMD_2_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Секс убивает', '14', '65-летнему пациенту нужна пересадка сердца. Хаусу удаётся уговорить мужа погибшей в автомобильной аварии женщины отдать так необходимое сердце. Но комиссия по трансплантации органов отказалась от этого сердца не просто так: мёртвая женщина была больна, и, чтобы провести операцию, команде доктора Хауса нужно сначала как можно быстрее вылечить её. Уилсон переезжает к Хаусу.', '2', 'HouseMD_2_14.mp4', 'HouseMD_2_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Невежество', '15', 'Мужчина начинает задыхаться во время сексуальных игр с женой, и Хаус подозревает, что их брак не такой безоблачный, как кажется.', '2', 'HouseMD_2_15.mp4', 'HouseMD_2_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Безопасность', '16', '16-летняя Мелинда на глазах её матери и друга неожиданно начинает задыхаться. Девушка страдает сильной аллергией и недавно перенесла пересадку сердца, но что могло вызвать приступ в стерильной комнате, в которой она находилась?', '2', 'HouseMD_2_16.mp4', 'HouseMD_2_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Ва-банк', '17', 'Команде Хауса приходится покинуть вечеринку, чтобы заняться диагностикой маленького мальчика с болями в животе. Скоро они понимают, что на самом деле Хаус пытается поставить диагноз давно умершей пациентке, загадочная болезнь которой многие годы не давала ему покоя.', '2', 'HouseMD_2_17.mp4', 'HouseMD_2_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Не будите спящую собаку', '18', 'Ханне, которая мучилась от бессонницы 10 дней, нужна пересадка печени, но согласится ли её подруга Макс стать донором? А команда Хауса занята этическими проблемами: Кэмерон обвиняет Формана в плагиате.', '2', 'HouseMD_2_18.mp4', 'HouseMD_2_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Хаус против Бога', '19', 'Хаус объявляет войну Богу после того как 15-летний проповедник «исцеляет» раковую больную. Между тем состояние самого подростка оставляет желать лучшего.', '2', 'HouseMD_2_19.mp4', 'HouseMD_2_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Эйфория. Часть 1', '20', 'Полицейский бросает погоню за преступником и начинает безумно смеяться. Вскоре симптомы становятся гораздо более страшными. Когда Хаус понимает, что Форман заразился, дело принимает критический оборот.', '2', 'HouseMD_2_20.mp4', 'HouseMD_2_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Эйфория. Часть 2', '21', 'Полицейский умер, а Форман смертельно напуган и готов на опасную процедуру, Хаус против.', '2', 'HouseMD_2_21.mp4', 'HouseMD_2_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Навечно', '22', 'По дороге на работу у мужчины начинается рвота, он возвращается домой и обнаруживает жену, которая чуть не утопила в ванне своего ребёнка во время приступа.', '2', 'HouseMD_2_22.mp4', 'HouseMD_2_22.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Кто твой папочка?', '23', 'Дилан Крэндалл, приятель Хауса, привозит в госпиталь свою дочь, которая страдает от галлюцинаций после урагана Катрина. Хаус уверен, что девочка обманывает Дилана. А Кадди озабочена поиском донора спермы для своего ребёнка.', '2', 'HouseMD_2_23.mp4', 'HouseMD_2_23.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Без причины', '24', 'В клинику приходит пациент с ужасно распухшим языком, и Хаус, придя в восторг от его забавного вида, поручает своей команде заняться этим случаем. Но, похоже, в этот раз Хаусу не суждено поставить диагноз: в его кабинет вбегает муж одной из его прежних пациенток и дважды стреляет в него. Очнувшись в больничной палате, Хаус пытается продолжить спасение пациента, симптомы которого приобретают всё более ужасный вид, но обнаруживает, что не может отличить реальность от собственного бреда.', '2', 'HouseMD_2_24.mp4', 'HouseMD_2_24.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Смысл', '1', 'Неизлечимо больной и парализованный Ричард направляет свою инвалидную коляску в бассейн. Все убеждены, что это несчастный случай, и только Хаус, думает, что он пытался покончить с собой. Хаус, который только что вернулся в госпиталь после операции и больше не мучается от болей в ноге, считает, что в его действиях скрывается какой-то иной смысл. Молодую девушку Карен неожиданно парализовало во время занятий йогой. На этот раз все считают, что Карен серьёзно больна, и только Хаус думает, что нет.', '3', 'HouseMD_3_1.mp4', 'HouseMD_3_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Каин и Авель', '2', 'Клэнси, 7-летний мальчик, который уверен, что его похищали пришельцы, попадает в госпиталь с ректальным кровотечением. Странный металлический предмет, который Чейз находит у него в шее, наводит команду Хауса на мысль, что фантазии мальчика — не такие уж и фантазии. Боли Хауса становятся сильнее, и Кадди решает сознаться во лжи: она скрывала от Хауса, что его последний пациент выздоровел.', '3', 'HouseMD_3_2.mp4', 'HouseMD_3_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Официальное согласие', '3', 'Знаменитый врач Эзра Пауэлл после бесплодных попыток поставить ему диагноз требует от Хауса помочь ему покончить с жизнью. А 17-летней дочери другого пациента понравился Хаус.', '3', 'HouseMD_3_3.mp4', 'HouseMD_3_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Линии на песке', '4', '10-летний мальчик, страдающий от аутизма, постоянно кричит, но не может объяснить, почему. Хаус не желает работать в своем кабинете, пока Кадди не вернёт туда залитый его кровью старый ковёр. Но так же, как и мальчик, он не может объяснить, зачем это ему понадобилось.', '3', 'HouseMD_3_4.mp4', 'HouseMD_3_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Любовь зла', '5', 'В госпиталь попадает девушка с болями в животе, а затем и её муж. Команда предполагает, что их заболевания могут быть связаны. Не зная об опасных последствиях, Хаус пренебрежительно относится к одному из пациентов клиники.', '3', 'HouseMD_3_5.mp4', 'HouseMD_3_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Будь, что будет', '6', 'Невероятно толстый пациент (220 кг) найден в коме. После того как он приходит в себя в госпитале, он требует, чтобы команда Хауса поставила диагноз, не связанный с его весом. Детектив Триттер объявляет Хаусу вендетту.', '3', 'HouseMD_3_6.mp4', 'HouseMD_3_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Сын коматозника', '7', 'Хаусу удается на время разбудить больного, долгие годы пролежавшего в коме, для того чтобы тот помог поставить диагноз собственному сыну. А Триттер пытается поссорить членов команды Хауса.', '3', 'HouseMD_3_7.mp4', 'HouseMD_3_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Из огня да в полымя', '8', '18-летний подросток, единственная опора младшего брата и сестры, попадает в госпиталь с сердечным приступом и рвотой. Хаус превращает лечение в игру, предлагая Кэмерон, Форману и Чейзу самим поставить диагноз. Однако игра быстро заканчивается, а поиск диагноза только начинается. А детектив Триттер пытается сделать существование Уилсона невыносимым.', '3', 'HouseMD_3_8.mp4', 'HouseMD_3_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('В поисках Иуды', '9', 'Разведённые родители маленькой Элис постоянно ссорятся, и Хаусу приходится обратиться в суд, чтобы начать лечение девочки. У Хауса кончился викодин, но никто не хочет выписывать ему новые рецепты.', '3', 'HouseMD_3_9.mp4', 'HouseMD_3_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Крошечное весёлое Рождество', '10', 'Хаус заинтересован пациенткой-карликом, недавно перенесшей коллапс лёгкого. Уилсон заключает сделку с Триттером, и через три дня Хаус должен начать реабилитационный курс, либо ему придётся сесть в тюрьму. Кадди шантажирует Хауса интересной пациенткой и викодином, заставляя пойти на сделку. Рождественский эпизод.', '3', 'HouseMD_3_10.mp4', 'HouseMD_3_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Слова и дела', '11', 'Пожарный доставлен в госпиталь с необычно высокой температурой и признаками дезориентации, а Хаус должен предстать перед судом — он обвиняется в незаконном хранении наркотиков.', '3', 'HouseMD_3_11.mp4', 'HouseMD_3_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Один день — одна комната', '12', 'Кадди пытается превратить работу Хауса в клинике в игру, но Хаус встречает там Еву, недавно пережившую изнасилование, и игры заканчиваются. А к Кэмерон попадает бездомный со сложной судьбой.', '3', 'HouseMD_3_12.mp4', 'HouseMD_3_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Иголка в стоге сена', '13', 'Команда Хауса сталкивается с множеством проблем, общаясь с семьей подростка-цыгана, попавшего в госпиталь с кровотечением и проблемами дыхания. А Хаус занят вызовом, который бросила ему Кадди.', '3', 'HouseMD_3_13.mp4', 'HouseMD_3_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Бесчувственность', '14', 'В День святого Валентина девушка, неспособная чувствовать боль, попадает в автомобильную аварию. Если бы не высокая температура, можно было бы предположить, что она почти не пострадала.', '3', 'HouseMD_3_14.mp4', 'HouseMD_3_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Недоумок', '15', 'Гениальный пианист неожиданно теряет способность играть. Но команда Хауса никак не может сосредоточиться на пациенте: они заподозрили, что у Хауса рак.', '3', 'HouseMD_3_15.mp4', 'HouseMD_3_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Совершенно секретно', '16', 'Морской пехотинец, вернувшийся из Ирака, страдает от синдрома Войны в заливе. А Хауса мучают странные сны и неспособность мочиться.', '3', 'HouseMD_3_16.mp4', 'HouseMD_3_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Положение плода', '17', 'У беременной женщины инсульт. Когда у неё начинают отказывать почки, Хаус рекомендует ей аборт, но доктор Кадди одержима желанием спасти ребёнка во что бы то ни стало.', '3', 'HouseMD_3_17.mp4', 'HouseMD_3_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('В воздухе', '18', 'Хаус и Кадди возвращаются с конференции в Сингапуре. На борту самолёта таинственная болезнь одного из пассажиров грозит превратиться в эпидемию, а Уилсон и команда Хауса в госпитале Принстон-Плейнсборо пытаются поставить диагноз пожилой женщине, страдающей от постоянных припадков.', '3', 'HouseMD_3_18.mp4', 'HouseMD_3_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Веди себя по-взрослому', '19', 'У 6-летней девочки симптомы совсем не детской болезни. Пока Хаус пытается поставить ей диагноз, симптомы появляются и у её брата.', '3', 'HouseMD_3_19.mp4', 'HouseMD_3_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Воспитание щенков', '20', 'Девушка неожиданно теряет способность принимать решения. Пока Хаус и его команда пытаются спасти её, у Формана появляются личные мотивы.', '3', 'HouseMD_3_20.mp4', 'HouseMD_3_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Семья', '21', '14-летний Ник страдает от лейкемии, и его последняя надежда — трансплантация костного мозга, донором которого должен был стать его младший брат Мэтт. Но Мэтт неожиданно начинает чихать, и Хаусу и его команде нужно как можно скорее выяснить, что случилось с Мэттом, или Мэтт и его брат умрут. Хаус между тем ведёт сражение с псом Уилсона Гектором, а Форман не может забыть об ошибке, которую сделал на прошлой неделе.', '3', 'HouseMD_3_21.mp4', 'HouseMD_3_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Увольнение', '22', '19-летняя студентка начинает кашлять кровью во время занятий по каратэ. Кажется, что Хаус готов убить пациентку, чтобы поставить ей диагноз. Форман заявляет о своем увольнении, а Хаус решает провести эксперимент над Уилсоном, не подозревая, что Уилсон замыслил то же самое.', '3', 'HouseMD_3_22.mp4', 'HouseMD_3_22.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Гадёныш', '23', 'Натан Харрисон, 16-летний шахматный гений, выиграв партию, неожиданно избивает своего соперника. Отвратительный характер пациента заставляет команду Хауса предположить, что болезнь влияет на его психику. Форман подозревает Хауса в попытке сорвать собеседование на новую работу.', '3', 'HouseMD_3_23.mp4', 'HouseMD_3_23.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Человеческий фактор', '24', 'На обследование к доктору Хаусу приезжает семейная пара с Кубы. Женщина тяжело больна, но все результаты её анализов и заключения врачей утонули во время кораблекрушения. Хаус считает, что лечение кубинских врачей было неправильным, и дело в чём-то другом. Так оно в результате и оказывается, но команда распадается, и что будет дальше, по ироничному выражению Хауса, «знает только Бог».', '3', 'HouseMD_3_24.mp4', 'HouseMD_3_24.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Один', '1', 'Молодая женщина сильно пострадала во время обрушения здания. Но ни одно из её многочисленных повреждений не может объяснить высокую температуру. Доктор Кадди пытается заставить Хауса взять на работу новых врачей в отделение диагностической медицины. Но Хаус против и заключает с Кадди сделку: если до конца дня он поставит пациентке диагноз, Кадди придется признать, что Хаус может работать один. Всем врачам госпиталя запрещено разговаривать с Хаусом и ему приходится пойти на крайнюю меру - взять себе в помощники уборщика.', '4', 'HouseMD_4_1.mp4', 'HouseMD_4_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('То, Что Нужно', '2', 'К Хаусу обращается женщина-пилот, которая скрывает от NASA свою болезнь, вызывающую дезориентацию. Сорок кандидатов на должность в диагностическом отделении должны не только вылечить её, но и сохранить её секрет.', '4', 'HouseMD_4_2.mp4', 'HouseMD_4_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('97 секунд', '3', 'Оставшиеся десять кандидатов соревнуются, стараясь как можно быстрее поставить диагноз парализованному пациенту, который медленно задыхается несмотря ни на что. Самого Хауса мучают мысли о странном пациенте, намеренно воткнувшем нож в розетку прямо на глазах у Хауса. Форман на новом месте работы сталкивается со старыми проблемами.', '4', 'HouseMD_4_3.mp4', 'HouseMD_4_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Ангелы-хранители', '4', 'Хаус и его команда лечат женщину, которая считает, что общается с призраками. По приказу Хауса испытуемые работники отправляются на задание на кладбище.', '4', 'HouseMD_4_4.mp4', 'HouseMD_4_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Зеркало, Зеркало', '5', 'Мужчина, попавший в госпиталь с проблемами дыхания, не может вспомнить, кто он. У потерявшего память пациента зеркальный синдром - он примеряет на себя личности тех, кого видит. Форман вынужден вновь вернуться в команду Хауса, где должен наблюдать за кандидатами.', '4', 'HouseMD_4_5.mp4', 'HouseMD_4_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Готовы На Всё', '6', 'ЦРУ приглашает Хауса поставить диагноз одному из агентов, но не желает давать почти никакой информации о нём, и на помощь готова только Самира, врач, работающий на ЦРУ. А Форман и оставшиеся шесть кандидатов заняты женщиной-автогонщиком, потерявшей сознание во время соревнований.', '4', 'HouseMD_4_6.mp4', 'HouseMD_4_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Урод', '7', 'Чейз готов провести операцию Кенни, подростку с сильно деформированным лицом, но у того происходит сердечный приступ. Команда Хауса оказывается под прицелом кинокамер телекомпании, снимающей фильм о Кенни, а сам Хаус начинает жалеть, что пригласил Самиру в госпиталь Принстон-Плейнсборо.', '4', 'HouseMD_4_7.mp4', 'HouseMD_4_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Лучше Тебе Не Знать', '8', 'Во время трюка с водяной камерой у фокусника начинается кровотечение, и Катнеру и Коулу приходится вытаскивать его из воды. А Хаус придумал команде новое испытание: тот, кто принесёт ему нижнее белье Кадди, получит иммунитет и сможет выбрать двоих кандидатов на выбывание.', '4', 'HouseMD_4_8.mp4', 'HouseMD_4_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Игры', '9', 'Хаус обещает, что тот кандидат, которому удастся поставить диагноз гитаристу-панку, останется в его команде. А Уилсон неожиданно обнаруживает, что один из его диагнозов был неверен, и его умирающий пациент будет жить.', '4', 'HouseMD_4_9.mp4', 'HouseMD_4_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Чудесная Ложь', '10', 'У женщины внезапный паралич рук, а Хаус начинает понимать, что ему, кажется, попался пациент, который не лжёт.', '4', 'HouseMD_4_10.mp4', 'HouseMD_4_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Во Льдах', '11', 'Врач, работающая на исследовательской станции на Южном полюсе, неожиданно чувствует сильные боли. Погода мешает покинуть станцию, но с госпиталем можно связаться с помощью видеоконференции. Перед Хаусом стоит сложная задача: набор лекарств на станции ограничен, из аппаратов под рукой только рентген, один из работников станции ранен, а молодая женщина отказывается принимать медикаменты, пока ей не будет поставлен диагноз. Хаусу приходится использовать всю свою изобретательность, чтобы придумать простейшие диагностические тесты, которые можно сделать даже на Южном полюсе.', '4', 'HouseMD_4_11.mp4', 'HouseMD_4_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Никогда Не Меняйся', '12', 'Женщина, недавно обратившаяся в хасидизм, теряет сознание на собственной свадьбе. Может быть, причины её болезни кроются в её прошлой жизни, которая была отнюдь не кошерной?', '4', 'HouseMD_4_12.mp4', 'HouseMD_4_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Прощайте, Мистер Добряк', '13', 'В госпитале забастовка медсестёр, и один пациент ждёт своей очереди уже целый день, но он ничуть не сердит. Хаус решает, что его добродушие — симптом какой-то болезни. Команда начинает подозревать, что на самом деле болен сам Хаус. А Хаусу и Эмбер удается установить совместную «опеку» над Уилсоном. Название - аллюзия на первую серию «Кошмаров Фредди»', '4', 'HouseMD_4_13.mp4', 'HouseMD_4_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Воплощая Мечты', '14', 'Доктор Хаус похищает актёра, играющего главную роль в его любимом сериале «Страсть по рецепту», потому что уверен, что тот болен. Но сам актёр считает, что с ним всё в порядке, так же считает и команда Хауса. В это время Кадди озабочена инспекцией, которая проходит в госпитале, а Уилсон мучается проблемой выбора нового матраса.', '4', 'HouseMD_4_14.mp4', 'HouseMD_4_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Голова Хауса', '15', 'В результате аварии автобуса Хаус получает серьёзную травму головы, которая приводит к частичной амнезии. Он помнит, что один из пассажиров чем-то опасно болен, и перед аварией он успел поставить ему диагноз, но не может вспомнить, кто этот пассажир.', '4', 'HouseMD_4_15.mp4', 'HouseMD_4_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Сердце Уилсона', '16', 'Эмбер при смерти, а Хаус всё ещё не может вспомнить, какой симптом он видел у неё перед аварией. Он поддаётся на уговоры Уилсона и соглашается на смертельно опасный эксперимент над собственным мозгом, который должен помочь вернуть ему память.', '4', 'HouseMD_4_16.mp4', 'HouseMD_4_16.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Смерть Меняет Всё', '1', 'В результате личной трагедии Уилсон уходит из больницы и разрывает свою дружбу с Хаусом. Уилсон больше не хочет с ним говорить. Кадди отчаянно пытается восстановить их отношения. Между тем Тринадцатая борется с личной проблемой, связанной с её диагнозом, а помогает ей пациентка с аналогичными трудностями.', '5', 'HouseMD_5_1.mp4', 'HouseMD_5_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Не Рак', '2', 'У нескольких совершенно не связанных между собой людей неожиданно начинают отказывать разные органы, и вскоре наступает смерть. Единственное, что их объединяет, — то, что в свое время всем им трансплантировали органы от одного и того же донора. Команда Хауса пытается найти причину и спасти единственную оставшуюся в живых пациентку, а сам Хаус в это время нанимает частного сыщика, для того чтобы следить за Уилсоном.', '5', 'HouseMD_5_2.mp4', 'HouseMD_5_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Осложнения', '3', 'В больницу поступает художник со странной болезнью, влияющей на его картины. Сам пациент уверяет, что уже пошёл на поправку, но Хаус относится к его словам скептически. Позже выясняется, что художник в тайне зарабатывал испытанием на себе экспериментальных препаратов, которые, возможно, и являются причиной его болезни. Тем временем частный сыщик Лукас шпионит для Хауса за его командой и Кадди.', '5', 'HouseMD_5_3.mp4', 'HouseMD_5_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Родовые Травмы', '4', 'Девушка отправляется в поисках своих биологических родителей в Китай, однако те почему-то при встрече лишь накричали на неё, утверждая что у них никогда не было дочери. Через некоторое время девушку начинает рвать кровью, после чего она оказывается пациенткой в Принстон - Плейнсборо. Тем временем Кадди, усыпив Хауса, обманом отправляет его вместе с Уилсоном на похороны отца Грегори, которые Хаус всеми силами старается пропустить, утверждая, что не является его родным сыном.', '5', 'HouseMD_5_4.mp4', 'HouseMD_5_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Счастливая Тринадцать', '5', 'В больницу после припадка попадает новая знакомая Тринадцатой. Постепенно выясняется, что пациентка познакомилась с ней только затем, чтобы попасть к Хаусу и быть им продиагностированной. Тем временем шпионские будни Лукаса продолжаются.', '5', 'HouseMD_5_5.mp4', 'HouseMD_5_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Джой', '6', 'В клинику попадает отец-одиночка, страдающий провалами в памяти. Похоже, они с дочерью не очень дружны… Тем временем Кадди встречается с девушкой, готовой отдать ей своего будущего ребенка на удочерение. Однако УЗИ показывает, что легкие ещё неродившейся девочки не до конца сформированы…', '5', 'HouseMD_5_6.mp4', 'HouseMD_5_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Зуд', '7', 'Мужчина, страдающий агорафобией (боязнью открытых пространств) серьёзно заболел, но при этом отказывается покинуть свой дом, чтобы поехать в больницу. Из-за чего Хаус и его команда отправляются прямо к нему домой, чтобы выяснить, чем тот на самом деле болен. Кэмерон берёт руководство диагностикой пациента на себя, так как уже лечила его в прошлом. Между тем, мужчине становится всё хуже и хуже, и Хаус планирует перевезти его в Принстон-Плейнсборо насильно…', '5', 'HouseMD_5_7.mp4', 'HouseMD_5_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Без Опеки', '8', 'Девушка-сирота, в шестнадцать лет освободившаяся от опеки, попадает в больницу с завода. Её диагностирование осложнено постоянной ложью. Хаус тем временем пытается понять, почему Уилсон никак не реагирует на развитие его и Кадди отношений. А Форман требует самостоятельности и занимается другим пациентом.', '5', 'HouseMD_5_8.mp4', 'HouseMD_5_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Крайний шаг', '9', 'В поисках диагноза вооруженный мужчина захватывает заложников в кабинете Кадди. Диагностику производит «случайно» оказавшийся в кабинете Хаус. Однако в лице пациента он встречает человека с таким же, как у себя, стремлением - для больного самым важным является нахождение ответа. После освобождения большинства заложников и разоружения террориста, Хаус возвращает ему оружие, а Тринадцатая осознает, насколько велика в ней жажда жизни…', '5', 'HouseMD_5_9.mp4', 'HouseMD_5_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Пусть едят торт', '10', 'Хаус вместе с командой принимаются за случай высококвалифицированного тренера по фитнесу, которой стало плохо во время съёмки рекламы. Многочисленные анализы и тестирование пациентки на астму, употребление стероидов и возможную нехватку витаминов, вызванную строгой диетой пациентки, не принесли никаких результатов, состояние пациентки только ухудшалось. Несмотря на то, что она рекламировала себя как фитнес-гуру, употребляющую только натуральные препараты, команда вскоре выясняет, что у неё есть секрет похудания, который, возможно, и вызвал ухудшение её здоровья.', '5', 'HouseMD_5_10.mp4', 'HouseMD_5_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Радость миру', '11', 'Хаус вместе с командой занимаются случаем девочки — тинэйджера, которой становится плохо во время Рождественской программы в школе. Вскоре они узнают, что в школе над девочкой издевались. По мере того как состояние тинэйджера продолжает ухудшаться, команда понимает, что вынуждена докопаться до истинной причины её таинственной болезни.', '5', 'HouseMD_5_11.mp4', 'HouseMD_5_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Без боли', '12', 'У Хауса проблемы с горячей водой в его доме, а Кадди пытается справиться в одиночку и с ребенком, и с работой, в то время как Форману нужно совершить трудный выбор. Хаус получает от Кэмерон пациента, пытавшегося покончить с собой из-за постоянных болей.', '5', 'HouseMD_5_12.mp4', 'HouseMD_5_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Большое дитя', '13', 'Кадди решает проводить больше времени дома, заботясь о своей приёмной дочери, и даёт Кэмерон возможность надзирать за Хаусом. Кажущаяся неотъемлемой доброта пациентки, которая занимается коррекционно-компенсирующим образованием детей, на самом деле является болезнью. Форман должен сделать потенциально опасное признание об участии Тринадцатой в испытании лекарства.', '5', 'HouseMD_5_13.mp4', 'HouseMD_5_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Высшее благо', '14', 'В 100-ом эпизоде Хаус и его команда принимает решение взять дело женщины, у которой произошёл припадок в классе по приготовлению еды. Когда они узнали, что пациент отказался от своей карьеры весьма известного исследователя рака с тем, чтобы жить в свое удовольствие, в команде возник вопрос о собственном счастье. Тем временем у Тринадцатой появляется побочный эффект от использования экспериментальных препаратов против её болезни Хантингтона.', '5', 'HouseMD_5_14.mp4', 'HouseMD_5_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Неверующий', '15', 'Хаус берётся за случай священника, который руководит приютом для бездомных (тому явился домой истекающий кровью и парящий над землёй Иисус). Вскоре выясняется, что из-за скандальной истории с мальчиком-прихожанином священник потерял работу, паству и веру.', '5', 'HouseMD_5_15.mp4', 'HouseMD_5_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Лучшая половина', '16', 'Хаус находится под действием метадона, который полностью избавляет его от боли. Он счастлив, и именно это становится причиной роковой врачебной ошибки. В результате практически здоровый пациент - тинейджер-гермафродит - заболевает и чуть не погибает.', '5', 'HouseMD_5_16.mp4', 'HouseMD_5_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Социальный контракт', '17', 'Хаусу и его команде достаётся случай издателя Ника Гринволда, который из-за повреждения мозга говорит вслух всё, что приходит ему в голову. Хаус обращает внимание на это, так как ему кажется забавным послушать всё то, что говорит пациент о персонале больницы и конкретно о его команде. Параллельно с распутыванием дела Хаус понимает, что Уилсон скрывает от него что-то важное, и старается как обычно вывести друга на чистую воду.', '5', 'HouseMD_5_17.mp4', 'HouseMD_5_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Сюда, Киска', '18', '35-летняя медсестра из дома престарелых симулирует припадок, чтобы попасть к Хаусу на диагностику. Выясняется, что ей предсказала скорую смерть кошка, известная своим безошибочным нюхом на потенциальных покойников. Хаус в такие глупости не верит, но вскоре медсестра серьёзно заболевает на самом деле…', '5', 'HouseMD_5_18.mp4', 'HouseMD_5_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Взаперти', '19', 'Хаус, разбивший свой мотоцикл во время тайной поездки, оказывается на соседней больничной койке с велосипедистом в состоянии псевдокомы, который уже рассматривается врачами как донор органов, и только благодаря решительному вмешательству Хауса оказывается спасён. Большая часть эпизода показана с точки зрения этого больного. Хаус перевозит своего интересного пациента в госпиталь Принстон-Плейнсборо, где выясняется, что не авария послужила причиной его состояния, а наоборот. Тауб пытается вернуть себе место, от которого он так неосмотрительно отказался в предыдущей серии.', '5', 'HouseMD_5_19.mp4', 'HouseMD_5_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Простое объяснение', '20', 'Шарлоту, немолодую женщину, которая последние шесть месяцев ухаживает за своим умирающим мужем Эдди, доставляют в Принстон-Плейнсборо после того, как у неё случился приступ удушья. Таинственность случая усиливается, когда состояние Эдди начинает улучшаться, а Шарлоты — ухудшаться. То, что невозможно было себе даже представить, происходит наяву: Шарлота может умереть раньше Эдди. Команда должна принять очень непростое решение, вдвойне сложное в связи с тем, что они, при невыясненных драматических обстоятельствах, теряют одного из своих сотрудников.', '5', 'HouseMD_5_20.mp4', 'HouseMD_5_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Спасители', '21', 'Сотрудники Принстон-Плейнсборо, а именно Хаус и его команда, пытаются справиться с недавней трагедией, тогда как Кэмерон знакомит Хауса со случаем защитника окружающей среды, который неожиданно потерял сознание во время акции протеста.', '5', 'HouseMD_5_21.mp4', 'HouseMD_5_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Разделение', '22', 'Сета, 14-летнего глухого подростка, доставляют в Принстон-Плейнсборо с борцовского ковра: в ходе состязания он стал слышать взрывы в своей голове. Диагностирование осложняется нежеланием пациента установить имплантат для восстановления слуха. Галлюцинация, вызванная длительной бессонницей, мешает Хаусу поставить верный диагноз, однако не препятствует руководить мальчишником накануне бракосочетания Чейза и Кэмерон.', '5', 'HouseMD_5_22.mp4', 'HouseMD_5_22.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('У Меня Под Кожей', '23', 'Пенелопа, 21-летняя звезда балета из Нью-Йорка, больше не может дышать. Из-за назначенных Хаусом антибиотиков у неё к тому же развивается токсикодермальный некроз, она теряет 80 % кожи, и ей грозит ампутация рук и ног. Тем временем Хаус признаётся Уилсону в том, что галлюцинирует и не понимает причину. Вместе они перебирают диагнозы: от шизофрении до рассеянного склероза, и Хаус отчаянно действует, но Эмбер продолжает являться в его сознании. Хаус прибегает к помощи Кадди.', '5', 'HouseMD_5_23.mp4', 'HouseMD_5_23.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Обе половинки вместе', '24', 'В больницу поступает пациент, левая рука которого взбесилась и делает все что ей «вздумается». Хаус после полной детоксикации перестает видеть Эмбер. Проведя ночь с Кадди, он встречает её на работе. Но та устанавливает правила, согласно которым с этого момента между ними могут быть только деловые отношения. Хаус с помощью Уилсона пытается выяснить причину подобного поведения, но чем сильнее он стремится к разгадке, тем больше встречает странностей и несоответствий и тем больше сомневается в своем излечении…', '5', 'HouseMD_5_24.mp4', 'HouseMD_5_24.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Сломанный', '1', 'Спустя некоторое время, проведенное в Мэйфилде, Хаус наконец-то очищается от викодина и перестаёт видеть галлюцинации. В этот же день он пытается выписаться из госпиталя. Однако доктор Нолан - главный врач Мэйфилда - соглашается написать рекомендацию на восстановление медицинской лицензии Хауса лишь в том случае, если тот останется для более длительного лечения.', '6', 'HouseMD_6_1.mp4', 'HouseMD_6_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Сломанный 2', '2', 'По мнению доктора, проблемы Грега находятся гораздо глубже, чем кажется, и викодин был далеко не основной из них. Хаус вынужден пойти на предложение Нолана. Тем не менее, он не оставляет попыток сбежать из госпиталя. С помощью Уилсона и своего нового соседа по палате Альви, Хаус надеется найти компромат на доктора Нолана. И даже после очередной неудачи он вновь готов строить план побега. И кажется, ничто не способно заставить Хауса изменить свое решение… До определенного момента…', '6', 'HouseMD_6_2.mp4', 'HouseMD_6_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Великое падение', '3', 'Форман пытается занять место Хауса в больнице «Принстон-Плейнсборо». В клинику попадает пациент со жгучей болью в ладонях. Он размещает свои симптомы в Интернете, предлагая вознаграждение за постановку правильного диагноза, в результате чего Форману приходится конкурировать со множеством виртуальных советчиков. В это время Хаус берет уроки кулинарного искусства, чтобы отвлечься от боли в ноге.', '6', 'HouseMD_6_3.mp4', 'HouseMD_6_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Тиран', '4', 'Чтобы развлечь себя, Хаус возвращается в диагностический отдел. Форман уговаривает Чейза и Кэмерон вернуться в отдел. В это время в больницу поступает с кровавой рвотой новый пациент - Дибала, президент-диктатор, которого многие обвиняют в геноциде. В то же время Хаусу докучает его сосед снизу, к выходкам которого Хаус старается относиться терпимо.', '6', 'HouseMD_6_4.mp4', 'HouseMD_6_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Карма в действии', '5', 'Миллиардер считает, что его сын заболел вследствие кармы, что финансовое благосостояние обратно пропорционально его семейному благополучию. Хаус предполагает неизлечимую болезнь и говорит, что мальчику осталось жить сутки. Тем временем Форман должен представить доклад на конференции по поводу смерти тирана Дибалы, но в поддельных тестах завышен холестерин. Тринадцатая планирует уехать в Таиланд, но её поездку кто-то срывает.', '6', 'HouseMD_6_5.mp4', 'HouseMD_6_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Храброе сердце', '6', 'В клинику попадает полицейский, отец, дед и прадед которого умерли в возрасте 40 лет от сердечного приступа. Пациент считает, что его постигнет та же участь. После ряда анализов и снимков, ничего не показавших, Хаус его выписывает, но вскоре выясняется, что проблема пациента реальна. Хаус в квартире у Уилсона переселяется из гостиной в спальню, где повсюду находятся фотографии и дипломы Эмбер, и вскоре начинает слышать странный шёпот по ночам. Чейз тем временем продолжает терзаться муками по поводу убийства Дибалы, даже исповедуется в церкви (безрезультатно), и в конечном счёте напивается.', '6', 'HouseMD_6_6.mp4', 'HouseMD_6_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Известные неизвестные', '7', 'После закрытой вечеринки в отеле в клинику попадает 16-летняя девушка с опухшими кистями и стопами. Команда Хауса начинает лечить её от булимии, но у пациентки развиваются внутренние кровотечения. Хаус отправляется вместе с Уилсоном и Кадди на конференцию по фармакологии, где Джеймс собирается прочитать доклад о своем последнем пациенте и признаться, что помог ему умереть. Это может стоить Уилсону карьеры, но тот не желает ничего слушать. В итоге Хаус накачивает Уилсона наркотиками и читает доклад вместо него.', '6', 'HouseMD_6_7.mp4', 'HouseMD_6_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Командная игра', '8', 'После восстановления медицинской лицензии Хаус вновь становится главой диагностического отдела. В это же время в госпиталь поступает Хэнк Хардвик, звезда порнофильмов, страдающий от пульсирующей боли в глазу. После того как Чейз рассказал Кэмерон, что это он убил Дибалу, они собираются уйти из клиники, чтобы начать новую жизнь, но Форман уговаривает их остаться до завершения дела. Хаус, предвидя скорое сокращение своей команды, пытается вернуть в неё Тауба и Тринадцатую, которые наотрез отказываются даже говорить об этом. Между тем, Кадди осознает, что «Принстон-Плейнсборо» не самое подходящее место для развития отношений.', '6', 'HouseMD_6_8.mp4', 'HouseMD_6_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Счастье в неведении', '9', 'В канун Дня благодарения Хаус и его команда ставят диагноз Джеймсу Сидису, исключительно выдающемуся физику и писателю, который променял карьеру на работу курьера. Для больного пациента интеллект стал тяжким бременем, которое вызвало депрессию и зависимость, а это, вместе с множеством странных симптомов, ставит команду Хауса в тупик. Тем временем Хаус пытается поссорить Кадди и Лукаса, а Чейз борется с надоедливыми утешениями коллег по поводу ухода Кэмерон.', '6', 'HouseMD_6_9.mp4', 'HouseMD_6_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Уилсон', '10', 'Когда у Такера, друга и бывшего пациента Уилсона, случается паралич руки, Уилсон берётся за его дело. И хотя сам Уилсон считает, что это не рак, Хаус с ним не соглашается и предлагает пари. Когда выясняется действительный диагноз Такера, Джеймсу приходится выбирать, кем ему лучше быть — другом или врачом. В это же время Кадди подыскивает новый дом для себя и Лукаса.', '6', 'HouseMD_6_10.mp4', 'HouseMD_6_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Под прикрытием', '11', 'В больницу, внезапно потеряв сознание на встрече наркодилеров, попадает Микки — агент, работающий под прикрытием. Постановка диагноза осложняется тем, что не может быть собран полный анамнез — пациент боится разоблачения. Хаус и Уилсон переезжают в новую квартиру, где соседка принимает их за любовную парочку, но не верит Уилсону, когда он пытается её в этом переубедить. Тогда Хаус решает воспользоваться этим и разыграть из себя гея, чтобы в итоге переспать с соседкой. Чтобы помешать этому, Уилсон даже делает предложение Хаусу на глазах у соседки.', '6', 'HouseMD_6_11.mp4', 'HouseMD_6_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Муки cовести', '12', 'В клинику попадает пациентка с пульсирующей болью в ушах. По ходу лечения выясняется, что она психопатка - не способна испытывать чувства - и изменяла мужу ради карьеры, а с мужем - несимпатичным и простоватым человеком - живёт только из-за его высокого финансового положения. Тринадцатая пытается открыть мужу глаза на сущность его жены, в результате чего вступает в жёсткую конфронтацию с бесчувственной и влиятельной стервой.', '6', 'HouseMD_6_12.mp4', 'HouseMD_6_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('На новый рубеж', '13', 'Дэрил, громадный чернокожий игрок в американский футбол, попадает в клинику после приступа ярости во время матча. Брат Формана выходит из тюрьмы, и Хаус нанимает его на работу своим помощником, преследуя при этом личные цели, а именно узнать правду о Формане. Дома у Хауса и Уилсона происходит ряд инцидентов, поначалу они подозревают в пакостничестве друг друга, но позже узнают, что это дело рук другого человека. На прием к Хаусу приходит парень, который не хочет идти на службу в армию.', '6', 'HouseMD_6_13.mp4', 'HouseMD_6_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('С 5 до 9', '14', 'Лиза Джей Кадди - холодная и жёсткая руководительница клиники - в то же время является нежной матерью и просто ранимой женщиной. Эпизод освещает один обычный день её жизни: уход за приболевшей дочкой, личные отношения с Лукасом, переговоры со страховой компанией о продлении контракта, расследование пропажи медикаментов, пациенты со странностями, и, конечно же, Хаус, который в присущем ему стиле хочет вылечить пациента с раковой опухолью, заразив его малярией.', '6', 'HouseMD_6_14.mp4', 'HouseMD_6_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Частная жизнь', '15', 'В клинику попадает блоггерша, у которой внезапно обнаружились синяки и открылось кровотечение. Прямо из больничной палаты, несмотря на разногласия с мужем, она пишет в свой блог о симптомах болезни, докторах, предполагаемых диагнозах, а также спрашивает у читателей совета по поводу наилучшего лечения. Хаус, Уилсон и Чейз посещают клуб пятиминутных свиданий для одиноких людей. Хаус узнает пикантные подробности из прошлого Уилсона и выставляет их на всеобщее обозрение. Джеймс пытается в отместку раскопать нечто подобное из личной жизни Грега.', '6', 'HouseMD_6_15.mp4', 'HouseMD_6_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Чёрная дыра', '16', 'Новым пациентом команды Хауса становится старшеклассница, внезапно потерявшая сознание во время лекции в планетарии. В больнице у девушки, помимо постоянных галлюцинаций, постепенно отказывают внутренние органы. После бесчисленных и безуспешных попыток диагностики ее загадочной болезни, Хаус, нарушая правила, предпринимает последнюю попытку: исследует когнитивную модель её мозга, чтобы найти новые зацепки и поставить диагноз. Уилсон по настоянию Хауса пытается самостоятельно, без посторонней помощи, купить в квартиру мебель, в то время как сам Хаус настойчиво пытается проникнуть в тайны личной жизни Тауба.', '6', 'HouseMD_6_16.mp4', 'HouseMD_6_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Под стражей', '17', 'В больнице у молодой семьи пропадает новорождённая девочка. В результате, до окончания поисков, заблокированы входы и выходы. Все герои оказываются запертыми в разных помещениях и, пока Кадди разыскивает пропавшего ребёнка, проводят время за разными занятиями. Уилсон с Тринадцатой в буфете играют в «Правду или вызов». Форман и Тауб в архиве, в попытке найти что-нибудь интересное в личном деле Хауса, узнают много нового друг о друге. Чейз выясняет отношения с Кэмерон, которая привезла документы на развод для подписи. Ну, а Хаус проводит несколько часов делясь сокровенным со смертельно больным пациентом, близким ему по духу.', '6', 'HouseMD_6_17.mp4', 'HouseMD_6_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Падение рыцарь', '18', 'У парня, участвующего в импровизированном рыцарском турнире, и на самом деле считающего себя рыцарем, происходит кровоизлияние в глаза, начинаются жуткие боли и выступает сыпь по всему телу. Команда в замешательстве и не может прийти к конкретному диагнозу. Уилсон начинает встречаться со своей бывшей женой Сэм, что дико не по душе Хаусу, который всеми средствами пытается помешать развитию их отношений считая что Сэм повторно разобьет сердце Джеймсу.', '6', 'HouseMD_6_18.mp4', 'HouseMD_6_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Открыть и закрыть', '19', 'У женщины, состоящей в открытом браке внезапно возникает острая боль в животе во время сексуальной прелюдии. Как и следовало ожидать, ни один из предположенных командой диагнозов не подходит. Хаус по прежнему пытается разрушить отношения Уилсона и Сэм. Тауб предлагает открытый брак своей жене под впечатлением от обсуждения личной жизни пациентки и флирта с одной из сотрудниц госпиталя.', '6', 'HouseMD_6_19.mp4', 'HouseMD_6_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Выбор', '20', 'Тэд попадает в руки команды Хауса после того как, в буквальном смысле, теряет дар речи во время собственной свадьбы. Установление диагноза затрудняют интимные подробности из личной жизни парня, в результате ему приходится сделать непростой выбор. Уилсон находит весьма оригинальный способ обеспечить себе вечера наедине с Самантой без присутствия Хауса.', '6', 'HouseMD_6_20.mp4', 'HouseMD_6_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Багаж', '21', 'В отношениях Хауса и Уилсона происходит весьма предсказуемый поворот. С целью разобраться в себе Хаус вновь посещает своего психоаналитика из Мэйфилда, попутно расследуя загадочные причины полной потери памяти у молодой девушки. Также его ожидает неожиданная встреча со старым хорошим знакомым.', '6', 'HouseMD_6_21.mp4', 'HouseMD_6_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Помоги мне', '22', 'Хаус и Кадди отправляются на место катастрофы, где в результате падения подъемного крана было разрушено здание и под завалами оказалось множество людей. По словам крановщика, за секунду до катастрофы он потерял сознание. Найдя это интересным, Хаус отправляет его со своей командой в Принстон - Плейнсборо и в дальнейшем руководит лечением по телефону. Вскоре он находит под завалами девушку по имени Ханна, которой придавило ногу железобетонной балкой. Хаус вынужден проводить с ней практически все время, чтобы не позволить Ханне впасть в панику.', '6', 'HouseMD_6_22.mp4', 'HouseMD_6_22.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Что теперь?', '1', 'Краткое описение серии - Как результат волнующего финала прошлого сезона, в котором Хаус и Кадди осознали свои чувства друг к другу, седьмой сезон начинается с выявления путей развития этих чувств и попыток выстроить настоящие отношения. А тем временем, вследствие болезни одного из работников, Принстон Плейнсборо осталась без нейрохирурга, что грозит закреплением за клиникой статуса Травматологического Центра Первого Уровня.', '7', 'HouseMD_7_1.mp4', 'HouseMD_7_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Эгоистичный', '2', 'Когда Делла (приглашенная звезда Стеонер), на вид здоровый и активный 14-летний ребенок, внезапно падает в обморок во время катания на скейтборде, Хаус и его команда пытаются изо всех сил диагностировать ее заболевание и успокоить ее родителей, которые уже и так сильно переживают из-за неизлечимой болезни их сына...', '7', 'HouseMD_7_2.mp4', 'HouseMD_7_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Ненаписанное', '3', 'Элиc (Эми Ирвинг), автор популярной серии детских книг, необъяснимо заболевает, после неудачной попытки покончить жизнь самоубийством. Команда госпиталя Принстон Плейнсборо сталкивается с трудностями в диагностировании ее заболевания, а также с ее нестабильным психическим состоянием. Дополнительной мотивацией для Хауса в преодолении трудностей с постановкой диагноза Элис служит тот факт, что Грегори является поклонником ее книг. Хаус убежден, что разгадка тайны заболевания Элис кроется на страницах ее нового романа.', '7', 'HouseMD_7_3.mp4', 'HouseMD_7_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Массажная терапия', '4', 'Пациентка по имени Маргарет Макферсон обращается в Принстон Плейсборо с жалобой на болезненную неконтролируемую рвоту. Хаус и команда, в процессе исследования болезни, делают неожиданные открытия о личности пациентки. Никаких улучшений у пациентки не наблюдается, и команда заглядывает в историю болезни, обнаруживая там еще более безумные вещи о прошлом больной. В то же время, Хаус организует холодный прием Чейзу, а визит массажиста заставляет Хауса и Кадди перестать осторожничать в их отношениях...', '7', 'HouseMD_7_4.mp4', 'HouseMD_7_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Незапланированное родительство', '5', ' новорожденного наблюдаются проблемы с дыханием и отказ печени. Чтобы решить эту проблему, Хаус и его команда исследуют историю болезни матери в поисках возможных подсказок. Команда Хауса раскрывает загадку, и теперь матери, Эбби, предстоит нелегкий выбор, который может отразиться и на здоровье ребенка, и на ее собственном. В то же время, согласно указаниям Кадди, Хаус заставляет Формана и Тауба нанять новую женщину-врача для команды. Также Кадди просит Хауса посидеть со своей дочерью, в результате чего Уилсон и Хаус учатся кое-чему важному о родительском долге.', '7', 'HouseMD_7_5.mp4', 'HouseMD_7_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Офисная политика', '6', 'Идут выборы, и в самый разгар тяжелой предвыборной борьбы управляющий предвыборной кампанией сенатора Нью-Джерси попадает в больницу с отказом печени и временным параличом. Кадди вынуждает Хауса нанять женщину-врача – третьекурсницу медицинского университета Марту Мастерс (ее сыграла Эмбер Тэмблин), на время отсутствия Тринадцатой. Хаус и его команда очень недоверчиво относятся к новенькой в связи с недостатком у нее медицинского опыта, но им приходится дать шанс ей проявить себя.', '7', 'HouseMD_7_6.mp4', 'HouseMD_7_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Чума на оба наши дома', '7', 'После того, как 200-летная медицинская колба, найденная в выброшенном на берег корабле, разбивается вдребезги в руках девушки-подростка, ее госпитализируют в ПП с симптомами, которые указывают на оспу. Доктор Дейв Брода из Центра Контроля за Заболеваниями объявляет карантин и отстраняет команду Хауса от постановки диагноза. У Мастерз возникают подозрения относительно его мотивов, она уверена, что у пациентки другое заболевание. Вскоре в ПП доставляют отца девушки с такими же симптомами и Хауз вынужден принять опасное решение, которая ставит под угрозу его собственную жизнь.', '7', 'HouseMD_7_7.mp4', 'HouseMD_7_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Маленькие жертвы', '8', 'В Принстон-Плейнсборо попадает мужчина, который начал задыхаться во время... Распятия на кресте! Как выяснилось, его дочь была больна раком, и жить ей оставалось не больше четырёх недель, но отец "заключил сделку с Богом". Рак у дочки исчез, но за каждый год её жизни отец претерпевал мучения Христа. Некоторое время после постановки правильного диагноза он отказывался от лечения стволовыми клетками, так как это "нарушает его договор с Богом". С помощью хитрой уловки Хаус убедил пациента начать лечение.', '7', 'HouseMD_7_8.mp4', 'HouseMD_7_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Больше, чем жизнь', '9', 'Мужчина хочет отдать донорские органы, чтобы спасти человека, попавшего под поезд в метро, но вдруг донора увозят в больницу, так как у него начинаются приступы. И больница, и город восхищаются альтруистским поступком донора, но в процессе изучения этого дела командой, оказывается, что кажущийся судьбоносным поступок главного героя не помог избавиться ему от старых привычек. В то же время Хаус пытается избежать званого ужина по случаю дня рождения Кадди, на котором должна присутствовать его потенциальная теща Арлин.', '7', 'HouseMD_7_9.mp4', 'HouseMD_7_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Морковь или палка', '10', 'Военный инструктор крепко полюбил своего ученика. У воспитанника исправительного учреждения для малолетних преступников начинают проявляется странные симптомы после интенсивного тренировочного курса. Внезапно оказывается, что такие же симптомы есть у сержанта, который его обучает. Команда не в состоянии найти причину, почему эти двое разделяют симптомы, поэтому начинает изучать семейную историю болезней воспитанника. Вдруг Мастерс и Хаус обнаруживают уникальную связь между сержантом и его воспитанником.', '7', 'HouseMD_7_10.mp4', 'HouseMD_7_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Семейная практика', '11', 'Мама Кадди поступила в Принстон Плейсборо после жалоб на необычные симптомы, но упертая Арлин настаивает, чтобы Хауса сняли с ее дела, вынуждая Хауса идти на противозаконные методы лечения пациентки. Хаус инструктирует команду как помогать ему, и после этого всплывает подробность личной жизни Арлин, которая держалась в секрете от Кадди и ее сестры Люсинды. В то же время, бывшая жена Тауба Рейчел сводит его с братом, который дает Таубу работу на стороне, так как Таубу сейчас нужны деньги.', '7', 'HouseMD_7_11.mp4', 'HouseMD_7_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Ты должен это помнить', '12', 'Когда официантка с идеальной памятью получает временный паралич, старшая сестра посещает её в больнице, но тогда у неё начинается стресс, и появляются осложнения. Тем временем, Форман решает помочь Таубу подготовиться к медицинскому обследованию, и Хаус, решая помочь Уилсону вернуться к знакомствам, обнаруживает его новую спутницу.', '7', 'HouseMD_7_12.mp4', 'HouseMD_7_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Две истории', '13', 'Хаус принимает участие в школьном «Дне карьеры» и нарушает несколько правил, подробно высказываясь об откровенных медицинских историях. У кабинета ректора он встречает двух студентов пятого курса (приглашенные актеры Остин Майкл Коулман и Хейли Паллос), которые помогают ему понять всю трагичность их отношений с Кадди и как его эгоистические выходки мешают ей в ее делах.', '7', 'HouseMD_7_13.mp4', 'HouseMD_7_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Противостоящий кризису', '14', 'Пациента доставляют в госпиталь с тяжелой формой сыпи, которая была вызвана едким химическим веществом во время его работы «синим воротничком». Обследуя пациента, команда выясняет, что он заставляет свою жену верить в то, что он все еще строит свою карьеру в некогда прибыльном бизнесе по недвижимости. Тем временем Кадди была удостоена премии и просит Хауса сопроводить ее на благотворительную акцию, где она должны была выступить. Но участие Хауса ставится под сомнение, когда пациент вынуждает его подвергнуть сомнению свою практику и собственное счастье.', '7', 'HouseMD_7_14.mp4', 'HouseMD_7_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Гром среди ясного неба', '15', 'Напряжение нарастает – Кадди узнает шокирующие новости, которые заставляют ее пересмотреть приоритеты. Пока Хаус отвлечен заботой о состоянии Кадди, команда лечит подростка, ухудшающее состояние которого и странные шрамы на теле указывают на нечто большее, чем физическую болезнь. Тауб обращает внимание на нездоровое эмоциональное и психологическое состояние пациента и понимает, что разгадку надо искать в личной жизни пациента: таким образом, он узнает про существование домашнего видео, которое ставит под угрозу жизни сверстников пациента.', '7', 'HouseMD_7_15.mp4', 'HouseMD_7_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Полный отрыв', '16', 'Команда лечит разбитое сердце в новой серии «Доктора Хауса». Молодой профессиональный наездник на быках, чемпион, поступил в госпиталь после нападения быка. После проведения нескольких неудачных анализов Хаус советует команде искать ответ вне больницы, пока сам он решает несколько проблем к делу не относящихся. В больнице здоровье наездника продолжает ухудшаться, а после того, как симптомы начали исчезать, а у пациента произошло несколько мини-приступов, команде приходится предпринять рискованную операцию на открытом сердце. В то же время Мастерс, к большому удивлению Тауба, ведет себя необычно по отношению к пациенту.', '7', 'HouseMD_7_16.mp4', 'HouseMD_7_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Грехопадение', '17', 'Тайна, угрожающая пациенту, будет раскрыта в новой серии «Доктора Хауса» Молодой бездомный человек, в прошлом – наркоман, найден в парке с признаками нарушения обоняния и ужасными шрамами и ожогами на груди. Личность пациента установить не смогли, и так как его состояние продолжало ухудшаться, команда пытается найти историю болезни пациента и записи о его семье, чтобы понять причину болезни. В то же время, Кадди в доверительной беседе с Уилсоном считает виновной себя в их разрыве с Хаусом. А как раз в тот момент, когда команда начинает успешно лечить пациента, они узнают ужасающую тайну о том, чью жизнь только что спасли.', '7', 'HouseMD_7_17.mp4', 'HouseMD_7_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Толчок', '18', 'Тринадцатая была в тюрьме весь прошлый год, но главная тайна для Хауса – за что она туда попала. Когда она выходит из тюрьмы, он уже ждет ее у выхода и сообщает ей, что следующие два дня они проведут вместе. Хаус пытается выпытать из Тринадцатой информацию о ее преступлении и для решения загадочной головоломки о лишении свободы они участвуют в ежегодном конкурсе стрельбы из картофельной пушки, где вместе противостоят молодой, уверенной команде, и в процессе борьбы узнают всю правду друг о друге.', '7', 'HouseMD_7_18.mp4', 'HouseMD_7_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Последнее искушение', '19', 'Мастерс сдает последний экзамен Хауса, а Тринадцатая возвращается в Принстон Плейсборо. Мастерс стоит перед тяжелейшим выбором, влияющим на судьбу ее карьеры в последние дни своего обучения в медицинском университете: ей предстоит выбрать – становится ли хирургом, или принять уникальное предложение стать членом команды Хауса. В то же время команда лечит 16-тилетнюю девочку, у которой внезапно случается приступ накануне кругосветного заезда на яхте. Несмотря на диагноз, перечеркивающий амбиции девушки, родители настаивают, чтобы дочь вернулась в море до начала претендующего на рекорд кругосветного путешествия.', '7', 'HouseMD_7_19.mp4', 'HouseMD_7_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Перемены', '20', 'Мама Кадди усложняет жизнь. Победитель в лотерею попадет в Принстон-Плейсборо на лечение у команды Хауса с частичным параличом, который возник после попыток больного найти настоящую любовь. Мама Кадди ложиться в больницу на профилактическое обследование и угрожает как Кадди так и Хаусу потерей их медицинских лицензий. Чейз и Форман заключают пари, которое проверяет каждого из них по-разному.', '7', 'HouseMD_7_20.mp4', 'HouseMD_7_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Спор', '21', 'Хаус и Уилсон идут на боксерский бой и не приходят к согласию касательно его результата. После этого Уилсон дает Хаусу один день чтобы доказать свою правоту или заплатить за проигрыш в споре. Хаус нажимает на боксера, чтобы добыть больше информации для победы в споре, а заканчивается это тем, что он дает новую надежду на победу давно проигрывавшему боксеру. В то же время команда начинает подозревать другой тип наркотической зависимости у Хауса.', '7', 'HouseMD_7_21.mp4', 'HouseMD_7_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('После работы', '22', 'Подруга Тринадцатой, Дарриан, с которой она делила камеру, внезапно приходит к ней домой, нуждаясь в срочной медицинской помощи. Когда Тринадцатая узнает, что ее подруга в бегах, обещает не отвозить ее в больницу, где полицейские могли бы найти ее, и вместо этого, в отчаянии обращается к Чейзу. Между тем, Хаус получает ужасающие известия, а Тауб старается справиться с неожиданными новостями.', '7', 'HouseMD_7_22.mp4', 'HouseMD_7_22.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Жизнь продолжается', '23', 'Хаус и команда занимаются случаем перфоманс-художницы, которая сознательно причиняет себе боль и превращает диагностику в очередной шедевр своего творчества. Команде предстоит решить, какие из её симптомов реальны, а какие нет, а Хаус тем временем решает изменить свою жизнь, но не может расстаться со старыми привычками.', '7', 'HouseMD_7_23.mp4', 'HouseMD_7_23.webp');

  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Двадцать Викодинок', '1', 'Прошел год после того, как Грегори Хаус влетел на своем автомобиле в дом Кадди. Хаус находится под пристальным наблюдением начальника тюрьмы в исправительном учреждении на востоке Нью-Джерси. Когда антагонистический лидер тюремной банды представляет серьезную угрозу, Хаус обращается за помощью к сокамернику. Но когда любопытство доктора пробуждают необычные медицинские симптомы другого заключенного, он должен придумать творческие способы лечения пациента, не выходя за рамки тюремных правил. В тюрьме Грегори знакомится девушкой по имени Джессика Адамс, молодым, умным и с сияющими глазами врача клиники.', '8', 'HouseMD_8_1.mp4', 'HouseMD_8_1.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Трансплантант', '2', 'Хаус возвращается в Принстон-Плейнсборо, чтобы помочь бывшей команде вылечить пациента под присмотром Уилсона. Он вынужден работать с робким молодым специалистом и должен, в конечном счете, выбрать, нарушить ли правила больницы, чтобы получить историю болезни пациента и найти лечение.', '8', 'HouseMD_8_2.mp4', 'HouseMD_8_2.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Благотворительные дела', '3', 'Хаус и Парк лечат пациента, Бенджамина, который упал в обморок после внесения удивительно большого благотворительного пожертвования и подозревают, что альтруистическое поведение является симптомом более глубокого расстройства. Между тем Хаус принимает на работу бывшего тюремного доктора Джессику Адамс. Но когда пациент предложил пожертвовать орган для другого пациента, врачи должны убедить доктора Адамса, чтобы помочь им подтвердить, действительно ли Бенджамин находится в здравом уме или нет. Тем временем, Адамс и Парк проверяют друг друга на великодушие и благодарность, а вина Тринадцать конфликтует с личным счастьем.', '8', 'HouseMD_8_3.mp4', 'HouseMD_8_3.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Рискованное дело', '4', 'Перед тем как подписать контракт, по которому вся рабочая сила компании перемещается в Китай, генеральный директор  компании неожиданно теряет сознание. Хаус пытается провести закулисную деловую сделку  его богатого пациента, но когда состояние больного ухудшается, команда должна работать круглосуточно, чтобы спасти ему жизнь. Тем временем Парк готовится к слушанию с Дисциплинарным Комитетом Принстон-Плейнсборо, под председательством Формана. А Адамс оценивает свою собственную этику, когда она узнает о планах их пациента.', '8', 'HouseMD_8_4.mp4', 'HouseMD_8_4.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Исповедь', '5', 'Человек, хорошо уважаемый в обществе внезапно падает в обморок, а в процессе диагностирования его состояния, команда обнаруживает, что пациент скрывал темные и не совсем четные поступки в  своей личной и профессиональной жизни. Но когда пациент открыто признается в своих правонарушениях своей семье и обществу, он ставит под угрозу свои возможности в получении надлежащего лечения. Тем временем Хаус не остановится ни перед чем, чтобы манипулировать Таубом по проведению теста на ДНК, который доказать, что он - отец двух шестимесячных дочерей.', '8', 'HouseMD_8_5.mp4', 'HouseMD_8_5.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Родители', '6', 'Хаус и его команда лечат подростка, который нуждается в трансплантации костного мозга. В процессе лечения открывается тревожная семейная тайна. Тем временем, Тауб пытается справиться с тем, что его бывшая жена хочет забрать их дочь и уехать. А Хаус принимает участие в боксерском поединке.', '8', 'HouseMD_8_6.mp4', 'HouseMD_8_6.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Мертвый и Погребенный', '7', 'Когда физические симптомы 14-летней пациентки ухудшаются, команда приходит к выводу, что страдания девушки больше, чем подростковая тоска. Несмотря на решительный протест Формана, Хаус становится одержимым решением специфического случая умершего четырехлетнего пациента, который вовлекает его в серьезные неприятности. Между тем, Парк пытается узнать у Чейза причину его недавней одержимостью уходом.', '8', 'HouseMD_8_7.mp4', 'HouseMD_8_7.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Опасности паранойи', '8', 'Обвинитель пострадал, по его мнению, от остановки сердца во время допроса свидетеля. Предварительный диагноз команды – гипер беспокойство, но когда Адамс и Парк исследуют дом пациента и находят скрытый арсенал огнестрельного оружия, они раскрывают более тревожные и глубинные психологические расстройства. Между тем Уилсон становится одержимым доказать, что Хаус скрывает что-то в его доме. Парк медленно вылезает из своей социальной раковины. А отсутствие у Формана романтических отношений возбуждает интерес Тауба и Чейза.', '8', 'HouseMD_8_8.mp4', 'HouseMD_8_8.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Лучшая половина', '9', 'Человек, страдающий от болезни Альцгеймера, выкашливает кровь и впадает в сильный гнев. В то время как команда пытается диагностировать его, Хаус проверяет пределы Формана. А Уилсон берет случай женщины, которая утверждает, что у нее и ее друга есть асексуальные отношения.', '8', 'HouseMD_8_9.mp4', 'HouseMD_8_9.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Беглецы', '10', 'Команда лечит несовершеннолетнюю и бездомную пациентку, но когда ее здоровье ухудшается и требуется хирургическое вмешательство, нужно получить согласие взрослых. Хаус и Адамс спорят по поводу того, следует ли обращаться к социальным услугам. Пациентка признается, что она убежала из дома, чтобы заботиться о своей матери, наркоманке в завязке. Но когда у постели пациентки появляется ее мать, открываются более сложные отношения, и мать с дочерью должны оставить все в прошлом и принять оптимальное решение для своей дочери. Тем временем, у Тауба сложное воссоединение с его грудными дочерями.', '8', 'HouseMD_8_10.mp4', 'HouseMD_8_10.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Никто не виноват', '11', 'Когда у инцидента с применением насилия, вовлекающего пациента, есть серьезные последствия для одного сотрудника, Хаус и команда находятся под наблюдением доктора Уолтера Кофилда, бывшим наставником Формана и нынешний глава неврологии. Поскольку Хаус и каждый член его команды рассказывают подробности драматических и опасных для жизни инцидентов, Кофилд должен взвесить нетрадиционный вид сотрудничества команды против их способности спасти жизни.', '8', 'HouseMD_8_11.mp4', 'HouseMD_8_11.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Чейз', '12', 'Чейз берет на себя пациентку, Мойру, уединенную монахиню, и  меняет ее взгляд на обеты и клятвы.  В процессе лечения, он и Мойра формируют уникальную связь, которая проверяет их веру и разум. Но когда состояние Мойры ухудшается и требуется провести опасную операцию, суждение Чейза под угрозой. Между тем, Хаус и Тауб пытаются  быть на шаг впереди в совершении  шалостей.', '8', 'HouseMD_8_12.mp4', 'HouseMD_8_12.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Человек Хауса', '13', 'Консультант по брачно-семейным отношениям падает в обморок во время важного разговора, когда он подвергнут нелестной оценке своих коллег, уловивших изменения в его поведении, которые идут в разрез его же собственному мнению по поводу ролей мужчин и женщин. Тем временем, Хаус и его украинская "жена" Доминика заключают сделку, чтобы убедить иммиграцию, что они – счастливая супружеская пара. Играя любящую пару, оба узнают кое-что о любви и браке. Кроме того Хаус решает назвать руководителя группы.', '8', 'HouseMD_8_13.mp4', 'HouseMD_8_13.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Любовь слепа', '14', 'Хаус и команда берутся за дело Уилла Вествуда, успешного и богатого, но слепого человека, который теряет сознание, делая предложение своей девушке. Хаус и команда пытаются поставить диагноз и спасти его, но в процессе они столкнутся с решением Уилла, которое изменит навсегда его жизнь. Между тем в Принстон-Плейнсборо появляется мать Хауса, чтобы рассказать сыну о новых отношениях.', '8', 'HouseMD_8_14.mp4', 'HouseMD_8_14.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Насвистывание', '15', 'Команда берется за лечение военного сухопутных войск США, поступившим с болями в животе, и обвиненного в измене, когда видео неудачной военной операции просочилось в Интернет. Его состояние становится все хуже и хуже, но ветеран отказывается от лечения, пока армия Штатов не предоставит ему файлы, которые докажут, что его отец, военный офицер, действительно умер. Тем временем, когда Адамс подозревает, что Хаус скрывает болезнь, которая может привести его к смертельному исходу, и просит Уилсон и товарищей по команде о помощи.', '8', 'HouseMD_8_15.mp4', 'HouseMD_8_15.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Внутренняя проверка', '16', 'Хаус и команда берут случай 22-летнего хоккеиста низшей лиги, который упал после борьбы на, выкашливая кровь, а затем потерял сознание.  Тем временем Хаус удивляет Уилсона некоторыми новостями, а Чейз предлагает Пак помощь в ее жилищных условиях.', '8', 'HouseMD_8_16.mp4', 'HouseMD_8_16.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Нужно быть мужиками', '17', 'Хаус и команда берут случай человека, которого начинает рвать кровью. Между тем Хаус объединяется с Доминикой, чтобы саботировать подающие надежды отношения Эмили - любимой проститутки Грегори, которая решила выйти замуж и бросить прежний бизнес.', '8', 'HouseMD_8_17.mp4', 'HouseMD_8_17.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Тело и душа', '18', 'Хаус и команда берутся за случай мальчика, который видит сон о том, что его душат, но когда он просыпается, то все еще не может вздохнуть. Тем временем Пак снятся сексуальные сны  с участием своих коллег, что заставляет остальных задуматься, а есть ли смысл в каждом из их сне. А Доминика узнает о тайне, которая может испортить ее отношения с Хаусом.', '8', 'HouseMD_8_18.mp4', 'HouseMD_8_18.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Слово на букву "Р"', '19', 'Когда команда берется за случай Эмили, шестилетней девочки, у которой многочисленные проблемы со здоровьем, они должны действовать сообща с матерью Эмили, которая, как оказалось сама врач, безуспешно пытающаяся вылечить дочь. Помимо этого команда сталкивается с постоянными пререканиями между матерью и отцом Эмили, мнения по поводу лечения дочери которых расходятся.  Ища в доме причины заболевания Эмили, команда понимает, что лечение, которое использует Элизабет для своей дочери, может оказаться и тем, что ее убивает. Тем временем Хаус и Уилсон берут небольшой отпуск.', '8', 'HouseMD_8_19.mp4', 'HouseMD_8_19.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('После смерти', '20', 'Команда берется за случай доктора Питера Трейбера, патолога из клиники Принстон-Плейнсборо, который знает слишком многое о больничном персонале, чтобы довериться любому из врачей. Единственный человек, которого он действительно уважает - Хаус, но он таинственно пропал без вести. В отсутствии Грегори Хауса команда должна Трейбера поверить в то, что все процедуры, которые они назначают, прописывает лично Хаус.', '8', 'HouseMD_8_20.mp4', 'HouseMD_8_20.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Держась', '21', 'Возвращается  бывшая сотрудница клиники Принстон Плейнсборо, Тринадцать. Команда берет случай Деррика, 19-летнего студента колледжа, у которого во время репетиции черлидинга пошла кровь носом, а позже обнаруживают, что его проблемы со здоровьем, вероятно, как физиологические, так и психологические. Возможно страдая от шизофрении, Деррик утверждает, что в голове услышал голос своего умершего брата. Тем временем Форман пробует другой подход к Хаусу.', '8', 'HouseMD_8_21.mp4', 'HouseMD_8_21.webp');
  INSERT INTO episode(name, number, description, season_id, video, photo) VALUES ('Все умирают', '22', 'Лечение пациента-наркомана заставляет Хауса пристальнее взглянуть на собственную жизнь, на будущее и на своих демонов.', '8', 'HouseMD_8_22.mp4', 'HouseMD_8_22.webp');
EOSQL
