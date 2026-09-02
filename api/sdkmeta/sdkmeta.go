package sdkmeta

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

//go:embed data/names.json
var namesJSON []byte

// Names is a map of SDK IDs to display names.
var Names map[string]string

//go:embed data/repos.json
var reposJSON []byte

// Repo contains the location of an SDK.
type Repo struct {
	// GitHub is the GitHub repo path including the owner and repo name (e.g. launchdarkly/js-core).
	GitHub string `json:"github"`
}

// Repos is a map of SDK IDs to repository information.
var Repos map[string]Repo

//go:embed data/languages.json
var languagesJSON []byte

// Languages is a map of SDK IDs to supported languages.
var Languages map[string][]string

//go:embed data/popularity.json
var popularityJSON []byte

// Popularity is a map of SDK IDs to popularity scores.
var Popularity map[string]int

//go:embed data/types.json
var typesJSON []byte

// Type represents the most common use-case for an SDK.
type Type string

const (
	// ClientSideType is an SDK that runs in a client context.
	ClientSideType Type = "client-side"
	// ServerSideType is an SDK that runs in a server context.
	ServerSideType Type = "server-side"
	// EdgeType is an SDK that runs in an edge deployment scenario.
	EdgeType Type = "edge"
	// AIType is an SDK that is primarily focused on AI/ML use cases.
	AIType Type = "ai"
	// OpenFeatureProviderType is an OpenFeature provider.
	OpenFeatureProviderType Type = "open-feature-provider"
	// RelayType is Relay Proxy.
	RelayType Type = "relay"
)

// Types is a map of SDK IDs to SDK types.
var Types map[string]Type

//go:embed data/releases.json
var releasesJSON []byte

type Release struct {
	Major int        `json:"major"`
	Minor int        `json:"minor"`
	Date  time.Time  `json:"date"`
	EOL   *time.Time `json:"eol"`
}

// MajorMinor returns a version string comprised of the major and minor version. For example,
// '2.1'.
func (r Release) MajorMinor() string {
	return fmt.Sprintf("%d.%d", r.Major, r.Minor)
}

// IsLatest returns true if the release is the latest release, meaning there is no EOL date set.
func (r Release) IsLatest() bool {
	return r.EOL == nil
}

// IsEOL returns true if the release is not the latest release and the current time is after the EOL date. The parameter
// represents the current time.
func (r Release) IsEOL(now time.Time) bool {
	return !r.IsLatest() && now.After(*r.EOL)
}

// IsApproachingEOL returns true if the release is not the latest release and the EOL date is within the time period
// from now to now + thresholdPrior. This is only valid if IsEOL() returns false.
func (r Release) IsApproachingEOL(now time.Time, thresholdPrior time.Duration) bool {
	return !r.IsLatest() && now.Add(thresholdPrior).After(*r.EOL)
}

// ReleaseList is an ordered list of releases. The first item should be the most recent release, while the
// last item is the oldest release.
type ReleaseList []Release

var Releases map[string]ReleaseList

// Earliest returns the earliest release.
func (r ReleaseList) Earliest() Release {
	return r[len(r)-1]
}

// Latest returns the latest release.
func (r ReleaseList) Latest() Release {
	return r[0]
}

//go:embed data/user_agents.json
var userAgentsJSON []byte

// SDKUserAgentMap contains user agent and wrapper information for an SDK
type SDKUserAgentMap struct {
	UserAgents   []string `json:"userAgents,omitempty"`
	WrapperNames []string `json:"wrapperNames,omitempty"`
}

// UserAgentMap maps SDK IDs to the user agents and wrapper names they report.
type UserAgentMap map[string]SDKUserAgentMap

// UserAgents is a map of SDK IDs to their user agent and wrapper information
var UserAgents UserAgentMap

// ResolveSDKID finds the SDK ID that reports the given wrapper name or user agent.
// Wrapper names are checked first, then user agents, in alphabetical order by SDK ID.
// Returns the SDK ID and true if found, empty string and false if not found.
//
// The returned ID is the key for Names, Types, Repos, and Releases.
//
// Some identifiers are reported by more than one SDK. For example both akamai-base and
// akamai-edgekv report "AkamaiEdgeSDK". In that case the first SDK ID in alphabetical
// order is returned.
func (m UserAgentMap) ResolveSDKID(identifier string) (string, bool) {
	// Get sorted SDK IDs to ensure consistent ordering
	var sdkIDs []string
	for sdkID := range m {
		sdkIDs = append(sdkIDs, sdkID)
	}
	sort.Strings(sdkIDs)

	// First check wrapper names
	for _, sdkID := range sdkIDs {
		for _, wrapper := range m[sdkID].WrapperNames {
			if wrapper == identifier {
				return sdkID, true
			}
		}
	}

	// Then check user agents
	for _, sdkID := range sdkIDs {
		for _, agent := range m[sdkID].UserAgents {
			if agent == identifier {
				return sdkID, true
			}
		}
	}

	return "", false
}

// GetSDKNameByWrapperOrUserAgent attempts to find an SDK name by first checking wrapper names,
// then user agents, in alphabetical order by SDK ID. Returns the SDK name and true if found,
// empty string and false if not found.
//
// Use UserAgents.ResolveSDKID when you need more than a display name.
func GetSDKNameByWrapperOrUserAgent(identifier string) (string, bool) {
	sdkID, ok := UserAgents.ResolveSDKID(identifier)
	if !ok {
		return "", false
	}
	return Names[sdkID], true
}

//go:embed data/ai_sdk_identifiers.json
var aiSdkIdentifiersJSON []byte

// AISDKIdentifier is one (name, language) pair that an AI SDK reports about itself
// in the $ld:ai:sdk:info custom event.
type AISDKIdentifier struct {
	// Name is the value reported as aiSdkName. This is the package name.
	Name string `json:"name"`
	// Language is the value reported as aiSdkLanguage, for example "python".
	Language string `json:"language"`
}

// AISDKIdentifierMap maps SDK IDs to the identifiers those AI SDKs report.
type AISDKIdentifierMap map[string][]AISDKIdentifier

// AISDKIdentifiers is a map of SDK IDs to the identifiers those AI SDKs report.
//
// An AI SDK wraps a client that the caller supplies, so it sends no user agent of its
// own. It reports itself in a custom event instead. Historical names are kept so that
// older deployments still resolve.
var AISDKIdentifiers AISDKIdentifierMap

// ResolveSDKID finds the SDK ID for an AI SDK from the aiSdkName and aiSdkLanguage it
// reports in the $ld:ai:sdk:info event. Returns the SDK ID and true if found, empty
// string and false if not found.
//
// Both arguments are required. The name alone is not unique: the Python and Ruby AI
// SDKs both report "launchdarkly-server-sdk-ai".
//
// The returned ID is the key for Names, Types, Repos, and Releases.
func (m AISDKIdentifierMap) ResolveSDKID(name string, language string) (string, bool) {
	// Get sorted SDK IDs to ensure consistent ordering
	var sdkIDs []string
	for sdkID := range m {
		sdkIDs = append(sdkIDs, sdkID)
	}
	sort.Strings(sdkIDs)

	for _, sdkID := range sdkIDs {
		for _, identifier := range m[sdkID] {
			if identifier.Name == name && identifier.Language == language {
				return sdkID, true
			}
		}
	}

	return "", false
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
	panicOnError(json.Unmarshal(popularityJSON, &Popularity))
	panicOnError(json.Unmarshal(userAgentsJSON, &UserAgents))
	panicOnError(json.Unmarshal(aiSdkIdentifiersJSON, &AISDKIdentifiers))
}
