CREATE SCHEMA IF NOT EXISTS culand_internship;

CREATE TABLE IF NOT EXISTS culand_internship.internships (
    id    SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL
);
