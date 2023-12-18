package account

import (
	"ginPratice/middlewares"
	"ginPratice/models"
	"ginPratice/utils"

	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AccountController struct {
}

func (con AccountController) CreateAccount(ctx *gin.Context) {
	db := utils.InitDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	IsSuperuser, err := strconv.Atoi(ctx.PostForm("isSuperuser"))
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(ctx, utils.UserError)
	}
	newUser := models.User{
		Name:        name,
		Telephone:   telephone,
		Password:    string(hasedPassword),
		IsSuperuser: int8(IsSuperuser),
	}
	db.Create(&newUser)
	utils.Ok(ctx)
}

func (con AccountController) Login(c *gin.Context) {
	db := utils.InitDB()
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user models.User
	db.Where("name = ?", username).First(&user)
	if user.ID == 0 {
		utils.Fail(c, utils.UserExist)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		utils.Fail(c, utils.UserError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": middlewares.GenerateToken(&middlewares.UserClaims{
			ID:             user.ID,
			Name:           user.Name,
			StandardClaims: jwt.StandardClaims{},
		}),
	})
}

func (con AccountController) UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	claims := user.(*middlewares.UserClaims)
	c.JSON(http.StatusOK, gin.H{
		"name": claims.Name,
	})
}
