package internal

func rootFlag(c *Command) *string {
	return c.String("root", rootDir(), "location of go")
}

func nFlag(c *Command) *int {
	return c.Int("n", 3, "")
}

func verboseFlag(c *Command) {
	c.verbose = c.Bool("v", false, "see more logs")
}
