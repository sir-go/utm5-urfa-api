package fn

// rpcf_user5_get_tariff_history
func xU15026(c conn, p Dict) Dict {
	c.Call(-0x15026)
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15026" name="rpcf_user5_get_tariff_history">
    <input>
      <integer name="account_id" />
    </input>
    <output>
	 <integer name="size" />
	 <for count="size" from="0" name="i">
	   <integer array_index="i" name="tariff_id" />
	   <integer array_index="i" name="link_date" />
	   <integer array_index="i" name="unlink_date" />
	   <string array_index="i" name="tariff_name" />
	 </for>
    </output>
  </function>



*/
