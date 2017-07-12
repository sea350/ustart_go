package types

type Course struct {
	CourseID   string   `json:CourseID`
	CourseName string   `json:CourseName`
	BoardIDs   []string `json:BoardIDs`
}
