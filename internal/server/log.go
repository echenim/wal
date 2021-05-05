package server

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	record.Offset = uint64(len(c.Records))
	c.Records = append(c.Records, record)
	return record.Offset, nil
}
