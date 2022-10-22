package fn

// rpcf_get_registry_form
func x104(c conn, p Dict) Dict {
	c.Call(0x104)
	putI(c, p, "group_id")
	putI(c, p, "object_id", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x104" name="rpcf_get_registry_form">
    <input>
      <integer name="group_id" />
      <integer default="0" name="object_id" />
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="setting_id_array" />
        <string array_index="i" name="description_array" />
        <integer array_index="i" name="value_type_array" />
        <integer array_index="i" name="attr_array" />
        <integer name="value_size" />
        <for count="value_size" from="0" name="j">
          <string array_index="i,j" name="value_array" />
        </for>
      </for>
    </output>
  </function>


*/
