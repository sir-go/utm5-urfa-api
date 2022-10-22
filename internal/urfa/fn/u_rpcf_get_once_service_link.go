package fn

// rpcf_get_once_service_link
func x2704(c conn, p Dict) Dict {
	c.Call(0x2704)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"discount_date": c.GetI(),
	}
}
