import sdkLanguages from './data/languages.json'
import sdkRepos from './data/repos.json'
import sdkNames from './data/names.json'
import sdkTypes from './data/types.json'
import sdkPopularity from './data/popularity.json'
import sdkReleases from './data/releases.json'
import sdkUserAgents from './data/user_agents.json'
import sdkAiSdkIdentifiers from './data/ai_sdk_identifiers.json'

export enum Type {
    // ClientSide is an SDK that runs in a client scenario.
    ClientSide = "client-side",
    // ServerSide is an SDK that runs in a server scenario.
    ServerSide = "server-side",
    // Edge is an SDK that runs in an edge deployment scenario.
    Edge = "edge",
    // AI is an SDK that is primarily focused on AI/ML use cases.
    AI = "ai",
    // OpenFeatureProvider is an OpenFeature provider.
    OpenFeatureProvider = "open-feature-provider",
    // Relay is Relay Proxy.
    Relay = "relay",
    // Unknown if the SDK's type is not recognized.
    Unknown = "unknown"
}

function isType(value: string): value is Type {
    return Object.values(Type).includes(value as Type);
}

export type Repo = {
    github: string;
}

export const Languages: Record<string, string[]> = sdkLanguages;
export const Names: Record<string, string> = sdkNames;
export const Repos: Record<string, Repo> = sdkRepos;
export const Popularity: Record<string, number> = sdkPopularity;
export const Releases: ReleaseList = Object.fromEntries(
    Object.entries(sdkReleases).map(([key, value]) => [
      key,
      value.map((release: any) => ({
          Major: release["major"],
          Minor: release["minor"],
          Date: new Date(release["date"]),
          EOL: release["eol"] ? new Date(release["eol"]) : null
      }))
    ]));

export const Types: Record<string, Type> = Object.fromEntries(
    Object.entries(sdkTypes).map(([key, value]) => [
      key,
      isType(value) ? value : Type.Unknown
    ]));


export interface Release {
    Major: number;
    Minor: number;
    Date: Date;
    EOL: Date | null;
}

export interface ReleaseList {
    [key: string]: Release[];
}

export namespace ReleaseHelpers {
    export const IsLatest = (release: Release) => release.EOL === null;
    export const IsEOL = (release: Release, now: Date) => !IsLatest(release) && now > release.EOL!;
    export const IsApproachingEOL = (release: Release, now: Date, thresholdPrior: number) =>
        !IsLatest(release) && now.getTime() + thresholdPrior > release.EOL!.getTime();

    export const Earliest = (releases: Release[]) => releases[releases.length - 1];
    export const Latest = (releases: Release[]) => releases[0];
}

export interface UserAgent {
    userAgents?: string[];
    wrapperNames?: string[];
}

export const UserAgents: Record<string, UserAgent> = sdkUserAgents;

export namespace UserAgentHelpers {
    /**
     * Attempts to find an SDK name by checking wrapper names and user agents.
     * First checks wrapper names, then user agents, in alphabetical order by SDK ID.
     * 
     * @param identifier - The wrapper name or user agent string to search for
     * @returns The SDK name if found, undefined if not found
     */
    export const getSDKNameByWrapperOrUserAgent = (identifier: string): string | undefined => {
        const sdkId = getSDKIDByWrapperOrUserAgent(identifier);
        return sdkId === undefined ? undefined : Names[sdkId];
    }

    /**
     * Attempts to find an SDK ID by checking wrapper names and user agents.
     * First checks wrapper names, then user agents, in alphabetical order by SDK ID.
     *
     * Prefer this over getSDKNameByWrapperOrUserAgent when you need more than a display
     * name. The ID is the key for Names, Types, Repos, and Releases.
     *
     * Some identifiers are reported by more than one SDK. For example both akamai-base
     * and akamai-edgekv report "AkamaiEdgeSDK". In that case the first SDK ID in
     * alphabetical order is returned.
     *
     * @param identifier - The wrapper name or user agent string to search for
     * @returns The SDK ID if found, undefined if not found
     */
    export const getSDKIDByWrapperOrUserAgent = (identifier: string): string | undefined => {
        // Sort the entries by SDK ID to ensure consistent ordering.
        const sortedEntries = Object.entries(UserAgents).sort(([a], [b]) => a.localeCompare(b));

        // First check wrapper names
        for (const [sdkId, info] of sortedEntries) {
            if (info.wrapperNames?.includes(identifier)) {
                return sdkId;
            }
        }

        // Then check user agents
        for (const [sdkId, info] of sortedEntries) {
            if (info.userAgents?.includes(identifier)) {
                return sdkId;
            }
        }

        return undefined;
    }
}

/**
 * One (name, language) pair that an AI SDK reports about itself in the
 * $ld:ai:sdk:info custom event.
 */
export interface AISDKIdentifier {
    /** The value reported as aiSdkName. This is the package name. */
    name: string;
    /** The value reported as aiSdkLanguage, for example "python". */
    language: string;
}

/**
 * A map of SDK IDs to the identifiers those AI SDKs report.
 *
 * An AI SDK wraps a client that the caller supplies, so it sends no user agent of its
 * own. It reports itself in a custom event instead. Historical names are kept so that
 * older deployments still resolve.
 */
export const AISDKIdentifiers: Record<string, AISDKIdentifier[]> = sdkAiSdkIdentifiers;

export namespace AISDKHelpers {
    /**
     * Finds the SDK ID for an AI SDK from the aiSdkName and aiSdkLanguage it reports in
     * the $ld:ai:sdk:info event.
     *
     * Both arguments are required. The name alone is not unique: the Python and Ruby AI
     * SDKs both report "launchdarkly-server-sdk-ai".
     *
     * The returned ID is the key for Names, Types, Repos, and Releases.
     *
     * @param name - The value reported as aiSdkName
     * @param language - The value reported as aiSdkLanguage
     * @returns The SDK ID if found, undefined if not found
     */
    export const resolveAISDK = (name: string, language: string): string | undefined => {
        // Sort the entries by SDK ID to ensure consistent ordering.
        const sortedEntries = Object.entries(AISDKIdentifiers).sort(([a], [b]) => a.localeCompare(b));

        for (const [sdkId, identifiers] of sortedEntries) {
            if (identifiers.some((i) => i.name === name && i.language === language)) {
                return sdkId;
            }
        }

        return undefined;
    }
}
