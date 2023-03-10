CREATE TABLE "devices" (
  "id" text PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "name" text NOT NULL,
  "location" text NOT NULL
);
CREATE TABLE "data_types" (
  "id" SERIAL PRIMARY KEY,
  "key" varchar(30) NOT NULL,
  "unit" varchar(30) NOT NULL,
  "device_id" text
);
CREATE TABLE "datas" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "data_type_id" int,
  "value" float NOT NULL
);
CREATE TABLE "api_keys" ("id" text PRIMARY KEY, "expiry_date" text);
ALTER TABLE "data_types"
ADD FOREIGN KEY ("device_id") REFERENCES "devices" ("id");
ALTER TABLE "datas"
ADD FOREIGN KEY ("data_type_id") REFERENCES "data_types" ("id");