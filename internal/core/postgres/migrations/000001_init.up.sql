CREATE SCHEMA IF NOT EXISTS internship;

CREATE TABLE IF NOT EXISTS internship.internships (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(100) NOT NULL,
    label       VARCHAR(50),
    description VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS internship.internship_skills (
    id            SERIAL PRIMARY KEY,
    internship_id INTEGER      NOT NULL,
    skill_name    VARCHAR(255) NOT NULL,
    CONSTRAINT fk_internship_skills FOREIGN KEY (internship_id)
        REFERENCES internship.internships (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS internship.internship_goals (
    id            SERIAL PRIMARY KEY,
    internship_id INTEGER      NOT NULL,
    goal_name     VARCHAR(255) NOT NULL,
    CONSTRAINT fk_internship_skills FOREIGN KEY (internship_id)
        REFERENCES internship.internships (id) ON DELETE CASCADE
);