CREATE TABLE IF NOT EXISTS Comment
(
    id               SERIAL PRIMARY KEY,
    user_id          INT       NOT NULL,
    parent_id        INT       NULL,
    post_id          INT       NOT NULL,
    body             TEXT      NOT NULL,
    timestamp        TIMESTAMP NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES Comment (id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES Post (id) ON DELETE CASCADE
);