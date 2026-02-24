package database

import (
	"fmt"
)

type GormConfig struct {
	MaxIdleConns int
	MaxOpenConns int

	dcn *DCN
}

// Return Data Source Name
func (c *GormConfig) GetDCN() string {
	return c.dcn.GetString()
}

// Data Source Name
//
// Used to connect to a database
type DCN struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	Sslmode  string
}

// Return Data Source Name to connect to db in string format:
//
// host=%s user=%s password=%s dbname=%s port=%d sslmode=%s
func (c *DCN) GetString() string {
	var mode string
	if c.Sslmode == "" {
		mode = "disable"
	} else {
		mode = c.Sslmode
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, mode,
	)
}
