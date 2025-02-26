CREATE TABLE IF NOT EXISTS Post
(
    id          SERIAL    PRIMARY KEY,
    body        INT       NOT NULL,
    userId      TEXT      NOT NULL,
    comments    BOOLEAN   NOT NULL
)