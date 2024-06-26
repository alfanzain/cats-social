CREATE TABLE IF NOT EXISTS public.users(
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR (200) NULL DEFAULT null UNIQUE,
    name VARCHAR (50) NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX email ON users (email);
