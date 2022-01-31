package zoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBestAnimal(t *testing.T) {
	t.Run("gopher is the best animal", func(t *testing.T) {
		got, err := IsBestAnimal("gopher")

		assert.Equal(t, got, true, "lol - it failed")
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("IsBestAnimal(%q) returned an error %v", "gopher", err)
		}
		if !got {
			t.Errorf("IsBestAnimal(%q) = %v, want %v", "gopher", got, true)
		}
	})

	t.Run("elephant is not the best animal", func(t *testing.T) {
		got, err := IsBestAnimal("elephant")
		if err != nil {
			t.Errorf("IsBestAnimal(%q) returned an error %v", "elephant", err)
		}
		if got {
			t.Errorf("IsBestAnimal(%q) = %v, want %v", "elephant", got, false)
		}
	})

	t.Run("empty animal is not the best animal", func(t *testing.T) {
		got, err := IsBestAnimal("")
		if err == nil {
			t.Errorf("IsBestAnimal(%q) did not return an error", "")
		}
		if got {
			t.Errorf("IsBestAnimal(%q) = %v, want %v", "", got, false)
		}
	})
}

// func TestIsBestAnimal(t *testing.T) {
// 	type args struct {
// 		animal string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "happy path - gopher",
// 			args: args{
// 				animal: "gopher",
// 			},
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name: "should fail - elephant",
// 			args: args{
// 				animal: "elephant",
// 			},
// 			want:    false,
// 			wantErr: false,
// 		},
// 		{
// 			name: "should fail - empty string",
// 			args: args{
// 				animal: "",
// 			},
// 			want:    false,
// 			wantErr: true,
// 		},
// 		{
// 			name:    "should fail - empty string 2",
// 			wantErr: true,
// 		},
// 		{
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// t.Parallel()
// 			got, err := IsBestAnimal(tt.args.animal)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("IsBestAnimal() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			// assert.Equal(t, tt.want, got)
// 			if got != tt.want {
// 				t.Errorf("IsBestAnimal() = %v, want %v", got, tt.want)
// 			}

// 		})
// 	}
// }
