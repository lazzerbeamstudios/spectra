-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "email" character varying NOT NULL, "password" character varying NOT NULL, "created" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
