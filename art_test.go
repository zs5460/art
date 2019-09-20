package art

import (
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "zs", args: args{text: "zs"}, want: `             
 ____   _____
/_  /  / ___/
 / /_ (__  ) 
/___//____/  
             `},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.text); got != tt.want {
				t.Errorf("String() = %s, want %s", got, tt.want)
			}
		})
	}
}
