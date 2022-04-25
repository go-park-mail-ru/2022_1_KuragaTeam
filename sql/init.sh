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
       is_movie bool not null,
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

  -- INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  -- VALUES('1+1', 'Intouchables.webp', true, '2011', '1 час 52 минуты', '16',
  --        'Пострадав в результате несчастного случая, богатый аристократ ' ||
  --        'Филипп нанимает в помощники человека, который менее всего подходит ' ||
  --        'для этой работы, – молодого жителя предместья Дрисса, только что освободившегося из тюрьмы. ' ||
  --        'Несмотря на то, что Филипп прикован к инвалидному креслу, Дриссу удается привнести в ' ||
  --        'размеренную жизнь аристократа дух приключений.', '8.8', 'Sometimes you have to reach into someone else''s world to find out what''s missing in your own', 'Intouchables.webp', 'Intouchables.mp4',
  --        'Intouchables.mp4') RETURNING id;
  --
  -- INSERT INTO movies(name, name_picture, is_movie, year, duration, age_limit, description, kinopoisk_rating, tagline, picture, video, trailer)
  -- VALUES('Игра в имитацию', 'TheImitationGame.webp', true, '2014', '1 час 54 минуты', '16',
  --        'Английский математик и логик Алан Тьюринг пытается взломать' ||
  --        ' код немецкой шифровальной машины Enigma во время Второй мировой войны.', '7.6', 'Основано на невероятной, но реальной истории',
  --        'TheImitationGame.webp', 'TheImitationGame.mp4', 'TheImitationGame.mp4') RETURNING id;

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

  DO
  $do$
      BEGIN
          FOR i IN 1..8 LOOP
                  INSERT INTO seasons(number, movie_id) VALUES (i, '14');
              END LOOP;
      END
  $do$;

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

EOSQL
