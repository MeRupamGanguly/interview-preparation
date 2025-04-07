# SQL

```bash
rupx@dev:~$ docker ps -a
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
rupx@dev:~$ 
rupx@dev:~$ docker volume create rupx-postgres-data
rupx-postgres-data
rupx@dev:~$ docker network create rupx-net
304fb9270e7ccfe5df0624a5aee34ff57a7b21a377bc2d0260466f6c655dc3e2
rupx@dev:~$ docker run -d \
  --name postgres \
  --network=rupx-net \
  -v rupx-postgres-data:/var/lib/postgresql/data \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  postgres:latest
Unable to find image 'postgres:latest' locally
latest: Pulling from library/postgres
6e909acdb790: Pull complete 
fec99121872b: Pull complete 
133acbc970df: Pull complete 
e02d97322fc6: Pull complete 
db9643c6baf3: Pull complete 
9bcedd9434e7: Pull complete 
fc8982ec96d9: Pull complete 
1824bd6b75d7: Pull complete 
fbad2bf2d5e6: Pull complete 
221788d72606: Pull complete 
e5f43b682bc0: Pull complete 
e7a2d9e24ab0: Pull complete 
a96cb29b0d13: Pull complete 
140970538145: Pull complete 
Digest: sha256:7f29c02ba9eeff4de9a9f414d803faa0e6fe5e8d15ebe217e3e418c82e652b35
Status: Downloaded newer image for postgres:latest
bb9ca88fb8d4d5b8e8a6a9219bb2e9c8719e40ae0aa82a2da650e1a60bce52f5
rupx@dev:~$ docker ps -a
CONTAINER ID   IMAGE             COMMAND                  CREATED          STATUS          PORTS                                       NAMES
bb9ca88fb8d4   postgres:latest   "docker-entrypoint.s…"   12 seconds ago   Up 11 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   postgres
rupx@dev:~$ 

```
![Image](https://github.com/user-attachments/assets/1ab8b786-5e86-4ec1-ba10-4b053f11497f)

```bash
docker exec -it postgres /bin/bash
psql -U myuser -d mydatabase

mydatabase=# \l
                                                  List of databases
    Name    | Owner  | Encoding | Locale Provider |  Collate   |   Ctype    | Locale | ICU Rules | Access privileges 
------------+--------+----------+-----------------+------------+------------+--------+-----------+-------------------
 mydatabase | myuser | UTF8     | libc            | en_US.utf8 | en_US.utf8 |        |           | 
 postgres   | myuser | UTF8     | libc            | en_US.utf8 | en_US.utf8 |        |           | 
 template0  | myuser | UTF8     | libc            | en_US.utf8 | en_US.utf8 |        |           | =c/myuser        +
            |        |          |                 |            |            |        |           | myuser=CTc/myuser
 template1  | myuser | UTF8     | libc            | en_US.utf8 | en_US.utf8 |        |           | =c/myuser        +
            |        |          |                 |            |            |        |           | myuser=CTc/myuser
(4 rows)

mydatabase=# \c mydatabase 
You are now connected to database "mydatabase" as user "myuser".
mydatabase=# \dt
Did not find any relations.
mydatabase=# CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
CREATE TABLE
mydatabase=# \dt
        List of relations
 Schema | Name  | Type  | Owner  
--------+-------+-------+--------
 public | users | table | myuser
(1 row)

mydatabase=# INSERT INTO users (name, email) VALUES ('John Doe', 'john@example.com');
INSERT 0 1
mydatabase=# \dt
        List of relations
 Schema | Name  | Type  | Owner  
--------+-------+-------+--------
 public | users | table | myuser
(1 row)

mydatabase=# SELECT * FROM users;
 id |   name   |      email       
----+----------+------------------
  1 | John Doe | john@example.com
(1 row)

mydatabase=# \d users
                                    Table "public.users"
 Column |          Type          | Collation | Nullable |              Default              
--------+------------------------+-----------+----------+-----------------------------------
 id     | integer                |           | not null | nextval('users_id_seq'::regclass)
 name   | character varying(100) |           |          | 
 email  | character varying(100) |           |          | 
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)

mydatabase=# 


```

![Image](https://github.com/user-attachments/assets/9b096e58-ec29-43d4-9f31-69d19955e318)

```bash

```