package fn

// rpcf_set_radius_attr
func x10105(c conn, p Dict) Dict {
	c.Call(0x10105)

	putI(c, p, "sid")
	putI(c, p, "st")

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

	return nil
}

/* xml:
<function id="0x10105" name="rpcf_set_radius_attr">
    <input>
      <integer name="sid" />
      <integer name="st" />
      <integer name="cnt" />
      <for count="cnt" from="0" name="i">
        <integer array_index="i" name="vendor" />
        <integer array_index="i" name="attr" />
        <integer array_index="i" name="tag" />
        <integer array_index="i" name="usage_flags" />
        <integer array_index="i" default="2000000000" name="expire_date" />
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
    </input>
    <output />
  </function>


*/
