fresh_db:
	cd data && rm -rf informer.db; \
	sqlite3 informer.db < migration.sql

migrate_db:
	cd data && sqlite3 informer.db < migration.sql;

rebuild:
	docker-compose build --force-rm

up: rebuild
	docker-compose up -d