CREATE SCHEMA IF NOT EXISTS "public";

CREATE TABLE "public"."categories" (
    "id" serial NOT NULL,
    "name" text NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."questions" (
    "id" serial NOT NULL,
    "category_id" integer,
    "question" text NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("category_id") REFERENCES "public"."categories"("id")
);

CREATE TABLE "public"."answers" (
    "id" serial NOT NULL,
    "question_id" integer,
    "answer" text NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("question_id") REFERENCES "public"."questions"("id")
);

