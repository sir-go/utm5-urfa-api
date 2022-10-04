package fn

func x4612(c conn, p Dict) Dict {
	c.Call(0x4612)

	putI(c, p, "contract_id")
	putI(c, p, "dont_donvert")
	c.Send()

	res := c.GetI()
	if res != 0 {
		return Dict{
			"result": res,
			"is_odt": c.GetI(),
			"data":   c.GetBin(),
		}
	}

	return Dict{"result": res}
}
