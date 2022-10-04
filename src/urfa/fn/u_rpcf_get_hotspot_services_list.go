package fn

// rpcf_get_hotspot_services_list
func x210f(c conn, _ Dict) Dict {
	c.Call(0x210f)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x210f" name="rpcf_get_hotspot_services_list">
    <input />
    <output>
      <integer name="services_size" />
      <for count="services_size" from="0" name="i">
        <integer array_index="i" name="service_id" />
        <if condition="ne" value="-1" variable="service_id">
        <string array_index="i" name="service_name" />
        <integer array_index="i" name="service_type" />
        <string array_index="i" name="commect" />
        </if>
      </for>
    </output>
  </function>


*/
