package window

func (win *Window) login() {
	loggedIn, err := win.Sess.Login()
	if err != nil {
		win.setStatus(err.Error(), true)
		return
	}
	if !loggedIn {
		return
	}

	win.setStatus("login succeeded", false)
	win.TopMenu.Right = topMenuRight2
	win.drawTopMenu()
	win.BottomMenu.Right = win.Sess.Username
	win.drawBottomMenu()

	if win.subreddit == "" {
		win.refreshPosts()
		win.drawPosts()
	}
}

func (win *Window) logout() {
	err := win.Sess.Logout()
	if err != nil {
		win.setStatus(err.Error(), true)
	}

	win.setStatus("logout succeeded", false)
	win.TopMenu.Right = topMenuRight1
	win.drawTopMenu()
	win.BottomMenu.Right = bottomMenuRight
	win.drawBottomMenu()
}
