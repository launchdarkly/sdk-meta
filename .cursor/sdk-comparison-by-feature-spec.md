# SDK Comparison by Feature - Design Document

## Project Goal

Create a feature-driven view for SDK comparison where users select one or more features and see which SDKs support them, with their version information and status (supported/deprecated/removed).

This is the inverse of the current view, which is SDK-driven (select SDKs, see their features).

## Architecture Decision: Single Page vs. Separate Pages

### Option 1: Single Page with View Toggle

Build both the SDK-driven and feature-driven views into one HTML page with a toggle/tab interface.

**Pros:**
- **Single URL to maintain** - Only one page to document, bookmark, and share
- **Shared code and data** - No duplication of HTML generation logic, CSS, or embedded JSON data
- **Easy view switching** - Users can toggle between perspectives without page reload
- **Consistent UI/UX** - Same theme, styling, and overall layout
- **Single deployment** - No additional workflow changes needed
- **DRY principle** - All the data processing and embedding happens once

**Cons:**
- **Increased complexity** - JavaScript must manage two different view states and interactions
- **Larger initial payload** - Page includes code for both views even if user only needs one
- **More complex state management** - Need to track which view is active, selections for both views
- **Potential for confusion** - Users might not understand they're looking at two different perspectives
- **Harder to optimize** - UI must work for both paradigms, potentially compromising each
- **Testing complexity** - Need to test interactions between views, mode switching, etc.

**Technical Implementation Notes:**
- Could use tabs or radio buttons to switch between "Compare by SDK" and "Compare by Feature" modes
- Would show/hide different sections based on active mode
- Could preserve selections when switching between modes (if applicable)
- URL hash or query parameter could indicate active view for bookmarking

---

### Option 2: Separate Pages

Create a second HTML page (`features.html` and `by-feature.html` or similar naming).

**Pros:**
- **Clear separation of concerns** - Each page has a single, focused purpose
- **Optimized UX per use case** - Each page can be designed specifically for its workflow
- **Smaller initial payload** - Each page only loads what it needs
- **Simpler JavaScript** - Each page has straightforward, linear logic
- **Independent evolution** - Can update one view without affecting the other
- **Better bookmarking** - Direct URLs to specific views (e.g., `/by-sdk.html` vs `/by-feature.html`)
- **Easier to test** - Each page tested independently
- **Clearer analytics** - Can track usage of each view separately
- **Easier onboarding** - New developers can understand one page at a time

**Cons:**
- **Code duplication** - Some shared UI components (header, theme toggle, legend) duplicated
- **Maintenance burden** - Need to keep both pages in sync for common elements
- **Multiple URLs** - Documentation must cover both pages
- **Discovery challenge** - Users might not know the other view exists
- **Slight deployment complexity** - Generate two HTML files, ensure both are deployed
- **Inconsistency risk** - Pages might diverge in styling or behavior over time

**Technical Implementation Notes:**
- Generate two separate HTML files from Go: `features.html` (current) and `by-feature.html` (new)
- Add navigation link on each page pointing to the other view
- Could use shared CSS by extracting to external file (though this breaks "single file" principle)
- Makefile would need to generate both: `make html-by-sdk` and `make html-by-feature`

---

## Comparison Matrix

| Aspect | Single Page | Separate Pages |
|--------|-------------|----------------|
| **URL Management** | One URL, uses hash/query | Two distinct URLs |
| **Code Duplication** | Minimal | Moderate (shared components) |
| **JavaScript Complexity** | High (state machine) | Low (linear flow) |
| **Page Load Size** | Larger (includes both) | Smaller (focused) |
| **User Discovery** | Toggle is visible | Requires navigation link |
| **Maintainability** | Complex but centralized | Simple but duplicated |
| **Bookmarking** | Harder (needs URL params) | Easy (distinct URLs) |
| **Testing** | Complex (mode switching) | Simple (independent) |
| **Analytics** | Single page view | Separate page views |
| **Build Complexity** | Same as current | Slightly increased |

---

## Recommendation

**Recommendation: Start with Separate Pages (Option 2)**

### Reasoning:

1. **Different mental models**: The SDK-driven and feature-driven views serve fundamentally different use cases:
   - SDK-driven: "I'm deciding between these SDKs, what features do they have?"
   - Feature-driven: "I need these features, which SDKs support them?"
   
   These are distinct questions that users ask at different times in their workflow.

2. **Optimization opportunity**: The feature-driven view will likely need different UI patterns:
   - Feature list might be long (requires search/filter)
   - SDK columns will be many (horizontal scrolling challenges)
   - May want to show only SDKs that support selected features
   - Sorting and grouping needs might differ

3. **Clearer user intent**: Separate URLs make it obvious which view you're looking at and allow sharing specific perspectives.

4. **Iterative development**: Easier to build and test the new view independently without risk of breaking the existing one.

5. **Code duplication is manageable**: With Go templates, we can potentially create shared template components to minimize duplication.

6. **Future extensibility**: If we later want a third view (e.g., "by SDK category" or "by deprecation status"), separate pages scale better than adding more modes to a single page.

### Mitigation for Discovery:

To address the "users won't find the other view" concern:
- Add a prominent navigation element on both pages (e.g., tabs or button in header)
- Include brief description of each view's purpose
- Add to documentation/README
- Consider landing page that explains both views (future enhancement)

---

## Implementation Plan

### Status: Steps 1-9 Complete! ✅

1. ✅ **Define detailed requirements** for the feature-driven view
2. ✅ **Design the UI/UX** - wireframes or mockups
3. ✅ **Decide on feature limits** - how many features can be selected at once?
4. ✅ **Plan the data structure** - how to efficiently query SDKs by feature
5. ✅ **Update the Go generator** - create second template or parameterize existing one
6. ✅ **Implement the JavaScript** - feature selection and table rendering logic
7. ✅ **Update Makefile** - add build target for second page
8. ✅ **Update GitHub Actions** - ensure both pages are deployed
9. ✅ **Write documentation** - update README and add user guide

**Implementation Complete!**

Both pages are now functional:
- `feature-comparison-by-sdk.html` - Select SDKs, view their features
- `sdk-comparison-by-feature.html` - Select features, view SDK support

Next step: Update README documentation

---

## Open Questions

1. ~~Should there be a limit on how many features can be selected at once? (Similar to the 3 SDK limit)~~ **ANSWERED**
2. ~~Should the feature view show ALL SDKs or only SDKs that support at least one selected feature?~~ **ANSWERED**
3. ~~How should features be organized? (Alphabetical, by category, by popularity?)~~ **ANSWERED**
4. ~~Should there be a search/filter for the feature list?~~ **ANSWERED**
5. ~~What's the desired file naming convention? (`by-sdk.html` + `by-feature.html`, or keep current name and add `features-by-sdk.html`?)~~ **ANSWERED**
6. ~~Should we maintain the client-side/server-side separation in the feature view?~~ **ANSWERED**

---

## Step 1: Detailed Requirements ✅

### Functional Requirements

**Page: sdk-comparison-by-feature.html**

1. **Feature Selection**
   - Display all features from `feature_info.json` as selectable items
   - Features organized alphabetically by feature name
   - Show feature name in selection area
   - Show feature description on hover (tooltip)
   - User can select up to 3 features
   - Visual feedback: selected features highlighted
   - Disable/gray out unselected features when limit reached
   - Show counter: "Selected: X/3"

2. **SDK Display Sections**
   - Two sections: "Client-side SDKs" and "Server-side SDKs"
   - Each section shows all SDKs of that type (from `types.json`)
   - SDKs displayed as columns in a table

3. **Comparison Table**
   - Rows: All SDKs in the category (client or server)
   - Columns: Selected features (up to 3)
   - Cell content for each feature/SDK intersection:
     - If supported: Version number when introduced (green text)
     - If deprecated: Version + "deprecated X.X" badge (yellow)
     - If removed: Version + "removed X.X" badge (red)
     - If not supported: "-" (gray text)
   - Table only visible when at least 1 feature is selected
   - Show "Select features above to compare SDKs" message when no features selected

4. **Navigation**
   - Header with navigation to switch between views
   - **Decision**: Tab-style buttons ("By SDK" | "By Feature")
   - Clear visual distinction between active and inactive tabs
   - Located prominently in header area

5. **Shared Features** (consistent with existing page)
   - Dark/light theme toggle
   - Responsive design
   - Sticky table headers
   - Horizontal scrolling for wide tables
   - LaunchDarkly branding and colors

### Non-Functional Requirements

1. **Performance**
   - Page should load quickly (single HTML file, no external dependencies)
   - Smooth interactions (no lag when selecting features)

2. **Accessibility**
   - Keyboard navigation support
   - Proper ARIA labels for screen readers
   - Color contrast meeting WCAG standards
   - Tooltips accessible via keyboard

3. **Browser Support**
   - Modern browsers (Chrome, Firefox, Safari, Edge)
   - No IE11 support required (consistent with current page)

4. **Mobile Responsiveness**
   - Usable on tablets (768px+)
   - Horizontal scroll on smaller screens
   - Collapsible elements where appropriate

---

## Step 2: UI/UX Design ✅

### Page Layout Structure

```
┌─────────────────────────────────────────────────────────────┐
│  ┌────────┬────────────┐                                    │
│  │ By SDK │ By Feature │  LaunchDarkly SDK Feature Comp 🌙 │
│  └────────┴────────────┘                                    │
├─────────────────────────────────────────────────────────────┤
│  Select up to 3 features to compare across SDKs             │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ Client-side SDKs                                            │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│ Select Features:  [0/3 selected]                            │
│                                                              │
│ ┌──────────────┐ ┌──────────────┐ ┌──────────────┐        │
│ │☐ allFlags    │ │☐ bigSegments │ │☐ contexts    │ ...    │
│ └──────────────┘ └──────────────┘ └──────────────┘        │
│                                                              │
│ ┌─────────────────────────────────────────────────────────┐│
│ │  SDK            │ Big segments │ Contexts │ Hooks      ││
│ ├─────────────────┼──────────────┼──────────┼────────────┤│
│ │  Android SDK    │     5.5      │   7.0    │     -      ││
│ │  Flutter SDK    │     3.2      │   5.0    │     -      ││
│ │  iOS SDK        │     6.0      │   8.0    │    9.0     ││
│ │  JavaScript SDK │     2.8      │   3.5    │    4.0     ││
│ │  ...            │     ...      │   ...    │    ...     ││
│ └─────────────────┴──────────────┴──────────┴────────────┘│
│                                                              │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ Server-side SDKs                                            │
├─────────────────────────────────────────────────────────────┤
│ [Similar structure to client-side]                          │
└─────────────────────────────────────────────────────────────┘
```

**Note**: Table orientation shows:
- **Rows**: All SDKs (alphabetically sorted by name)
- **Columns**: Selected features (up to 3)

### Feature Selection UI

**Visual Design:**
- Similar to current SDK selector
- Grid layout with flex-wrap for responsiveness
- Each feature as a clickable card/button
- Hover shows tooltip with description
- Selected state: purple background (LaunchDarkly brand color)
- Disabled state: reduced opacity when limit reached

**Interaction:**
1. Click feature name → toggles selection
2. Hover feature name → shows description tooltip
3. When 3 selected → remaining features become disabled
4. Deselect any feature → re-enables all features

### Comparison Table UI

**Design Considerations:**

1. **Table Dimensions**:
   - Many rows (all SDKs - could be 20-30 SDKs per section)
   - Few columns (up to 3 selected features)
   - Vertical scrolling for long SDK lists
   - Minimal horizontal scrolling (only 3-4 columns total)

2. **SDK Name Column** (First column):
   - Left-aligned
   - SDK name in regular weight
   - Min-width to ensure readability
   - Sticky column when scrolling horizontally (if needed)

3. **Feature Columns** (2-4 columns depending on selection):
   - Header shows feature name (with description on hover)
   - Center-aligned values
   - Min-width for version numbers + badges
   - Color-coded status as per existing page

4. **Cell Content**:
   - Version numbers centered
   - Badges inline (deprecated/removed)
   - "-" for unsupported features
   - Consider: Highlight cells with deprecated/removed status

5. **Empty State**:
   - Friendly message: "Select features above to compare SDKs"
   - Optional: Show illustration or icon

**Table Benefits with This Orientation:**
- Shows all SDKs at once (complete landscape view)
- Easier to scan down SDK list
- Less horizontal scrolling (max 3 feature columns)
- Better for mobile/tablet (fewer columns)

### Navigation Design

**Tab-Style Navigation (Selected):**

```
┌──────────┬──────────────┐
│  By SDK  │  By Feature  │
└──────────┴──────────────┘
   (active)    (inactive)
```

**Visual Specifications:**
- Active tab: 
  - Solid bottom border (3px, LaunchDarkly purple #405BFF)
  - Bold text
  - Full opacity
  - Background: transparent or subtle highlight
- Inactive tab: 
  - No border
  - Regular weight text
  - Slightly reduced opacity (0.7)
  - Background: transparent
- On hover (inactive): 
  - Opacity increases to 0.85
  - Optional: slight background color
  - Cursor: pointer
- Location: 
  - Left side of header
  - Before the page title
  - Horizontal layout

**Implementation Notes:**
- Each tab is a clickable link/button
- "By SDK" tab links to `feature-comparison-by-sdk.html`
- "By Feature" tab links to `sdk-comparison-by-feature.html`
- Current page's tab shows active state
- No JavaScript needed for navigation (pure links)
- Maintains current theme (dark/light) when switching pages via localStorage

### Tooltip Design

For feature descriptions on hover:

- **Trigger**: Hover over feature name for 300ms
- **Position**: Above or below feature button (auto-adjust for viewport)
- **Style**: 
  - Dark background with white text (or inverse in light mode)
  - Subtle drop shadow
  - Max-width: 300px
  - Padding: 8px 12px
  - Border radius: 4px
- **Content**: Feature description from `feature_info.json`
- **Accessibility**: Also show on keyboard focus

---

## Step 3: Feature Limits ✅

**Decision: Maximum 3 features can be selected at once**

### Rationale:

1. **Consistency**: Matches the existing 3 SDK limit
2. **UI Constraint**: More than 3 rows keeps table readable
3. **Use Case Alignment**: Users typically evaluate a small set of key features
4. **Performance**: Keeps DOM updates manageable

### Implementation:

- Counter display: "Selected: X/3"
- Disable unselected features when limit reached
- Clear visual feedback (opacity, cursor change)
- No need for warning message; UI prevents over-selection

---

## Step 4: Data Structure Planning ✅

### Current Data Flow (Existing Page)

```
Input JSON Files → Go Generator → Embedded JSON in HTML → JavaScript Tables
```

**For SDK-driven view:**
- Go reads: features.json, types.json, names.json, feature_info.json
- Go embeds: SDKs array + features object
- JavaScript iterates: SDKs → show their features

### New Data Flow (Feature-Driven Page)

**Same input files, different organization:**

**Go Generator:**
1. Read same JSON files
2. Build inverted index structure:
   ```go
   // For each feature, list all SDKs and their support status
   type FeatureSDKSupport struct {
       FeatureID   string
       FeatureName string
       FeatureDesc string
       ClientSDKs  map[string]FeatureSupport  // sdk-id → support info
       ServerSDKs  map[string]FeatureSupport  // sdk-id → support info
   }
   ```

3. Embed in HTML template:
   - All features (sorted alphabetically)
   - For each feature: complete SDK support matrix
   - SDK metadata (names, types)

**JavaScript Structure:**

```javascript
// Embedded data structure
const featuresData = {
  "allFlags": {
    name: "Generate bootstrap details",
    description: "Return flag variations...",
    clientSDKs: {
      "android": { introduced: "5.5", deprecated: null, removed: null },
      "ios": { introduced: "6.0", deprecated: null, removed: null },
      // ... all client SDKs
    },
    serverSDKs: {
      "go-server-sdk": { introduced: "2.0", deprecated: null, removed: null },
      // ... all server SDKs
    }
  },
  // ... all features
};

const sdkNames = {
  "android": "Android SDK",
  "go-server-sdk": "Go Server SDK",
  // ... all SDK names
};

const sdkTypes = {
  "android": "client-side",
  "go-server-sdk": "server-side",
  // ... all SDK types
};
```

### Data Organization Strategy

**Option A: Feature-First Index (Recommended)**
- Top-level keys are feature IDs
- Each feature contains all SDK support data
- Efficient for feature selection workflow
- Example above

**Option B: Dual Index**
- Maintain both SDK→features AND feature→SDKs
- More data duplication
- Flexible but heavier payload

**Option C: Flat List with Filtering**
- Array of all feature-SDK combinations
- Filter in JavaScript based on selection
- Simplest Go code but more JS work

**Recommendation**: Option A (Feature-First Index)
- Most efficient for the feature-selection workflow
- Clear data structure
- Easy to iterate and render

### Data Processing in Go

```go
// Pseudocode for data transformation
func buildFeatureIndex(
    features map[string]map[string]FeatureSupport,  // from features.json
    types map[string]string,                         // from types.json
    names map[string]string,                         // from names.json
    featureInfo map[string]FeatureInfo,             // from feature_info.json
) map[string]FeatureSDKSupport {
    
    result := make(map[string]FeatureSDKSupport)
    
    // Get all unique features
    allFeatures := getAllUniqueFeatures(features)
    
    for _, featureID := range allFeatures {
        fss := FeatureSDKSupport{
            FeatureID:   featureID,
            FeatureName: featureInfo[featureID].Name,
            FeatureDesc: featureInfo[featureID].Description,
            ClientSDKs:  make(map[string]FeatureSupport),
            ServerSDKs:  make(map[string]FeatureSupport),
        }
        
        // For each SDK, check if it supports this feature
        for sdkID, sdkFeatures := range features {
            support := sdkFeatures[featureID]  // might be empty
            
            if types[sdkID] == "client-side" {
                fss.ClientSDKs[sdkID] = support
            } else if types[sdkID] == "server-side" {
                fss.ServerSDKs[sdkID] = support
            }
        }
        
        result[featureID] = fss
    }
    
    return result
}
```

### JavaScript Rendering Logic

```javascript
// Pseudocode for table rendering
function updateTable(section) {  // 'client' or 'server'
    const selectedFeatures = getSelectedFeatures();  // Array of feature IDs
    
    if (selectedFeatures.length === 0) {
        showEmptyState();
        return;
    }
    
    // Get all SDKs for this section
    const sdks = section === 'client' 
        ? getClientSDKs()
        : getServerSDKs();
    
    // Sort SDKs by name
    sdks.sort((a, b) => sdkNames[a].localeCompare(sdkNames[b]));
    
    // Build table header - SDK column + feature columns
    let headerHTML = '<th class="sdk-col">SDK</th>';
    selectedFeatures.forEach(featureID => {
        const feature = featuresData[featureID];
        headerHTML += `<th class="feature-col" title="${feature.description}">${feature.name}</th>`;
    });
    
    // Build table rows - one row per SDK
    let rowsHTML = '';
    sdks.forEach(sdkID => {
        rowsHTML += `<tr>
            <td class="sdk-name">${sdkNames[sdkID]}</td>`;
        
        selectedFeatures.forEach(featureID => {
            const feature = featuresData[featureID];
            const sdkData = section === 'client' ? feature.clientSDKs : feature.serverSDKs;
            const support = sdkData[sdkID];
            
            rowsHTML += `<td>${formatSupport(support)}</td>`;
        });
        
        rowsHTML += '</tr>';
    });
    
    // Inject into DOM
    renderTable(headerHTML, rowsHTML);
}

function getClientSDKs() {
    // Return all client-side SDK IDs
    return Object.keys(sdkTypes).filter(id => sdkTypes[id] === 'client-side');
}

function getServerSDKs() {
    // Return all server-side SDK IDs
    return Object.keys(sdkTypes).filter(id => sdkTypes[id] === 'server-side');
}
```

---

## Step 5: Go Generator Updates ⏸️

**PAUSING HERE FOR REVIEW AND APPROVAL**

### Decisions Made:

1. ✅ **Make target**: `make html` generates both pages automatically
2. ✅ **Go code organization**: Add `--view=by-sdk|by-feature` flag to existing `genhtml` command (Option B)
3. ✅ **Coding standards**: Use idiomatic Go and best practices; no unit tests for POC

### Planned Implementation Approach:

**1. Command-Line Interface:**
```go
// Add new flag
view := flag.String("view", "by-sdk", "View type: by-sdk or by-feature")

// Validate flag value
if *view != "by-sdk" && *view != "by-feature" {
    return fmt.Errorf("invalid view: %s (must be 'by-sdk' or 'by-feature')", *view)
}
```

**2. Directory Structure:**
```
tool/cmd/genhtml/
├── main.go              - Go code (~300 lines instead of 758)
└── templates/
    ├── by-sdk.html      - HTML template for SDK-driven view (~750 lines)
    └── by-feature.html  - HTML template for feature-driven view (~750 lines)
```

**3. Code Structure Changes in main.go:**
```go
package main

import (
    _ "embed"  // Required for go:embed
    ...
)

// Embed templates at compile time
//go:embed templates/by-sdk.html
var htmlTemplateBySDK string

//go:embed templates/by-feature.html
var htmlTemplateByFeature string

// Functions:
├── main() - parse flags, call run()
├── run() - orchestrate data reading and HTML generation
├── Data reading functions (existing):
│   ├── readFeaturesJSON()
│   ├── readTypesJSON()
│   ├── readNamesJSON()
│   └── readFeatureInfoJSON()
├── Data transformation functions:
│   ├── buildSDKIndex() - existing logic for by-sdk view (refactor from run())
│   └── buildFeatureIndex() - NEW - for by-feature view
└── Template execution:
    ├── generateBySDK() - uses htmlTemplateBySDK
    └── generateByFeature() - uses htmlTemplateByFeature
```

**3. New Data Transformation Function:**
```go
// buildFeatureIndex creates a feature-first index for by-feature view
func buildFeatureIndex(
    features map[string]map[string]FeatureSupport,
    types map[string]string,
    names map[string]string,
    featureInfo map[string]FeatureInfo,
) ([]FeatureData, map[string]string, map[string]string) {
    // Returns:
    // 1. Array of FeatureData (sorted alphabetically)
    // 2. SDK names map
    // 3. SDK types map
}

type FeatureData struct {
    ID          string
    Name        string
    Description string
    ClientSDKs  map[string]FeatureSupport
    ServerSDKs  map[string]FeatureSupport
}
```

**4. Template Strategy (using go:embed):**
- Extract existing template to `templates/by-sdk.html`
- Create new template at `templates/by-feature.html`
- Both templates embedded at compile time using `go:embed`
- Benefits:
  - Proper HTML syntax highlighting and validation
  - Easier to edit and maintain
  - Clean separation of Go code and HTML
  - Still produces single binary (no runtime file loading)
- Shared CSS/JS patterns where possible
- Each template optimized for its view

**5. Makefile Update:**
```makefile
.PHONY: html
html: #! Generate SDK features HTML comparison pages (both views)
	cd tool && go run ./cmd/genhtml --view=by-sdk --output ../products/feature-comparison-by-sdk.html --data ../products
	cd tool && go run ./cmd/genhtml --view=by-feature --output ../products/sdk-comparison-by-feature.html --data ../products
```

**6. GitHub Actions Update:**
```yaml
- name: Generate HTML
  run: |
    cd tool && go run ./cmd/genhtml --view=by-sdk --output ../products/feature-comparison-by-sdk.html --data ../products
    cd tool && go run ./cmd/genhtml --view=by-feature --output ../products/sdk-comparison-by-feature.html --data ../products

- name: Prepare publish directory
  run: |
    mkdir -p _site
    cp products/feature-comparison-by-sdk.html _site/index.html
    cp products/sdk-comparison-by-feature.html _site/by-feature.html
```

**7. Navigation Links:**
- Both pages will include tab navigation
- "By SDK" tab → links to `index.html` (or `feature-comparison-by-sdk.html`)
- "By Feature" tab → links to `by-feature.html` (or `sdk-comparison-by-feature.html`)

### File Outputs:

1. `products/feature-comparison-by-sdk.html` - Existing page renamed
2. `products/sdk-comparison-by-feature.html` - New feature-driven page

### Backward Compatibility:

- Current `features.html` location - options:
  - Option A: Rename to new naming convention
  - Option B: Keep as symlink/redirect for backward compatibility
  - Option C: Generate both old and new paths during transition

**Recommended**: Option A (clean break with new naming) since this is in development branch

### Implementation Steps for Step 5:

1. ✅ Add `go:embed` imports to main.go
2. ✅ Create `tool/cmd/genhtml/templates/` directory
3. ✅ Extract current template to `templates/by-sdk.html`
4. ✅ Add `--view` flag with validation
5. ✅ Refactor `run()` to handle both views:
   - Read all JSON files (same for both views)
   - Branch based on `--view` flag
   - Call appropriate data transformation and template execution
6. ✅ Create `buildFeatureIndex()` function
7. ✅ Create `templates/by-feature.html` template
8. ✅ Add tab navigation to both templates
9. ✅ Test both views locally

**Progress Update:**
- ✅ Successfully refactored Go code to use `go:embed`
- ✅ Both by-sdk and by-feature views tested and working
- ✅ Tab navigation added to both templates
- ✅ Makefile updated to generate both pages
- ✅ GitHub Actions workflow updated to deploy both pages

**Step 5 Complete!**

---

## Step 6: JavaScript Implementation ✅

The JavaScript for both views has been implemented as part of the templates:

**by-sdk.html JavaScript:**
- SDK selection with 3-item limit per section
- Dynamic table generation (features as rows, SDKs as columns)
- Status formatting with color coding and badges
- Theme toggle functionality

**by-feature.html JavaScript:**
- Feature selection with 3-item limit per section
- Feature tooltips on hover showing descriptions
- Dynamic table generation (SDKs as rows, features as columns)
- Status formatting with color coding and badges
- Theme toggle functionality

---

## Step 7: Makefile Update ✅

Updated `Makefile` to generate both pages with `make html` command:
```makefile
html: #! Generate SDK features HTML comparison pages (both views)
	cd tool && go run ./cmd/genhtml --view=by-sdk --output ../products/feature-comparison-by-sdk.html --data ../products
	cd tool && go run ./cmd/genhtml --view=by-feature --output ../products/sdk-comparison-by-feature.html --data ../products
```

Tested and confirmed working.

---

## Step 8: GitHub Actions Update ✅

Updated `.github/workflows/publish-pages.yml` to:
1. Generate both HTML pages
2. Copy by-sdk page to `index.html` (landing page)
3. Copy by-feature page to `by-feature.html`

Changes made to:
- "Generate HTML" step
- "Prepare publish directory" step

---

## Step 9: Local Testing ✅

Both views successfully tested:
- ✅ by-sdk view generates correctly
- ✅ by-feature view generates correctly
- ✅ `make html` generates both pages
- ✅ Templates use `go:embed` properly
- ✅ Tab navigation links work between pages

### Benefits of go:embed Approach:

- **Developer Experience**: Edit HTML with full IDE support (syntax highlighting, auto-complete, validation)
- **Maintainability**: Clear separation between Go logic and HTML presentation
- **Deployment**: Still a single binary - templates embedded at compile time
- **Modern Go**: Using Go 1.16+ best practices
- **File Organization**: Clean project structure with dedicated templates directory



| Date | Decision | Rationale |
|------|----------|-----------|
| 2026-02-10 | Use separate pages approach (Option 2) | Different mental models, optimization freedom, simpler implementation |
| 2026-02-10 | 3 feature selection limit | Consistent with SDK limit, keeps UI manageable |
| 2026-02-10 | Show all SDKs, use "-" for unsupported | Complete view of SDK landscape |
| 2026-02-10 | Alphabetical by feature name | Simple, predictable ordering |
| 2026-02-10 | Show description on hover only | Keeps feature list compact |
| 2026-02-10 | No search/filter initially | Start simple, add if needed based on feedback |
| 2026-02-10 | File naming: `feature-comparison-by-sdk.html` and `sdk-comparison-by-feature.html` | Clear naming convention |
| 2026-02-10 | Maintain client/server separation | Consistent with existing page, can combine later if needed |
| 2026-02-10 | "By SDK" page as landing page | Existing page becomes default, with navigation to new view |
| 2026-02-10 | Table orientation: SDKs as rows, features as columns | Better for viewing all SDKs, less horizontal scrolling |
| 2026-02-10 | Tab-style navigation | Most discoverable and conventional for web |
| 2026-02-10 | `make html` generates both pages | Single command for convenience |
| 2026-02-10 | Use `--view` flag in genhtml | Single command with view parameter |
| 2026-02-10 | Idiomatic Go, no unit tests for POC | Following best practices, keeping POC lightweight |
| 2026-02-10 | Use `go:embed` with separate template files | Modern best practice, easier maintenance, single binary |

---

## Implementation Summary

### What Was Built

Successfully implemented a feature-driven comparison view for LaunchDarkly SDKs as a companion to the existing SDK-driven view.

### Files Created/Modified

**New Files:**
- `tool/cmd/genhtml/templates/by-sdk.html` - SDK-driven view template
- `tool/cmd/genhtml/templates/by-feature.html` - Feature-driven view template
- `.cursor/sdk-comparison-by-feature.md` - This design document

**Modified Files:**
- `tool/cmd/genhtml/main.go` - Refactored to use `go:embed`, added `--view` flag and by-feature logic
- `Makefile` - Updated to generate both pages
- `.github/workflows/publish-pages.yml` - Updated to deploy both pages
- `README.md` - Added SDK Feature Comparison Tool documentation

**Generated Files:**
- `products/feature-comparison-by-sdk.html` - SDK-driven comparison page
- `products/sdk-comparison-by-feature.html` - Feature-driven comparison page

### Key Features Implemented

1. **Feature Selection UI**
   - Select up to 3 features per section (client/server)
   - Hover tooltips showing feature descriptions
   - Visual feedback for selection state
   - Alphabetically sorted feature list

2. **Comparison Table**
   - SDKs displayed as rows, features as columns
   - All SDKs shown (not just those with selected features)
   - Color-coded status indicators
   - Version information with deprecated/removed badges

3. **Navigation**
   - Tab-style navigation between By SDK and By Feature views
   - Consistent header across both pages
   - Theme toggle (dark/light mode)

4. **Technical Architecture**
   - Go templates separated from code using `go:embed`
   - Feature-first data indexing for efficient lookups
   - Single command with `--view` flag
   - Automated build and deployment pipeline

### Testing Results

✅ Both views generate successfully
✅ `make html` produces both pages
✅ Tab navigation works between views
✅ Dark/light theme toggle functions properly
✅ Feature selection and table rendering work correctly
✅ Responsive design adapts to different screen sizes

### Deployment

Pages are deployed to GitHub Pages:
- `index.html` - By SDK view (landing page)
- `by-feature.html` - By Feature view

### Future Enhancements

Potential improvements mentioned during planning:
- Search/filter for features (if list becomes too long)
- Sortable table columns
- Deep linking with URL parameters to save selections
- Export functionality (CSV/PDF)
- Combine client/server sections (depending on feedback)
- Remove 3-item selection limit (if UI handles it well)

### Documentation

- Design decisions documented in this file
- README updated with tool description and usage
- Code comments added for clarity
- Original functionality documented in `.cursor/sdk-comparison.md`

