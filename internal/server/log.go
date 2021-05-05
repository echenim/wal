package server

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {

}
