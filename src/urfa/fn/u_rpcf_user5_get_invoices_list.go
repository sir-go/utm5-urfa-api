package fn

// rpcf_user5_get_invoices_list
func xU4047(c conn, p Dict) Dict {
	c.Call(-0x4047)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4047" name="rpcf_user5_get_invoices_list">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="currency_id" />
      <string name="currency_name" />
      <integer name="accts_size" />
      <for count="accts_size" from="0" name="i">
        <integer name="count_of_invoice" />
        <set dst="count_of_invoice_array" dst_index="i" src="count_of_invoice" />
        <for count="count_of_invoice" from="0" name="j">
            <integer array_index="i,j" name="id" />
            <string array_index="i,j" name="ext_num" />
            <integer array_index="i,j" name="invoice_date" />
            <integer array_index="i,j" name="is_payed" />
            <integer array_index="i,j" name="account_id" />
            <double array_index="i,j" name="total_sum" />
            <double array_index="i,j" name="total_tax" />
            <double array_index="i,j" name="total_sum_plus_total_tax" />
        </for>
      </for>
    </output>
  </function>


*/
