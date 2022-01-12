package errors

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("Error: Website Not Found")
	ErrUnauthorized = errors.New("Error: You are not allowed to perform this action")
	ErrSitemapMissing = errors.New("Error: Sitemap missing!")
	ErrTimeout = errors.New("Error: Timeout")
	ErrRobotsMissing = errors.New("Error: Robots.txt missing")
)
