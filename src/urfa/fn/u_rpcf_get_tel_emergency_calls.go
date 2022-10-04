package fn

// rpcf_get_tel_emergency_calls
func x15032(c conn, _ Dict) Dict {
	c.Call(0x15032)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x15032" name="rpcf_get_tel_emergency_calls">
      <input />
      <output>
        <integer name="service_count" />
        <for count="service_count" from="0" name="i">
          <integer name="service_id" />
          <integer name="zone_count" />
          <for count="zone_count" from="0" name="j">
            <integer array_index="i" name="zone_id_array" />
          </for>
        </for>
      </output>
    </function>


*/
