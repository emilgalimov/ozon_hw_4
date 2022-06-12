package storageService

import (
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/storage"
	"testing"
)

func Test_equalProductSets(t *testing.T) {
	type args struct {
		productIDs []int
		products   []storage.StoreItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal when both empty",
			args: args{
				productIDs: []int{},
				products:   []storage.StoreItem{},
			},
			want: true,
		},
		{
			name: "compares incorrect when equals",
			args: args{
				productIDs: []int{
					1,
					2,
				},
				products: []storage.StoreItem{},
			},
			want: false,
		},
		{
			name: "compares correct when equals",
			args: args{
				productIDs: []int{
					1,
					2,
				},
				products: []storage.StoreItem{
					{
						ProductID:  1,
						OrderID:    1,
						IsReserved: false,
					},
					{
						ProductID:  2,
						OrderID:    1,
						IsReserved: false,
					},
				},
			},
			want: true,
		},
		{
			name: "not equal",
			args: args{
				productIDs: []int{
					1,
					2,
				},
				products: []storage.StoreItem{
					{
						ProductID:  1,
						OrderID:    1,
						IsReserved: false,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalProductSets(tt.args.productIDs, tt.args.products); got != tt.want {
				t.Errorf("equalProductSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
