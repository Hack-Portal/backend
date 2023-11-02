CREATE TABLE "account_tags" (
  "account_id" varchar NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "accounts" (
  "account_id" varchar PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "locate_id" int NOT NULL,
  "rate" int NOT NULL,
  "icon" text,
  "explanatory_text" text,
  "character" int,
  "show_locate" boolean NOT NULL,
  "show_rate" boolean NOT NULL,
  "twitter_link" varchar,
  "github_link" varchar,
  "discord_link" varchar,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT (now()),
  "is_delete" boolean NOT NULL DEFAULT false
);

CREATE TABLE "rate_entities" (
  "account_id" varchar NOT NULL,
  "rate" int NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "locates" (
  "locate_id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "hackathons" (
  "hackathon_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "icon" text NOT NULL,
  "link" varchar NOT NULL,
  "expired" date NOT NULL,
  "start_date" date NOT NULL,
  "term" int NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "is_delete" bool NOT NULL
);

CREATE TABLE "hackathon_status_tags" (
  "hackathon_id" varchar NOT NULL,
  "status_id" int NOT NULL
);

CREATE TABLE "status_tags" (
  "status_id" SERIAL PRIMARY KEY,
  "status" varchar NOT NULL
);

CREATE TABLE "follows" (
  "to_account_id" varchar NOT NULL,
  "from_account_id" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rooms" (
  "room_id" varchar PRIMARY KEY,
  "hackathon_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "member_limit" int NOT NULL,
  "is_closing" boolean NOT NULL DEFAULT 'false',
  "include_rate" boolean NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT (now()),
  "is_delete" boolean NOT NULL DEFAULT false
);

CREATE TABLE "tech_tags" (
  "tech_tag_id" serial PRIMARY KEY,
  "language" varchar NOT NULL,
  "icon" varchar NOT NULL
);

CREATE TABLE "rooms_accounts" (
  "rooms_account_id" varchar PRIMARY KEY,
  "account_id" varchar NOT NULL,
  "room_id" varchar NOT NULL,
  "is_owner" boolean NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rooms_accounts_roles" (
  "rooms_account_id" varchar NOT NULL,
  "role_id" int NOT NULL
);

CREATE TABLE "roles" (
  "role_id" SERIAL PRIMARY KEY,
  "role" varchar NOT NULL
);

CREATE TABLE "frameworks" (
  "framework_id" SERIAL PRIMARY KEY,
  "tech_tag_id" int NOT NULL,
  "framework" varchar NOT NULL,
  "icon" text NOT NULL
);

CREATE TABLE "account_frameworks" (
  "account_id" varchar NOT NULL,
  "framework_id" int NOT NULL
);

CREATE TABLE "room_chat" (
  "chat_id" varchar PRIMARY KEY,
  "room_id" varchar NOT NULL,
  "account_id" varchar NOT NULL,
  "message" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "proposal_hackathons" (
  "proposal_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "icon" text NOT NULL,
  "link" varchar NOT NULL,
  "expired" date NOT NULL,
  "start_date" date NOT NULL,
  "term" int NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "proposal_hackathon_status_tags" (
  "proposal_id" varchar NOT NULL,
  "status_id" int NOT NULL
);

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "rate_entities" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "account_tags" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "account_frameworks" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "account_frameworks" ADD FOREIGN KEY ("framework_id") REFERENCES "frameworks" ("framework_id");

ALTER TABLE "account_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "frameworks" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("locate_id") REFERENCES "locates" ("locate_id");

ALTER TABLE "rooms_accounts_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "rooms_accounts_roles" ADD FOREIGN KEY ("rooms_account_id") REFERENCES "rooms_accounts" ("rooms_account_id");

ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");

ALTER TABLE "room_chat" ADD FOREIGN KEY ("chat_id") REFERENCES "rooms" ("room_id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "proposal_hackathon_status_tags" ADD FOREIGN KEY ("proposal_id") REFERENCES "proposal_hackathons" ("proposal_id");

ALTER TABLE "proposal_hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");
