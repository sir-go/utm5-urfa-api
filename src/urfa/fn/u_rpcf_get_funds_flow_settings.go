package fn

// rpcf_get_funds_flow_settings
func x15024(c conn, _ Dict) Dict {
	c.Call(0x15024)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x15024" name="rpcf_get_funds_flow_settings">
    <input>
    </input>
    <output>
    <integer name="count" />
    <for count="count" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="group_id" />
        <integer array_index="i" name="priopity" />
        <integer array_index="i" name="is_enabled" />

    </for>
    </output>
  </function>


*/
