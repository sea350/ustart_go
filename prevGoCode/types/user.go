package types

type User struct {
	UserID     string   	`json:UserID`
	Username   string   	`json:Username`
	Password   string   	`json:Password` // Maybe we shouldn't keep it in plain text later?
	Privileges []Privilege 	`json:Privileges`
}
