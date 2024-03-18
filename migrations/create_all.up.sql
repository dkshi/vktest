CREATE TABLE actors (
    actor_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    birth_date DATE NOT NULL
);

CREATE TABLE films (
    film_id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT,
    release_date DATE,
    rating NUMERIC(3,1) CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE actor_film (
    actor_id INT REFERENCES actors(actor_id),
    film_id INT REFERENCES films(film_id),
    PRIMARY KEY (actor_id, film_id)
);

CREATE TABLE admins (
    admin_id SERIAL PRIMARY KEY,
    adminname VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);