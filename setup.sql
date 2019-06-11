--
-- Run from snaily user in madast database
--

---- START OF VERSION 1

create schema madast;

create table madast.users
(
    id        text not null
        constraint users_pk
            primary key,
    email     text
        constraint users_email_key
            unique,
    firstname text,
    lastname  text
);

alter table madast.users
    owner to snaily;


alter table madast.users
    add column password text;

---- END OF VERSION 1

---- START OF VERSION 2

create table madast.items
(
    id text not null
        constraint items_pk
            primary key,
    parent_id text,
    title text,
    body text,
    created_date timestamp with time zone,
    created_by text,
    waiting_for text,
    organization_id text,
    waiting_for_done boolean,
    waiting_for_done_date timestamp with time zone,
    created_by_done boolean,
    created_by_done_date timestamp with time zone
);

alter table madast.items owner to snaily;

create index items_created_by_index
    on madast.items (created_by);

create index items_waiting_for_index
    on madast.items (waiting_for);

---- END OF VERSION 2

