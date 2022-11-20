package aoc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"
)

func Input(year, day int) (io.ReadCloser, error) {
	session, err := os.ReadFile("aoc-session")
	if err != nil {
		return nil, err
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: strings.TrimSpace(string(session)),

		Path:    "/",
		Domain:  "adventofcode.com",
		Expires: time.Now().Add(time.Hour),

		MaxAge:   0,
		Secure:   true,
		HttpOnly: false,
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest(http.MethodGet, inputURL(year, day), nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&sessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		errorText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorText))
	}
	return resp.Body, nil
}

func inputURL(year, day int) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
}
