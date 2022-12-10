package go_logger

type Config struct {
	Director    string
	Level       string
	FileExt     string
	FileName    string
	LinkName    string
	Format      string
	WithConsole bool
	MaxAge      int
}
