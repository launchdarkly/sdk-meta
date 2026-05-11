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

//go:embed templates/spec-support.html
var matrixTemplate string

//go:embed templates/spec-support-by-sdk.html
var bySDKTemplate string

func runHTML(args []string) error {
	fs := flag.NewFlagSet("html", flag.ExitOnError)
	specsJSON := fs.String("specs-json", "products/specs.json", "Path to specs.json (input).")
	supportJSON := fs.String("support-json", "products/spec_support.json", "Path to spec_support.json (input).")
	typesJSON := fs.String("types-json", "products/types.json", "Path to types.json (input).")
	namesJSON := fs.String("names-json", "products/names.json", "Path to names.json (input).")
	outDir := fs.String("out-dir", "_site", "Output directory for generated HTML.")
	if err := fs.Parse(args); err != nil {
		return err
	}

	specs, err := loadJSON[SpecsProduct](*specsJSON)
	if err != nil {
		return fmt.Errorf("loading specs.json: %w", err)
	}
	support, err := loadJSON[SpecSupportProduct](*supportJSON)
	if err != nil {
		return fmt.Errorf("loading spec_support.json: %w", err)
	}
	types, err := loadJSON[map[string]string](*typesJSON)
	if err != nil {
		return fmt.Errorf("loading types.json: %w", err)
	}
	names, err := loadJSON[map[string]string](*namesJSON)
	if err != nil {
		return fmt.Errorf("loading names.json: %w", err)
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		return err
	}

	specOrder := sortedKeys(specs.Specs)
	sdkOrder := make([]string, 0, len(types))
	for id := range types {
		sdkOrder = append(sdkOrder, id)
	}
	sort.Slice(sdkOrder, func(i, j int) bool {
		ni, nj := names[sdkOrder[i]], names[sdkOrder[j]]
		if ni == "" {
			ni = sdkOrder[i]
		}
		if nj == "" {
			nj = sdkOrder[j]
		}
		return ni < nj
	})

	payload := map[string]any{
		"specs":     specs.Specs,
		"sdkOrder":  sdkOrder,
		"specOrder": specOrder,
		"types":     types,
		"names":     names,
		"cells":     support.SDKs,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	type tmplData struct {
		GeneratedAt   string
		SpecsCommit   string
		HarnessCommit string
		Model         string
		PromptVersion string
		PayloadJSON   template.JS
	}
	td := tmplData{
		GeneratedAt:   support.GeneratedAt,
		SpecsCommit:   support.SpecsCommit,
		HarnessCommit: support.HarnessCommit,
		Model:         support.Model,
		PromptVersion: support.PromptVersion,
		PayloadJSON:   template.JS(payloadJSON),
	}

	funcs := template.FuncMap{
		"shortCommit": func(c string) string {
			if len(c) > 8 {
				return c[:8]
			}
			return c
		},
	}

	if err := renderTemplate(matrixTemplate, funcs, td, filepath.Join(*outDir, "spec-support.html")); err != nil {
		return fmt.Errorf("rendering matrix view: %w", err)
	}
	if err := renderTemplate(bySDKTemplate, funcs, td, filepath.Join(*outDir, "spec-support-by-sdk.html")); err != nil {
		return fmt.Errorf("rendering by-sdk view: %w", err)
	}
	fmt.Fprintf(os.Stderr, "Wrote %s and %s\n",
		filepath.Join(*outDir, "spec-support.html"),
		filepath.Join(*outDir, "spec-support-by-sdk.html"))
	return nil
}

func renderTemplate(src string, funcs template.FuncMap, td any, outPath string) error {
	tmpl, err := template.New("page").Funcs(funcs).Parse(src)
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, td)
}
