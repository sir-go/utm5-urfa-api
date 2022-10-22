package fn

// rpcf_dealer_get_currency_list
func x7000003e(c conn, _ Dict) Dict {
	c.Call(0x7000003e)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000003e" name="rpcf_dealer_get_currency_list">
      <input>
      </input>
      <output>
	 <integer name="currency_size" />
	 <for count="currency_size" from="0" name="i">
	   <integer array_index="i" name="id" />
	   <string array_index="i" name="currency_brief_name" />
	   <string array_index="i" name="currency_full_name" />
	   <double array_index="i" name="percent" />
	   <double array_index="i" name="rates" />
         </for>
      </output>
    </function>


*/
