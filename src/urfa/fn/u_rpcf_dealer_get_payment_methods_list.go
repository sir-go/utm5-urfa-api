package fn

// rpcf_dealer_get_payment_methods_list
func x70000019(c conn, _ Dict) Dict {
	c.Call(0x70000019)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000019" name="rpcf_dealer_get_payment_methods_list">
      <input>
      </input>
      <output>
          <integer name="payments_list_count" />
          <for count="payments_list_count" from="0" name="i">
	     <integer array_index="i" name="id" />
	     <string array_index="i" name="name" />
          </for>
      </output>
    </function>


*/
