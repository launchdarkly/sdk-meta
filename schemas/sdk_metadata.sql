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


INSERT INTO sdk_type_info (type, description) VALUES
                                              ('client-side', 'Primarily used for user-facing application.'),
                                              ('server-side', 'Primarily used for server-side applications.'),
                                              ('edge', 'Primarily used to delivery flag payloads to edge services.');

INSERT INTO sdk_feature_info (id, name, description) VALUES
    ('appConfig', 'Application metadata', 'Specify application and application version information.'),
    ('autoEnvAttrs', 'Automatic environment attributes', 'Automatically include device and application data in each evaluated context.'),
    ('bigSegments', 'Big segments', 'Configure a persistent store to hold segments that are either synced from external tools, or that contain an arbitrarily large number of contexts of any one context kind.'),
    ('bootstrapping', 'Bootstrapping', 'Provide an initial set of flag values that are immediately available during client initialization.'),
    ('contexts', 'Contexts and context kinds', 'Evaluate flags based on contexts. A context is a generalized way of referring to the people, services, machines, or other resources that encounter feature flags. SDKs without this feature only support flag evaluation for users.'),
    ('experimentation', 'Experimentation', 'Connect a flag with one or more metrics to measure end-user behavior for different variations of a flag. Requires minimum SDK versions, but no SDK configuration.'),
    ('variationDetail', 'Flag evaluation reasons', 'Receive information about how a flag variation was calculated, for example, because it matched a specific targeting rule.'),
    ('allFlags', 'Getting all flags', 'Return the flag variations for all feature flags for a given context.'),
    ('hooks', 'Hooks', 'Define callbacks that are executed by the SDK at various points of interest, usually for metrics or tracing.'),
    ('migrations', 'Migration flags', 'Configure how to read and write data for an old and new system, determine which stage of a migration the application is in, execute the correct read and write calls for each stage.'),
    ('multiEnv', 'Multiple environments', 'Evaluate flags from multiple environments using a single client instance'),
    ('offlineMode', 'Offline mode', 'Close the SDK''s connection to LaunchDarkly. Use cached or fallback values for each flag evaluation.'),
    ('otel', 'OpenTelemetry', 'Add flag evaluation information to OpenTelemetry spans.'),
    ('privateAttrs', 'Private attributes', 'Use context attribute values for targeting, but do not send them to LaunchDarkly.'),
    ('fileDataSource', 'Reading flags from a file', 'Use flag values, specified in JSON or YAML files, for all flag evaluations. Useful for testing or prototyping; do not use in production.'),
    ('relayProxyDaemon', 'Relay Proxy in daemon mode', 'Configure the SDK to connect to a persistent store, such as the Relay Proxy''s data store.'),
    ('relayProxyProxy', 'Relay Proxy in proxy mode', 'Configure the SDK to connect to alternate service endpoints, such as the Relay Proxy.'),
    ('secureMode', 'Secure mode', 'For clent-side SDKs, require a hash, signed with the SDK key for the LaunchDarkly environment, to evaluate flag variations. For server-side or edge SDKs, generate a secure mode hash.'),
    ('track', 'Sending custom events', 'Record actions taken in your application as events. You can connect to these events to metrics for use in experiments.'),
    ('storingData', 'Storing data', 'Configure an external database as a feature store. Persist flag data across application restarts.'),
    ('flagChanges', 'Subscribing to flag changes', 'Use a listener pattern to subscribe to flag change notifications.'),
    ('testDataSource', 'Test data sources', 'Mock data of an SDK. Useful for unit tests; cannot be used in production.'),
    ('webProxy', 'Web proxy configuration', 'Configure the SDK to connect to LaunchDarkly through a web proxy.');

INSERT INTO sdk_language_info (language) VALUES
    ('Apex'),
    ('BrightScript'),
    ('JavaScript'),
    ('TypeScript'),
    ('Python'),
    ('Ruby'),
    ('C++'),
    ('C'),
    ('C#'),
    ('Java'),
    ('Kotlin'),
    ('Go'),
    ('Swift'),
    ('Rust'),
    ('PHP'),
    ('Haskell'),
    ('Erlang'),
    ('Elixir'),
    ('Lua'),
    ('Dart'),
    ('Objective-C');
