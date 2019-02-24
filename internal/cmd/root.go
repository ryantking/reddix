package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/RyanTKing/reddix/internal/window"

	termbox "github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "reddix [subreddit]",
	Short: "reddix is a command line utility for browsing reddit",
	Run: func(cmd *cobra.Command, args []string) {
		win, err := window.New()
		if err != nil {
			die(err)
		}

		defer func() {
			if r := recover(); r != nil {
				switch rval := r.(type) {
				case error:
					die(rval)
				default:
					die(errors.New(fmt.Sprint(rval)))
				}
			}
		}()

		win.Run()
		win.Close()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "the user to browse reddit as")
	rootCmd.PersistentFlags().BoolP("anonymous", "a", false, "don't attempt to log in in as a user")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("anonymous", rootCmd.PersistentFlags().Lookup("anonymous"))
	viper.BindEnv("username", "REDDIX_USERNAME")
}

// Execute runs the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		die(err)
	}
}

func die(err error) {
	termbox.Close()
	fmt.Println(err.Error())
	os.Exit(1)
}
