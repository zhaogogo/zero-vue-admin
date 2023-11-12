package chanhandle

type SlienceChan struct {
	name           string
	c              int
	len            int
	headChanHandle ChanHandler
	execed         []string
}

func (s *SlienceChan) Exec(sliencechan *SlienceChan) error {
	return sliencechan.headChanHandle.Exec(sliencechan)
}

func NewSlienceChan(actions ...ChanHandler) *SlienceChan {
	s := &SlienceChan{
		name:   "slience_chan",
		c:      0,
		len:    len(actions),
		execed: make([]string, 0, len(actions)),
	}
	var handler ChanHandler = &headChanHandler{Next{Name: "1111"}}
	s.headChanHandle = handler
	for _, action := range actions {
		s.execed = append(s.execed, action.Named())
		handler = handler.SetNext(action)
	}
	return s
}
