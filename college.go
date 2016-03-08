//
// college.go
// 

package main

type College struct {
	votes   int
	dem2008 bool
}

var college = map[string]College{
	"AL": College{50, false},
	"AK": College{28, false},
	"AZ": College{58, false},
	"AR": College{40, false},
	"CA": College{172, true},
	"CO": College{37, true},
	"CT": College{28, true},
	"DE": College{16, true},
	"DC": College{19, true},
	"FL": College{99, true},
	"GA": College{76, false},
	"HI": College{19, true},
	"ID": College{32, false},
	"IL": College{69, true},
	"IN": College{57, true},
	"IA": College{30, true},
	"KS": College{40, false},
	"KY": College{46, false},
	"LA": College{46, false},
	"ME": College{23, true},
	"MD": College{38, true},
	"MA": College{42, true},
	"MI": College{59, true},
	"MN": College{38, true},
	"MS": College{40, false},
	"MO": College{52, false},
	"MT": College{27, false},
	"NE": College{36, false},
	"NV": College{30, true},
	"NH": College{23, true},
	"NJ": College{51, true},
	"NM": College{24, true},
	"NY": College{95, true},
	"NC": College{72, true},
	"ND": College{28, false},
	"OH": College{66, true},
	"OK": College{43, false},
	"OR": College{28, true},
	"PA": College{71, true},
	"RI": College{19, true},
	"SC": College{50, false},
	"SD": College{29, false},
	"TN": College{58, false},
	"TX": College{155, false},
	"UT": College{40, false},
	"VT": College{16, true},
	"VA": College{49, true},
	"WA": College{44, true},
	"WV": College{34, false},
	"WI": College{42, true},
	"WY": College{29, false},
}
