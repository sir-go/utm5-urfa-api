package fn

// rpcf_dealer_delete_user
func x70000003(c conn, p Dict) Dict {
	c.Call(0x70000003)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
