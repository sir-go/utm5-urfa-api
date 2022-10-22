package fn

// rpcf_get_currency_list
func x2910(c conn, _ Dict) Dict {
	c.Call(0x2910)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2910" name="rpcf_get_currency_list">
    <input>
    </input>
    <output>
      <integer name="currency_size" />
      <for count="currency_size" from="0" name="i">
        <integer name="id" />
        <string array_index="i" name="currency_brief_name" />
        <string array_index="i" name="currency_full_name" />
        <double array_index="i" name="percent" />
        <double array_index="i" name="rates" />
      </for>
    </output>
  </function>


*/
