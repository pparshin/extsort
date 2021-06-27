package merge

import (
	"bufio"
	"errors"
	"os"

	"github.com/pparshin/extsort/internal/ioutil"
)

var (
	ErrNothingMerge = errors.New("merger: no files to merge")
)

type Merger struct {
	files []string
}

func NewMerger() *Merger {
	return &Merger{
		files: make([]string, 0),
	}
}

func (m *Merger) Add(filename string) {
	m.files = append(m.files, filename)
}

func (m *Merger) Merge() (string, error) {
	if len(m.files) == 0 {
		return "", ErrNothingMerge
	}

	for len(m.files) > 1 {
		err := m.mergeTwo()
		if err != nil {
			return "", err
		}
	}

	return m.files[0], nil
}

func (m *Merger) mergeTwo() error {
	merged := make([]string, 0)
	for i := 0; i < len(m.files)-1; i += 2 {
		res, err := mergeSorted(m.files[i], m.files[i+1])
		if err != nil {
			return err
		}
		merged = append(merged, res)
	}

	if len(m.files)%2 != 0 {
		merged = append(merged, m.files[len(m.files)-1])
	}

	m.files = merged

	return nil
}

func mergeSorted(filename1, filename2 string) (filename string, err error) {
	f1, err := os.Open(filename1)
	if err != nil {
		return
	}
	defer f1.Close()

	f2, err := os.Open(filename2)
	if err != nil {
		return
	}
	defer f2.Close()

	fRes, err := ioutil.CreateTempFile()
	if err != nil {
		return
	}
	defer fRes.Close()

	w := bufio.NewWriter(fRes)

	leftScanner := bufio.NewScanner(f1)
	rightScanner := bufio.NewScanner(f2)

	var left, right string
	if leftScanner.Scan() {
		left = leftScanner.Text()
	}
	if rightScanner.Scan() {
		right = rightScanner.Text()
	}

	for left != "" && right != "" {
		var next string
		if left < right {
			next = left
			if leftScanner.Scan() {
				left = leftScanner.Text()
			} else {
				left = ""
				next += "\n" + right
			}
		} else {
			next = right
			if rightScanner.Scan() {
				right = rightScanner.Text()
			} else {
				right = ""
				next += "\n" + left
			}
		}

		_, err = w.WriteString(next + "\n")
		if err != nil {
			return
		}
	}

	for leftScanner.Scan() {
		left = leftScanner.Text()
		_, err = w.WriteString(left + "\n")
		if err != nil {
			return
		}
	}

	for rightScanner.Scan() {
		right = rightScanner.Text()
		_, err = w.WriteString(right + "\n")
		if err != nil {
			return
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}

	return fRes.Name(), nil
}
