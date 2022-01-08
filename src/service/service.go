package service

import (
	"fmt"
	"sort"
	"strings"
)

// service data model
type service struct {
}

type Words struct {
	Occurances int
	Data       string
}

type ITopten interface {
	GetTopTenWords(s string) []Words
	PrintWords(topTenWords []Words)
}

func NewService() ITopten {
	svc := service{}

	return &svc
}

// GetTopTenWords retrieve the top ten wors used in string given
func (svc *service) GetTopTenWords(s string) []Words {
	str := strings.TrimSpace(s)
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, ".", "")
	hashMap := make(map[string]int)

	myWords := []Words{}
	if len(str) > 0 {
		sliceStr := strings.Split(str, " ") //create slice of strings by separating it by space
		for i := 0; i < len(sliceStr); i++ {
			stringData := strings.ToLower(sliceStr[i])
			if val, exist := hashMap[stringData]; exist {
				hashMap[stringData] = val + 1
				continue
			}
			hashMap[stringData] = 1

		}

		for key, val := range hashMap { // get the occurances and the word and store it in the Word struct
			singleWord := Words{}
			singleWord.Occurances = val
			singleWord.Data = key
			myWords = append(myWords, singleWord)
		}

		sort.Slice(myWords, func(i, j int) bool { //sort slice to decending
			return myWords[i].Occurances > myWords[j].Occurances
		})
	}
	return myWords
}

// GetWords prints the word and occurances
func (svc *service) PrintWords(topTenWords []Words) {
	if len(topTenWords) == 0 {
		fmt.Println("No data")
		return
	} else {
		for i := 0; i < len(topTenWords); i++ {
			word := topTenWords[i]
			fmt.Println(i+1, fmt.Sprintf(" Word:%v  Occurances:%v", word.Data, word.Occurances))
			if i == 9 {
				break
			}
		}
	}

}
