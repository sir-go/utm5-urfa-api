package fn

// rpcf_dealer_get_banks
func x7000000e(c conn, _ Dict) Dict {
	c.Call(0x7000000e)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000000e" name="rpcf_dealer_get_banks">
      <input>
      </input>
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
