package fn

// rpcf_user5_get_invoice
func xU15033(c conn, p Dict) Dict {
	c.Call(-0x15033)
	putI(c, p, "invoice_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15033" name="rpcf_user5_get_invoice">
    <input>
      <integer name="invoice_id" />
    </input>
    <output>
      <integer name="invoice_id" />
      <if condition="ne" value="0" variable="invoice_id">
        <integer name="account_id" />
        <integer name="invoice_date" />
        <integer name="period_start" />
        <integer name="period_end" />
        <double name="balance_on_set" />
        <integer name="entries_size" />
        <for count="entries_size" from="0" name="i">
          <string name="ent_name" />
          <double name="qnt" />
          <double name="base_cost" />
          <double name="sum_cost" />
          <double name="tax_amount" />
        </for>
      </if>
    </output>
  </function>


*/
