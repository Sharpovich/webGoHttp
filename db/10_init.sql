CREATE TABLE Users (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    firstname varchar(30) NOT NULL,
    lastname varchar(30) NOT NULL,
    city varchar(30) NOT NULL
);