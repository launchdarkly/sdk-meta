CREATE TABLE sdk_names (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE sdk_repos (
    id TEXT PRIMARY KEY,
    github TEXT NOT NULL
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
    FOREIGN KEY (feature) REFERENCES sdk_feature_info(feature)
);

CREATE TABLE sdk_feature_info (
    feature TEXT PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE sdk_releases (
    id TEXT NOT NULL,
    major TEXT NOT NULL,
    minor TEXT NOT NULL,
    date TEXT NOT NULL,
    eol TEXT,
    PRIMARY KEY(id, major, minor)
);


INSERT INTO sdk_type_info (type, description) VALUES
                                              ('client-side', 'Primarily used for user-facing application.'),
                                              ('server-side', 'Primarily used for server-side applications.'),
                                              ('edge', 'Primarily used to delivery flag payloads to edge services.');

INSERT INTO sdk_feature_info (feature, description) VALUES
    ('u2c', 'The concept of Users is replaced with Contexts, which can be used to represent users, devices, or other entities.'),
    ('hooks', 'Hooks are collections of user-defined callbacks that are executed by the SDK at various points of interest, usually for metrics or tracing.');

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
    ('Lua'),
    ('Dart'),
    ('Objective-C');
