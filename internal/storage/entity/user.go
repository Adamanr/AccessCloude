package storage

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	// 	lainRole    Role = "Lain"
	// 	adminRole   Role = "admin"
	memberRole string = "member"
	// 	mentorRole  Role = "mentor"
	// 	teacherRole Role = "teacher"
)

type User struct {
	Id          int64     `json:"id,omitempty"`
	Login       string    `json:"login"`
	Email       string    `json:"email"`
	AvatarId    string    `json:"avatar_id,omitempty"`
	Description string    `json:"description,omitempty"`
	Role        []string  `json:"role,omitempty"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Userer interface {
	GetByID(ctx context.Context, db *pgx.Conn, id int) error
	SignIn(ctx context.Context, db *pgx.Conn, salt string) error
	SignUp(ctx context.Context, db *pgx.Conn, salt string) error
	Update(ctx context.Context, db *pgx.Conn, id int) error
	Delete(ctx context.Context, db *pgx.Conn, id int) error
	ForgetPassword(ctx context.Context, db *pgx.Conn, id int) (string, error)
}

func (u *User) String() string {
	return fmt.Sprintf("{ID: %d, Login: %s, Email: %s, AvatarId: %s, Description: %s, Role: %v, Password: %s}",
		u.Id, u.Login, u.Email, u.AvatarId, u.Description, u.Role, u.Password)
}

var _ Userer = (*User)(nil)

func GetUsers(ctx context.Context, db *pgx.Conn) ([]*User, error) {
	query := `SELECT * FROM users`

	rows, err := db.Query(ctx, query)
	if err != nil {
		comment := fmt.Errorf("error getting all users: %s", err.Error())
		slog.Error(comment.Error())
		return nil, comment
	}

	users := make([]*User, 0)
	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.Description, &u.AvatarId, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
			comment := fmt.Errorf("error scanning user: %s", err.Error())
			slog.Error(comment.Error())
			return nil, comment
		}

		users = append(users, u)
	}

	return users, nil
}

func (u *User) SignIn(ctx context.Context, db *pgx.Conn, salt string) error {
	var (
		password string
	)

	query := `SELECT * FROM users WHERE email = $1`
	if err := db.QueryRow(ctx, query, u.Email).Scan(&u.Id, &u.Login, &u.Email, &password, &u.Description, &u.AvatarId, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
		comment := fmt.Errorf("error get user by login: %s", err.Error())
		return comment
	}

	hashP := CheckPasswordWithSalt(u.Password, salt, password)
	fmt.Print("P", hashP, " ", u.Password, "\n")

	if !hashP {
		comomment := fmt.Errorf("incorrect password")
		return comomment
	}

	return nil
}

func (u *User) SignUp(ctx context.Context, db *pgx.Conn, salt string) error {
	password, err := HashPasswordWithSalt(u.Password, salt)
	if err != nil {
		commment := fmt.Errorf("error hashing password: %s", err.Error())
		return commment
	}

	if u.Role == nil {
		u.Role = append(u.Role, memberRole)
	}

	query := `INSERT INTO users (login, email, avatar_id, description, role, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if err := db.QueryRow(ctx, query, u.Login, u.Email, u.AvatarId, u.Description, u.Role, password).Scan(&u.Id); err != nil {
		comment := fmt.Errorf("error inserting user: %s", err.Error())
		return comment
	}

	return nil

}

func (u *User) Delete(ctx context.Context, db *pgx.Conn, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	if _, err := db.Exec(ctx, query, id); err != nil {
		comment := fmt.Errorf("error deleting user: %s", err.Error())
		return comment
	}

	return nil
}

// TODO: Сделать работу с нотификациями
func (u *User) ForgetPassword(ctx context.Context, db *pgx.Conn, id int) (string, error) {
	return "", nil
}

func (u *User) GetByID(ctx context.Context, db *pgx.Conn, id int) error {
	query := `SELECT * FROM users WHERE id = $1`
	if err := db.QueryRow(ctx, query, id).Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.Description, &u.AvatarId, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
		comment := fmt.Errorf("error getting user: %s", err.Error())
		slog.Error(comment.Error())

		return comment
	}

	return nil
}

func (u *User) Update(ctx context.Context, db *pgx.Conn, id int) error {

	args := make([]interface{}, 0)

	query := checkFieldUpdate(*u, &args, id)
	if args == nil {
		return nil
	}

	if _, err := db.Exec(ctx, query, args...); err != nil {
		comment := fmt.Errorf("error updating user: %s", err.Error())
		return comment
	}

	return nil
}

func checkFieldUpdate(u User, args *[]interface{}, id int) string {
	i := 1
	query := "UPDATE users SET"

	if u.Login != "" {
		query += fmt.Sprintf(" login = $%d, ", i)
		*args = append(*args, u.Login)
		i++
	}

	if u.Email != "" {
		query += fmt.Sprintf(" email = $%d, ", i)
		*args = append(*args, u.Email)
		i++
	}

	if u.Password != "" {
		query += fmt.Sprintf(" password = $%d, ", i)
		*args = append(*args, u.Password)
		i++
	}

	if u.Description != "" {
		query += fmt.Sprintf(" description = $%d, ", i)
		*args = append(*args, u.Description)
		i++
	}

	if u.AvatarId != "" {
		query += fmt.Sprintf(" avatar_id = $%d, ", i)
		*args = append(*args, u.AvatarId)
		i++
	}

	if u.Role != nil {
		query += fmt.Sprintf(" role = $%d, ", i)
		*args = append(*args, u.Role)
		i++
	}

	if i != 1 {
		*args = append(*args, id)
		query = query[:len(query)-2]
		query += fmt.Sprintf(" WHERE id = $%d", len(*args))
	}

	return query
}

func HashPasswordWithSalt(password, salt string) (string, error) {
	passwordWithSalt := password + salt

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordWithSalt(password, salt, hashedPassword string) bool {
	passwordWithSalt := password + salt

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithSalt))

	return err == nil
}
