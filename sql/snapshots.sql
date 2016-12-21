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

 Date: 12/10/2016 10:37:47 AM
*/

-- ----------------------------
--  Table structure for snapshots
-- ----------------------------
CREATE TABLE "public"."snapshots" (
	"event_source_id" uuid NOT NULL,
	"version" int8,
	"time_stamp" timestamp(6) NOT NULL,
	"source_type" varchar(255) NOT NULL COLLATE "default",
	"data" json NOT NULL,
	"id" uuid NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."snapshots" OWNER TO "admin";

-- ----------------------------
--  Primary key structure for table snapshots
-- ----------------------------
ALTER TABLE "public"."snapshots" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table snapshots
-- ----------------------------
CREATE INDEX  "snapshots_event_source_id_idx" ON "public"."snapshots" USING btree(event_source_id "pg_catalog"."uuid_ops" ASC NULLS LAST);

-- ----------------------------
--  Foreign keys structure for table snapshots
-- ----------------------------
ALTER TABLE "public"."snapshots" ADD CONSTRAINT "snapshot_event_source_id_fk" FOREIGN KEY ("event_source_id") REFERENCES "public"."event_sources" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

