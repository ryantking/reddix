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

	Text: TextTheme{
		Text:  NewStyle(ColorWhite),
		Error: NewStyle(ColorRed),
	},

	Post: PostTheme{
		Title:    NewStyle(ColorWhite, ColorBlack, ModifierBold),
		Subtitle: NewStyle(ColorWhite),
		UpVote:   NewStyle(ColorRed),
		DownVote: NewStyle(ColorBlue),
		Score:    NewStyle(ColorWhite),
		Link:     NewStyle(ColorWhite, ColorBlack, ModifierUnderline),
	},

	PostNumber: PostNumberTheme{
		Num: NewStyle(ColorWhite, ColorBlack, ModifierBold),
	},
}
