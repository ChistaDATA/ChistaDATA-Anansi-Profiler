package parsers

import "github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"

type IParser interface {
	SetInfoCorpus(infoCorpus *stucts.InfoCorpus)
	GetInfoCorpus() *stucts.InfoCorpus
	Parse(logLine string)
	IsUsable(version string, database string) bool
}
