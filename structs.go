package asciiart

type Art struct {
	Output     string
	Standard   string
	Shadow     string
	Thinkertoy string
	Input      string
	Color      string
	BackColor  string
	Year       int
	Status     bool
	ErrorMsg   string
	ErrorCode  int
}

type Err struct {
	Status    bool
	ErrorMsg  string
	ErrorCode int
	Year      int
}
