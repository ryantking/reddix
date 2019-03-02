package main

import (
	"github.com/RyanTKing/reddix/internal/term"
	"github.com/RyanTKing/reddix/pkg/reddix"
	"github.com/RyanTKing/reddix/pkg/window"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "reddix [subreddit]",
	Short: "reddix is a command line utility for browsing reddit",
	Run: func(cmd *cobra.Command, args []string) {
		term, err := term.New()
		if err != nil {
			die(err)
		}
		reddix := newReddix(term)
		reddix.Run()
		term.Close()
	},
}

func newReddix(win window.Window) *reddix.Reddix {
	if viper.GetBool("anonymous") {
		return reddix.NewAnonymous(win)
	}
	username := viper.GetString("username")
	if username != "" {
		return reddix.NewWithUser(win, username)
	}

	return reddix.New(win)
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "the user to browse reddit as")
	rootCmd.PersistentFlags().BoolP("anonymous", "a", false, "don't attempt to log in in as a user")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("anonymous", rootCmd.PersistentFlags().Lookup("anonymous"))
	viper.BindEnv("username", "REDDIX_USERNAME")
}
