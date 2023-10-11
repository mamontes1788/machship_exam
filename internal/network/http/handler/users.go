package handler

import (
	"machship/internal/core/domain"
	"machship/internal/core/port"
	"machship/internal/core/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	userService port.UserService
}

func NewUsersHandler(userService port.UserService) UsersHandler {
	handler := new(UsersHandler)
	handler.userService = userService
	return *handler
}

func (n *UsersHandler) GetUsers(ctx *gin.Context) {
	util.Infoln(ctx, "### Start GetUsers() handler")

	var request domain.RetrieveUsersRequest
	if err := ctx.BindJSON(&request); err != nil {
		util.Errorf(err.Error())
		ctx.JSON(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		)
	}

	res, err := n.userService.GetUsers(ctx, request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)

	util.Infoln(ctx, "### End GetUsers() handler")
}
