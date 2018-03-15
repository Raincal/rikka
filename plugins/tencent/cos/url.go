package cos

import (
	"errors"
	"strings"

	"github.com/Raincal/rikka/api"
	"github.com/Raincal/rikka/plugins"
)

func (plugin tccosPlugin) URLRequestHandle(q *plugins.URLRequest) (*api.URL, error) {
	if bucketHost == "" {
		l.Error("Request URL of task", q.TaskID, "before state become to finish")
		return nil, errors.New("Get url before task finish")
	}

	if !strings.Contains(bucketHost, taskIDPlaceholder) {
		l.Fatal("Error! Unable to get image url from bucket host:", bucketHost)
	}

	return &api.URL{
		URL: strings.Replace(bucketHost, taskIDPlaceholder, q.TaskID, -1),
	}, nil
}
