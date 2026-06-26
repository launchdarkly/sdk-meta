/*
 * Stub of <launchdarkly/integrations/file_data.h> from the v2.x C server
 * SDK (launchdarkly/c-server-sdk). Declares the file data source
 * constructor that flags-from-files doc fragments reference. The inline
 * definition is a no-op — this header is for parse + type-check
 * validation only, never linked into a running binary.
 */
#ifndef LAUNCHDARKLY_INTEGRATIONS_FILE_DATA_H
#define LAUNCHDARKLY_INTEGRATIONS_FILE_DATA_H

#include <launchdarkly/api.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Mirrors the real v2 header: returns a data source that reads flag
 * data from the given files, for use with LDConfigSetDataSource. */
static inline struct LDDataSource *LDFileDataInit(int fileCount, const char **filenames) {
    (void)fileCount;
    (void)filenames;
    return (struct LDDataSource *)0;
}

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_INTEGRATIONS_FILE_DATA_H */
