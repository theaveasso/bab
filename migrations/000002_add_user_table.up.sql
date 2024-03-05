create table if not exists users (
    username varchar(255) primary key,
    created_at timestamptz not null default now(),
    hashed_password varchar(255) not null,
    full_name varchar(255) not null,
    email varchar(255) unique not null,
    password_changed_at timestamptz not null default '0001-01-01 00:00:00Z'
);

alter table accounts
add foreign key (username) references users (username);

alter table accounts
add constraint accounts_currency_key
unique (username, currency);
