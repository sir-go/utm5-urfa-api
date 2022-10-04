package fn

// rpcf_get_access_card_status
func x450a(c conn, p Dict) Dict {
	c.Call(0x450a)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"status": c.GetI(),
	}
}
