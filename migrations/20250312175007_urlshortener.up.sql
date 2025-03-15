CREATE TABLE urlshortener (
    id bigserial NOT NULL PRIMARY KEY,
    url TEXT NOT NULL,
    name VARCHAR NOT NULL UNIQUE
);
CREATE INDEX idx_urlshortener_name ON urlshortener (name);