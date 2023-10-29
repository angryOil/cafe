SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."ban"
(
    id          SERIAL PRIMARY KEY,
    user_id     bigint not null,
    member_id   bigint not null,
    cafe_id     bigint not null,
    description varchar(2000),
    created_at  timestamptz
);
