package fn

// rpcf_get_streets_list
func x2816(c conn, _ Dict) Dict {
	c.Call(0x2816)

	return Dict{
		"streets": getMapOf(c, func() Dict {
			return Dict{
				"street_id": c.GetI(),
				"country":   c.GetS(),
				"region":    c.GetS(),
				"city":      c.GetS(),
				"street":    c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2816" name="rpcf_get_streets_list">
    <input />
    <output>
      <integer name="streets_size" />
      <for count="streets_size" from="0" name="i">
        <integer array_index="i" name="street_id" />
        <string array_index="i" name="country" />
        <string array_index="i" name="region" />
        <string array_index="i" name="city" />
        <string array_index="i" name="street" />
      </for>
    </output>
  </function>


*/
