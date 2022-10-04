package fn

func xTmpl(c conn, p Dict) Dict {
	// c.Call(0x2812)
	forEachDict(c, p, "attrs", func(ai Dict) {
		putI(c, ai, "vendor")
		putI(c, ai, "attr")
		putI(c, ai, "tag")
		putI(c, ai, "usage_flags")
		putI(c, ai, "expire_date", TimeMax)

		putI(c, ai, "type")
		switch ai["type"].(float64) {
		case 0:
			putI(c, ai, "val")
		case 1:
			putS(c, ai, "val")
		case 2:
			putI(c, ai, "val")
		case 3:
			putS(c, ai, "val")
		}

		forEachDict(c, ai, "props", func(bi Dict) {
			putI(c, bi, "type")
			putS(c, bi, "value")
		})
	})
	c.Send()

	return Dict{
		"houses": getMapOf(c, func() Dict {
			return Dict{
				"house_id":     c.GetI(),
				"ip_zone_id":   c.GetI(),
				"connect_date": c.GetI(),
				"post_code":    c.GetS(),
				"country":      c.GetS(),
				"region":       c.GetS(),
				"city":         c.GetS(),
				"street":       c.GetS(),
				"number":       c.GetS(),
				"building":     c.GetS(),
			}
		}),
	}
}
// Dict{"error": Dict{"code": 13, "comment": "unable to delete service link"}}