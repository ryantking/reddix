package main

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
	"github.com/jzelinskie/geddit"
)

func valid(err error) bool {
	if err != nil {
		if err != gocui.ErrUnknownView {
			return false
		}
	}

	return true
}

func frontpage(g *gocui.Gui) error {
	w, h := g.Size()
	if v, err := g.SetView("nav", 0, 0, w-1, 2); err != gocui.ErrUnknownView {
		if err != nil {
			return err
		}

		v.Frame = false
		fmt.Fprintln(v, "q:Quit l:Login r:Subreddit u:User")
		v.BgColor = gocui.ColorWhite
		v.FgColor = gocui.ColorBlack
	}

	sesh := geddit.NewSession("default")
	subs, err := sesh.DefaultFrontpage(geddit.DefaultPopularity, geddit.ListingOptions{})
	if err != nil {
		return err
	}

	post_id := 0
	for i := 2; i < h-1; i += 4 {
		if v, err := g.SetView("up-"+strconv.Itoa(post_id), 0, i, 3, i+2); err != gocui.ErrUnknownView {
			if err != nil {
				return err
			}

			v.Frame = false
			fmt.Fprintln(v, "\u21E7\n")
			v.FgColor = gocui.ColorRed
		}

		if v, err := g.SetView("score-"+strconv.Itoa(post_id), 0, i+1, 3, i+3); err != gocui.ErrUnknownView {
			if err != nil {
				return err
			}
			v.Frame = false
			fmt.Fprintln(v, strconv.Itoa(subs[post_id].Score))
			v.FgColor = gocui.ColorWhite
		}

		if v, err := g.SetView("down-"+strconv.Itoa(post_id), 0, i+2, 3, i+4); err != gocui.ErrUnknownView {
			if err != nil {
				return err
			}

			v.Frame = false
			fmt.Fprintln(v, "\u21E9\n")
			v.FgColor = gocui.ColorBlue
		}
		if v, err := g.SetView("post-"+strconv.Itoa(post_id), 3, i, w-1, i+4); err != gocui.ErrUnknownView {
			if err != nil {
				return err
			}

			v.Frame = false
			fmt.Fprintln(v, "\033[1m"+subs[post_id].Title+"\033[0m")
			fmt.Fprintln(v, "by "+subs[post_id].Author)
			fmt.Fprintln(v, strconv.Itoa(subs[post_id].NumComments)+" Comments")
		}
		post_id += 1
	}

	return nil
}
