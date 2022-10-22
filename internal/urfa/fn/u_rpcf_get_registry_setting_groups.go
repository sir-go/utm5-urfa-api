package fn

// rpcf_get_registry_setting_groups
func x103(c conn, p Dict) Dict {
	c.Call(0x103)
	putI(c, p, "component_id", 1)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x103" name="rpcf_get_registry_setting_groups">
    <input>
      <integer default="1" name="component_id" />
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="group_id_array" />
        <integer array_index="i" name="reserved_array" />
        <string array_index="i" name="description_array" />
      </for>
    </output>
  </function>


*/
