package gogger

type Color string

type Colors struct {
	Yellow Color
	Red    Color
	Green  Color
	Cyan   Color
	Gray   Color
	White  Color
	Reset  Color
}

type Config struct {
	Time   bool
	Bold   bool
	Colors Colors
}

type Logger struct {
	Config
}
