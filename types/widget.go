package types

type Widget struct{
	WidgetType	string	`json:WidgetType`
	Position	int		`json:Position`
	Content		string	`json:Content`

}