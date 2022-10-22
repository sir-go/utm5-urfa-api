package fn

// rpcf_is_account_external_id_used
func x203a(c conn, p Dict) Dict {
	c.Call(0x203a)
	putS(c, p, "external_id")
	c.Send()

	return Dict{
		"aid": c.GetI(),
	}
}
