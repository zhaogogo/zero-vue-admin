package chanhandle

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type One struct {
	Next
}

func (o *One) Do(s *SlienceChan) error {
	fmt.Println("one do...")
	return nil
}

type Two struct {
	Next
}

func (o *Two) Do(s *SlienceChan) error {
	fmt.Println("Two do")
	return nil
}

type Thread struct {
	Next
}

func (o *Thread) Do(s *SlienceChan) error {
	fmt.Println("Thread do")
	return errors.New("Thread error")
}

type Four struct {
	Next
}

func (o *Four) Do(s *SlienceChan) error {
	fmt.Println("Four do")
	return nil
}

type Five struct {
	Next
}

func (o *Five) Do(s *SlienceChan) error {
	fmt.Println("Five do")
	return nil
}

func TestSlience(t *testing.T) {
	a := NewSlienceChan(&One{Next{Name: "one"}}, &Two{Next{Name: "two"}}, &Thread{Next{Name: "thread"}}, &Four{Next{Name: "for"}}, &Five{Next{Name: "five"}})
	err := a.Exec(a)
	if err != nil {
		t.Log(a.execed)
		t.Error(err)
		return
	}
	t.Log(a.execed)
}
