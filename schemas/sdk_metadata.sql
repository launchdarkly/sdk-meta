CREATE TABLE sdk_names (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE sdk_repos (
    id TEXT,
    github TEXT NOT NULL,
    PRIMARY KEY (id, github)
);

CREATE TABLE sdk_languages (
    id TEXT,
    language TEXT NOT NULL,
    PRIMARY KEY (id, language),
    FOREIGN KEY (language) REFERENCES sdk_language_info(language)
);

CREATE TABLE sdk_language_info (
    language TEXT PRIMARY KEY
);

CREATE TABLE sdk_types (
    id TEXT PRIMARY KEY,
    type TEXT NOT NULL,
    FOREIGN KEY (type) REFERENCES sdk_type_info(type)
);


CREATE TABLE sdk_type_info (
    type TEXT PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE sdk_features (
    id TEXT NOT NULL,
    feature TEXT NOT NULL,
    introduced TEXT NOT NULL,
    deprecated TEXT,
    removed TEXT,
    PRIMARY KEY (id, feature),
    FOREIGN KEY (feature) REFERENCES sdk_feature_info(id)
);

CREATE TABLE sdk_feature_info (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);


CREATE TABLE sdk_releases (
    id TEXT NOT NULL,
    major INTEGER NOT NULL,
    minor INTEGER NOT NULL,
    patch INTEGER NOT NULL,
    date TEXT NOT NULL,
    PRIMARY KEY(id, major, minor, patch)
);


CREATE TABLE sdk_popularity (
    id TEXT NOT NULL,
    popularity INTEGER NOT NULL,
    PRIMARY KEY(id, popularity)
);

CREATE TABLE sdk_user_agents (
    id TEXT NOT NULL,
    userAgent TEXT
);

CREATE TABLE sdk_wrappers (
    id TEXT NOT NULL,
    wrapper TEXT
);

INSERT INTO sdk_popularity (id, popularity) VALUES
    ('react-client-sdk', 1),
    ('node-server', 2),
    ('python-server-sdk', 3),
    ('java-server-sdk', 4),
    ('dotnet-server-sdk', 5),
    ('js-client-sdk', 6),
    ('vue', 7),
    ('swift-client-sdk', 8),
    ('go-server-sdk', 9),
    ('android', 10),
    ('react-native', 11),
    ('ruby-server-sdk', 12),
    ('flutter-client-sdk', 13),
    ('dotnet-client-sdk', 14),
    ('erlang-server-sdk', 15),
    ('rust-server-sdk', 16),
    ('cpp-client-sdk', 17),
    ('roku', 18),
    ('node-client-sdk', 19),
    ('cpp-server-sdk', 20),
    ('lua-server-sdk', 21),
    ('haskell-server-sdk', 22),
    ('php-server-sdk', 23);


INSERT INTO sdk_type_info (type, description) VALUES
                                              ('client-side', 'Primarily used for user-facing application.'),
                                              ('edge', 'Primarily used to delivery flag payloads to edge services.'),
                                              ('ai', 'Used for AI/ML integrations.'),
                                              ('open-feature-provider', 'Used for OpenFeature providers.'),
                                              ('relay', 'Special case for Relay Proxy.'),
                                              ('server-side', 'Primarily used for server-side applications.');

INSERT INTO sdk_feature_info (id, name, description) VALUES
    ('aiAgentConfig', 'AI Agent Configuration', 'Use an autonomous agent to perform multi-step tasks. Agents use a combination of available tools and strategies to complete goals like analyzing data, generating reports, or triggering actions.'),
    ('aiCompletionConfig', 'AI Completion Configuration', 'Send prompts to an LLM to generate a single response.'),
    ('aiJudgeConfig', 'AI Judge Configuration', 'AI Configs automatically score responses using LLM-as-a-Judge, providing actionable data points for quality protection.'),
    ('aiTrackMetrics', 'Tracking AI metrics', 'Collect and report metrics specific to AI/ML workloads and flag evaluations.'),
    ('allFlags', 'Generate bootstrap details', 'Return the flag variations for all feature flags for a given context.'),
    ('appMetadata', 'Application metadata', 'Specify application and application version information.'),
    ('autoEnvAttrs', 'Automatic environment attributes', 'Automatically include device and application data in each evaluated context.'),
    ('bigSegments', 'Big segments', 'Configure a persistent store to hold segments that are either synced from external tools, or that contain an arbitrarily large number of contexts of any one context kind.'),
    ('bootstrapping', 'Bootstrapping', 'Provide an initial set of flag values that are immediately available during client initialization.'),
    ('contexts', 'Contexts and context kinds', 'Evaluate flags based on contexts. A context is a generalized way of referring to the people, services, machines, or other resources that encounter feature flags. SDKs without this feature only support flag evaluation for users.'),
    ('eventCompression', 'Event Compression', 'Optionally support compressing event payloads sent to LaunchDarkly.'),
    ('experimentation', 'Experimentation', 'Connect a flag with one or more metrics to measure end-user behavior for different variations of a flag. Requires minimum SDK versions, but no SDK configuration.'),
    ('fdv2', 'Data Saving Mode', 'Our second generation flag delivery protocol.'),
    ('fileDataSource', 'Reading flags from a file', 'Use flag values, specified in JSON or YAML files, for all flag evaluations. Useful for testing or prototyping; do not use in production.'),
    ('flagChanges', 'Subscribing to flag changes', 'Use a listener pattern to subscribe to flag change notifications.'),
    ('hooks', 'Hooks', 'Define callbacks that are executed by the SDK at various points of interest, usually for metrics or tracing.'),
    ('inlineContextCustomEvents', 'Inline Context In Custom Events', 'The SDK will inline the context inside custom events.'),
    ('inlineContextEvaluationEvents', 'Inline Context In Evaluation Events', 'The SDK will inline the context inside evaluation events.'),
    ('migrations', 'Migration flags', 'Configure how to read and write data for an old and new system, determine which stage of a migration the application is in, execute the correct read and write calls for each stage.'),
    ('multiEnv', 'Multiple environments', 'Evaluate flags from multiple environments using a single client instance'),
    ('offlineMode', 'Offline mode', 'Close the SDK''s connection to LaunchDarkly. Use cached or fallback values for each flag evaluation.'),
    ('omitAnonContexts', 'Omit Anonymous Contexts from Events', 'Optionally omit anonymous contexts from events.'),
    ('otel', 'OpenTelemetry', 'Add flag evaluation information to OpenTelemetry spans.'),
    ('perContextSummaryEvents', 'Per-Context Summary Events', 'The SDK supports emitting a summary event per-context.'),
    ('pluginSupport', 'Plugin Support', 'Provides SDK extensibility through the use of plugins.'),
    ('pollingResponseCompression', 'Polling Response Compression', 'Supports receiving compressed polling responses.'),
    ('privateAttrs', 'Private attributes', 'Use context attribute values for targeting, but do not send them to LaunchDarkly.'),
    ('relayProxyDaemon', 'Relay Proxy in daemon mode', 'Configure the SDK to connect to a persistent store, such as the Relay Proxy''s data store.'),
    ('relayProxyProxy', 'Relay Proxy in proxy mode', 'Configure the SDK to connect to alternate service endpoints, such as the Relay Proxy.'),
    ('secureMode', 'Secure mode', 'For clent-side SDKs, require a hash, signed with the SDK key for the LaunchDarkly environment, to evaluate flag variations. For server-side or edge SDKs, generate a secure mode hash.'),
    ('storingData', 'Storing data', 'Configure an external database as a feature store. Persist flag data across application restarts.'),
    ('storingDataConsul', 'Storing data (Consul)', 'Configure Consul as a feature store. Persist flag data across application restarts.'),
    ('storingDataDynamodb', 'Storing data (DynamoDB)', 'Configure DynamoDB as a feature store. Persist flag data across application restarts.'),
    ('storingDataRedis', 'Storing data (Redis)', 'Configure Redis as a feature store. Persist flag data across application restarts.'),
    ('testDataSource', 'Test data sources', 'Mock data of an SDK. Useful for unit tests; cannot be used in production.'),
    ('track', 'Sending custom events', 'Record actions taken in your application as events. You can connect to these events to metrics for use in experiments.'),
    ('variationDetail', 'Flag evaluation reasons', 'Receive information about how a flag variation was calculated, for example, because it matched a specific targeting rule.'),
    ('webProxy', 'Web proxy configuration', 'Configure the SDK to connect to LaunchDarkly through a web proxy.');

INSERT INTO sdk_language_info (language) VALUES
    ('Apex'),
    ('BrightScript'),
    ('C'),
    ('C#'),
    ('C++'),
    ('Dart'),
    ('Elixir'),
    ('Erlang'),
    ('Go'),
    ('Haskell'),
    ('Java'),
    ('JavaScript'),
    ('Kotlin'),
    ('Lua'),
    ('Objective-C'),
    ('PHP'),
    ('Python'),
    ('Ruby'),
    ('Rust'),
    ('Swift'),
    ('TypeScript');
