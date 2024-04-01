#!/bin/bash

gh api --paginate graphql -f query='{
  search(
    type: REPOSITORY
    query: """topic:launchdarkly-sdk -topic:launchdarkly-sdk-component -topic:sdk-examples org:launchdarkly is:public"""
    first: 100
  ) {
    repositoryCount
    nodes {
      ... on Repository {
        nameWithOwner
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
}' --jq '.data.search.nodes[] | .nameWithOwner'
