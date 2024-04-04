#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <sqlite3 database file>"
  exit 1
fi

sqlite3 -json "$1" "
SELECT
  r.id,
  r.major,
  r.minor,
  r.date,
  (SELECT DATE(MIN(r2.date), '+1 year')
    FROM sdk_releases r2
      WHERE r2.id = r.id
        AND (r2.major > r.major OR (r2.major = r.major AND r2.minor > r.minor))
) AS eol
  FROM sdk_releases r
  WHERE NOT EXISTS (
    SELECT 1
    FROM sdk_releases r_inner
    WHERE r_inner.id = r.id
      AND r_inner.major = r.major
      AND r_inner.minor = r.minor
      AND r_inner.patch < r.patch
) ORDER BY r.id ASC, r.major DESC, r.minor DESC;"
