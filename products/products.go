package products

import (
	_ "embed"
	"encoding/json"
	"time"
)

//go:embed names.json
var namesJSON []byte
var Names map[string]string

//go:embed repos.json
var reposJSON []byte

type Repo struct {
	Github string `json:"github"`
}

var Repos map[string]Repo

//go:embed languages.json
var languagesJSON []byte

var Languages map[string][]string

//go:embed types.json
var typesJSON []byte

type Type string

const (
	ClientSideType Type = "client-side"
	ServerSideType Type = "server-side"
	EdgeType       Type = "edge"
	RelayType      Type = "relay"
)

var Types map[string]Type

//go:embed releases.json
var releasesJSON []byte

type Release struct {
	Major int        `json:"major"`
	Minor int        `json:"minor"`
	Date  time.Time  `json:"date"`
	EOL   *time.Time `json:"eol"`
}

type ReleaseList []Release

var Releases map[string]ReleaseList

func (r ReleaseList) Earliest() Release {
	return r[len(r)-1]
}

func (r ReleaseList) Latest() Release {
	return r[0]
}

func panicOnError(err error) {
	if err != nil {
		panic("couldn't initialize SDK Metadata module: " + err.Error())
	}
}

func init() {
	panicOnError(json.Unmarshal(namesJSON, &Names))
	panicOnError(json.Unmarshal(reposJSON, &Repos))
	panicOnError(json.Unmarshal(languagesJSON, &Languages))
	panicOnError(json.Unmarshal(typesJSON, &Types))
	panicOnError(json.Unmarshal(releasesJSON, &Releases))
}
