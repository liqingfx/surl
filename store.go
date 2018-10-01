package surl

import (
	"net/url"
)

type Store interface {
	NextID() (int64, error)
	Insert(int64, *url.URL) error
	Query(int64) (*url.URL, error)
}
