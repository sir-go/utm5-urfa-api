package fn

// rpcf_upload_contract
func x4614(c conn, p Dict) Dict {
	c.Call(0x4614)
	putI(c, p, "user_id")
	putS(c, p, "name")
	c.Send()

	sendBin(c, p, "data")

	return Dict{
		"contract_id": c.GetI(),
	}
}
