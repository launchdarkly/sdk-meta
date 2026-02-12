# SDK Feature Comparison Tool

## Overview

The SDK Feature Comparison Tool is a web-based interface that allows users to compare feature support across LaunchDarkly's various SDKs. The tool provides two complementary views to help teams make informed decisions about which SDK to use for their projects:

1. **By SDK View** (landing page) - Select SDKs and view their features
2. **By Feature View** - Select features and see which SDKs support them

Both views display which features are supported, deprecated, or removed across different SDK versions. The comparison tool is deployed to GitHub Pages and provides interactive table-based comparisons with dark mode support and responsive design.

## Architecture

### Data Sources

The tool consumes JSON files from the `products/` directory, which are generated from the SQLite database (`metadata.sqlite3`):

1. **features.json** - Contains feature support data for each SDK
   - Structure: `{ "sdk-id": { "feature-id": { "introduced": "version", "deprecated": "version", "removed": "version" } } }`
   - Tracks when features were introduced, deprecated, or removed

2. **feature_info.json** - Contains human-readable feature information
   - Structure: `{ "feature-id": { "name": "Display Name", "description": "Detailed description" } }`
   - Provides feature names and descriptions for the UI

3. **types.json** - Categorizes each SDK
   - Structure: `{ "sdk-id": "client-side" | "server-side" | "edge" }`
   - Determines which section (Client-side or Server-side) each SDK appears in

4. **names.json** - Maps SDK IDs to display names
   - Structure: `{ "sdk-id": "Friendly Display Name" }`
   - Used for showing user-friendly SDK names in the UI

### Components

#### 1. HTML Generator (`tool/cmd/genhtml/`)

A Go application that reads the JSON data files and generates static HTML pages. The generator supports two views controlled by the `--view` flag.

**Project Structure:**
```
tool/cmd/genhtml/
├── main.go                      - Go code (~380 lines)
└── templates/
    ├── by-sdk.html             - By SDK view template
    └── by-feature.html         - By Feature view template
```

**Key Constants:**

- `MaxSDKSelection = 3` - Maximum SDKs selectable in by-sdk view
- `MaxFeatureSelection = 5` - Maximum features selectable in by-feature view

**Key Functions:**

- **Data Reading**:
  - `readFeaturesJSON()` - Loads feature support data
  - `readTypesJSON()` - Loads SDK type categories
  - `readNamesJSON()` - Loads SDK display names
  - `readFeatureInfoJSON()` - Loads feature descriptions

- **Data Processing**:
  - `generateBySDKView()` - Builds SDK-first data index for by-sdk view
  - `generateByFeatureView()` - Builds feature-first data index for by-feature view
  - `buildFeatureIndex()` - Creates inverted index (feature → SDKs)

- **Template Execution**:
  - Uses Go's `embed` directive to embed HTML templates at compile time
  - Templates accessed via `htmlTemplateBySDK` and `htmlTemplateByFeature` variables
  - Produces self-contained HTML files with no external dependencies

**Command-Line Flags:**

- `--view`: View type (`by-sdk` or `by-feature`) - required
- `--output`: Output HTML file path
- `--data`: Path to products directory containing JSON files (default: `products`)

**Data Structures:**

```go
// Common across both views
type FeatureSupport struct {
    Introduced *string  // Version when feature was introduced
    Deprecated *string  // Version when feature was deprecated (optional)
    Removed    *string  // Version when feature was removed (optional)
}

type FeatureInfo struct {
    Name        string  // Display name of the feature
    Description string  // Detailed description
}

// For by-sdk view
type SDK struct {
    ID       string                        // SDK identifier (e.g., "go-server-sdk")
    Name     string                        // Display name (e.g., "Go Server SDK")
    Type     string                        // "client-side" or "server-side"
    Features map[string]FeatureSupport     // Feature support mapping
}

type TemplateData struct {
    ClientSDKs    []SDK
    ServerSDKs    []SDK
    AllFeatures   []string
    FeatureInfo   map[string]FeatureInfo
    FeaturesJSON  template.JS
    SDKDataJSON   template.JS
    MaxSelection  int
}

// For by-feature view
type FeatureData struct {
    ID          string
    Name        string
    Description string
    ClientSDKs  map[string]FeatureSupport
    ServerSDKs  map[string]FeatureSupport
}

type FeatureTemplateData struct {
    Features         []FeatureData
    SDKNamesJSON     template.JS
    SDKTypesJSON     template.JS
    FeaturesDataJSON template.JS
    MaxSelection     int
}
```

#### 2. HTML Templates

The tool generates two separate HTML pages, each optimized for its comparison paradigm. Both are self-contained single-page applications with no external dependencies.

**Shared Features (both views):**

- **Navigation Tabs**:
  - Tab-style navigation to switch between views
  - Active tab highlighted with purple bottom border
  - Links: "By SDK" (index.html) and "By Feature" (by-feature.html)

- **Dark Mode Support**:
  - Automatic detection of system preference
  - Manual toggle button
  - Persists user preference in localStorage
  - CSS custom properties for theme switching

- **Responsive Design**:
  - Mobile-friendly layout
  - Scrollable tables when needed
  - Sticky table headers
  - Collapsible elements on small screens

**By SDK View** (`by-sdk.html`):

- **SDK Selection**:
  - Two sections: Client-side SDKs and Server-side SDKs
  - Maximum of 3 SDKs can be selected per section
  - Checkbox-based selection with visual feedback
  - Disables unselected SDKs when limit is reached

- **Comparison Table**:
  - **Rows**: All features (alphabetically sorted by feature name)
  - **Columns**: Selected SDKs (up to 3)
  - Feature name and description shown in first column
  - Cell content: version info with status badges

**By Feature View** (`by-feature.html`):

- **Feature Selection**:
  - Two sections: Client-side SDKs and Server-side SDKs
  - Maximum of 5 features can be selected per section
  - Checkbox-based selection with visual feedback
  - Hover tooltips show feature descriptions (smart positioning)
  - Features sorted alphabetically by name
  - Disables unselected features when limit is reached

- **Comparison Table**:
  - **Rows**: All SDKs in category (alphabetically sorted by SDK name)
  - **Columns**: Selected features (up to 5)
  - SDK name shown in first column
  - Cell content: version info with status badges

**Common Table Features (both views):**

- Each cell displays:
  - Version when introduced (green text)
  - Badge for "deprecated" status (yellow)
  - Badge for "removed" status (red)
  - "—" for unsupported features (gray)
- Color-coded status indicators:
  - Green: Supported
  - Yellow: Deprecated
  - Red: Removed
  - Gray: Not supported

**JavaScript Functionality:**

- Selection management with configurable limits
- Dynamic table generation
- Status formatting with badges
- Theme toggle functionality
- Smart tooltip positioning (by-feature view only)

### Build Process

#### Local Build

The `Makefile` provides build commands:

```bash
make html
```

This command:
1. Creates the `_site/` directory
2. Generates `_site/feature-comparison-by-sdk.html` (by-sdk view)
3. Generates `_site/sdk-comparison-by-feature.html` (by-feature view)
4. Copies by-sdk to `_site/index.html` (landing page)
5. Copies by-feature to `_site/by-feature.html` (shorter URL)

**File Outputs:**
- `_site/feature-comparison-by-sdk.html` - Descriptive filename for by-sdk view
- `_site/sdk-comparison-by-feature.html` - Descriptive filename for by-feature view
- `_site/index.html` - Copy of by-sdk (canonical landing page URL)
- `_site/by-feature.html` - Copy of by-feature (canonical navigation URL)

#### Full Pipeline

```bash
make all
```

This runs the complete pipeline:
1. `crawl` - Updates the SQLite database with latest repo metadata (requires `GITHUB_TOKEN`)
2. `products` - Generates JSON products from the database
3. `html` - Generates both comparison HTML pages to `_site/`

### Deployment

The tool is deployed to GitHub Pages via the workflow defined in `.github/workflows/publish-pages.yml`.

**Trigger Conditions:**
- Push to `main` branch
- Push to `aaronz/feature-comparison-tool` branch (development branch)
- Manual workflow dispatch

**Deployment Steps:**

1. **Checkout**: Clone the repository
2. **Setup Go**: Install Go 1.22
3. **Generate HTML Pages**: Run `make html` to generate both views to `_site/`
4. **Create gh-pages Branch**: Ensure the `gh-pages` branch exists (one-time setup)
5. **Publish to GitHub Pages**: Deploy `_site/` directory to GitHub Pages

**Deployment URLs:**
- `https://<org>.github.io/<repo>/` - By SDK view (landing page)
- `https://<org>.github.io/<repo>/by-feature.html` - By Feature view
- `https://<org>.github.io/<repo>/feature-comparison-by-sdk.html` - By SDK view (descriptive URL)
- `https://<org>.github.io/<repo>/sdk-comparison-by-feature.html` - By Feature view (descriptive URL)

## Feature Support Status Types

The tool tracks three lifecycle stages for each feature:

1. **Introduced** (required): The SDK version when the feature was first added
2. **Deprecated** (optional): The SDK version when the feature was marked as deprecated
3. **Removed** (optional): The SDK version when the feature was completely removed

## User Workflows

### By SDK View Workflow

1. Visit the GitHub Pages URL (landing page)
2. Select up to 3 SDKs from the "Client-side SDKs" section
3. View the comparison table showing all features for selected SDKs
4. Select up to 3 SDKs from the "Server-side SDKs" section
5. View the comparison table showing all features for selected SDKs
6. Toggle between light and dark modes using the theme button in the header
7. Navigate to "By Feature" view using the tab navigation

### By Feature View Workflow

1. Navigate to the "By Feature" tab from the landing page
2. Hover over feature names to see descriptions
3. Select up to 5 features from the "Client-side SDKs" section
4. View the comparison table showing all client-side SDKs and their support for selected features
5. Select up to 5 features from the "Server-side SDKs" section
6. View the comparison table showing all server-side SDKs and their support for selected features
7. Scroll vertically to view all SDKs (tables show all SDKs, not just those with selected features)
8. Navigate back to "By SDK" view using the tab navigation

## Current Limitations

### By SDK View
1. **Maximum Selection**: Users can only compare 3 SDKs at a time per section
2. **Static Data**: The page is regenerated on each deployment; there's no real-time data fetching
3. **Separate Sections**: Client-side and server-side SDKs cannot be compared directly
4. **No Deep Linking**: Selected SDKs are not preserved in the URL, so comparisons can't be bookmarked or shared
5. **No Export**: Users cannot export comparison data (e.g., to CSV or PDF)

### By Feature View
1. **Maximum Selection**: Users can only compare 5 features at a time per section
2. **Static Data**: The page is regenerated on each deployment; there's no real-time data fetching
3. **Separate Sections**: Client-side and server-side features must be compared separately
4. **No Search/Filter**: All features must be browsed manually (though tooltips help with discovery)
5. **No Deep Linking**: Selected features are not preserved in the URL
6. **No Export**: Users cannot export comparison data

### General Limitations
1. **No Cross-Category Comparison**: Cannot compare client-side and server-side SDKs together
2. **No Sorting**: Table rows/columns cannot be reordered by user preference
3. **No Filtering**: Cannot hide SDKs or features based on criteria

## File Locations

- **Generator Source**: `tool/cmd/genhtml/main.go`
- **Templates**: `tool/cmd/genhtml/templates/`
  - `by-sdk.html` - By SDK view template
  - `by-feature.html` - By Feature view template
- **Generated Output**: `_site/` directory (gitignored)
  - `feature-comparison-by-sdk.html` - By SDK view (descriptive name)
  - `sdk-comparison-by-feature.html` - By Feature view (descriptive name)
  - `index.html` - By SDK view (landing page)
  - `by-feature.html` - By Feature view (navigation URL)
- **Build Command**: `Makefile` (target: `html`)
- **Deployment Workflow**: `.github/workflows/publish-pages.yml`
- **Data Sources**: `products/*.json`
- **Database**: `metadata.sqlite3`

## Technologies Used

- **Backend**: Go 1.22
- **Frontend**: Vanilla HTML/CSS/JavaScript (no frameworks)
- **Build Tool**: Make
- **CI/CD**: GitHub Actions
- **Hosting**: GitHub Pages
- **Template Engine**: Go's `html/template` package with `go:embed` directive
- **Architecture**: Single-file deployment with embedded templates

## Maintenance

To update the comparison tool:

1. **Update Data**: Run `make crawl` to fetch latest SDK metadata (requires `GITHUB_TOKEN`)
2. **Generate Products**: Run `make products` to regenerate JSON files
3. **Generate HTML**: Run `make html` to regenerate both comparison pages
4. **Test Locally**: Open `_site/index.html` in your browser to verify
5. **Deploy**: Push changes to `main` branch to trigger automatic deployment

Or simply run `make all` to execute all steps in sequence.

### Adjusting Selection Limits

To change the maximum number of SDKs or features that can be selected, edit the constants in `tool/cmd/genhtml/main.go`:

```go
const (
    MaxSDKSelection     = 3  // Max SDKs in by-sdk view
    MaxFeatureSelection = 5  // Max features in by-feature view
)
```

After changing, run `make html` to regenerate the pages.

## Related Documentation

- **Design Specification**: `.cursor/sdk-comparison-by-feature-spec.md` - Detailed design decisions and implementation spec for the by-feature view
- **Main README**: `README.md` - General repository information and tool overview
