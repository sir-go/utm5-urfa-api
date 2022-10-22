package fn

// rpcf_get_account_external_id
func x2039(c conn, p Dict) Dict {
	c.Call(0x2039)
	putI(c, p, "aid")
	c.Send()

	return Dict{
		"external_id": c.GetS(),
	}
}
