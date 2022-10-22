package fn

// rpcf_remove_contract
func x4615(c conn, p Dict) Dict {
	c.Call(0x4615)
	putI(c, p, "contract_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
