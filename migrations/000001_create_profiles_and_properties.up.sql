CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE profiles (
                          id UUID PRIMARY KEY,
                          name TEXT NOT NULL,
                          age INT NOT NULL CHECK (age > 0),
                          email TEXT NOT NULL,
                          phone TEXT,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE properties (
                            id SERIAL PRIMARY KEY,
                            profile_id UUID NOT NULL,
                            tags TEXT[] NOT NULL,
                            created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                            CONSTRAINT fk_profile
                                FOREIGN KEY (profile_id)
                                    REFERENCES profiles(id)
                                    ON DELETE CASCADE
);