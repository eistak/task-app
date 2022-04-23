CREATE TABLE tasks (
    id SERIAL NOT NULL,
    content VARCHAR(255) NOT NULL,
    done BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
