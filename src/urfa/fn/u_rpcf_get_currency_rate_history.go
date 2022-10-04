package fn

// rpcf_get_currency_rate_history
func x2913(c conn, p Dict) Dict {
	c.Call(0x2913)
	putI(c, p, "id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2913" name="rpcf_get_currency_rate_history">
    <input>
      <integer name="id" />
    </input>
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="time" />
        <double array_index="i" name="value" />
      </for>
    </output>
  </function>


*/
