package config

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDBTeste() *gorm.DB {
	// Usar e mudar as infos pois o outro ta no gitignore
	dsn := "sqlserver://{{user}}:{{senha}}@localhost:{{porta ex: 1433}}?database={{nomeDatabase}}"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	return db
}
