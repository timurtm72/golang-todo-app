CREATE TABLE users
(
    id              bigserial primary key,
    name            varchar(100) not null,
    username        varchar(100) not null unique ,
    password_hash   varchar(255)
);

CREATE TABLE todo_lists
(
    id              bigserial primary key,
    title           varchar(255) not null,
    description     varchar(255)
);

CREATE TABLE users_lists
(
    id              bigserial primary key,
    user_id         bigint not null,
    list_id         bigint not null,
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (list_id) references todo_lists(id) on delete cascade
);

CREATE TABLE todo_items
(
    id              bigserial primary key,
    title           varchar(255) not null,
    descriptions    varchar(255),
    done            boolean not null default false
);

CREATE TABLE lists_items
(
    id              bigserial   primary key,
    item_id         bigint not null,
    list_id         bigint  not null,
    foreign key (item_id) references todo_items(id) on delete cascade,
    foreign key (list_id) references todo_lists(id) on delete cascade
);
