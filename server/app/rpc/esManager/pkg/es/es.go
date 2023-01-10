package es

type ES interface {
	Ping() (res interface{}, err error)
}

func NewES(url string, user string, password string, version int64) ES {
	switch version {
	case 7:
		return &es7{
			URL:      url,
			User:     user,
			PassWord: password,
		}
	default:
		return nil
	}
}
