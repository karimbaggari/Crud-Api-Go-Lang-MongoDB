package controller

import( 
	"github.com/gin-gonic/gin"
	"scriptology/app/controllers/User/service"
	"scriptology/app/models"
	"net/http"
	
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService  service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}



func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	err := u.UserService.CreateUser(&user)
	if err!= nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


func (u *UserController) GetUser(ctx *gin.Context) {
	firstName := ctx.Param("firstName")
	_, err := u.UserService.GetUser(&firstName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
}

func (u *UserController) GetAll(ctx *gin.Context) {
	users, err := u.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200,users)
}


func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err!= nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
    }
	err := u.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
}



func (u *UserController) DeleteUser(ctx *gin.Context) {
    firstName := ctx.Param("first_name") // Correct assignment using :=
    err := u.UserService.DeleteUser(&firstName)
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
        return // Add return to exit the function if there's an error
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"}) // Return success message on successful deletion
}


func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:firstName", uc.GetUser)
	userRoute.GET("/getall", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:firstName", uc.DeleteUser)
}

