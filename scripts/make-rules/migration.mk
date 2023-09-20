##@ Migration

MIGRATION_DIR="database"

.PHONY: migrate
migrate:## 运行数据库迁移
	migrate -path $(MIGRATION_DIR)/migrations -database "$(DATABASE_DRIVER)://$(DATABASE_SOURCE)" -verbose up

.PHONY: migrate.create
migrate.create:## 创建数据库迁移文件 (make migrate.create name=XXX)
	migrate create -ext sql -dir $(MIGRATION_DIR)/migrations $(name)

.PHONY: migrate.rollback
migrate.rollback:## 回滚一次迁移
	migrate -path $(MIGRATION_DIR)/migrations -database "$(DATABASE_DRIVER)://$(DATABASE_SOURCE)" -verbose down 1

.PHONY: migrate.rollback.all
migrate.rollback.all:## 回滚所有迁移
	migrate -path $(MIGRATION_DIR)/migrations -database "$(DATABASE_DRIVER)://$(DATABASE_SOURCE)" -verbose down -all

.PHONY: migrate.force
migrate.force:## 强制迁移特定的版本
	migrate -path $(MIGRATION_DIR)/migrations -database "$(DATABASE_DRIVER)://$(DATABASE_SOURCE)" -verbose force $(version)

.PHONY: migrate.version
migrate.version:## 显示 migrate 的版本
	@migrate -version