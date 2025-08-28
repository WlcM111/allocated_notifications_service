package db_service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type Notification struct {
	ID      int
	UserID  int
	Channel string
	Message string
	Status  string
}

type Repository struct {
	DB *sql.DB
}

func (r *Repository) SaveNewNotification(n Notification) error {
	_, err := r.DB.Exec(
		"INSERT INTO notifications(user_id, channel, message, status) VALUES($1, $2, $3, $4)",
		n.UserID, n.Channel, n.Message, n.Status,
	)

	return err
}

func (r *Repository) GetUser(userID int) (*User, error) {
	rw := r.DB.QueryRow("SELECT id, name, email, phone FROM users WHERE id = $1", userID)
	u := &User{}

	if err := rw.Scan(&u.ID, &u.Name, &u.Email, &u.Phone); err != nil {
		return nil, err
	}

	return u, nil
}

func CreateNewDB(cfgHost string, cfgPort int, user, password, dbname string) *Repository {
	cn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfgHost, cfgPort, user, password, dbname)

	db, err := sql.Open("postgres", cn)

	if err != nil {
		log.Fatalf("cannot connect db: %v", err)
	}

	return &Repository{DB: db}
}
