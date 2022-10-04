package fn

// rpcf_get_tel_zone
func x10307(c conn, p Dict) Dict {
	c.Call(0x10307)
	putI(c, p, "zone_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x10307" name="rpcf_get_tel_zone">
    <input>
        <integer name="zone_id" />
    </input>
    <output>
        <integer name="zone_id" />
        <if condition="ne" value="-1" variable="zone_id">
            <integer name="supplier_id" />
            <string name="name" />
            <integer name="create_date" />
            <integer name="update_date" />
            <integer name="dir_count" />
            <for count="dir_count" from="0" name="i">
                <integer array_index="i" name="dir_id_array" />
                <string array_index="i" name="dir_name_array" />
            </for>
        </if>
    </output>
  </function>


*/
