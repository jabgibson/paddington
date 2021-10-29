package paddington

// Now having created a second solution to the same problem... rather than using a second set of tests,
// I decided to share the tests since functionally both solutions shouldn't care about anything other than
// the signature
type WNPTest struct {
	name string
	args WNPArgs
	want string
}

type WNPArgs struct {
	item string
	pad  int
}

var wnpTests []WNPTest

func init() {
	wnpTests = []WNPTest{
		{
			name: "Give me James Bond",
			args: WNPArgs{
				item: "James Bond 7",
				pad:  3,
			},
			want: "James Bond 007",
		},
		{
			name: "Pi",
			args: WNPArgs{
				item: "PI=3.14",
				pad:  2,
			},
			want: "PI=03.14",
		},
		{
			name: "Trick Pi",
			args: WNPArgs{
				item: "PI=3.14",
				pad:  3,
			},
			want: "PI=003.14",
		},
		{
			name: "clock time",
			args: WNPArgs{
				item: "It's 3:13pm",
				pad:  2,
			},
			want: "It's 03:13pm",
		},
		{
			name: "clock time 2",
			args: WNPArgs{
				item: "It's 12:13pm",
				pad:  2,
			},
			want: "It's 12:13pm",
		},
		{
			name: "nonsense",
			args: WNPArgs{
				item: "99UR1337",
				pad:  6,
			},
			want: "000099UR001337",
		},
	}
}
