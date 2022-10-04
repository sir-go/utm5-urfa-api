package fn

// rpcf_get_registry_settings_list
func x5070(c conn, _ Dict) Dict {
	c.Call(0x5070)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5070" name="rpcf_get_registry_settings_list">
    <input />
    <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
              <integer array_index="i" name="id" />
              <string array_index="i" name="name" />
              <string array_index="i" name="description" />
              <integer array_index="i" name="validation_type" />
              <integer array_index="i" name="object_size" />
                  <for count="object_size" from="0" name="j">
                      <integer array_index="i,j" name="object_id" />
                      <string array_index="i,j" name="value" />
                  </for>
            </for>
    </output>
  </function>


*/
