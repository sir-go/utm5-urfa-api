package fn

// rpcf_get_ipzones_list
func x2800(c conn, _ Dict) Dict {
	c.Call(0x2800)

	return Dict{
		"ip_zones": getMapOf(c, func() Dict {
			return Dict{
				"zone_id":   c.GetI(),
				"zone_name": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2800" name="rpcf_get_ipzones_list">
    <input />
    <output>
      <integer name="zones_count" />
      <for count="zones_count" from="0" name="i">
        <integer array_index="i" name="zone_id" />
        <string array_index="i" name="zone_name" />
      </for>
    </output>
  </function>


*/
