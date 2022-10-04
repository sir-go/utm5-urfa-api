package fn

// rpcf_tel_zones_list
func x10309(c conn, _ Dict) Dict {
	c.Call(0x10309)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x10309" name="rpcf_tel_zones_list">
    <input />
    <output>
        <integer name="zone_count" />
        <for count="zone_count" from="0" name="i">
            <integer array_index="i" name="zone_id_array" />
            <integer array_index="i" name="supplier_id_array" />
            <string array_index="i" name="name_array" />
            <integer array_index="i" name="create_date_array" />
            <integer array_index="i" name="update_date_array" />
        </for>
    </output>
  </function>


*/
