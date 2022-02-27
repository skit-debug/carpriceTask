FROM mysql:5.7
COPY ./src/db/test_data.sql /docker-entrypoint-initdb.d/