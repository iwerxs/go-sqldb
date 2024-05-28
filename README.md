## Golang with  Postgres SQL DB with Docker container

docker exec -ti pg-container psql -U postgres

postgres=# \c gopgtest

c -> connect to db gopgtest

gopgtest=# \dt

dt -> display tables in connected db

gopgtest=# SELECT * FROM product;