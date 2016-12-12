#!/bin/sh
PGPASSWORD=password psql -U admin eventstore_dev <<OMG
BEGIN;

\i event_sources.sql

\i events.sql

\i snapshots.sql

COMMIT;
OMG
