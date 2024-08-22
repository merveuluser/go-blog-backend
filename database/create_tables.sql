CREATE TABLE IF NOT EXISTS "posts" (
    "id" SERIAL PRIMARY KEY,
    "title" text NOT NULL,
    "content" text NOT NULL,
    "summary" text,
    "url" text,
    "author_id" integer,
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS "authors" (
    "id" SERIAL PRIMARY KEY,
    "username" text NOT NULL,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS "comments" (
    "id" SERIAL PRIMARY KEY,
    "content" text NOT NULL,
    "post_id" integer,
    "author_id" integer,
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS "categories" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL
);

CREATE TABLE IF NOT EXISTS "post_categories" (
    "post_id" integer,
    "category_id" integer
);

ALTER TABLE "posts" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "post_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "post_categories" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");