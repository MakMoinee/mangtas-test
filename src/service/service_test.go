package service

import (
	"reflect"
	"testing"
)

func Test_service_GetWords(t *testing.T) {
	word := Words{}
	list := []Words{}
	word.Occurances = 2
	word.Data = "michael is the best michael because he is michael."
	list = append(list, word)
	type args struct {
		topTenWords []Words
	}
	tests := []struct {
		name string
		svc  *service
		args args
	}{
		{"Scenario1", &service{}, args{[]Words{}}},
		{"Scenario2", &service{}, args{list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.svc.PrintWords(tt.args.topTenWords)
		})
	}
}

func Test_service_GetTopTenWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		svc  *service
		args args
		want []Words
	}{
		{"Scenario1", &service{}, args{""}, []Words{}},
		{"Scenario1", &service{}, args{"michael is the best michael because he is michael"}, []Words{{3, "michael"}, {2, "is"}, {1, "the"}, {1, "best"}, {1, "because"}, {1, "he"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.svc.GetTopTenWords(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetTopTenWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
