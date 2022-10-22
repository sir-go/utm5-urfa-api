package fn

// rpcf_delete_promised_payment_settings
func x15023(c conn, p Dict) Dict {
	c.Call(0x15023)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
