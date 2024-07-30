package api

import (
	"context"
	"fmt"
	"net/http"

	entity "accessCloude/internal/storage/entity"
)

func (ac *AccessCloude) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	var (
		user entity.User
		ctx  = context.Background()
	)

	if err := user.Delete(ctx, ac.DB.Conn, id); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	Response(w, "user deleted!", 200)
}

func (ac *AccessCloude) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	var (
		user entity.User
		ctx  = context.Background()
	)

	fmt.Print(id, "\n")
	if err := user.GetByID(ctx, ac.DB.Conn, id); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	Response(w, user, 200)
}

func (ac *AccessCloude) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	ctx := context.Background()

	users, err := entity.GetUsers(ctx, ac.DB.Conn)
	if err != nil {
		Response(w, err.Error(), 500)
		return
	}

	Response(w, users, 200)
}

func (acu *AccessCloude) Pong(w http.ResponseWriter, r *http.Request) {
	Response(w, "ping", 200)
}

func (ac *AccessCloude) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	var (
		user entity.User
		ctx  = context.Background()
	)

	if err := UnmarshalObject(r, &user); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	if err := user.Update(ctx, ac.DB.Conn, id); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	Response(w, "user updated!", 200)

}

func (ac *AccessCloude) UploadUserAvatar(w http.ResponseWriter, r *http.Request) {

}

func (ac *AccessCloude) UserSignIn(w http.ResponseWriter, r *http.Request) {
	var (
		user entity.User
		ctx  = context.Background()
	)

	if err := UnmarshalObject(r, &user); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	if err := user.SignIn(ctx, ac.DB.Conn, ac.DB.Salt); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	Response(w, user, 200)
}

func (ac *AccessCloude) UserSignUp(w http.ResponseWriter, r *http.Request) {
	var (
		user entity.User
		ctx  = context.Background()
	)

	if err := UnmarshalObject(r, &user); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	if err := user.SignUp(ctx, ac.DB.Conn, ac.DB.Salt); err != nil {
		Response(w, err.Error(), 500)
		return
	}

	Response(w, user, 200)
}
