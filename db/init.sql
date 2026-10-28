create database if not exists api_database;

use api_database;

create table almbums (
    id int primary key auto_increment,,
    release_date date,
    category_id int,
    title varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)

create table category(
    id int primary key auto_increment,
    name varchar(10),
    xreated_at timestamp default current_timestamp
)
