package util

import "time"

func TimeNowStd() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RFC3339ToStd(t string) (string, error) {
	a, e := time.Parse(time.RFC3339, t)
	if e != nil {
		return "", e
	}
	return a.Format("2006-01-02 15:04:05"), nil
}
