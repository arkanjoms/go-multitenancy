create table todo
(
    id          SMALLSERIAL primary key,
    description VARCHAR(255) not null,
    completed   BOOLEAN      not null
);
