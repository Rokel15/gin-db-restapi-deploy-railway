-- +migrate Up
-- +migrate StatementBegin

create table afkeur(
    ID  INTEGER PRIMARY KEY,
    jumlahAyam INTEGER,
    startDate varchar(20),
    endDate varchar(20),
    mortalitas INTEGER
    -- first_name varchar(256),
    -- last_name varchar(256),
);

-- +migrate StatementEnd