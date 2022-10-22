package fn

// rpcf_delete_voluntary_suspension_settings
func x15013(c conn, p Dict) Dict {
	c.Call(0x15013)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
