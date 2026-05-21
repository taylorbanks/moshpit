package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// Theme defines all color slots used throughout the application UI.
type Theme struct {
	Name string

	// Base backgrounds
	Base     tcell.Color // main background
	Mantle   tcell.Color // header/status bar bg
	Crust    tcell.Color // deepest bg
	Surface0 tcell.Color // contrast bg / input fields
	Surface1 tcell.Color // borders
	Surface2 tcell.Color // inactive elements

	// Text
	Text     tcell.Color // primary text
	Subtext1 tcell.Color // secondary text
	Subtext0 tcell.Color // tertiary/muted
	Overlay0 tcell.Color // very muted text

	// Accents
	Blue     tcell.Color // selection highlight, links
	Green    tcell.Color // success, active status
	Red      tcell.Color // error, unavailable
	Yellow   tcell.Color // section headers, warnings
	Mauve    tcell.Color // commit tag
	Sapphire tcell.Color // tag chips bg
	Teal     tcell.Color // branding accent
	Peach    tcell.Color // warning fallback
	Lavender tcell.Color // secondary accent
}

// ActiveTheme is the current theme used by all UI components.
var ActiveTheme = DefaultTheme()

// Themes is the registry of all available themes.
var Themes = []Theme{
	DefaultTheme(),
	CatppuccinMocha(),
	CatppuccinLatte(),
	SolarizedDark(),
	SolarizedLight(),
	Dracula(),
	TokyoNight(),
	Nord(),
	GruvboxDark(),
	GruvboxLight(),
	Ristretto(),
	RosePine(),
}

// ThemeByName returns the theme with the given name, or DefaultTheme if not found.
func ThemeByName(name string) Theme {
	for _, t := range Themes {
		if t.Name == name {
			return t
		}
	}
	return DefaultTheme()
}

// SetActiveTheme sets the active theme by name and re-applies tview globals.
func SetActiveTheme(name string) {
	ActiveTheme = ThemeByName(name)
}

// DefaultTheme returns a bland neutral dark theme.
func DefaultTheme() Theme {
	return Theme{
		Name:     "Default",
		Base:     tcell.NewHexColor(0x1c1c1c),
		Mantle:   tcell.NewHexColor(0x171717),
		Crust:    tcell.NewHexColor(0x121212),
		Surface0: tcell.NewHexColor(0x2a2a2a),
		Surface1: tcell.NewHexColor(0x3a3a3a),
		Surface2: tcell.NewHexColor(0x4a4a4a),

		Text:     tcell.NewHexColor(0xd4d4d4),
		Subtext1: tcell.NewHexColor(0xb0b0b0),
		Subtext0: tcell.NewHexColor(0x909090),
		Overlay0: tcell.NewHexColor(0x606060),

		Blue:     tcell.NewHexColor(0x6da2d4),
		Green:    tcell.NewHexColor(0x7dae7d),
		Red:      tcell.NewHexColor(0xd47d7d),
		Yellow:   tcell.NewHexColor(0xd4c47d),
		Mauve:    tcell.NewHexColor(0xa67dba),
		Sapphire: tcell.NewHexColor(0x6da2b8),
		Teal:     tcell.NewHexColor(0x7db8a6),
		Peach:    tcell.NewHexColor(0xd4a67d),
		Lavender: tcell.NewHexColor(0x9d9dd4),
	}
}

// CatppuccinMocha returns the Catppuccin Mocha color theme.
func CatppuccinMocha() Theme {
	return Theme{
		Name:     "Catppuccin Mocha",
		Base:     tcell.NewHexColor(0x1e1e2e),
		Mantle:   tcell.NewHexColor(0x181825),
		Crust:    tcell.NewHexColor(0x11111b),
		Surface0: tcell.NewHexColor(0x313244),
		Surface1: tcell.NewHexColor(0x45475a),
		Surface2: tcell.NewHexColor(0x585b70),

		Text:     tcell.NewHexColor(0xcdd6f4),
		Subtext1: tcell.NewHexColor(0xbac2de),
		Subtext0: tcell.NewHexColor(0xa6adc8),
		Overlay0: tcell.NewHexColor(0x6c7086),

		Blue:     tcell.NewHexColor(0x89b4fa),
		Green:    tcell.NewHexColor(0xa6e3a1),
		Red:      tcell.NewHexColor(0xf38ba8),
		Yellow:   tcell.NewHexColor(0xf9e2af),
		Mauve:    tcell.NewHexColor(0xcba6f7),
		Sapphire: tcell.NewHexColor(0x74c7ec),
		Teal:     tcell.NewHexColor(0x94e2d5),
		Peach:    tcell.NewHexColor(0xfab387),
		Lavender: tcell.NewHexColor(0xb4befe),
	}
}

// CatppuccinLatte returns the Catppuccin Latte light theme.
func CatppuccinLatte() Theme {
	return Theme{
		Name:     "Catppuccin Latte",
		Base:     tcell.NewHexColor(0xeff1f5),
		Mantle:   tcell.NewHexColor(0xe6e9ef),
		Crust:    tcell.NewHexColor(0xdce0e8),
		Surface0: tcell.NewHexColor(0xccd0da),
		Surface1: tcell.NewHexColor(0xbcc0cc),
		Surface2: tcell.NewHexColor(0xacb0be),

		Text:     tcell.NewHexColor(0x4c4f69),
		Subtext1: tcell.NewHexColor(0x5c5f77),
		Subtext0: tcell.NewHexColor(0x6c6f85),
		Overlay0: tcell.NewHexColor(0x9ca0b0),

		Blue:     tcell.NewHexColor(0x1e66f5),
		Green:    tcell.NewHexColor(0x40a02b),
		Red:      tcell.NewHexColor(0xd20f39),
		Yellow:   tcell.NewHexColor(0xdf8e1d),
		Mauve:    tcell.NewHexColor(0x8839ef),
		Sapphire: tcell.NewHexColor(0x209fb5),
		Teal:     tcell.NewHexColor(0x179299),
		Peach:    tcell.NewHexColor(0xfe640b),
		Lavender: tcell.NewHexColor(0x7287fd),
	}
}

// SolarizedDark returns the Solarized Dark color theme.
// Based on Ethan Schoonover's official Solarized palette.
// Dark uses base03/base02 as bg, base0/base1 as fg.
func SolarizedDark() Theme {
	return Theme{
		Name:     "Solarized Dark",
		Base:     tcell.NewHexColor(0x002b36), // base03
		Mantle:   tcell.NewHexColor(0x00212b), // derived darker
		Crust:    tcell.NewHexColor(0x001920), // derived darkest
		Surface0: tcell.NewHexColor(0x073642), // base02
		Surface1: tcell.NewHexColor(0x094352), // between base02 and base01
		Surface2: tcell.NewHexColor(0x586e75), // base01

		Text:     tcell.NewHexColor(0x839496), // base0
		Subtext1: tcell.NewHexColor(0x93a1a1), // base1
		Subtext0: tcell.NewHexColor(0x657b83), // base00
		Overlay0: tcell.NewHexColor(0x586e75), // base01

		Blue:     tcell.NewHexColor(0x268bd2), // blue
		Green:    tcell.NewHexColor(0x859900), // green
		Red:      tcell.NewHexColor(0xdc322f), // red
		Yellow:   tcell.NewHexColor(0xb58900), // yellow
		Mauve:    tcell.NewHexColor(0x6c71c4), // violet
		Sapphire: tcell.NewHexColor(0x2aa198), // cyan
		Teal:     tcell.NewHexColor(0x2aa198), // cyan
		Peach:    tcell.NewHexColor(0xcb4b16), // orange
		Lavender: tcell.NewHexColor(0xd33682), // magenta
	}
}

// SolarizedLight returns the Solarized Light color theme.
// Based on Ethan Schoonover's official Solarized palette.
// Light uses base3/base2 as bg, base00/base01 as fg.
func SolarizedLight() Theme {
	return Theme{
		Name:     "Solarized Light",
		Base:     tcell.NewHexColor(0xfdf6e3), // base3
		Mantle:   tcell.NewHexColor(0xeee8d5), // base2
		Crust:    tcell.NewHexColor(0xe4ddca), // derived slightly darker
		Surface0: tcell.NewHexColor(0xe4ddc8), // between base2 and base1
		Surface1: tcell.NewHexColor(0xd3cdb9), // derived mid
		Surface2: tcell.NewHexColor(0x93a1a1), // base1

		Text:     tcell.NewHexColor(0x657b83), // base00
		Subtext1: tcell.NewHexColor(0x586e75), // base01
		Subtext0: tcell.NewHexColor(0x73878d), // between base00 and base0
		Overlay0: tcell.NewHexColor(0x93a1a1), // base1

		Blue:     tcell.NewHexColor(0x268bd2), // blue
		Green:    tcell.NewHexColor(0x859900), // green
		Red:      tcell.NewHexColor(0xdc322f), // red
		Yellow:   tcell.NewHexColor(0xb58900), // yellow
		Mauve:    tcell.NewHexColor(0x6c71c4), // violet
		Sapphire: tcell.NewHexColor(0x2aa198), // cyan
		Teal:     tcell.NewHexColor(0x2aa198), // cyan
		Peach:    tcell.NewHexColor(0xcb4b16), // orange
		Lavender: tcell.NewHexColor(0xd33682), // magenta
	}
}

// Dracula returns the Dracula color theme.
func Dracula() Theme {
	return Theme{
		Name:     "Dracula",
		Base:     tcell.NewHexColor(0x282a36),
		Mantle:   tcell.NewHexColor(0x21222c),
		Crust:    tcell.NewHexColor(0x191a21),
		Surface0: tcell.NewHexColor(0x44475a),
		Surface1: tcell.NewHexColor(0x565868),
		Surface2: tcell.NewHexColor(0x6272a4),

		Text:     tcell.NewHexColor(0xf8f8f2),
		Subtext1: tcell.NewHexColor(0xe0e0da),
		Subtext0: tcell.NewHexColor(0xbfbfb9),
		Overlay0: tcell.NewHexColor(0x6272a4),

		Blue:     tcell.NewHexColor(0x8be9fd),
		Green:    tcell.NewHexColor(0x50fa7b),
		Red:      tcell.NewHexColor(0xff5555),
		Yellow:   tcell.NewHexColor(0xf1fa8c),
		Mauve:    tcell.NewHexColor(0xbd93f9),
		Sapphire: tcell.NewHexColor(0x8be9fd),
		Teal:     tcell.NewHexColor(0x8be9fd),
		Peach:    tcell.NewHexColor(0xffb86c),
		Lavender: tcell.NewHexColor(0xff79c6),
	}
}

// TokyoNight returns the Tokyo Night color theme.
// Based on the official tokyo-night-vscode-theme.
func TokyoNight() Theme {
	return Theme{
		Name:     "Tokyo Night",
		Base:     tcell.NewHexColor(0x1a1b26), // bg
		Mantle:   tcell.NewHexColor(0x16161e), // bg_dark
		Crust:    tcell.NewHexColor(0x101014), // derived darkest
		Surface0: tcell.NewHexColor(0x292e42), // bg_highlight
		Surface1: tcell.NewHexColor(0x3b4261), // fg_gutter
		Surface2: tcell.NewHexColor(0x545c7e), // dark3

		Text:     tcell.NewHexColor(0xc0caf5), // fg
		Subtext1: tcell.NewHexColor(0xa9b1d6), // fg_dark
		Subtext0: tcell.NewHexColor(0x9aa5ce), // fg_dark variant
		Overlay0: tcell.NewHexColor(0x565f89), // comment

		Blue:     tcell.NewHexColor(0x7aa2f7), // blue
		Green:    tcell.NewHexColor(0x9ece6a), // green
		Red:      tcell.NewHexColor(0xf7768e), // red
		Yellow:   tcell.NewHexColor(0xe0af68), // yellow
		Mauve:    tcell.NewHexColor(0xbb9af7), // magenta
		Sapphire: tcell.NewHexColor(0x7dcfff), // cyan
		Teal:     tcell.NewHexColor(0x73daca), // teal
		Peach:    tcell.NewHexColor(0xff9e64), // orange
		Lavender: tcell.NewHexColor(0xb4f9f8), // terminal cyan bright
	}
}

// Nord returns the Nord color theme.
func Nord() Theme {
	return Theme{
		Name:     "Nord",
		Base:     tcell.NewHexColor(0x2e3440),
		Mantle:   tcell.NewHexColor(0x272c36),
		Crust:    tcell.NewHexColor(0x20242d),
		Surface0: tcell.NewHexColor(0x3b4252),
		Surface1: tcell.NewHexColor(0x434c5e),
		Surface2: tcell.NewHexColor(0x4c566a),

		Text:     tcell.NewHexColor(0xeceff4),
		Subtext1: tcell.NewHexColor(0xe5e9f0),
		Subtext0: tcell.NewHexColor(0xd8dee9),
		Overlay0: tcell.NewHexColor(0x4c566a),

		Blue:     tcell.NewHexColor(0x81a1c1),
		Green:    tcell.NewHexColor(0xa3be8c),
		Red:      tcell.NewHexColor(0xbf616a),
		Yellow:   tcell.NewHexColor(0xebcb8b),
		Mauve:    tcell.NewHexColor(0xb48ead),
		Sapphire: tcell.NewHexColor(0x88c0d0),
		Teal:     tcell.NewHexColor(0x8fbcbb),
		Peach:    tcell.NewHexColor(0xd08770),
		Lavender: tcell.NewHexColor(0x5e81ac),
	}
}

// GruvboxDark returns the Gruvbox Dark color theme.
// Based on the official morhetz/gruvbox palette.
func GruvboxDark() Theme {
	return Theme{
		Name:     "Gruvbox Dark",
		Base:     tcell.NewHexColor(0x282828), // bg0
		Mantle:   tcell.NewHexColor(0x1d2021), // bg0_hard
		Crust:    tcell.NewHexColor(0x141617), // derived darkest
		Surface0: tcell.NewHexColor(0x3c3836), // bg1
		Surface1: tcell.NewHexColor(0x504945), // bg2
		Surface2: tcell.NewHexColor(0x665c54), // bg3

		Text:     tcell.NewHexColor(0xebdbb2), // fg1
		Subtext1: tcell.NewHexColor(0xd5c4a1), // fg2
		Subtext0: tcell.NewHexColor(0xbdae93), // fg3
		Overlay0: tcell.NewHexColor(0x928374), // gray

		Blue:     tcell.NewHexColor(0x83a598), // bright_blue
		Green:    tcell.NewHexColor(0xb8bb26), // bright_green
		Red:      tcell.NewHexColor(0xfb4934), // bright_red
		Yellow:   tcell.NewHexColor(0xfabd2f), // bright_yellow
		Mauve:    tcell.NewHexColor(0xd3869b), // bright_purple
		Sapphire: tcell.NewHexColor(0x83a598), // bright_blue
		Teal:     tcell.NewHexColor(0x8ec07c), // bright_aqua
		Peach:    tcell.NewHexColor(0xfe8019), // bright_orange
		Lavender: tcell.NewHexColor(0xb16286), // neutral_purple
	}
}

// GruvboxLight returns the Gruvbox Light color theme.
func GruvboxLight() Theme {
	return Theme{
		Name:     "Gruvbox Light",
		Base:     tcell.NewHexColor(0xfbf1c7),
		Mantle:   tcell.NewHexColor(0xf2e5bc),
		Crust:    tcell.NewHexColor(0xe8d8ad),
		Surface0: tcell.NewHexColor(0xebdbb2),
		Surface1: tcell.NewHexColor(0xd5c4a1),
		Surface2: tcell.NewHexColor(0xbdae93),

		Text:     tcell.NewHexColor(0x3c3836),
		Subtext1: tcell.NewHexColor(0x504945),
		Subtext0: tcell.NewHexColor(0x665c54),
		Overlay0: tcell.NewHexColor(0x928374),

		Blue:     tcell.NewHexColor(0x076678),
		Green:    tcell.NewHexColor(0x79740e),
		Red:      tcell.NewHexColor(0x9d0006),
		Yellow:   tcell.NewHexColor(0xb57614),
		Mauve:    tcell.NewHexColor(0x8f3f71),
		Sapphire: tcell.NewHexColor(0x076678),
		Teal:     tcell.NewHexColor(0x427b58),
		Peach:    tcell.NewHexColor(0xaf3a03),
		Lavender: tcell.NewHexColor(0x8f3f71),
	}
}

// Ristretto returns the Ristretto (coffee-inspired warm dark) color theme.
func Ristretto() Theme {
	return Theme{
		Name:     "Ristretto",
		Base:     tcell.NewHexColor(0x2a211c), // dark espresso
		Mantle:   tcell.NewHexColor(0x221a16), // darker roast
		Crust:    tcell.NewHexColor(0x1a1310), // darkest bean
		Surface0: tcell.NewHexColor(0x3b302a), // medium roast
		Surface1: tcell.NewHexColor(0x4d4038), // light roast
		Surface2: tcell.NewHexColor(0x5f5048), // tan

		Text:     tcell.NewHexColor(0xfff8e7), // cream
		Subtext1: tcell.NewHexColor(0xe8dfd0), // steamed milk
		Subtext0: tcell.NewHexColor(0xc8bfb0), // latte foam
		Overlay0: tcell.NewHexColor(0x78685c), // cocoa

		Blue:     tcell.NewHexColor(0x6db5c4), // iced coffee
		Green:    tcell.NewHexColor(0xa8c97d), // matcha
		Red:      tcell.NewHexColor(0xd47070), // cherry
		Yellow:   tcell.NewHexColor(0xd4a76a), // caramel
		Mauve:    tcell.NewHexColor(0xb08fad), // lavender latte
		Sapphire: tcell.NewHexColor(0x6db5c4), // iced
		Teal:     tcell.NewHexColor(0x87b9a5), // mint
		Peach:    tcell.NewHexColor(0xd49060), // cinnamon
		Lavender: tcell.NewHexColor(0xc4a0b8), // berry
	}
}

// RosePine returns the Rosé Pine color theme.
// Based on the official rosepinetheme.com palette (main variant).
func RosePine() Theme {
	return Theme{
		Name:     "Rosé Pine",
		Base:     tcell.NewHexColor(0x191724), // base
		Mantle:   tcell.NewHexColor(0x14121f), // derived darker than base
		Crust:    tcell.NewHexColor(0x0f0d19), // derived darkest
		Surface0: tcell.NewHexColor(0x1f1d2e), // surface
		Surface1: tcell.NewHexColor(0x26233a), // overlay
		Surface2: tcell.NewHexColor(0x403d52), // highlight-med

		Text:     tcell.NewHexColor(0xe0def4), // text
		Subtext1: tcell.NewHexColor(0x908caa), // subtle
		Subtext0: tcell.NewHexColor(0x817da0), // between subtle and muted
		Overlay0: tcell.NewHexColor(0x6e6a86), // muted

		Blue:     tcell.NewHexColor(0x31748f), // pine
		Green:    tcell.NewHexColor(0x9ccfd8), // foam
		Red:      tcell.NewHexColor(0xeb6f92), // love
		Yellow:   tcell.NewHexColor(0xf6c177), // gold
		Mauve:    tcell.NewHexColor(0xc4a7e7), // iris
		Sapphire: tcell.NewHexColor(0x31748f), // pine
		Teal:     tcell.NewHexColor(0x9ccfd8), // foam
		Peach:    tcell.NewHexColor(0xebbcba), // rose
		Lavender: tcell.NewHexColor(0xc4a7e7), // iris
	}
}

// Hex returns the tview-compatible hex color string for a tcell.Color.
func Hex(c tcell.Color) string {
	r, g, b := c.RGB()
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
