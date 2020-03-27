package accounts

import (
	"MagicPotal/controllers/accounts/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

func CreateAccount(c *gin.Context) {
	log.Println("Creating account")

	txn := c.MustGet("Txn").(*gorm.DB)
	var data models.AccountInput
	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var account models.Account
	account.FirstName = data.FirstName
	account.LastName = data.LastName
	account.Nickname = data.Nickname
	account.Email = data.Email
	account.Password = data.Password
	account.RoleID = data.RoleID
	account.CreatedAt = time.Now()

	if  err = txn.Create(&account).Error; err != nil {
		log.Println("error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateAccount(c *gin.Context) {

	txn := c.MustGet("Txn").(*gorm.DB)
	aid := c.Param("aid")

	var data models.AccountUpdateInput
	var account models.Account

	query := txn.Table("account").
		Where("id = ?", aid).
		First(&account)

	if query.RecordNotFound() {
		log.Println("Account not found: ", query.Error.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, "account with id " + aid + " was not found")
		return
	}

	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	account.FirstName = data.FirstName
	account.LastName = data.LastName
	account.Nickname = data.Nickname
	account.Email = data.Email

	if  err = txn.Save(&account).Error; err != nil {
		log.Println("error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func GetAccounts(c *gin.Context) {
	txn := c.MustGet("Txn").(*gorm.DB)
	var accounts []models.BaseAccount
	query := txn.Table("account").Find(&accounts)

	if query.Error != nil {
		log.Println("Error on query: ", query.Error)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, query.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": accounts})
}

func GetSingleAccount(c *gin.Context) {
	txn := c.MustGet("Txn").(*gorm.DB)
	aid := c.Param("aid")

	var account models.Account

	query := txn.Table("account").Where("id = ?", aid).First(&account)

	if query.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, query.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": account})
}

func ResetPassword(c *gin.Context) {

}
