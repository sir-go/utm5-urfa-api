package fn

// rpcf_delete_supplier
func x8017(c conn, p Dict) Dict {
	c.Call(0x8017)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
