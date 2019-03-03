-- ----------------------------
-- Table structure for u1
-- ----------------------------
DROP TABLE IF EXISTS "public"."u1";
CREATE TABLE "public"."u1" (
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "age" int4,
  "address" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."u1" OWNER TO "postgres";