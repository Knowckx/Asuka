CREATE TABLE IF NOT EXISTS "azure_cost_report" (
  "id" int8 NOT NULL DEFAULT nextval('azure_resgroup_cost_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "date_type" varchar,
  "date" timestamptz(6),
  "project" varchar,
  "res_group_name" varchar ,
  "res_type" varchar,
  "qty" int4,
  "cost" money
);

