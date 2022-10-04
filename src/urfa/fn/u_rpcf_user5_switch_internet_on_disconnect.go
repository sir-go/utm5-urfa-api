package fn

// rpcf_user5_switch_internet_on_disconnect
func xU4030(c conn, p Dict) Dict {
	c.Call(-0x4030)
	putI(c, p, "on")
	c.Send()

	return nil
}
