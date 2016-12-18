/*
 Navicat Premium Data Transfer

 Source Server         : eventsource
 Source Server Type    : PostgreSQL
 Source Server Version : 90503
 Source Host           : localhost
 Source Database       : eventstore_dev
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90503
 File Encoding         : utf-8

 Date: 12/18/2016 09:07:02 AM
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
	"stream_name" varchar(255) NOT NULL COLLATE "default",
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
CREATE INDEX  "events_stream_name_idx" ON "public"."events" USING btree(stream_name COLLATE "default" "pg_catalog"."text_ops" ASC NULLS LAST);

-- ----------------------------
--  Foreign keys structure for table events
-- ----------------------------
ALTER TABLE "public"."events" ADD CONSTRAINT "events_stream_name_fk" FOREIGN KEY ("stream_name") REFERENCES "public"."event_sources" ("stream_name") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

