package model

func QueryTitleByBourse(bourse string) map[string][]string {
	var bourses []Exchange
	db.Where("bourse = ?", bourse).Find(&bourses)

	result := make(map[string][]string)

	for _, value := range bourses {

		//do something here
		result[value.Name] = append(result[value.Name], value.Contract)

	}

	return result
}
