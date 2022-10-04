package fn

// rpcf_get_uaparam_list
func x440b(c conn, _ Dict) Dict {
	c.Call(0x440b)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x440b" name="rpcf_get_uaparam_list">
    <input />
    <output>
      <integer name="uparam_size" />
      <for count="uparam_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
        <string array_index="i" name="display_name" />
        <integer array_index="i" name="visible" />
      </for>
    </output>
  </function>


*/
