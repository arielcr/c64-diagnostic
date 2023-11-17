package diagnostics

type Diagnostic struct {
	ID    string `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type Step struct {
	Step        int     `json:"step"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Test        string  `json:"test"`
	Error       Error   `json:"error"`
	Success     Success `json:"success"`
}

type Error struct {
	Next       Next   `json:"next"`
	Diagnostic string `json:"diagnostic"`
}

type Success struct {
	Next       Next   `json:"next"`
	Diagnostic string `json:"diagnostic"`
}

type Next struct {
	Category string `json:"category"`
	Step     int    `json:"step"`
}

type DiagnosticStatus struct {
	Diagnostic string  `json:"diagnostic" validate:"required"`
	Step       int     `json:"step" validate:"required"`
	Result     string  `json:"result"`
	Results    Results `json:"results"`
}

type Results struct {
	Precheck          Precheck          `json:"precheck"`
	SystemPowerSupply SystemPowerSupply `json:"system_power_supply"`
	SystemReset       SystemReset       `json:"system_reset"`
	SystemClocks      SystemClocks      `json:"system_clocks"`
	SystemVideo       SystemVideo       `json:"system_video"`
	SystemRAM         SystemRAM         `json:"system_ram"`
	SystemRomIo       SystemRomIo       `json:"system_rom_io"`
}

type Precheck struct {
	Num1  string `json:"1"`
	Num2  string `json:"2"`
	Num3  string `json:"3"`
	Num4  string `json:"4"`
	Num5  string `json:"5"`
	Num6  string `json:"6"`
	Num7  string `json:"7"`
	Num8  string `json:"8"`
	Num9  string `json:"9"`
	Num10 string `json:"10"`
	Num11 string `json:"11"`
}

type SystemPowerSupply struct {
	Num1  string `json:"1"`
	Num2  string `json:"2"`
	Num3  string `json:"3"`
	Num4  string `json:"4"`
	Num5  string `json:"5"`
	Num6  string `json:"6"`
	Num7  string `json:"7"`
	Num8  string `json:"8"`
	Num9  string `json:"9"`
	Num10 string `json:"10"`
}

type SystemReset struct {
	Num1 string `json:"1"`
	Num2 string `json:"2"`
}

type SystemClocks struct {
	Num1  string `json:"1"`
	Num2  string `json:"2"`
	Num3  string `json:"3"`
	Num4  string `json:"4"`
	Num5  string `json:"5"`
	Num6  string `json:"6"`
	Num7  string `json:"7"`
	Num8  string `json:"8"`
	Num9  string `json:"9"`
	Num10 string `json:"10"`
}

type SystemVideo struct {
	Num1 string `json:"1"`
	Num2 string `json:"2"`
	Num3 string `json:"3"`
	Num4 string `json:"4"`
	Num5 string `json:"5"`
	Num6 string `json:"6"`
	Num7 string `json:"7"`
}

type SystemRAM struct {
	Num1   string `json:"1"`
	Num2   string `json:"2"`
	Num3   string `json:"3"`
	Num4   string `json:"4"`
	Num5   string `json:"5"`
	Num6   string `json:"6"`
	Num7   string `json:"7"`
	Num8   string `json:"8"`
	Num9   string `json:"9"`
	Three1 string `json:"3.1"`
	Seven1 string `json:"7.1"`
	Seven2 string `json:"7.2"`
	Seven3 string `json:"7.3"`
	Seven4 string `json:"7.4"`
	Nine1  string `json:"9.1"`
	Nine2  string `json:"9.2"`
}

type SystemRomIo struct {
	Num1   string `json:"1"`
	Num2   string `json:"2"`
	Num3   string `json:"3"`
	Num4   string `json:"4"`
	Num5   string `json:"5"`
	Num6   string `json:"6"`
	Num7   string `json:"7"`
	Num8   string `json:"8"`
	Num9   string `json:"9"`
	Num10  string `json:"10"`
	Num11  string `json:"11"`
	Num12  string `json:"12"`
	Num13  string `json:"13"`
	Eight1 string `json:"8.1"`
}
