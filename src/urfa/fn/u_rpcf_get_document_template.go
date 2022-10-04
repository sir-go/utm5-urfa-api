package fn

func x4605(c conn, p Dict) Dict {
	c.Call(0x4605)

	putI(c, p, "template_id")
	c.Send()

	res := c.GetI()
	if res != 0 {
		return Dict{
			"result":   res,
			"doc_type": c.GetI(),
			"created":  c.GetI(),
			"modified": c.GetI(),
			"name":     c.GetS(),
			"path":     c.GetS(),
			"odt_data": c.GetBin(),
		}
	}

	return Dict{"result": res}
}
