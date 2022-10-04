package fn

// rpcf_delete_block
func x300d(c conn, p Dict) Dict {
	c.Call(0x300d)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
