package fn

// rpcf_new_activation_code
func x450c(c conn, p Dict) Dict {
	c.Call(0x450c)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
