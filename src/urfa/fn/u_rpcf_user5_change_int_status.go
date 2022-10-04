package fn

// rpcf_user5_change_int_status
func xU4007(c conn, p Dict) Dict {
	c.Call(-0x4007)
	putI(c, p, "deprecated")
	c.Send()

	return nil
}
