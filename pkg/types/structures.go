package types

import validation "github.com/go-ozzo/ozzo-validation"

// response struct | marshalled into json fromat from struct
type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.Author,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(1, 50)))
}
//author request struct
type AuthorRequest struct{
	ID        uint `json:"id"`
	Name string	`json:"name"`
	Address string `json:"address"`
	Phone string	`json:"phone"`
}
//author validation
func (author AuthorRequest) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.Name,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(1, 50)))
}

//user request struct
type UserRequest struct{
	ID        uint `json:"id"`
	Username string	`json:"username"`
	Password string `json:"password"`
	Name string	`json:"name"`
	Email string	`json:"email"`
	Address string	`json:"address"`
}
//user validation
func (user UserRequest) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username,
			validation.Required.Error("Username cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&user.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&user.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(1, 50)),
		// validation.Field(&user.Email,
		// 	validation.Required.Error("Email cannot be empty"),
		// 	validation.Length(1, 50)),
		// validation.Field(&user.Address,
		// 	validation.Required.Error("Address cannot be empty"),
		// 	validation.Length(1, 50)))
	)
}
//login request struct
type LoginRequest struct{
	Username string	`json:"username"`
	Password string `json:"password"`
}
//login validation
func (login LoginRequest) Validate() error {
	return validation.ValidateStruct(&login,
		validation.Field(&login.Username,
			validation.Required.Error("Username cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&login.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(1, 50)))
}
