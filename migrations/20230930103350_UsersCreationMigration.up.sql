CREATE TABLE users (
    id int not null primary key,
    name varchar not null,
    email varchar not null unique,
    age int not null,
    password_hash varchar not null
);