package fn

// rpcf_user5_get_user_lifestream_tariff_list
func xU15050(c conn, _ Dict) Dict {
	c.Call(-0x15050)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15050" name="rpcf_user5_get_user_lifestream_tariff_list">
    <input>
    </input>
    <output>
      <integer name="user_tariffs_size" />
      <for count="user_tariffs_size" from="0" name="i">
        <integer array_index="i" name="tariff_id_array" />
        <string array_index="i" name="tariff_name_array" />
        <string array_index="i" name="tariff_comments_array" />
        <double array_index="i" name="cost" />
      </for>
    </output>
  </function>


*/
