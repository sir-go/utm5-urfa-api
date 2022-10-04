package fn

// rpcf_del_time_range
func x10353(c conn, p Dict) Dict {
	c.Call(0x10353)
	putI(c, p, "time_range_id")
	c.Send()

	return nil
}
