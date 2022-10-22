package fn

// rpcf_user5_get_accounts
func xU4055(c conn, _ Dict) Dict {
	c.Call(-0x4055)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4055" name="rpcf_user5_get_accounts">
    <input />
    <output>
      <integer name="accounts_size" />
      <for count="accounts_size" from="0" name="i">
        <integer array_index="i" name="account_id" />
        <double array_index="i" name="balance" />
        <double array_index="i" name="credit" />
      </for>
    </output>
  </function>


*/
