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
	users := uc.UserService.List()
	c.JSON(200, schemas.NewAPIResponse("List of users", users))
}

func (uc *UserController) GetByID(c *gin.Context) {
	user := uc.UserService.GetByID(c.Param("id"))
	if user.ID != "" {
		c.JSON(200, schemas.NewAPIResponse("User found", user))
		return
	}
	c.JSON(404, schemas.NewAPIMessageResponse("User not found"))
}

func (uc *UserController) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, schemas.NewAPIMessageResponse("Invalid request"))
		return
	}
	user = uc.UserService.Create(user)
	c.JSON(201, schemas.NewAPIResponse("User created", user))
}

func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, schemas.NewAPIMessageResponse("Invalid request"))
		return
	}
	user = uc.UserService.Update(id, user)
	if user.ID != "" {
		c.JSON(200, schemas.NewAPIResponse("User updated successfully", user))
		return
	}
	c.JSON(404, schemas.NewAPIMessageResponse("User not found"))
}

func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if uc.UserService.Delete(id) {
		c.JSON(200, schemas.NewAPIMessageResponse("User deleted successfully"))
		return
	}
	c.JSON(404, schemas.NewAPIMessageResponse("User not found"))
}
