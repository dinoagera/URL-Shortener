CREATE TABLE urlshort (
    id bigserial NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    url TEXT NOT NULL
);
CREATE INDEX idx_urlshort_name ON urlshort (name);