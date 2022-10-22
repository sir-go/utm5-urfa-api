package fn

// rpcf_get_switch_tariff_settings
func x15014(c conn, _ Dict) Dict {
	c.Call(0x15014)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x15014" name="rpcf_get_switch_tariff_settings">
    <input>
    </input>
    <output>
       <integer name="count" />
       <for count="count" from="0" name="i">
            <integer array_index="i" name="id" />
            <integer array_index="i" name="group_id" />
            <integer array_index="i" name="priopity" />
            <integer array_index="i" name="is_enabled" />
            <integer array_index="i" name="instant_change" />
            <integer array_index="i" name="instant_change_count" />
            <integer name="tp_size" />
            <set dst="tp_size_array" dst_index="i" src="tp_size" />
            <for count="tp_size" from="0" name="j">
                <integer array_index="i,j" name="tp_id" />
                <double array_index="i,j" name="min_balance" />
                <integer array_index="i,j" name="use_min_balance" />
                <double array_index="i,j" name="free_balance" />
                <integer array_index="i,j" name="use_free_balance" />
                <integer array_index="i,j" name="service_id" />
            </for>
       </for>
    </output>
  </function>


*/
