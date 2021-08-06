CREATE TABLE websites (
    id   SERIAL       PRIMARY KEY,
    host VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE resources (
    id         SERIAL PRIMARY KEY,
    website_id INT NOT NULL,
    path       VARCHAR(50) NOT NULL,

    CONSTRAINT fk_websites FOREIGN KEY(website_id) REFERENCES websites(id)
);

CREATE TABLE contents (
    id        SERIAL      PRIMARY KEY,
    content   TEXT        UNIQUE NOT NULL,
    type      VARCHAR(20) NOT NULL,
    processed BOOL        DEFAULT FALSE
);

CREATE EXTENSION hstore;

CREATE TABLE requests (
    id          SERIAL      PRIMARY KEY,
    website_id  INT         NOT NULL,
    resource_id INT         NOT NULL,
    content_id  INT         NOT NULL,
    started     TIMESTAMPTZ NOT NULL,
    ended       TIMESTAMPTZ NOT NULL,
    headers     hstore      NOT NULL,
    status_code INT         NOT NULL,

    CONSTRAINT fk_websites FOREIGN KEY(website_id) REFERENCES websites(id),
    CONSTRAINT fk_resources FOREIGN KEY(resource_id) REFERENCES resources(id),
    CONSTRAINT fk_contents FOREIGN KEY(content_id) REFERENCES contents(id)
);

CREATE VIEW urls AS (
    SELECT
        host || path AS url,
        r.website_id,
        r.id as resource_id
    FROM
       websites AS ws
    INNER JOIN
       resources r on ws.id = r.website_id
)
