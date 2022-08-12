package users

import (
	"libs/modules/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Adhar string `json:"aadharNumber" binding:"required"`
}

var UsersList = make(map[string]common.User)

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": UsersList,
	})
}

func RetrieveUser(ctx *gin.Context) {
	id := common.GetSHA1([]byte(ctx.Param("id")))
	if _, OK := UsersList[id]; !OK {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data": "",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": UsersList[id],
	})
}

func AddUser(ctx *gin.Context) {
	body := Body{}
	if err := ctx.BindJSON(&body); err != nil {
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 	"status": false,
		// 	"error":  "Improper request data",
		// })
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}
	id := common.GetSHA1([]byte(body.Adhar))
	var u = common.User{
		Name:  body.Name,
		Age:   body.Age,
		Adhar: body.Adhar,
	}
	if _, OK := UsersList[id]; !OK {
		UsersList[id] = u
		ctx.JSON(http.StatusAccepted, UsersList[id])
	} else {
		ctx.AbortWithStatusJSON(http.StatusAlreadyReported, gin.H{
			"status": false,
			"error":  "Already Exists",
		})
	}
}

func DeleteUser(ctx *gin.Context) {
	id := common.GetSHA1([]byte(ctx.Param("id")))
	if _, OK := UsersList[id]; !OK {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data": "",
		})
		return
	}
	delete(UsersList, id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": UsersList,
	})
}
