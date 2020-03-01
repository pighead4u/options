package model

import "sort"

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

func QueryContentByContractAndDateAndType(contract, date, contractType string) SortedContentVOS {
	var contents []Content
	db.Where("contract_code = ? AND transaction_date = ? AND transaction_type = ?", contract, date, contractType).Find(&contents)
	//db.Where("contract_code = ? AND transaction_date = ? AND transaction_type = ? AND ranking != ?", contract, date, contractType, "合计").Find(&contents)
	//db.Where("contract_code = ? AND transaction_date = ? AND transaction_type = ? AND ranking = ?", contract, date, contractType, "合计").Find(&heji)
	var tmp SortedContentVOS
	var vos []ContentVO
	for _, value := range contents {
		vos = append(vos, value.translate())
	}
	tmp = vos
	sort.Sort(tmp)

	return tmp
}
