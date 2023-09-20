package data

import "gorm.io/gorm"

var defaultDBClient *gorm.DB

// GetDefaultDBClient 获取默认的 DB 实例
func GetDefaultDBClient() *gorm.DB {
	return defaultDBClient
}

// SetDefaultDBClient 设置默认的MySQL实例
func SetDefaultDBClient(client *gorm.DB) {
	defaultDBClient = client
}
