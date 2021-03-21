package users

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"graphqlgo/internal/pkg/db/mysql"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	// NOTE: Store password as plain text.
	Password string `json:"password"`
}

func (user *User) Create() {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	print(stmt)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = stmt.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("SELECT ID from Users WHERE Username=?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return 0, err
	}

	return Id, nil
}

func (user *User) Authenticate() bool {
	stmt, err := database.Db.Prepare("SELECT Password from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err != sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}
