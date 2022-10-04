package fn

// rpcf_set_tel_emergency_calls
func x15033(c conn, p Dict) Dict {
	c.Call(0x15033)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x15033" name="rpcf_set_tel_emergency_calls">
      <input>
        <integer name="service_id" />
        <integer default="size(zone_id_array)" name="zone_count" />
        <for count="size(zone_id_array)" from="0" name="i">
          <integer array_index="i" name="zone_id_array" />
        </for>
      </input>
      <output>
        <integer name="result" />
      </output>
    </function>


*/
