package ui

// Theme is the theme used throught the entire program
var Theme = &RootTheme{
	Default: NewStyle(ColorWhite),

	Block: BlockTheme{
		Title:  NewStyle(ColorWhite),
		Border: NewStyle(ColorWhite),
	},

	Menu: MenuTheme{
		Text: NewStyle(ColorBlack, ColorWhite),
	},

	TextEntry: TextEntryTheme{
		Text: NewStyle(ColorWhite),
	},

	Error: ErrorTheme{
		Text: NewStyle(ColorRed),
	},

	Posts: PostsTheme{
		Title:    NewStyle(ColorWhite, ColorBlack, ModifierBold),
		Subtitle: NewStyle(ColorWhite),
		UpVote:   NewStyle(ColorRed),
		DownVote: NewStyle(ColorBlue),
		Score:    NewStyle(ColorWhite),
		Link:     NewStyle(ColorWhite, ColorBlack, ModifierUnderline),
		Selected: NewStyle(ColorWhite, ColorBlack, ModifierBold|ModifierUnderline),
	},
}
