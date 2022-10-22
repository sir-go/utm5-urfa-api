package fn

// rpcf_get_radius_attr
func x10106(c conn, p Dict) Dict {
	c.Call(0x10106)
	putI(c, p, "sid")
	putI(c, p, "st")
	c.Send()

	aSize := c.GetI()
	a := make([]Dict, aSize)

	for i := 0; i < aSize; i++ {
		a[i] = Dict{
			"vendor":      c.GetI(),
			"attr":        c.GetI(),
			"tag":         c.GetI(),
			"usage_flags": c.GetI(),
			"expire_date": c.GetI(),
			"type":        c.GetI(),
		}

		switch a[i]["type"] {
		case 0:
			a[i]["val"] = c.GetI()
		case 1:
			a[i]["val"] = c.GetS()
		case 2:
			a[i]["val"] = c.GetI()
		case 3:
			a[i]["val"] = c.GetS()
		}

		bSize := c.GetI()
		a[i]["props"] = make([]Dict, bSize)

		for j := 0; j < bSize; j++ {
			a[i]["props"].([]Dict)[j] = Dict{
				"type":  c.GetI(),
				"value": c.GetS(),
			}
		}
	}

	return Dict{"attrs": a}
}

/* xml:
<function id="0x10106" name="rpcf_get_radius_attr">
    <input>
      <integer name="sid" />
      <integer name="st" />
    </input>
    <output>
      <integer name="radius_data_size" />
      <for count="radius_data_size" from="0" name="i">
        <integer array_index="i" name="vendor" />
        <integer array_index="i" name="attr" />
        <integer array_index="i" name="tag" />
        <integer array_index="i" name="usage_flags" />
        <integer array_index="i" name="expire_date" />
        <integer array_index="i" name="param1" />
        <set dst="tmp_type" src="param1" src_index="i" />

        <if condition="eq" value="0" variable="tmp_type">
          <integer array_index="i" name="val" />
        </if>

        <if condition="eq" value="1" variable="tmp_type">
          <string array_index="i" name="val" />
        </if>

        <if condition="eq" value="2" variable="tmp_type">
          <integer array_index="i" name="val" />
        </if>

        <if condition="eq" value="3" variable="tmp_type">
          <string array_index="i" name="val" />
        </if>
        <integer array_index="i" name="prop_size" />
        <for count="prop_size" from="0" name="p">
          <integer array_index="i,p" name="type" />
          <string array_index="i,p" name="value" />
        </for>
      </for>
    </output>
  </function>


*/
