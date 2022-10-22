package fn

// rpcf_get_voluntary_suspension_settings
func x15010(c conn, _ Dict) Dict {
	c.Call(0x15010)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x15010" name="rpcf_get_voluntary_suspension_settings">
    <input>
    </input>
    <output>
    <integer name="count" />
    <for count="count" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="group_id" />
        <integer array_index="i" name="priopity" />
        <integer array_index="i" name="is_enabled" />
        <integer array_index="i" name="min_duration" />
        <integer array_index="i" name="max_duration" />
        <integer array_index="i" name="interval_duration" />
        <integer array_index="i" name="block_type" />
        <integer array_index="i" name="self_unlock" />
        <double array_index="i" name="min_balance" />
        <integer array_index="i" name="use_min_balance" />
        <double array_index="i" name="free_balance" />
        <integer array_index="i" name="use_free_balance" />
        <integer array_index="i" name="service_id" />
    </for>
    </output>
  </function>


*/
