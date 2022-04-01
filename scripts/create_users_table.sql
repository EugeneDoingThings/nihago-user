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
    company_id    integer
        constraint users_companies_id_fk
            references companies
            on delete set null,
    photo         varchar,
    created_at    date default CURRENT_DATE not null,
    updated_at    date default CURRENT_DATE not null
);

alter table users
    owner to postgres;

create unique index users_id_uindex
    on users (id);
