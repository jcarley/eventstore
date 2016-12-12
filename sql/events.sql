/*
 Navicat Premium Data Transfer

 Source Server         : eventstore
 Source Server Type    : PostgreSQL
 Source Server Version : 90503
 Source Host           : localhost
 Source Database       : eventstore_dev
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90503
 File Encoding         : utf-8

 Date: 12/10/2016 09:28:06 AM
*/

-- ----------------------------
--  Table structure for events
-- ----------------------------
DROP TABLE IF EXISTS "public"."events";
CREATE TABLE "public"."events" (
	"id" uuid NOT NULL,
	"time_stamp" timestamp(6) NOT NULL,
	"name" varchar NOT NULL COLLATE "default",
	"version" varchar NOT NULL COLLATE "default",
	"event_source_id" uuid NOT NULL,
	"sequence" int8,
	"data" json NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."events" OWNER TO "admin";

-- ----------------------------
--  Primary key structure for table events
-- ----------------------------
ALTER TABLE "public"."events" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table events
-- ----------------------------
CREATE INDEX  "events_event_source_id_idx" ON "public"."events" USING btree(event_source_id "pg_catalog"."uuid_ops" ASC NULLS LAST);

-- ----------------------------
--  Foreign keys structure for table events
-- ----------------------------
ALTER TABLE "public"."events" ADD CONSTRAINT "events_event_source_id_fk" FOREIGN KEY ("event_source_id") REFERENCES "public"."event_sources" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

