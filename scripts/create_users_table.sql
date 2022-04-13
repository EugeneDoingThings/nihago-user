create table users
(
    id            serial                    not null
        constraint users_pk
            primary key,
    firstname     varchar                   not null,
    lastname      varchar                   not null,
    patronymic    varchar,
    date_of_birth date                      not null,
    about         varchar,
    company_id    integer,
    photo         varchar,
    created_at timestamp(0) with time zone default current_timestamp not null,
    updated_at timestamp(0) with time zone default current_timestamp not null
);

alter table users
    owner to postgres;

create unique index users_id_uindex
    on users (id);
