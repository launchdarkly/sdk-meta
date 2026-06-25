/*
 * Stub of <launchdarkly/integrations/test_data.h> from the v2.x C server
 * SDK (launchdarkly/c-server-sdk). Declares the test-data-source surface
 * that doc fragments reference. Unlike api.h, this header carries plain
 * prototypes instead of inline no-op definitions: the harness compiles
 * with `gcc -c` (no link step), and doc fragments `#include` this header
 * from inside a function body, where block-scope declarations are legal
 * but nested inline function definitions are not. (The scaffold also
 * includes it at file scope, so the include guard makes any in-body
 * re-include a no-op.)
 */
#ifndef LAUNCHDARKLY_INTEGRATIONS_TEST_DATA_H
#define LAUNCHDARKLY_INTEGRATIONS_TEST_DATA_H

#include <launchdarkly/api.h>

#ifdef __cplusplus
extern "C" {
#endif

struct LDTestData;
struct LDFlagBuilder;
struct LDFlagRuleBuilder;

struct LDTestData *LDTestDataInit(void);
void LDTestDataFree(struct LDTestData *testData);
struct LDDataSource *LDTestDataCreateDataSource(struct LDTestData *testData);
struct LDFlagBuilder *LDTestDataFlag(struct LDTestData *testData, const char *key);
LDBoolean LDTestDataUpdate(struct LDTestData *testData, struct LDFlagBuilder *flagBuilder);

LDBoolean LDFlagBuilderBooleanFlag(struct LDFlagBuilder *flagBuilder);
void LDFlagBuilderOn(struct LDFlagBuilder *flagBuilder, LDBoolean on);
LDBoolean LDFlagBuilderVariations(struct LDFlagBuilder *flagBuilder, struct LDJSON *variations);
LDBoolean LDFlagBuilderVariationForAllUsersBoolean(struct LDFlagBuilder *flagBuilder, LDBoolean value);
LDBoolean LDFlagBuilderVariationForUserBoolean(struct LDFlagBuilder *flagBuilder, const char *userKey, LDBoolean value);
LDBoolean LDFlagBuilderFallthroughVariationBoolean(struct LDFlagBuilder *flagBuilder, LDBoolean value);
void LDFlagBuilderFallthroughVariation(struct LDFlagBuilder *flagBuilder, int variationIndex);
LDBoolean LDFlagBuilderOffVariation(struct LDFlagBuilder *flagBuilder, int variationIndex);
struct LDFlagRuleBuilder *LDFlagBuilderIfMatch(struct LDFlagBuilder *flagBuilder, const char *attribute, struct LDJSON *value);
struct LDFlagRuleBuilder *LDFlagBuilderIfNotMatch(struct LDFlagBuilder *flagBuilder, const char *attribute, struct LDJSON *value);
void LDFlagRuleBuilderThenReturn(struct LDFlagRuleBuilder *ruleBuilder, int variationIndex);

#ifdef __cplusplus
}
#endif

#endif /* LAUNCHDARKLY_INTEGRATIONS_TEST_DATA_H */
