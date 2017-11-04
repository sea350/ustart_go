package universal

import "errors"

//FindIndex ... FINDS THE FIRST LOCATION OF AN ELEMENT IN AN ARRAY
//Requires an array and the specific object you want to find
//Returns the index of that element, will return -1 if doesnt exist
func FindIndex(slice []interface{}, element interface{}) int {
	index := -1
	for i := range slice {
		if slice[i] == element {
			return i
		}
	}
	return index
}

//RemoveByIndex ... REMOVES A SINGLE ELEMENT FROM AN ARRAY
//Requires an array and the index of the element that needs to be removed
//Returns the cleaned array and an error
func RemoveByIndex(slice []interface{}, index int) ([]interface{}, error) {

	if (len(slice)-1) <= index || index < 0 {
		return nil, errors.New("index is out of bounds")
	}
	return append(slice[:index], slice[index+1:]...), nil
}
