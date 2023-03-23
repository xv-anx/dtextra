package modules

type CheatData struct {
	Name         string
	StartAddr    string
	Note         string
	Values       []string
	ReadOnlyAddr []string
	ReadOnlyVal  []string
}

type Cheats struct {
	Data []*CheatData
}
