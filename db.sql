CREATE TABLE IF NOT EXISTS "users" ("id" bigserial, "branch_id" bigint, "description" text, "no" text, "password" text, "real_name" text, "state" text, "phone" text, "username" text, "add_time" timestamptz, PRIMARY KEY ("id")) 


--##

CREATE TABLE IF NOT EXISTS "groups" ("id" text, "name" text, "roles" text[], "desc" text, PRIMARY KEY ("id")) 

--##


CREATE TABLE IF NOT EXISTS "user_groups" ("user_id" bigint, "group_id" text, FOREIGN KEY ("user_id") REFERENCES "users" ("id"), FOREIGN KEY ("group_id") REFERENCES "groups" ("id")) 