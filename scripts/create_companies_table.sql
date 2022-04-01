create table companies
(
    id serial not null
        constraint companies_pk
            primary key,
    name varchar not null,
    created_at date default current_date not null,
    updated_at date default current_date not null
);

create unique index companies_name_uindex
    on companies (name);
