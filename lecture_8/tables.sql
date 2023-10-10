create table if not exists users (
    id serial primary key,
    first_name varchar(100) not null,
    last_name varchar(100),
    email varchar(100) not null,
    password varchar(100) not null
);