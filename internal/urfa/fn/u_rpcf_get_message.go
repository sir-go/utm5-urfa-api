package fn

// rpcf_get_message
func x500c(c conn, p Dict) Dict {
	c.Call(0x500c)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"subject":       c.GetS(),
		"message":       c.GetS(),
		"mime":          c.GetS(),
		"send_date":     c.GetI(),
		"sender_id":     c.GetI(),
		"receiver_id":   c.GetI(),
		"receiver_type": c.GetI(),
	}
}
