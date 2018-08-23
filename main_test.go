package db

import "testing"

func TestBuildSQL(t *testing.T) {
	type args struct {
		userID    int
		startDate string
		endDate   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1548",
			args: args{11, "2018-05-01", "2018-08-01"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuildSQL(tt.args.userID, tt.args.startDate, tt.args.endDate)
		})
	}
}
