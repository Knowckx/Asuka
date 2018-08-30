package main

import (
	"fmt"
	"reflect"
	"testing"

	goqu "gopkg.in/doug-martin/goqu.v3"
)

func Test_buildTraderInfoSet(t *testing.T) {
	type args struct {
		userID    int
		startDate string
		endDate   string
	}
	tests := []struct {
		name string
		args args
		want *goqu.Dataset
	}{
		{
			name: "1548",
			args: args{107041, "2018-04-01", "2018-07-01"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTraderInfoSet(tt.args.userID, tt.args.startDate, tt.args.endDate)
			SQL, _, _ := got.ToSql()
			fmt.Println(SQL)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTraderInfoSet() = %v, want %v", got, tt.want)
			}

		})
	}
}
