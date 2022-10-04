package fn

// rpcf_remove_document_profile
func x4609(c conn, p Dict) Dict {
	c.Call(0x4609)
	putI(c, p, "profile_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
