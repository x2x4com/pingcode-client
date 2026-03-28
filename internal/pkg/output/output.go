package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Format string

const (
	FormatJSON     Format = "json"
	FormatYAML     Format = "yaml"
	FormatMarkdown Format = "markdown"
	FormatTable    Format = "table"
)

type Options struct {
	Format Format
	Raw    bool
}

var stdout io.Writer = os.Stdout

func SetOutput(w io.Writer) {
	stdout = w
}

func FormatAndPrint(data interface{}, opts Options) error {
	if opts.Raw {
		return formatRaw(data, opts.Format)
	}
	return formatDisplay(data, opts.Format)
}

func formatRaw(data interface{}, format Format) error {
	var output []byte
	var err error

	switch format {
	case FormatJSON, FormatTable, "":
		output, err = json.MarshalIndent(data, "", "  ")
	case FormatYAML:
		output, err = yaml.Marshal(data)
	case FormatMarkdown:
		output, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, "```json")
		fmt.Fprintln(stdout, string(output))
		fmt.Fprintln(stdout, "```")
		return nil
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}

	if err != nil {
		return err
	}
	fmt.Fprintln(stdout, string(output))
	return nil
}

func formatDisplay(data interface{}, format Format) error {
	switch format {
	case FormatJSON:
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(data)
	case FormatYAML:
		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(output))
		return nil
	case FormatMarkdown:
		return formatMarkdown(data)
	case FormatTable, "":
		return formatTable(data)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

func formatMarkdown(data interface{}) error {
	switch v := data.(type) {
	case []interface{}:
		return formatSliceAsMarkdown(v)
	case []map[string]interface{}:
		return formatSliceAsMapMarkdown(v)
	case map[string]interface{}:
		return formatMapAsMarkdown(v)
	default:
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, "```json")
		fmt.Fprintln(stdout, string(b))
		fmt.Fprintln(stdout, "```")
		return nil
	}
}

func formatSliceAsMarkdown(items []interface{}) error {
	if len(items) == 0 {
		fmt.Fprintln(stdout, "No data")
		return nil
	}

	// Get keys from first item
	first, ok := items[0].(map[string]interface{})
	if !ok {
		b, err := json.MarshalIndent(items, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(b))
		return nil
	}

	keys := make([]string, 0)
	for k := range first {
		keys = append(keys, k)
	}

	// Build markdown table
	fmt.Fprint(stdout, "|")
	for _, k := range keys {
		fmt.Fprintf(stdout, " %s |", k)
	}
	fmt.Fprintln(stdout, "")
	fmt.Fprint(stdout, "|")
	for range keys {
		fmt.Fprint(stdout, "---|")
	}
	fmt.Fprintln(stdout, "")

	for _, item := range items {
		if m, ok := item.(map[string]interface{}); ok {
			fmt.Fprint(stdout, "|")
			for _, k := range keys {
				val := fmt.Sprintf("%v", m[k])
				fmt.Fprintf(stdout, " %s |", val)
			}
			fmt.Fprintln(stdout, "")
		}
	}
	return nil
}

func formatSliceAsMapMarkdown(items []map[string]interface{}) error {
	if len(items) == 0 {
		fmt.Fprintln(stdout, "No data")
		return nil
	}

	keys := make([]string, 0)
	for k := range items[0] {
		keys = append(keys, k)
	}

	// Build markdown table
	fmt.Fprint(stdout, "|")
	for _, k := range keys {
		fmt.Fprintf(stdout, " %s |", k)
	}
	fmt.Fprintln(stdout, "")
	fmt.Fprint(stdout, "|")
	for range keys {
		fmt.Fprint(stdout, "---|")
	}
	fmt.Fprintln(stdout, "")

	for _, item := range items {
		fmt.Fprint(stdout, "|")
		for _, k := range keys {
			val := fmt.Sprintf("%v", item[k])
			fmt.Fprintf(stdout, " %s |", val)
		}
		fmt.Fprintln(stdout, "")
	}
	return nil
}

func formatMapAsMarkdown(m map[string]interface{}) error {
	for k, v := range m {
		fmt.Fprintf(stdout, "**%s**: %v\n", k, v)
	}
	return nil
}

func formatTable(data interface{}) error {
	switch v := data.(type) {
	case []interface{}:
		return formatSliceAsTable(v)
	case []map[string]interface{}:
		return formatSliceAsMapTable(v)
	case map[string]interface{}:
		return formatMapAsTable(v)
	default:
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(b))
		return nil
	}
}

func formatSliceAsTable(items []interface{}) error {
	if len(items) == 0 {
		fmt.Fprintln(stdout, "No data")
		return nil
	}

	// Get keys from first item
	first, ok := items[0].(map[string]interface{})
	if !ok {
		b, err := json.MarshalIndent(items, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(b))
		return nil
	}

	// Determine column widths
	keys := make([]string, 0)
	width := make(map[string]int)
	for k := range first {
		keys = append(keys, k)
		width[k] = len(k)
	}

	for _, item := range items {
		if m, ok := item.(map[string]interface{}); ok {
			for _, k := range keys {
				val := fmt.Sprintf("%v", m[k])
				if len(val) > width[k] {
					width[k] = len(val)
				}
			}
		}
	}

	// Print header
	for _, k := range keys {
		fmt.Fprintf(stdout, "  %-*s", width[k], k)
	}
	fmt.Fprintln(stdout, "")

	// Print separator
	for _, k := range keys {
		fmt.Fprintf(stdout, "  %-*s", width[k], "---")
	}
	fmt.Fprintln(stdout, "")

	// Print rows
	for _, item := range items {
		if m, ok := item.(map[string]interface{}); ok {
			for _, k := range keys {
				val := fmt.Sprintf("%v", m[k])
				fmt.Fprintf(stdout, "  %-*s", width[k], val)
			}
			fmt.Fprintln(stdout, "")
		}
	}
	return nil
}

func formatMapAsTable(m map[string]interface{}) error {
	for k, v := range m {
		fmt.Fprintf(stdout, "  %s: %v\n", k, v)
	}
	return nil
}

func formatSliceAsMapTable(items []map[string]interface{}) error {
	if len(items) == 0 {
		fmt.Fprintln(stdout, "No data")
		return nil
	}

	// Determine column widths
	keys := make([]string, 0)
	width := make(map[string]int)
	for k := range items[0] {
		keys = append(keys, k)
		width[k] = len(k)
	}

	for _, item := range items {
		for _, k := range keys {
			val := fmt.Sprintf("%v", item[k])
			if len(val) > width[k] {
				width[k] = len(val)
			}
		}
	}

	// Print header
	for _, k := range keys {
		fmt.Fprintf(stdout, "  %-*s", width[k], k)
	}
	fmt.Fprintln(stdout, "")

	// Print separator
	for _, k := range keys {
		fmt.Fprintf(stdout, "  %-*s", width[k], "---")
	}
	fmt.Fprintln(stdout, "")

	// Print rows
	for _, item := range items {
		for _, k := range keys {
			val := fmt.Sprintf("%v", item[k])
			fmt.Fprintf(stdout, "  %-*s", width[k], val)
		}
		fmt.Fprintln(stdout, "")
	}
	return nil
}
