package fn

// rpcf_get_house
func x2812(c conn, p Dict) Dict {
	c.Call(0x2812)
	putI(c, p, "house_id")
	c.Send()

	return Dict{
		"house_id":     c.GetI(),
		"connect_date": c.GetI(),
		"post_code":    c.GetS(),
		"country":      c.GetS(),
		"region":       c.GetS(),
		"city":         c.GetS(),
		"street":       c.GetS(),
		"number":       c.GetS(),
		"building":     c.GetS(),
		"ip_zones": getMapOf(c, func() Dict {
			return Dict{
				"ipzone_id":   c.GetI(),
				"ipzone_name": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2812" name="rpcf_get_house">
    <input>
      <integer name="house_id" />
    </input>
    <output>
      <integer name="house_id" />
      <integer name="connect_date" />
      <string name="post_code" />
      <string name="country" />
      <string name="region" />
      <string name="city" />
      <string name="street" />
      <string name="number" />
      <string name="building" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="ipzone_id" />
        <string array_index="i" name="ipzone_name" />
      </for>
    </output>
  </function>


*/
