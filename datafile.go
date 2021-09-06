package datafile

import (
	"bufio"
	"errors"
	"os"
)

type Date struct {
	year  int
	month int
	day   int
}

type Event struct {
	Title string
	Date
}

func (d *Date) SetYear(year int) error {

	if year < 1 {
		return errors.New("invalid")
	}

	d.year = year

	return nil
}

func (d *Date) SetMonth(month int) error {

	if month < 1 || month > 12 {
		return errors.New("invalid")
	}

	d.month = month

	return nil
}

func (d *Date) SetDay(day int) error {

	if day < 1 || day > 31 {
		return errors.New("invalid")
	}

	d.day = day

	return nil
}

/* Test time 11.02 */
func GetStrings(filename string) ([]string, error) {

	var lines []string

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
