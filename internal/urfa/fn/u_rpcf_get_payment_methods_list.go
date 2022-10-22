package fn

// rpcf_get_payment_methods_list
func x3100(c conn, _ Dict) Dict {
	c.Call(0x3100)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3100" name="rpcf_get_payment_methods_list">
    <input />
    <output>
      <integer name="payments_list_count" />
      <for count="payments_list_count" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
      </for>
    </output>
  </function>


*/
