-- +migrate Up
    CREATE TABLE courses
    (
    id serial primary key,
    code text not null,
    name text not null,
    description text
    );

    CREATE TABLE themas
    (
    id serial primary key,
    code text not null,
    name text not null,
    description text
    );

    CREATE TABLE questions
    (
    id serial primary key,
    code text not null,
    name text not null,
    description text
    );

CREATE TABLE answers
(
    id serial primary key,
    code text not null,
    name text not null,
    description text
);

CREATE TABLE courses_vs_themas
(
    id serial primary key,
    course_id integer not null references courses,
    thema_id integer not null references themas

);


