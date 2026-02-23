CREATE SCHEMA IF NOT EXISTS internship;

DO $$
    BEGIN
        CREATE TYPE internship.status AS ENUM ('draft', 'active', 'archive');
    EXCEPTION
        WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS internship.internships (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(100) NOT NULL,
    label       VARCHAR(50),
    description VARCHAR(1000),
    status      internship.status NOT NULL DEFAULT 'draft'
);

CREATE TABLE IF NOT EXISTS internship.internship_skills (
    id            SERIAL PRIMARY KEY,
    internship_id INTEGER      NOT NULL,
    skill_name    VARCHAR(200) NOT NULL,
    CONSTRAINT fk_internship_skills FOREIGN KEY (internship_id)
        REFERENCES internship.internships (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_skills_internship_id ON internship.internship_skills(internship_id);

CREATE TABLE IF NOT EXISTS internship.internship_goals (
    id            SERIAL PRIMARY KEY,
    internship_id INTEGER      NOT NULL,
    goal_name     VARCHAR(200) NOT NULL,
    CONSTRAINT fk_internship_goals FOREIGN KEY (internship_id)
        REFERENCES internship.internships (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_goals_internship_id ON internship.internship_goals(internship_id);
