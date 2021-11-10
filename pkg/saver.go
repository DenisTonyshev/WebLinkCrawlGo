package pkg

type SaverFunc func(u string) error

type Saver interface {
	Save(val string) error
}

func NewFileSaver(value string) *FileSaver {
	return &FileSaver{value: value}
}

type FileSaver struct {
	value     string
	ResultMap map[string]bool
	tempMap   map[string]bool
}

func (fs *FileSaver) Save(value string) error {
	if !fs.ResultMap[value] {
		fs.ResultMap[value] = true
		fs.tempMap[value] = true
	}
	return nil
}
