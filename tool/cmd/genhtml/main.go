package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
)

// FeatureSupport represents the support status of a feature in an SDK
type FeatureSupport struct {
	Introduced *string `json:"introduced"`
	Deprecated *string `json:"deprecated"`
	Removed    *string `json:"removed"`
}

// FeatureInfo contains the name and description of a feature
type FeatureInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SDK represents an SDK with its features
type SDK struct {
	ID       string
	Name     string
	Type     string
	Features map[string]FeatureSupport
}

// TemplateData contains all data passed to the HTML template
type TemplateData struct {
	ClientSDKs    []SDK
	ServerSDKs    []SDK
	AllFeatures   []string
	FeatureInfo   map[string]FeatureInfo
	FeaturesJSON  template.JS
	SDKDataJSON   template.JS
}

func main() {
	outputPath := flag.String("output", "products/features.html", "Output HTML file path")
	dataPath := flag.String("data", "products", "Path to products directory")
	flag.Parse()

	if err := run(*dataPath, *outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s\n", *outputPath)
}

func run(dataPath, outputPath string) error {
	// Read all JSON data files
	features, err := readFeaturesJSON(filepath.Join(dataPath, "features.json"))
	if err != nil {
		return fmt.Errorf("reading features.json: %w", err)
	}

	types, err := readTypesJSON(filepath.Join(dataPath, "types.json"))
	if err != nil {
		return fmt.Errorf("reading types.json: %w", err)
	}

	names, err := readNamesJSON(filepath.Join(dataPath, "names.json"))
	if err != nil {
		return fmt.Errorf("reading names.json: %w", err)
	}

	featureInfo, err := readFeatureInfoJSON(filepath.Join(dataPath, "feature_info.json"))
	if err != nil {
		return fmt.Errorf("reading feature_info.json: %w", err)
	}

	// Build SDK list
	var clientSDKs, serverSDKs []SDK
	allFeaturesMap := make(map[string]bool)

	for sdkID, sdkType := range types {
		sdk := SDK{
			ID:       sdkID,
			Name:     names[sdkID],
			Type:     sdkType,
			Features: features[sdkID],
		}
		if sdk.Name == "" {
			sdk.Name = sdkID
		}
		if sdk.Features == nil {
			sdk.Features = make(map[string]FeatureSupport)
		}

		// Collect all features
		for feature := range sdk.Features {
			allFeaturesMap[feature] = true
		}

		switch sdkType {
		case "client-side":
			clientSDKs = append(clientSDKs, sdk)
		case "server-side":
			serverSDKs = append(serverSDKs, sdk)
		}
	}

	// Sort SDKs by name
	sort.Slice(clientSDKs, func(i, j int) bool {
		return clientSDKs[i].Name < clientSDKs[j].Name
	})
	sort.Slice(serverSDKs, func(i, j int) bool {
		return serverSDKs[i].Name < serverSDKs[j].Name
	})

	// Sort features alphabetically
	var allFeatures []string
	for feature := range allFeaturesMap {
		allFeatures = append(allFeatures, feature)
	}
	sort.Strings(allFeatures)

	// Create JSON for JavaScript
	sdkData := make(map[string]map[string]FeatureSupport)
	for _, sdk := range append(clientSDKs, serverSDKs...) {
		sdkData[sdk.ID] = sdk.Features
	}
	sdkDataBytes, err := json.Marshal(sdkData)
	if err != nil {
		return fmt.Errorf("marshaling SDK data: %w", err)
	}

	featureInfoBytes, err := json.Marshal(featureInfo)
	if err != nil {
		return fmt.Errorf("marshaling feature info: %w", err)
	}

	// Prepare template data
	data := TemplateData{
		ClientSDKs:   clientSDKs,
		ServerSDKs:   serverSDKs,
		AllFeatures:  allFeatures,
		FeatureInfo:  featureInfo,
		FeaturesJSON: template.JS(featureInfoBytes),
		SDKDataJSON:  template.JS(sdkDataBytes),
	}

	// Generate HTML
	tmpl, err := template.New("features").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}

func readFeaturesJSON(path string) (map[string]map[string]FeatureSupport, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]map[string]FeatureSupport
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func readTypesJSON(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func readNamesJSON(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func readFeatureInfoJSON(path string) (map[string]FeatureInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]FeatureInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

var htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>LaunchDarkly SDK Feature Comparison</title>
  <style>
    :root {
      --ld-purple: #405BFF;
      --ld-dark: #282828;
      --ld-light: #F5F5F5;
      --bg-primary: #F5F5F5;
      --bg-secondary: #FFFFFF;
      --text-primary: #282828;
      --text-secondary: #6C757D;
      --supported-green: #28A745;
      --deprecated-yellow: #FFC107;
      --removed-red: #DC3545;
      --not-supported-gray: #6C757D;
      --border-color: #DEE2E6;
      --table-header-bg: #282828;
      --table-header-text: #FFFFFF;
      --row-hover: #f8f9fa;
    }

    @media (prefers-color-scheme: dark) {
      :root:not([data-theme="light"]) {
        --bg-primary: #1a1a1a;
        --bg-secondary: #2d2d2d;
        --text-primary: #e0e0e0;
        --text-secondary: #a0a0a0;
        --border-color: #404040;
        --table-header-bg: #1a1a1a;
        --table-header-text: #e0e0e0;
        --row-hover: #363636;
        --not-supported-gray: #808080;
      }
    }

    [data-theme="dark"] {
      --bg-primary: #1a1a1a;
      --bg-secondary: #2d2d2d;
      --text-primary: #e0e0e0;
      --text-secondary: #a0a0a0;
      --border-color: #404040;
      --table-header-bg: #1a1a1a;
      --table-header-text: #e0e0e0;
      --row-hover: #363636;
      --not-supported-gray: #808080;
    }

    * {
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
      margin: 0;
      padding: 20px;
      background-color: var(--bg-primary);
      color: var(--text-primary);
      transition: background-color 0.3s, color 0.3s;
    }

    .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 10px;
    }

    .header h1 {
      margin: 0;
      flex: 1;
      text-align: center;
    }

    .theme-toggle {
      background: var(--bg-secondary);
      border: 1px solid var(--border-color);
      border-radius: 6px;
      padding: 8px 12px;
      cursor: pointer;
      font-size: 1.2em;
      transition: all 0.2s;
      display: flex;
      align-items: center;
      gap: 6px;
    }

    .theme-toggle:hover {
      border-color: var(--ld-purple);
    }

    .theme-toggle .icon {
      font-size: 1.1em;
    }

    h1 {
      color: var(--ld-purple);
      text-align: center;
      margin-bottom: 10px;
    }

    .subtitle {
      text-align: center;
      color: var(--text-secondary);
      margin-bottom: 30px;
    }

    section {
      background: var(--bg-secondary);
      border-radius: 8px;
      padding: 20px;
      margin-bottom: 30px;
      box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }

    h2 {
      color: var(--text-primary);
      border-bottom: 2px solid var(--ld-purple);
      padding-bottom: 10px;
      margin-top: 0;
    }

    .sdk-selector {
      display: flex;
      flex-wrap: wrap;
      gap: 10px;
      margin-bottom: 20px;
      padding: 15px;
      background-color: var(--bg-primary);
      border-radius: 6px;
    }

    .sdk-selector label {
      display: flex;
      align-items: center;
      gap: 5px;
      padding: 8px 12px;
      background: var(--bg-secondary);
      border: 1px solid var(--border-color);
      border-radius: 4px;
      cursor: pointer;
      transition: all 0.2s;
    }

    .sdk-selector label:hover {
      border-color: var(--ld-purple);
    }

    .sdk-selector label.selected {
      background-color: var(--ld-purple);
      color: white;
      border-color: var(--ld-purple);
    }

    .sdk-selector label.disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    .sdk-selector input[type="checkbox"] {
      display: none;
    }

    .selection-info {
      width: 100%;
      font-size: 0.9em;
      color: var(--text-secondary);
      margin-top: 5px;
    }

    .table-container {
      overflow-x: auto;
    }

    table {
      width: 100%;
      border-collapse: collapse;
      font-size: 14px;
    }

    th, td {
      padding: 12px 15px;
      text-align: left;
      border-bottom: 1px solid var(--border-color);
    }

    th {
      background-color: var(--table-header-bg);
      color: var(--table-header-text);
      font-weight: 600;
      position: sticky;
      top: 0;
    }

    th.feature-col {
      min-width: 200px;
    }

    th.sdk-col {
      min-width: 150px;
    }

    tr:hover {
      background-color: var(--row-hover);
    }

    .feature-name {
      font-weight: 500;
    }

    .feature-description {
      font-size: 0.85em;
      color: var(--text-secondary);
      margin-top: 4px;
    }

    .status {
      display: inline-flex;
      align-items: center;
      gap: 5px;
    }

    .status-supported {
      color: var(--supported-green);
    }

    .status-deprecated {
      color: var(--deprecated-yellow);
    }

    .status-removed {
      color: var(--removed-red);
    }

    .status-not-supported {
      color: var(--not-supported-gray);
    }

    .badge {
      font-size: 0.75em;
      padding: 2px 6px;
      border-radius: 3px;
      font-weight: 600;
      text-transform: uppercase;
    }

    .badge-deprecated {
      background-color: var(--deprecated-yellow);
      color: var(--ld-dark);
    }

    .badge-removed {
      background-color: var(--removed-red);
      color: white;
    }

    .no-selection {
      text-align: center;
      padding: 40px;
      color: var(--text-secondary);
    }

    @media (max-width: 768px) {
      body {
        padding: 10px;
      }

      .sdk-selector label {
        font-size: 0.9em;
        padding: 6px 10px;
      }

      th, td {
        padding: 8px 10px;
        font-size: 13px;
      }

      .header h1 {
        font-size: 1.3em;
      }

      .theme-toggle span:not(.icon) {
        display: none;
      }
    }
  </style>
</head>
<body>
  <div class="header">
    <div style="width: 100px;"></div>
    <h1>LaunchDarkly SDK Feature Comparison</h1>
    <button class="theme-toggle" id="theme-toggle" title="Toggle dark mode">
      <span class="icon">🌙</span>
      <span>Dark</span>
    </button>
  </div>
  <p class="subtitle">Select up to 3 SDKs in each section to compare their features</p>

  <section id="client-side">
    <h2>Client-side SDKs</h2>
    <div class="sdk-selector" data-section="client">
      {{range .ClientSDKs}}
      <label data-sdk="{{.ID}}">
        <input type="checkbox" value="{{.ID}}" data-name="{{.Name}}">
        {{.Name}}
      </label>
      {{end}}
      <div class="selection-info">Selected: <span class="count">0</span>/3</div>
    </div>
    <div class="table-container">
      <table class="comparison-table" data-section="client">
        <thead>
          <tr>
            <th class="feature-col">Feature</th>
          </tr>
        </thead>
        <tbody>
        </tbody>
      </table>
    </div>
    <div class="no-selection">Select SDKs above to compare features</div>
  </section>

  <section id="server-side">
    <h2>Server-side SDKs</h2>
    <div class="sdk-selector" data-section="server">
      {{range .ServerSDKs}}
      <label data-sdk="{{.ID}}">
        <input type="checkbox" value="{{.ID}}" data-name="{{.Name}}">
        {{.Name}}
      </label>
      {{end}}
      <div class="selection-info">Selected: <span class="count">0</span>/3</div>
    </div>
    <div class="table-container">
      <table class="comparison-table" data-section="server">
        <thead>
          <tr>
            <th class="feature-col">Feature</th>
          </tr>
        </thead>
        <tbody>
        </tbody>
      </table>
    </div>
    <div class="no-selection">Select SDKs above to compare features</div>
  </section>

  <script>
    const sdkData = {{.SDKDataJSON}};
    const featureInfo = {{.FeaturesJSON}};
    const MAX_SELECTION = 3;

    // Track selected SDKs per section
    const selectedSDKs = {
      client: [],
      server: []
    };

    // Initialize
    document.querySelectorAll('.sdk-selector').forEach(selector => {
      const section = selector.dataset.section;

      selector.querySelectorAll('input[type="checkbox"]').forEach(checkbox => {
        checkbox.addEventListener('change', (e) => {
          const sdkId = e.target.value;
          const sdkName = e.target.dataset.name;
          const label = e.target.closest('label');

          if (e.target.checked) {
            if (selectedSDKs[section].length < MAX_SELECTION) {
              selectedSDKs[section].push({ id: sdkId, name: sdkName });
              label.classList.add('selected');
            } else {
              e.target.checked = false;
              return;
            }
          } else {
            selectedSDKs[section] = selectedSDKs[section].filter(sdk => sdk.id !== sdkId);
            label.classList.remove('selected');
          }

          updateSelectorState(section);
          updateTable(section);
        });
      });
    });

    function updateSelectorState(section) {
      const selector = document.querySelector(` + "`" + `.sdk-selector[data-section="${section}"]` + "`" + `);
      const count = selectedSDKs[section].length;

      // Update count display
      selector.querySelector('.count').textContent = count;

      // Enable/disable checkboxes
      selector.querySelectorAll('label').forEach(label => {
        const checkbox = label.querySelector('input');
        if (!checkbox.checked && count >= MAX_SELECTION) {
          label.classList.add('disabled');
          checkbox.disabled = true;
        } else {
          label.classList.remove('disabled');
          checkbox.disabled = false;
        }
      });
    }

    function updateTable(section) {
      const sdks = selectedSDKs[section];
      const sectionEl = document.getElementById(section + '-side');
      const table = sectionEl.querySelector('.comparison-table');
      const noSelection = sectionEl.querySelector('.no-selection');

      if (sdks.length === 0) {
        table.style.display = 'none';
        noSelection.style.display = 'block';
        return;
      }

      table.style.display = 'table';
      noSelection.style.display = 'none';

      // Build header
      const thead = table.querySelector('thead tr');
      thead.innerHTML = '<th class="feature-col">Feature</th>';
      sdks.forEach(sdk => {
        thead.innerHTML += ` + "`" + `<th class="sdk-col">${sdk.name}</th>` + "`" + `;
      });

      // Show all features from featureInfo
      const features = Object.keys(featureInfo).sort((a, b) => {
        const nameA = featureInfo[a]?.name || a;
        const nameB = featureInfo[b]?.name || b;
        return nameA.localeCompare(nameB);
      });

      // Build rows
      const tbody = table.querySelector('tbody');
      tbody.innerHTML = '';

      features.forEach(feature => {
        const info = featureInfo[feature] || { name: feature, description: '' };
        let row = ` + "`" + `<tr>
          <td>
            <div class="feature-name">${info.name}</div>
            <div class="feature-description">${info.description}</div>
          </td>` + "`" + `;

        sdks.forEach(sdk => {
          const support = sdkData[sdk.id]?.[feature];
          row += ` + "`" + `<td>${formatSupport(support)}</td>` + "`" + `;
        });

        row += '</tr>';
        tbody.innerHTML += row;
      });
    }

    function formatSupport(support) {
      if (!support || !support.introduced) {
        return '<span class="status status-not-supported">—</span>';
      }

      let html = '<span class="status ';

      if (support.removed) {
        html += 'status-removed">';
        html += support.introduced;
        html += ' <span class="badge badge-removed">removed ' + support.removed + '</span>';
      } else if (support.deprecated) {
        html += 'status-deprecated">';
        html += support.introduced;
        html += ' <span class="badge badge-deprecated">deprecated ' + support.deprecated + '</span>';
      } else {
        html += 'status-supported">';
        html += support.introduced;
      }

      html += '</span>';
      return html;
    }

    // Theme toggle functionality
    const themeToggle = document.getElementById('theme-toggle');
    const html = document.documentElement;

    function updateToggleButton(isDark) {
      const icon = themeToggle.querySelector('.icon');
      const text = themeToggle.querySelector('span:not(.icon)');
      if (isDark) {
        icon.textContent = '☀️';
        text.textContent = 'Light';
      } else {
        icon.textContent = '🌙';
        text.textContent = 'Dark';
      }
    }

    function getSystemTheme() {
      return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }

    function getCurrentTheme() {
      const stored = localStorage.getItem('theme');
      if (stored) return stored;
      return getSystemTheme();
    }

    function setTheme(theme) {
      if (theme === 'dark') {
        html.setAttribute('data-theme', 'dark');
      } else {
        html.setAttribute('data-theme', 'light');
      }
      localStorage.setItem('theme', theme);
      updateToggleButton(theme === 'dark');
    }

    // Initialize theme
    const initialTheme = getCurrentTheme();
    setTheme(initialTheme);

    // Toggle handler
    themeToggle.addEventListener('click', () => {
      const current = html.getAttribute('data-theme') || getSystemTheme();
      setTheme(current === 'dark' ? 'light' : 'dark');
    });

    // Listen for system theme changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
      if (!localStorage.getItem('theme')) {
        setTheme(e.matches ? 'dark' : 'light');
      }
    });
  </script>
</body>
</html>
`
