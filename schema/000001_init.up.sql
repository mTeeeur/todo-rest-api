CREATE TABLE users
(
    id            serializable not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
)

CREATE TABLE todo_lists
(
    id          serializable not null unique,
    title       varchar(255) not null,
    description varchar(255)
)

CREATE TABLE users_list
(
    id      serializable                                     not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null,
)

CREATE TABLE todo_items
(
    id          serializable not null unique,
    title       varchar(255) not null,
    description varchar(255),
    compilted   Boolean null default false
)

CREATE TABLE list_items
(
    id      serializable                                     not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
)
