package fn

// rpcf_save_cashier_settings
func x212d(c conn, p Dict) Dict {
	c.Call(0x212d)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x212d" name="rpcf_save_cashier_settings">
    <input>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <string array_index="i" name="attribute" />
	<string array_index="i" name="value" />
      </for>
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
