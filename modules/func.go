package modules

import (
	"strings"

	"github.com/fatih/color"
)

// CheatData from modules/data.go

func NewCheats() *Cheats {
	return &Cheats{
		Data: []*CheatData{},
	}
}

func (c *Cheats) AddCheatData(name string) *CheatData {
	cheat := &CheatData{
		Name:         name,
		StartAddr:    "",
		Values:       []string{},
		ReadOnlyAddr: []string{},
		ReadOnlyVal:  []string{},
	}
	c.Data = append(c.Data, cheat)
	return cheat
}

func (c *Cheats) AddValue(val string) {
	cheat := c.Data[len(c.Data)-1]
	cheat.Values = append(cheat.Values, val)
}

func (c *Cheats) AddReadonlyData(addr, val string) {
	cheat := c.Data[len(c.Data)-1]
	cheat.ReadOnlyAddr = append(cheat.ReadOnlyAddr, addr)
	cheat.ReadOnlyVal = append(cheat.ReadOnlyVal, val)
}

func (c *CheatData) PrintCheatData() {
	border := "-------------------------------------------------------------------"
	space := "    "
	color.White("void"+space+"%s(MenuEntry* entry)\n{", c.Name)
	color.Green(space+"// %s\n", c.Note)
	if c.StartAddr != "" && len(c.Values) != 0 {
		color.White(space+"u32 offset = %s;\n", c.StartAddr)
		color.White(space+"const std::vector<u32> data =\n"+space+"{\n"+space+space+"%s"+space+"\n", strings.Join(dataGroupe(c.Values, 4), ", \n"+space+space+"")+"\n"+space+"};")
		color.HiMagenta(space+"memcpy(reinterpret_cast<void*>(offset), &data[0], data.size() * sizeof(u32));\n")
	}
	
	for i := range c.ReadOnlyAddr {
		color.HiYellow(space+"Process::Write32(%s, %s);\n", c.ReadOnlyAddr[i], c.ReadOnlyVal[i])
	}
	color.White("}")
	color.Blue(border)
}

func dataGroupe(slice []string, dataSize int) []string {
	var datas []string
	for i := 0; i < len(slice); i += dataSize {
		end := i + dataSize
		if end > len(slice) {
			end = len(slice)
		}
		datas = append(datas, strings.Join(slice[i:end], ", "))
	}
	return datas
}
