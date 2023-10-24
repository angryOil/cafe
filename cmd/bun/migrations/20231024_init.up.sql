SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."cafe"
(
    id          SERIAL PRIMARY KEY,
    owner_id    bigint             not null,
    name        VARCHAR(60) unique not null,
    description varchar(2000),
    created_at  timestamptz
);
