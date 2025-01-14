package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter items from a list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filterCmd).Standalone()

	filterCmd.Flags().String("cursor-text.align", "", "Text Alignment")
	filterCmd.Flags().String("cursor-text.background", "", "Background Color")
	filterCmd.Flags().Bool("cursor-text.bold", false, "Bold text")
	filterCmd.Flags().String("cursor-text.border", "", "Border Style")
	filterCmd.Flags().String("cursor-text.border-background", "", "Border Background Color")
	filterCmd.Flags().String("cursor-text.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("cursor-text.faint", false, "Faint text")
	filterCmd.Flags().String("cursor-text.foreground", "", "Foreground Color")
	filterCmd.Flags().String("cursor-text.height", "", "Text height")
	filterCmd.Flags().Bool("cursor-text.italic", false, "Italicize text")
	filterCmd.Flags().String("cursor-text.margin", "", "Text margin")
	filterCmd.Flags().String("cursor-text.padding", "", "Text padding")
	filterCmd.Flags().Bool("cursor-text.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("cursor-text.underline", false, "Underline text")
	filterCmd.Flags().String("cursor-text.width", "", "Text width")
	filterCmd.Flags().Bool("fuzzy", false, "Enable fuzzy matching; otherwise match from start of word")
	filterCmd.Flags().Bool("fuzzy-sort", false, "Sort fuzzy results by their scores")
	filterCmd.Flags().String("header", "", "Header value")
	filterCmd.Flags().String("header.align", "", "Text Alignment")
	filterCmd.Flags().String("header.background", "", "Background Color")
	filterCmd.Flags().Bool("header.bold", false, "Bold text")
	filterCmd.Flags().String("header.border", "", "Border Style")
	filterCmd.Flags().String("header.border-background", "", "Border Background Color")
	filterCmd.Flags().String("header.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("header.faint", false, "Faint text")
	filterCmd.Flags().String("header.foreground", "", "Foreground Color")
	filterCmd.Flags().String("header.height", "", "Text height")
	filterCmd.Flags().Bool("header.italic", false, "Italicize text")
	filterCmd.Flags().String("header.margin", "", "Text margin")
	filterCmd.Flags().String("header.padding", "", "Text padding")
	filterCmd.Flags().Bool("header.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("header.underline", false, "Underline text")
	filterCmd.Flags().String("header.width", "", "Text width")
	filterCmd.Flags().String("height", "", "Input height")
	filterCmd.Flags().String("indicator", "", "Character for selection")
	filterCmd.Flags().String("indicator.align", "", "Text Alignment")
	filterCmd.Flags().String("indicator.background", "", "Background Color")
	filterCmd.Flags().Bool("indicator.bold", false, "Bold text")
	filterCmd.Flags().String("indicator.border", "", "Border Style")
	filterCmd.Flags().String("indicator.border-background", "", "Border Background Color")
	filterCmd.Flags().String("indicator.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("indicator.faint", false, "Faint text")
	filterCmd.Flags().String("indicator.foreground", "", "Foreground Color")
	filterCmd.Flags().String("indicator.height", "", "Text height")
	filterCmd.Flags().Bool("indicator.italic", false, "Italicize text")
	filterCmd.Flags().String("indicator.margin", "", "Text margin")
	filterCmd.Flags().String("indicator.padding", "", "Text padding")
	filterCmd.Flags().Bool("indicator.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("indicator.underline", false, "Underline text")
	filterCmd.Flags().String("indicator.width", "", "Text width")
	filterCmd.Flags().String("input-delimiter", "", "Option delimiter when reading from STDIN")
	filterCmd.Flags().String("limit", "", "Maximum number of options to pick")
	filterCmd.Flags().String("match.align", "", "Text Alignment")
	filterCmd.Flags().String("match.background", "", "Background Color")
	filterCmd.Flags().Bool("match.bold", false, "Bold text")
	filterCmd.Flags().String("match.border", "", "Border Style")
	filterCmd.Flags().String("match.border-background", "", "Border Background Color")
	filterCmd.Flags().String("match.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("match.faint", false, "Faint text")
	filterCmd.Flags().String("match.foreground", "", "Foreground Color")
	filterCmd.Flags().String("match.height", "", "Text height")
	filterCmd.Flags().Bool("match.italic", false, "Italicize text")
	filterCmd.Flags().String("match.margin", "", "Text margin")
	filterCmd.Flags().String("match.padding", "", "Text padding")
	filterCmd.Flags().Bool("match.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("match.underline", false, "Underline text")
	filterCmd.Flags().String("match.width", "", "Text width")
	filterCmd.Flags().Bool("no-limit", false, "Pick unlimited number of options (ignores limit)")
	filterCmd.Flags().String("output-delimiter", "", "Option delimiter when writing to STDOUT")
	filterCmd.Flags().String("placeholder", "", "Placeholder value")
	filterCmd.Flags().String("placeholder.align", "", "Text Alignment")
	filterCmd.Flags().String("placeholder.background", "", "Background Color")
	filterCmd.Flags().Bool("placeholder.bold", false, "Bold text")
	filterCmd.Flags().String("placeholder.border", "", "Border Style")
	filterCmd.Flags().String("placeholder.border-background", "", "Border Background Color")
	filterCmd.Flags().String("placeholder.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("placeholder.faint", false, "Faint text")
	filterCmd.Flags().String("placeholder.foreground", "", "Foreground Color")
	filterCmd.Flags().String("placeholder.height", "", "Text height")
	filterCmd.Flags().Bool("placeholder.italic", false, "Italicize text")
	filterCmd.Flags().String("placeholder.margin", "", "Text margin")
	filterCmd.Flags().String("placeholder.padding", "", "Text padding")
	filterCmd.Flags().Bool("placeholder.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("placeholder.underline", false, "Underline text")
	filterCmd.Flags().String("placeholder.width", "", "Text width")
	filterCmd.Flags().String("prompt", "", "Prompt to display")
	filterCmd.Flags().String("prompt.align", "", "Text Alignment")
	filterCmd.Flags().String("prompt.background", "", "Background Color")
	filterCmd.Flags().Bool("prompt.bold", false, "Bold text")
	filterCmd.Flags().String("prompt.border", "", "Border Style")
	filterCmd.Flags().String("prompt.border-background", "", "Border Background Color")
	filterCmd.Flags().String("prompt.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("prompt.faint", false, "Faint text")
	filterCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	filterCmd.Flags().String("prompt.height", "", "Text height")
	filterCmd.Flags().Bool("prompt.italic", false, "Italicize text")
	filterCmd.Flags().String("prompt.margin", "", "Text margin")
	filterCmd.Flags().String("prompt.padding", "", "Text padding")
	filterCmd.Flags().Bool("prompt.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("prompt.underline", false, "Underline text")
	filterCmd.Flags().String("prompt.width", "", "Text width")
	filterCmd.Flags().Bool("reverse", false, "Display from the bottom of the screen")
	filterCmd.Flags().Bool("select-if-one", false, "Select the given option if there is only one")
	filterCmd.Flags().StringSlice("selected", []string{}, "Options that should start as selected (selects all if given '*')")
	filterCmd.Flags().String("selected-indicator.align", "", "Text Alignment")
	filterCmd.Flags().String("selected-indicator.background", "", "Background Color")
	filterCmd.Flags().Bool("selected-indicator.bold", false, "Bold text")
	filterCmd.Flags().String("selected-indicator.border", "", "Border Style")
	filterCmd.Flags().String("selected-indicator.border-background", "", "Border Background Color")
	filterCmd.Flags().String("selected-indicator.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("selected-indicator.faint", false, "Faint text")
	filterCmd.Flags().String("selected-indicator.foreground", "", "Foreground Color")
	filterCmd.Flags().String("selected-indicator.height", "", "Text height")
	filterCmd.Flags().Bool("selected-indicator.italic", false, "Italicize text")
	filterCmd.Flags().String("selected-indicator.margin", "", "Text margin")
	filterCmd.Flags().String("selected-indicator.padding", "", "Text padding")
	filterCmd.Flags().Bool("selected-indicator.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("selected-indicator.underline", false, "Underline text")
	filterCmd.Flags().String("selected-indicator.width", "", "Text width")
	filterCmd.Flags().String("selected-prefix", "", "Character to indicate selected items (hidden if limit is 1)")
	filterCmd.Flags().Bool("show-help", false, "Show help keybinds")
	filterCmd.Flags().Bool("sort", false, "Sort fuzzy results by their scores")
	filterCmd.Flags().Bool("strict", false, "Only returns if anything matched. Otherwise return Filter")
	filterCmd.Flags().Bool("strip-ansi", false, "Strip ANSI sequences when reading from STDIN")
	filterCmd.Flags().String("text.align", "", "Text Alignment")
	filterCmd.Flags().String("text.background", "", "Background Color")
	filterCmd.Flags().Bool("text.bold", false, "Bold text")
	filterCmd.Flags().String("text.border", "", "Border Style")
	filterCmd.Flags().String("text.border-background", "", "Border Background Color")
	filterCmd.Flags().String("text.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("text.faint", false, "Faint text")
	filterCmd.Flags().String("text.foreground", "", "Foreground Color")
	filterCmd.Flags().String("text.height", "", "Text height")
	filterCmd.Flags().Bool("text.italic", false, "Italicize text")
	filterCmd.Flags().String("text.margin", "", "Text margin")
	filterCmd.Flags().String("text.padding", "", "Text padding")
	filterCmd.Flags().Bool("text.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("text.underline", false, "Underline text")
	filterCmd.Flags().String("text.width", "", "Text width")
	filterCmd.Flags().String("timeout", "", "Timeout until filter command aborts")
	filterCmd.Flags().String("unselected-prefix", "", "Character to indicate unselected items (hidden if limit is 1)")
	filterCmd.Flags().String("unselected-prefix.align", "", "Text Alignment")
	filterCmd.Flags().String("unselected-prefix.background", "", "Background Color")
	filterCmd.Flags().Bool("unselected-prefix.bold", false, "Bold text")
	filterCmd.Flags().String("unselected-prefix.border", "", "Border Style")
	filterCmd.Flags().String("unselected-prefix.border-background", "", "Border Background Color")
	filterCmd.Flags().String("unselected-prefix.border-foreground", "", "Border Foreground Color")
	filterCmd.Flags().Bool("unselected-prefix.faint", false, "Faint text")
	filterCmd.Flags().String("unselected-prefix.foreground", "", "Foreground Color")
	filterCmd.Flags().String("unselected-prefix.height", "", "Text height")
	filterCmd.Flags().Bool("unselected-prefix.italic", false, "Italicize text")
	filterCmd.Flags().String("unselected-prefix.margin", "", "Text margin")
	filterCmd.Flags().String("unselected-prefix.padding", "", "Text padding")
	filterCmd.Flags().Bool("unselected-prefix.strikethrough", false, "Strikethrough text")
	filterCmd.Flags().Bool("unselected-prefix.underline", false, "Underline text")
	filterCmd.Flags().String("unselected-prefix.width", "", "Text width")
	filterCmd.Flags().String("value", "", "Initial filter value")
	filterCmd.Flags().String("width", "", "Input width")
	rootCmd.AddCommand(filterCmd)

	carapace.Gen(filterCmd).FlagCompletion(carapace.ActionMap{
		"cursor-text.align":                    gum.ActionAlignments(),
		"cursor-text.background":               gum.ActionColors(),
		"cursor-text.border":                   gum.ActionBorders(),
		"cursor-text.border-background":        gum.ActionColors(),
		"cursor-text.border-foreground":        gum.ActionColors(),
		"cursor-text.foreground":               gum.ActionColors(),
		"header.align":                         gum.ActionAlignments(),
		"header.background":                    gum.ActionColors(),
		"header.border":                        gum.ActionBorders(),
		"header.border-background":             gum.ActionColors(),
		"header.border-foreground":             gum.ActionColors(),
		"header.foreground":                    gum.ActionColors(),
		"indicator.align":                      gum.ActionAlignments(),
		"indicator.background":                 gum.ActionColors(),
		"indicator.border":                     gum.ActionBorders(),
		"indicator.border-background":          gum.ActionColors(),
		"indicator.border-foreground":          gum.ActionColors(),
		"indicator.foreground":                 gum.ActionColors(),
		"match.align":                          gum.ActionAlignments(),
		"match.background":                     gum.ActionColors(),
		"match.border":                         gum.ActionBorders(),
		"match.border-background":              gum.ActionColors(),
		"match.border-foreground":              gum.ActionColors(),
		"match.foreground":                     gum.ActionColors(),
		"placeholder.align":                    gum.ActionAlignments(),
		"placeholder.background":               gum.ActionColors(),
		"placeholder.border":                   gum.ActionBorders(),
		"placeholder.border-background":        gum.ActionColors(),
		"placeholder.border-foreground":        gum.ActionColors(),
		"placeholder.foreground":               gum.ActionColors(),
		"prompt.align":                         gum.ActionAlignments(),
		"prompt.background":                    gum.ActionColors(),
		"prompt.border":                        gum.ActionBorders(),
		"prompt.border-background":             gum.ActionColors(),
		"prompt.border-foreground":             gum.ActionColors(),
		"prompt.foreground":                    gum.ActionColors(),
		"selected-indicator.align":             gum.ActionAlignments(),
		"selected-indicator.background":        gum.ActionColors(),
		"selected-indicator.border":            gum.ActionBorders(),
		"selected-indicator.border-background": gum.ActionColors(),
		"selected-indicator.border-foreground": gum.ActionColors(),
		"selected-indicator.foreground":        gum.ActionColors(),
		"text.align":                           gum.ActionAlignments(),
		"text.background":                      gum.ActionColors(),
		"text.border":                          gum.ActionBorders(),
		"text.border-background":               gum.ActionColors(),
		"text.border-foreground":               gum.ActionColors(),
		"text.foreground":                      gum.ActionColors(),
		"unselected-prefix.align":              gum.ActionAlignments(),
		"unselected-prefix.background":         gum.ActionColors(),
		"unselected-prefix.border":             gum.ActionBorders(),
		"unselected-prefix.border-background":  gum.ActionColors(),
		"unselected-prefix.border-foreground":  gum.ActionColors(),
		"unselected-prefix.foreground":         gum.ActionColors(),
	})
}
