package csv

import (
	"encoding/csv"
	"os"
)

type Csv struct {
	Fields   []string
	Filename string
	Writer   *csv.Writer
	IsTitle  bool
}

func NewCsv(name string, fields []string) (*Csv, error) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(f)
	return &Csv{
		Fields:   fields,
		Filename: name,
		Writer:   w,
		IsTitle:  false,
	}, nil
}

func (s *Csv) WriteLines(d [][]string) error {
	if !s.IsTitle {
		s.WriteHead()
		s.IsTitle = true
	}
	err := s.WriteAll(d)
	if err != nil {
		return err
	}
	return s.Writer.Error()
}

func (s *Csv) WriteHead() error {
	err := s.Writer.Write(s.Fields)
	return err
}

func (s *Csv) WriteAll(d [][]string) error {
	return s.Writer.WriteAll(d)
}
