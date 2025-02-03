GRANT ALL PRIVILEGES ON DATABASE postgres TO admin;
GRANT ALL PRIVILEGES ON SCHEMA public TO admin;

CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    refresh_token VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    last_login TIMESTAMP,
);

CREATE INDEX idx_email ON users (email);

CREATE TABLE IF NOT EXISTS public.characters (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    level INTEGER NOT NULL,
);

ALTER TABLE characters ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS public.stats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
);

CREATE TABLE IF NOT EXISTS public.character_stats (
    character_id INTEGER NOT NULL,
    stat_id INTEGER NOT NULL,
    value INTEGER NOT NULL,
);

ALTER TABLE character_stats ADD FOREIGN KEY (character_id) REFERENCES characters (id);
ALTER TABLE character_stats ADD FOREIGN KEY (stat_id) REFERENCES stats (id);
