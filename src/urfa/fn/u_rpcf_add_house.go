package fn

// rpcf_add_house
func x2811(c conn, p Dict) Dict {
	c.Call(0x2811)
	putI(c, p, "house_id", 0)
	putI(c, p, "connect_date", TimeNow())
	putS(c, p, "post_code", "")
	putS(c, p, "country", "")
	putS(c, p, "region", "")
	putS(c, p, "city", "")
	putS(c, p, "street", "")
	putS(c, p, "number", "")
	putS(c, p, "building", "")

	forEachDict(c, p, "ip_zones", func(ai Dict) {
		putI(c, ai, "ipzone_id")
	})
	c.Send()

	return nil
}

/* xml:
<function id="0x2811" name="rpcf_add_house">
    <input>
      <integer name="house_id" />
      <integer name="connect_date" />
      <string name="post_code" />
      <string name="country" />
      <string name="region" />
      <string name="city" />
      <string name="street" />
      <string name="number" />
      <string name="building" />
      <integer default="size(ipzone_id)" name="count" />
      <for count="size(ipzone_id)" from="0" name="i">
        <integer array_index="i" name="ipzone_id" />
      </for>
    </input>
    <output />
  </function>


*/
