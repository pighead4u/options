package model

import (
	"fmt"
	"strconv"
)

type TitleVO struct {
	Name      string `json:"name"`
	Contracts string `json:"contracts"`
}

type ContentVO struct {
	Ranking         string `json:"ranking"`
	Company         string `json:"company"`
	Volumn          string `json:"volumn"`
	Change          string `json:"change"`
	TransactionType string `json:"transactionType"`
	Contract        string `json:"contract"`
	Date            string `json:"date"`
}

type SortedContentVOS []ContentVO

// 获取此 slice 的长度
func (p SortedContentVOS) Len() int { return len(p) }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p SortedContentVOS) Less(i, j int) bool {
	if len(p[i].Ranking) == len(p[i].Ranking) {

		i1, err := strconv.Atoi(p[i].Ranking)
		if err != nil {
			fmt.Println(err)
			i1 = 21
		}

		i2, err2 := strconv.Atoi(p[j].Ranking)
		if err2 != nil {
			fmt.Println(err)
			i2 = 21
		}
		return (i1 - i2) < 0
	}

	if len(p[i].Ranking) > len(p[i].Ranking) {
		return false
	}

	return true
}

// 交换数据
func (p SortedContentVOS) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
