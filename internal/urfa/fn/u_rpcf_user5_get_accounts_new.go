package fn

// rpcf_user5_get_accounts_new
func xU15028(c conn, _ Dict) Dict {
	c.Call(-0x15028)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15028" name="rpcf_user5_get_accounts_new">
    <input />
    <output>
      <integer name="accounts_size" />
      <for count="accounts_size" from="0" name="i">
        <integer array_index="i" name="account_id" />
        <double array_index="i" name="balance" />
        <double array_index="i" name="credit" />
        <integer array_index="i" name="internet_status" />
        <integer array_index="i" name="block_status" />
        <double array_index="i" name="vat_rate" />
        <double array_index="i" name="locked_in_funds" />
      </for>
    </output>
  </function>


*/
