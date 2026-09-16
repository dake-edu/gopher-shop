-- Initial schema for the local legacy API. Not a migration system.
CREATE TABLE IF NOT EXISTS books (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL CHECK (btrim(title) <> ''),
    author text NOT NULL CHECK (btrim(author) <> ''),
    price numeric(12, 2) NOT NULL CHECK (price > 0 AND price <> 'NaN'::numeric),
    image_url text NOT NULL DEFAULT ''
);
