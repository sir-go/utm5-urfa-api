package fn

// rpcf_user5_get_message
func xU4042(c conn, p Dict) Dict {
	c.Call(-0x4042)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"subject":   c.GetS(),
		"message":   c.GetS(),
		"mime":      c.GetS(),
		"send_date": c.GetI(),
		"sender_id": c.GetI(),
	}
}
