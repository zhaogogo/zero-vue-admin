package chanhandle

import (
	"fmt"
	"github.com/pkg/errors"
)

type Next struct {
	Name        string
	nextHandler ChanHandler
}

type ChanHandler interface {
	SetNext(handler ChanHandler) ChanHandler
	Exec(*SlienceChan) error
	Do(*SlienceChan) error
	Named() string
}

func (n *Next) Named() string {
	return n.Name
}

func (n *Next) SetNext(handler ChanHandler) ChanHandler {
	n.nextHandler = handler
	return n.nextHandler
}

func (n *Next) Exec(s *SlienceChan) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(s); err != nil {
			s.execed[s.c] = s.execed[s.c] + "(false)"
			return errors.Wrapf(err, "slience chan name %s execute failed", s.execed[s.c])
		}
		s.execed[s.c] = s.execed[s.c] + "(ok)"
		s.c++
		return n.nextHandler.Exec(s)
	}
	return errors.New(fmt.Sprintf("%s-%s nexHandler is null", s.name, n.Named()))
}

type headChanHandler struct {
	Next
}

func (h *headChanHandler) Do(s *SlienceChan) error {
	return nil
}
