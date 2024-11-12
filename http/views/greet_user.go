package views

type GreetUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type GreetUserResponse struct {
	Message string `json:"message"`
}
