package middlewares

import (
	"MagicPotal/server/middlewares/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var dbs = map[string]*gorm.DB{
	"defaultDB": nil,
}

func TXNInit(dbType string) gin.HandlerFunc {

	return func(c *gin.Context) {

		log.Println("transaction begin")
		txn := NewTransaction(dbType)
		if txn.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, txn.Error)
			log.Println("Transaction begin error", txn.Error)
			return
		}

		c.Set("Txn", txn)

		log.Println("next")
		c.Next()
		log.Println("transaction management")

		if txn != nil && (txn.Error != nil || c.IsAborted()) {
			txn.Rollback()
			log.Println("Transaction rollback", txn.Error)
		} else if txn != nil {
			log.Println("Transaction committed")
			txn.Commit()
		}
		txn = nil
	}
}

func DBInit(dbType string) gin.HandlerFunc {
	var err error

	log.Println("start connection")
	dbs[dbType], err = helpers.GetDB(dbType)
	return func(c *gin.Context) {
		log.Println("internal DB fun")
		if err != nil {
			log.Println("no connection", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, "No database connection")
			return
		}
		c.Next()
	}
}

func NewTransaction(dbType string) *gorm.DB {
	db := dbs[dbType].New()
	return db.Begin()
}
