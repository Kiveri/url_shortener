package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type URL struct {
	ShortUrl string
	LongUrl  string
}

func NewURL(shortUrl, longURL string) *URL {
	return &URL{
		ShortUrl: shortUrl,
		LongUrl:  longURL,
	}
}

func (u *URL) Value() (driver.Value, error) {
	url, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("error marshalling URL: %s", err)
	}

	return string(url), nil
}

func (u *URL) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("not string")
	}

	return json.Unmarshal(b, &u)
}

//func GenerateShortURL(length int) string {
//	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
//	b := make([]byte, length)
//	for i := range b {
//		b[i] = charset[rand.Intn(len(charset))]
//	}
//	return string(b)
//}
