package fn

// rpcf_get_banks
func x6002(c conn, _ Dict) Dict {
	c.Call(0x6002)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x6002" name="rpcf_get_banks">
    <input />
    <output>
      <integer name="banks_size" />
      <for count="banks_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="bic" />
        <string array_index="i" name="name" />
        <string array_index="i" name="city" />
        <string array_index="i" name="kschet" />
      </for>
    </output>
  </function>


*/
