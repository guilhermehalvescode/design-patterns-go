package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/schemas"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/service"
)

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) List(c *gin.Context) {
	users, err := uc.UserService.List()
	if err != nil {
		c.JSON(500, schemas.NewAPIMessageResponse("Internal server error"))
		return
	}
	c.JSON(200, schemas.NewAPIResponse("List of users", users))
}

func (uc *UserController) GetByID(c *gin.Context) {
	user, err := uc.UserService.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(404, schemas.NewAPIMessageResponse(err.Error()))
		return
	}
	c.JSON(200, schemas.NewAPIResponse("User found", user))
}

func (uc *UserController) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, schemas.NewAPIMessageResponse("Invalid request"))
		return
	}
	user, err := uc.UserService.Create(user)
	if err != nil {
		c.JSON(400, schemas.NewAPIMessageResponse(err.Error()))
		return
	}
	c.JSON(201, schemas.NewAPIResponse("User created", user))
}

func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, schemas.NewAPIMessageResponse("Invalid request"))
		return
	}
	newUser, err := uc.UserService.Update(id, user)
	if err != nil {
		c.JSON(404, schemas.NewAPIMessageResponse(err.Error()))
		return
	}
	c.JSON(200, schemas.NewAPIResponse("User updated successfully", newUser))
}

func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := uc.UserService.Delete(id)
	if err != nil {
		c.JSON(404, schemas.NewAPIMessageResponse(err.Error()))
		return
	}
	c.JSON(200, schemas.NewAPIMessageResponse("User deleted successfully"))
}
