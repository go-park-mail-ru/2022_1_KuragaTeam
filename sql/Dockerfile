FROM postgres:latest

RUN DEBIAN_FRONTEND=noninteractive apt-get update \
    &&  apt-get install -y git \
    &&  apt-get install -y build-essential \
    &&  apt-get install -y postgresql-server-dev-14 \
    &&  rm -rf /var/lib/apt/lists/*

RUN mkdir -p hunspell \
    && cd hunspell \
    && git clone https://github.com/postgrespro/hunspell_dicts

WORKDIR hunspell/hunspell_dicts/hunspell_ru_ru

RUN make USE_PGXS=1 && make USE_PGXS=1 install