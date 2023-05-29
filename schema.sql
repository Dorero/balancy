create table accounts
(
    id      serial
        primary key,
    user_id integer           not null,
    money   integer default 0 not null
);

alter table accounts
    owner to postgres;

create table reserves
(
    id         serial
        primary key,
    user_id    integer not null,
    id_service integer not null,
    id_order   integer not null,
    money      integer not null
);

alter table reserves
    owner to postgres;

