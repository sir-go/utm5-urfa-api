package fn

// rpcf_set_registry_setting
func x105(c conn, p Dict) Dict {
	c.Call(0x105)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x105" name="rpcf_set_registry_setting">
    <input>
      <integer name="setting_id" />
      <integer name="value_size" />
      <for count="value_size" from="0" name="i">
        <string array_index="i" name="value_array" />
      </for>
      <integer default="0" name="object_id" />
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>




*/
