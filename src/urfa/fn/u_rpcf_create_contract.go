package fn

// rpcf_create_contract
func x4613(c conn, p Dict) Dict {
	c.Call(0x4613)
	putI(c, p, "user_id")
	putI(c, p, "template_id")
	c.Send()

	return Dict{
		"contract_id": c.GetI(),
	}
}
