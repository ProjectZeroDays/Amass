// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package sources

type LoCArchive struct {
	BaseDataSource
	baseURL string
}

func NewLoCArchive() DataSource {
	la := &LoCArchive{baseURL: "http://webarchive.loc.gov/all"}

	la.BaseDataSource = *NewBaseDataSource(ARCHIVE, "LoC Archive")
	return la
}

func (la *LoCArchive) Query(domain, sub string) []string {
	if sub == "" {
		return []string{}
	}
	return runArchiveCrawler(la.baseURL, domain, sub, la)
}

func (la *LoCArchive) Subdomains() bool {
	return true
}
