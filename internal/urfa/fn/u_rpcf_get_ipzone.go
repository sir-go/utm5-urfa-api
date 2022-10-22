package fn

// rpcf_get_ipzone
func x2802(c conn, p Dict) Dict {
	c.Call(0x2802)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"name": c.GetS(),
		"ip_zones": getMapOf(c, func() Dict {
			return Dict{
				"net":       c.GetA(),
				"mask":      c.GetI(),
				"zone_name": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2802" name="rpcf_get_ipzone">
    <input>
      <integer name="id" />
    </input>
    <output>
      <string name="name" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="net" />
        <integer array_index="i" name="mask" />
        <integer array_index="i" name="gateaway" />
      </for>
    </output>
  </function>


*/
