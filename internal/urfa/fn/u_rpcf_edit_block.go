package fn

// rpcf_edit_block
func x300c(c conn, p Dict) Dict {
	c.Call(0x300c)
	putI(c, p, "id")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	putI(c, p, "unabon")
	putI(c, p, "unprepay")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
