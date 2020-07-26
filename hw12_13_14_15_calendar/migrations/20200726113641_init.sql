-- +goose Up
CREATE table users (
    id serial primary key,
    name varchar(50)
);

-- CREATE table events (
--     id  serial primary key,
--     title varchar(50),
--     start_at timestamp,
--     end_at timestamp,
--     description varchar(1000),
--     user_id serial,
--     notify_at timestamp,
--     constraint fk_user foreign key(user_id) references users(id)
-- );

INSERT INTO users (name)
VALUES ('Anna'), ('Bob'), ('Carl');

-- SQL in this section is executed when the migration is applied.

-- +goose Down
drop table events;
drop table users;
-- SQL in this section is executed when the migration is rolled back.