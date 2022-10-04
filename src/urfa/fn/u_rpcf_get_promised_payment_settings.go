package fn

// rpcf_get_promised_payment_settings
func x15034(c conn, _ Dict) Dict {
	c.Call(0x15034)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x15034" name="rpcf_get_promised_payment_settings">
    <input>
    </input>
    <output>
    <integer name="count" />
    <for count="count" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="group_id" />
        <integer array_index="i" name="priopity" />
        <integer array_index="i" name="is_enabled" />
        <double array_index="i" name="max_value" />
        <integer array_index="i" name="max_duration" />
        <integer array_index="i" name="interval_duration" />
        <double array_index="i" name="min_balance" />
        <integer array_index="i" name="use_min_balance" />
        <double array_index="i" name="free_balance" />
        <integer array_index="i" name="use_free_balance" />
        <integer array_index="i" name="service_id" />
        <integer array_index="i" name="flags" />
    </for>
    </output>
  </function>


*/
