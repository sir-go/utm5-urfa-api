package fn

// rpcf_get_document_profile_for_user
func x4610(c conn, p Dict) Dict {
	c.Call(0x4610)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"profile_id": c.GetI(),
	}
}
