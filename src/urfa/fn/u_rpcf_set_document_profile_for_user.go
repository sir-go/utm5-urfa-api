package fn

// rpcf_set_document_profile_for_user
func x460f(c conn, p Dict) Dict {
	c.Call(0x460f)
	putI(c, p, "user_id")
	putI(c, p, "profile_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
