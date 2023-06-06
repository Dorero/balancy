create table balance
(
    id      integer default nextval('accounts_id_seq'::regclass) not null
        constraint accounts_pkey
            primary key,
    user_id integer                                              not null,
    money   numeric                                              not null
);

alter table balance
    owner to postgres;

create table reserves
(
    id         serial
        primary key,
    amount     numeric not null,
    order_id   bigint,
    service_id bigint,
    user_id    bigint
);

alter table reserves
    owner to postgres;

create table payments
(
    id         bigserial
        primary key,
    balance_id bigint
        constraint payments_balance_id_fk
            references balance
            on delete cascade,
    amount     numeric not null
);

alter table payments
    owner to postgres;

create table unreserves
(
    id         serial
        primary key,
    amount     numeric,
    order_id   bigint,
    service_id integer,
    user_id    bigint
);

alter table unreserves
    owner to postgres;



