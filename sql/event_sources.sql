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

 Date: 12/18/2016 09:06:51 AM
*/

-- ----------------------------
--  Table structure for event_sources
-- ----------------------------
DROP TABLE IF EXISTS "public"."event_sources";
CREATE TABLE "public"."event_sources" (
	"id" uuid NOT NULL,
	"stream_name" varchar(255) NOT NULL COLLATE "default",
	"version" int4 NOT NULL,
	"created_at" timestamp(6) NOT NULL,
	"updated_at" timestamp(6) NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."event_sources" OWNER TO "admin";

-- ----------------------------
--  Primary key structure for table event_sources
-- ----------------------------
ALTER TABLE "public"."event_sources" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table event_sources
-- ----------------------------
CREATE UNIQUE INDEX  "event_sources_id_key" ON "public"."event_sources" USING btree("id" "pg_catalog"."uuid_ops" ASC NULLS LAST);
CREATE UNIQUE INDEX  "event_sources_stream_name_idx" ON "public"."event_sources" USING btree(stream_name COLLATE "default" "pg_catalog"."text_ops" ASC NULLS LAST);

