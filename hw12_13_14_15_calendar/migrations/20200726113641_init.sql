-- +goose Up
CREATE TABLE post (
    id int NOT NULL,
    title text,
    body text,
    PRIMARY KEY(id)
);

INSERT INTO post(id, title, body)
VALUES (1, '1', '1'),
        (2, '2', '2'),
        (3, '3', '3');

-- +goose Down
DROP TABLE post;