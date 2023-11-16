package diagnostics

type Diagnostic struct {
	ID    string `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type Step struct {
	Step        int     `json:"step"`
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
	CurrentDiagnostic string  `json:"current_diagnostic" validate:"required"`
	CurrentStep       int     `json:"current_step" validate:"required"`
	InProgress        bool    `json:"in_progress" validate:"required"`
	Results           Results `json:"results"`
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
	Num1  bool `json:"1"`
	Num2  bool `json:"2"`
	Num3  bool `json:"3"`
	Num4  bool `json:"4"`
	Num5  bool `json:"5"`
	Num6  bool `json:"6"`
	Num7  bool `json:"7"`
	Num8  bool `json:"8"`
	Num9  bool `json:"9"`
	Num10 bool `json:"10"`
	Num11 bool `json:"11"`
}

type SystemPowerSupply struct {
	Num1  bool `json:"1"`
	Num2  bool `json:"2"`
	Num3  bool `json:"3"`
	Num4  bool `json:"4"`
	Num5  bool `json:"5"`
	Num6  bool `json:"6"`
	Num7  bool `json:"7"`
	Num8  bool `json:"8"`
	Num9  bool `json:"9"`
	Num10 bool `json:"10"`
}

type SystemReset struct {
	Num1 bool `json:"1"`
	Num2 bool `json:"2"`
}

type SystemClocks struct {
	Num1  bool `json:"1"`
	Num2  bool `json:"2"`
	Num3  bool `json:"3"`
	Num4  bool `json:"4"`
	Num5  bool `json:"5"`
	Num6  bool `json:"6"`
	Num7  bool `json:"7"`
	Num8  bool `json:"8"`
	Num9  bool `json:"9"`
	Num10 bool `json:"10"`
}

type SystemVideo struct {
	Num1 bool `json:"1"`
	Num2 bool `json:"2"`
	Num3 bool `json:"3"`
	Num4 bool `json:"4"`
	Num5 bool `json:"5"`
	Num6 bool `json:"6"`
	Num7 bool `json:"7"`
}

type SystemRAM struct {
	Num1   bool `json:"1"`
	Num2   bool `json:"2"`
	Num3   bool `json:"3"`
	Num4   bool `json:"4"`
	Num5   bool `json:"5"`
	Num6   bool `json:"6"`
	Num7   bool `json:"7"`
	Num8   bool `json:"8"`
	Num9   bool `json:"9"`
	Three1 bool `json:"3.1"`
	Seven1 bool `json:"7.1"`
	Seven2 bool `json:"7.2"`
	Seven3 bool `json:"7.3"`
	Seven4 bool `json:"7.4"`
	Nine1  bool `json:"9.1"`
	Nine2  bool `json:"9.2"`
}

type SystemRomIo struct {
	Num1   bool `json:"1"`
	Num2   bool `json:"2"`
	Num3   bool `json:"3"`
	Num4   bool `json:"4"`
	Num5   bool `json:"5"`
	Num6   bool `json:"6"`
	Num7   bool `json:"7"`
	Num8   bool `json:"8"`
	Num9   bool `json:"9"`
	Num10  bool `json:"10"`
	Num11  bool `json:"11"`
	Num12  bool `json:"12"`
	Num13  bool `json:"13"`
	Eight1 bool `json:"8.1"`
}
