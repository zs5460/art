package art

import (
	"strings"
	"testing"
)

func Test_getArt(t *testing.T) {
	type args struct {
		char byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name:"a",args:args{char:'a'},want:[]string{
			"        ",
			"  ____ _",
			" / __ `/",
			"/ /_/ / ",
			"\\__,_/  ",
			"        ",
		}},
		{name:"b",args:args{char:'b'},want:[]string{
			"    __  ",
			"   / /_ ",
			"  / __ \\",
			" / /_/ /",
			"/_.___/ ",
			"        ",
		}},
		{name:"a",args:args{char:'a'},want:[]string{
			"        ",
			"  ____ _",
			" / __ `/",
			"/ /_/ / ",
			"\\__,_/  ",
			"        ",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getArt(tt.args.char); strings.Join(got,"")!=strings.Join(tt.want,"") {
				t.Errorf("getArt() = %v, want %v", got, tt.want)
			}
		})
	}
}