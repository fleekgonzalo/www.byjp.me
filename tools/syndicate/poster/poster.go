package poster

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type poster struct {
	services map[string]shared.Service
	done     map[shared.SyndicationID]string
}

func New(services map[string]shared.Service) *poster {
	return &poster{
		services: services,
		done:     make(map[shared.SyndicationID]string),
	}
}

func (p *poster) Connect(sname string) error {
	return p.services[sname].Connect(false)
}

func (p *poster) Post(sid shared.SyndicationID, post shared.Post) error {
	if _, ok := p.done[sid]; ok {
		return nil
	}

	url, err := p.services[sid.Source].Post(post)
	if err != nil {
		return err
	}

	p.done[sid] = url

	return nil
}

func (p *poster) PostedURL(sid shared.SyndicationID) string {
	if url, ok := p.done[sid]; ok {
		return url
	} else {
		return ""
	}
}

func (p *poster) ReplaceReferences(fname string, tagMatcher *regexp.Regexp) error {
	f, err := os.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("couldn't read %s to replace posts within it: %w", fname, err)
	}

	tags := tagMatcher.FindAllSubmatch(f, -1)
	if tags == nil {
		return nil
	}

	var replacedURLs []string
	for _, tag := range tags {
		sid := shared.SyndicationID{Source: string(tag[1]), ID: string(tag[2])}
		if url, ok := p.done[sid]; ok {
			replacedURLs = append(replacedURLs, url)
			f = bytes.ReplaceAll(f, sid.Bytes(), []byte(url))
		}
	}

	if len(replacedURLs) == 0 {
		return nil
	}

	if err := os.WriteFile(fname, f, 0644); err != nil {
		return fmt.Errorf("couldn't write %s after replacing post URLs (%s) within it: %w", fname, strings.Join(replacedURLs, ", "), err)
	}

	return nil
}
