package fn

// rpcf_user5_get_account_external_id
func xU15030(c conn, p Dict) Dict {
	c.Call(-0x15030)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"external_id": c.GetS(),
	}
}
