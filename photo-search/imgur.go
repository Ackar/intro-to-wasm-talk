package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type Imgur struct {
	clientID string
}

func NewImgur(clientID string) *Imgur {
	return &Imgur{clientID: clientID}
}

type searchRes struct {
	Data []searchItem
}

type searchItem struct {
	Images []itemImages
}

type itemImages struct {
	Link string
}

func (i *Imgur) Search(q string) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.imgur.com/3/gallery/search/?q="+url.PathEscape(q), nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}

	h := req.Header
	h.Add("Authorization", "Client-ID "+i.clientID)
	req.Header = h
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to request imgur")
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 300 {
		return nil, errors.Errorf("imgur returned an error %d", res.StatusCode)
	}

	var r searchRes
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, errors.Wrap(err, "unable decode response")
	}

	var links []string
	for _, si := range r.Data {
		if len(si.Images) == 0 {
			continue
		}
		l := si.Images[0].Link
		if !strings.HasSuffix(l, "jpg") {
			continue
		}
		links = append(links, l)
	}

	return links, nil
}
