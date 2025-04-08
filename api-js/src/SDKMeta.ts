import sdkLanguages from './data/languages.json'
import sdkRepos from './data/repos.json'
import sdkNames from './data/names.json'
import sdkTypes from './data/types.json'
import sdkPopularity from './data/popularity.json'
import sdkReleases from './data/releases.json'
import sdkUserAgents from './data/user_agents.json'

export enum Type {
    // ClientSide is an SDK that runs in a client scenario.
    ClientSide = "client-side",
    // ServerSide is an SDK that runs in a server scenario.
    ServerSide = "server-side",
    // Edge is an SDK that runs in an edge deployment scenario.
    Edge = "edge",
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
    export const getSDKNameByWrapperOrUserAgent = (identifier: string): string | undefined => {
        // First check wrapper names
        for (const [sdkId, info] of Object.entries(UserAgents)) {
            if (info.wrapperNames?.includes(identifier)) {
                return Names[sdkId];
            }
        }

        // Then check user agents
        for (const [sdkId, info] of Object.entries(UserAgents)) {
            if (info.userAgents?.includes(identifier)) {
                return Names[sdkId];
            }
        }

        return undefined;
    }
}
