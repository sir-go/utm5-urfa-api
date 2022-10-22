package fn

// rpcf_tel_directions_list
func x1031d(c conn, p Dict) Dict {
	c.Call(0x1031d)

	// fixme: input has a complex param
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1031d" name="rpcf_tel_directions_list">
    <input>
        <integer name="start" />
        <integer name="count" />
        <integer name="skip_ids_cnt" />
        <for count="skip_ids_cnt" from="0" name="i">
            <integer array_index="i" name="skip_id" />
        </for>
    </input>
    <output>
        <integer name="dir_count" />
        <for count="dir_count" from="0" name="i">
            <integer array_index="i" name="dir_id_array" />
            <integer array_index="i" name="zone_id_array" />
            <integer array_index="i" name="supplier_id_array" />
            <string array_index="i" name="dir_name_array" />
            <string array_index="i" name="calling_prefix_array" />
            <string array_index="i" name="called_prefix_array" />
            <integer array_index="i" name="create_date_array" />
            <integer array_index="i" name="update_date_array" />
        </for>
    </output>
  </function>


*/
