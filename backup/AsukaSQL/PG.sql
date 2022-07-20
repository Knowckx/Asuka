
CREATE TABLE IF NOT EXISTS "azure_cost_report" (
  "id" SERIAL primary key,  -- pg的SERIAL就是自增

  "date_type" varchar,
  "date" timestamptz,
  "project" varchar,
  "subscription" varchar,
  "res_group_name" varchar,
  "dajosn" JSONB,
  "qty" int4,
  "cost" money,
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);


CREATE INDEX "idx_date" ON "azure_cost_report" USING btree (  -- 加索引模板
  "date"  ASC NULLS LAST
);

CREATE UNIQUE INDEX "idx_date" ON "azure_cost_report" USING btree (  -- 加唯一索引
  "date"  ASC NULLS LAST
);




--删
ALTER TABLE azure_cost_report DROP COLUMN "Subscription"

-- 类型 
    字符串类型三种  char varchar  text
    在pg中，这三种类型之间没有性能差别！
    在大多数情况下，应该使用text或者varchar