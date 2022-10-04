package fn

// rpcf_get_tel_direction
func x10306(c conn, p Dict) Dict {
	c.Call(0x10306)
	putI(c, p, "dir_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x10306" name="rpcf_get_tel_direction">
    <input>
        <integer name="dir_id" />
    </input>
    <output>
        <integer name="dir_id" />
        <if condition="ne" value="-1" variable="dir_id">
            <integer name="zone_id" />
            <integer name="supplier_id" />
            <string name="name" />
            <string name="calling_prefix" />
            <string name="called_prefix" />
            <string name="incoming_trunk" />
            <string name="outgoing_trunk" />
            <string name="pbx_id" />
            <integer name="calling_prefix_regexp" />
            <integer name="called_prefix_regexp" />
            <integer name="skip" />
            <integer name="create_date" />
            <integer name="update_date" />
        </if>
    </output>
  </function>


*/
