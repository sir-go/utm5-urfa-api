package fn

// rpcf_get_houses_list
func x2810(c conn, _ Dict) Dict {
	c.Call(0x2810)

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

/* xml:
<function id="0x2810" name="rpcf_get_houses_list">
    <input />
    <output>
      <integer name="houses_size" />
      <for count="houses_size" from="0" name="i">
        <integer array_index="i" name="house_id" />
        <integer array_index="i" name="ip_zone_id" />
        <integer array_index="i" name="connect_date" />
        <string array_index="i" name="post_code" />
        <string array_index="i" name="country" />
        <string array_index="i" name="region" />
        <string array_index="i" name="city" />
        <string array_index="i" name="street" />
        <string array_index="i" name="number" />
        <string array_index="i" name="building" />
      </for>
    </output>
  </function>


*/
