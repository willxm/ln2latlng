package ln2latlng

import "testing"
import "fmt"

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
			"location",
			args{"33032619950707471X"},
			"浙江",
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
