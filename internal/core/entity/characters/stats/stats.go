package stats

type Level uint16

type Experience uint64

type Combat struct {
	Level Level
	Exp   Experience
}

type Craft struct {
	Level Level
	Exp   Experience
}

type Gather struct {
	Level Level
	Exp   Experience
}
