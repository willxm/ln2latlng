package ln2latlng

import (
	"fmt"
	"testing"
)

func TestIdCrad2LatLng(t *testing.T) {
	type args struct {
		idCard string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   float64
		want2   float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{"33032619950707471X"},
			"??",
			120.15,
			30.28,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := IdCrad2LatLng(tt.args.idCard)
			fmt.Println(got, got1, got2, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdCrad2LngLat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IdCrad2LngLat() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IdCrad2LngLat() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("IdCrad2LngLat() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestIdCard2Area(t *testing.T) {
	type args struct {
		idCard string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				idCard: "33032619xxxx",
			},
			want:    "",
			want1:   "",
			want2:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := IdCard2Area(tt.args.idCard)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdCard2Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IdCard2Area() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IdCard2Area() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("IdCard2Area() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
