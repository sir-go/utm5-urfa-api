package fn

// rpcf_create_access_card
func x4509(c conn, p Dict) Dict {
	c.Call(0x4509)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
