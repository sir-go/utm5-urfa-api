package fn

// rpcf_edit_switch_tariff_settings
func x15016(c conn, p Dict) Dict {
	c.Call(0x15016)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x15016" name="rpcf_edit_switch_tariff_settings">
    <input>
            <integer name="id" />
            <integer name="group_id" />
            <integer name="priopity" />
            <integer name="is_enabled" />
            <integer name="instant_change" />
            <integer name="instant_change_count" />
            <integer name="tp_size" />
            <for count="tp_size" from="0" name="i">
                <integer array_index="i" name="tp_id" />
                <double array_index="i" name="min_balance" />
                <integer array_index="i" name="use_min_balance" />
                <double array_index="i" name="free_balance" />
                <integer array_index="i" name="use_free_balance" />
                <integer array_index="i" name="service_id" />
            </for>
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
