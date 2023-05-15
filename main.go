package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	u := User{ID: 10}
	if err := db.Take(&u).Error; err != nil {
		panic(err)
	}
	prettyPrint("got user:", u)

	// NOTE: This will return the id, even though it is on conflict.
	// If we just do `INSERT IGNORE`, we won't get the id.
	u2 := User{
		Name: "jessie",
	}
	if err := db.Select("name").
		Clauses(clause.OnConflict{
			// This might not be needed? Check with compound UK.
			//Columns: []clause.Column{{Name: "name"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"updated_at": gorm.Expr("current_timestamp()"),
				"age":        40,
			}),
		}).
		Create(&u2).Error; err != nil {
		panic(err)
	}

	prettyPrint("upsert:", u2)

	// This will not return the user id.
	u3 := User{
		Name: "jessie",
	}
	if err := db.Select("name").
		Clauses(clause.Insert{Modifier: "IGNORE"}).
		Create(&u3).Error; err != nil {
		panic(err)
	}
	prettyPrint("upsert:", u3)
}

func prettyPrint(msg string, v any) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		panic(err)
	}

	log.Println(msg, string(b))
}

type User struct {
	ID        int64
	Name      string
	Age       int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
