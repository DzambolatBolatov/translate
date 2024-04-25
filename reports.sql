CREATE TABLE reports (
    id     SERIAL PRIMARY KEY,
    word_id int not null,
    title       VARCHAR(50) UNIQUE,
    description VARCHAR(255)
);

