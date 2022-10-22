package fn

// rpcf_set_account_external_id
func x2038(c conn, p Dict) Dict {
	c.Call(0x2038)
	putI(c, p, "aid")
	putS(c, p, "external_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
