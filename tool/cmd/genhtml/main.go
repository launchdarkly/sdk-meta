package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
)

//go:embed templates/by-sdk.html
var htmlTemplateBySDK string

//go:embed templates/by-feature.html
var htmlTemplateByFeature string

const (
	// MaxSDKSelection is the maximum number of SDKs that can be selected for comparison
	MaxSDKSelection = 3
	// MaxFeatureSelection is the maximum number of features that can be selected for comparison
	MaxFeatureSelection = 5
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

// TemplateData contains all data passed to the HTML template for by-sdk view
type TemplateData struct {
	ClientSDKs    []SDK
	ServerSDKs    []SDK
	AllFeatures   []string
	FeatureInfo   map[string]FeatureInfo
	FeaturesJSON  template.JS
	SDKDataJSON   template.JS
	MaxSelection  int
}

// FeatureTemplateData contains all data passed to the HTML template for by-feature view
type FeatureTemplateData struct {
	Features         []FeatureData
	SDKNamesJSON     template.JS
	SDKTypesJSON     template.JS
	FeaturesDataJSON template.JS
	MaxSelection     int
}

// FeatureData represents a feature with its SDK support information
type FeatureData struct {
	ID          string                        `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	ClientSDKs  map[string]FeatureSupport     `json:"clientSDKs"`
	ServerSDKs  map[string]FeatureSupport     `json:"serverSDKs"`
}

func main() {
	outputPath := flag.String("output", "products/features.html", "Output HTML file path")
	dataPath := flag.String("data", "products", "Path to products directory")
	view := flag.String("view", "by-sdk", "View type: by-sdk or by-feature")
	flag.Parse()

	// Validate view parameter
	if *view != "by-sdk" && *view != "by-feature" {
		fmt.Fprintf(os.Stderr, "Error: invalid view '%s' (must be 'by-sdk' or 'by-feature')\n", *view)
		os.Exit(1)
	}

	if err := run(*dataPath, *outputPath, *view); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s (%s view)\n", *outputPath, *view)
}

func run(dataPath, outputPath, view string) error {
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

	// Branch based on view type
	if view == "by-sdk" {
		return generateBySDKView(features, types, names, featureInfo, outputPath)
	}
	return generateByFeatureView(features, types, names, featureInfo, outputPath)
}

func generateBySDKView(
	features map[string]map[string]FeatureSupport,
	types map[string]string,
	names map[string]string,
	featureInfo map[string]FeatureInfo,
	outputPath string,
) error {
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
		MaxSelection: MaxSDKSelection,
	}

	// Generate HTML
	tmpl, err := template.New("by-sdk").Parse(htmlTemplateBySDK)
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

func generateByFeatureView(
	features map[string]map[string]FeatureSupport,
	types map[string]string,
	names map[string]string,
	featureInfo map[string]FeatureInfo,
	outputPath string,
) error {
	// Build feature-first index
	featuresData := buildFeatureIndex(features, types, names, featureInfo)

	// Sort features alphabetically by name
	sort.Slice(featuresData, func(i, j int) bool {
		return featuresData[i].Name < featuresData[j].Name
	})

	// Convert to JSON for JavaScript
	featuresDataBytes, err := json.Marshal(featuresData)
	if err != nil {
		return fmt.Errorf("marshaling features data: %w", err)
	}

	sdkNamesBytes, err := json.Marshal(names)
	if err != nil {
		return fmt.Errorf("marshaling SDK names: %w", err)
	}

	sdkTypesBytes, err := json.Marshal(types)
	if err != nil {
		return fmt.Errorf("marshaling SDK types: %w", err)
	}

	// Prepare template data
	data := FeatureTemplateData{
		Features:         featuresData,
		SDKNamesJSON:     template.JS(sdkNamesBytes),
		SDKTypesJSON:     template.JS(sdkTypesBytes),
		FeaturesDataJSON: template.JS(featuresDataBytes),
		MaxSelection:     MaxFeatureSelection,
	}

	// Generate HTML
	tmpl, err := template.New("by-feature").Parse(htmlTemplateByFeature)
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

// buildFeatureIndex creates a feature-first index for by-feature view
func buildFeatureIndex(
	features map[string]map[string]FeatureSupport,
	types map[string]string,
	names map[string]string,
	featureInfo map[string]FeatureInfo,
) []FeatureData {
	// Collect all unique features
	allFeaturesMap := make(map[string]bool)
	for _, sdkFeatures := range features {
		for featureID := range sdkFeatures {
			allFeaturesMap[featureID] = true
		}
	}

	// Build feature data for each feature
	var result []FeatureData
	for featureID := range allFeaturesMap {
		info := featureInfo[featureID]
		if info.Name == "" {
			info.Name = featureID
		}

		fd := FeatureData{
			ID:          featureID,
			Name:        info.Name,
			Description: info.Description,
			ClientSDKs:  make(map[string]FeatureSupport),
			ServerSDKs:  make(map[string]FeatureSupport),
		}

		// For each SDK, check if it supports this feature
		for sdkID, sdkType := range types {
			support := FeatureSupport{}
			if sdkFeatures, exists := features[sdkID]; exists {
				if featureSupport, hasFeature := sdkFeatures[featureID]; hasFeature {
					support = featureSupport
				}
			}

			// Add to appropriate category
			if sdkType == "client-side" {
				fd.ClientSDKs[sdkID] = support
			} else if sdkType == "server-side" {
				fd.ServerSDKs[sdkID] = support
			}
		}

		result = append(result, fd)
	}

	return result
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

