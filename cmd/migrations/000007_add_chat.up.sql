CREATE TABLE "room_chat" (
  "chat_id" varchar PRIMARY KEY,
  "room_id" varchar NOT NULL,
  "account_id" varchar NOT NULL,
  "message" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "room_chat" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");
ALTER TABLE "room_chat" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");
