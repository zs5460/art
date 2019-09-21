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
		{name: "not exist", args: args{text: "\n"}, want: " \n \n \n \n \n "},
		{name: "null", args: args{text: ""}, want: ""},
		{name: "zs", args: args{text: "zs"}, want: `             
 ____   _____
/_  /  / ___/
 / /_ (__  ) 
/___//____/  
             `},
		{name: "sign", args: args{text: "+_"}, want: `          
    __    
 __/ /_   
/_  __/   
 /_/______
   /_____/`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.text); got != tt.want {
				t.Errorf("String() = %s, want %s", got, tt.want)
			}
		})
	}
}
