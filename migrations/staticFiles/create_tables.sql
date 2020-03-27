
create table role (
    id uuid primary key,
    role_name varchar(80) unique not null
);

create table account (
    id uuid primary key,
    first_name varchar(80) not null,
    last_name varchar(80) not null,
    nickname varchar(20),
    email varchar(180) unique not null,
    password varchar (10) not null,
    created_at timestamp not null,
    deleted_ad timestamp,
    role_id uuid not null,

    constraint role_id_rl foreign key (role_id)
    references role (id) match simple
    on update no action
    on delete  no action
);

create table articles (
    id uuid primary key,
    title varchar(250) not null,
    subtitle varchar(250),
    category varchar (80) not null,
    content text not null,
    account_id uuid not null,

    constraint account_id_rl foreign key (account_id)
    references account (id) match simple
    on update no action
    on delete  no action
);

create table category (
    id uuid primary key,
    name varchar(80) unique not null
);