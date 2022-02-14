#!/bin/sh

DB = example
SCHEMA = schema.sql
USER = root

sqldef_dry:
	./mysqldef -h localhost -u $(USER) $(DB) --dry-run < $(SCHEMA)

sqldef_apply:
	./mysqldef -h localhost -u $(USER) $(DB) < $(SCHEMA)

sqlboiler_apply:
	./sqlboiler mysql --no-tests --no-hooks --wipe
