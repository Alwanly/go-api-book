-- Create "users" table
CREATE TABLE "users" ("id" character varying(36) NOT NULL, "username" character varying(36) NOT NULL, "password" integer NOT NULL, "created_at" bigint NULL, "created_by" character varying(36) NULL, "updated_at" bigint NULL, "updated_by" character varying(36) NULL, PRIMARY KEY ("id"));
