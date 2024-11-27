package controller

import (
	"time"

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
	c.JSON(200, schemas.APIResponse[[]domain.User]{Message: "List of users", Timestamp: time.Now().Unix(), Data: users})
}

func (uc *UserController) GetByID(c *gin.Context) {
	user := uc.UserService.GetByID(c.Param("id"))
	if user.ID != "" {
		c.JSON(200, schemas.APIResponse[domain.User]{Message: "User found", Timestamp: time.Now().Unix(), Data: user})
		return
	}
	c.JSON(404, schemas.APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
}

func (uc *UserController) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, schemas.APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
		return
	}
	user = uc.UserService.Create(user)
	c.JSON(201, schemas.APIResponse[domain.User]{Message: "User created", Timestamp: time.Now().Unix(), Data: user})
}

func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, schemas.APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
		return
	}
	user = uc.UserService.Update(id, user)
	if user.ID != "" {
		c.JSON(200, schemas.APIResponse[domain.User]{Message: "User updated successfully", Timestamp: time.Now().Unix(), Data: user})
		return
	}
	c.JSON(404, schemas.APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
}

func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if uc.UserService.Delete(id) {
		c.JSON(200, schemas.APIResponse[any]{Message: "User deleted successfully", Timestamp: time.Now().Unix()})
		return
	}
	c.JSON(404, schemas.APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
}
