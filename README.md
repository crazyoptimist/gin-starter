# Gin Starter

## TODO
- [ ] Write tests
- [ ] Setup [goose](https://github.com/pressly/goose) for migrations
- [ ] Setup Swagger

## Authentication middleware
```go
// building authentication middleware
authMiddleware := middlewares.AuthMiddleware{
    Jwt: &jwtService,
    DB:  db,
}

// initializing route with authentication middleware
router.GET("/me", authMiddleware.Validate(jwt.AccessToken), authCtrl.GetMe())

// getting user in controller
user := c.MustGet("user").(*models.UserModel)
```

## Http response builder
```go
lib.HttpResponse(200).Message("User registered successfully").Send(c)
```
Response
```json
{
     "code": "200",
     "success": true,
     "message": "User registered successfully"
}
```

## Request body validation with validation middleware
```go
// dto
type RegisterDto struct {
	Email    string `json:"email" form:"email" binding:"required,email,max=100"`
	Password string `json:"password" form:"password" binding:"required,max=100,min=8"`
	Name     string `json:"name" form:"name" binding:"required,max=100"`
	Status   string `json:"status" form:"status" binding:"required,max=150"`
}

// adding validation middleware in route
router.POST("/login", middlewares.Validate(&LoginDto{}), authCtrl.Login())

// retreiving the dto struct in controller
dto := c.MustGet("data").(*LoginDto)
```
