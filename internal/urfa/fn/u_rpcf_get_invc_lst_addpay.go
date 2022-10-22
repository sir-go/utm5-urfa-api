package fn

// rpcf_get_invc_lst_addpay
func x8002(c conn, p Dict) Dict {
	c.Call(0x8002)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x8002" name="rpcf_get_invc_lst_addpay">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
    </input>
    <output>
      <integer name="accounts_size" />
      <for count="accounts_size" from="0" name="i">
        <integer array_index="i" name="aid" />
        <integer array_index="i" name="binded_currency" />
        <integer array_index="i" name="inv_size" />
        <for count="inv_size" from="0" name="j">

            <integer array_index="i,j" name="aid" />
            <integer array_index="i,j" name="id" />
          <string array_index="i,j" name="ext_num" />
          <integer array_index="i,j" name="invoice_date" />
          <double array_index="i,j" name="system_cur_sum" />
          <double array_index="i,j" name="binded_cur_sum" />
          <double array_index="i,j" name="total" />

        </for>
      </for>
    </output>
  </function>


*/
