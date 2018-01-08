package get

//EmailToUsername ...
func EmailToUsername(email string) string {

	var usr []rune //in case of special characters, use runes

	for _, element := range email { //iterate through email string
		if element != '@' { //remove '@'
			usr = append(usr, element)
		} else {
			usr = append(usr, '.') //replace '@' with '.'
		}
	}

	retUsr := string(usr) //converts to string for username

	return retUsr //returns username

}
