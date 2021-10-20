package paddington

import "testing"

func TestWholeNumPad(t *testing.T) {
	type args struct {
		item    string
		padding int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Give me James Bond",
			args: args{
				item:    "James Bond 7",
				padding: 3,
			},
			want: "James Bond 007",
		},
		{
			name: "Pi",
			args: args{
				item:    "PI=3.14",
				padding: 2,
			},
			want: "PI=03.14",
		},
		{
			name: "clock time",
			args: args{
				item:    "It's 3:13pm",
				padding: 2,
			},
			want: "It's 03:13pm",
		},
		{
			name: "clock time 2",
			args: args{
				item:    "It's 12:13pm",
				padding: 2,
			},
			want: "It's 12:13pm",
		},
		{
			name: "nonsense",
			args: args{
				item:    "99UR1337",
				padding: 6,
			},
			want: "000099UR001337",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WholeNumPad(tt.args.item, tt.args.padding); got != tt.want {
				t.Errorf("WholeNumPad() = %v, want %v", got, tt.want)
			}
		})
	}
}
