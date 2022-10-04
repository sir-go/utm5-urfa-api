package fn

// rpcf_put_banks
func x6001(c conn, p Dict) Dict {
	c.Call(0x6001)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x6001" name="rpcf_put_banks">
    <input>
      <integer default="size(bic)" name="banks_size" />
      <for count="size(bic)" from="0" name="i">
        <string array_index="i" name="bic" />
        <string array_index="i" name="name" />
        <string array_index="i" name="city" />
        <string array_index="i" name="kschet" />
      </for>
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
