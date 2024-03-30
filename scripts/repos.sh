#!/bin/bash

gh api --paginate graphql -f query='{
  search(
    type: REPOSITORY
    query: """
    topic:launchdarkly-sdk
    AND -topic:launchdarkly-sdk-component
    AND -topic:sdk-examples
    AND org:launchdarkly
    """
    first: 100
  ) {
    repositoryCount
    nodes {
      ... on Repository {
        nameWithOwner
        isPrivate
        repositoryTopics(first: 100) {
          nodes {
            topic {
              name
            }
          }
        }
      }
    }
  }
}' --jq '.data.search.nodes[] | select(.isPrivate | not) | .nameWithOwner'
