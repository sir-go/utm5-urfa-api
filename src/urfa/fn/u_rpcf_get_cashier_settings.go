package fn

// rpcf_get_cashier_settings
func x212c(c conn, _ Dict) Dict {
	c.Call(0x212c)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x212c" name="rpcf_get_cashier_settings">
    <input />
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <string array_index="i" name="attribute" />
	<string array_index="i" name="value" />
      </for>
    </output>
  </function>


*/
