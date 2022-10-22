package fn

// rpcf_delete_isg_profile
func x1404(c conn, p Dict) Dict {
	c.Call(0x1404)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
