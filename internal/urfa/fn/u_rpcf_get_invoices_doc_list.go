package fn

// rpcf_get_invoices_doc_list
func x8030(c conn, p Dict) Dict {
	c.Call(0x8030)
	putI(c, p, "aid")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	putI(c, p, "gid")
	putI(c, p, "doc_type")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x8030" name="rpcf_get_invoices_doc_list">
    <input>
      <integer name="aid" />
      <integer name="time_start" />
      <integer name="time_end" />
      <integer name="gid" />
      <integer name="doc_type" />
    </input>
    <output>
      <integer name="accts_size" />
      <for count="accts_size" from="0" name="i">
        <integer name="count_of_invoice" />
        <set dst="count_of_invoice_array" dst_index="i" src="count_of_invoice" />
        <if condition="ne" value="0" variable="count_of_invoice">
          <integer array_index="i" name="currency_id" />
          <string array_index="i" name="currency_name" />
          <string array_index="i" name="payment_rule" />
        </if>
        <for count="count_of_invoice" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <string array_index="i,j" name="alt_num" />
          <string array_index="i,j" name="ext_num" />
          <integer array_index="i,j" name="invoice_date" />
          <integer array_index="i,j" name="uid" />
          <integer array_index="i,j" name="payment_transaction_id" />
          <integer array_index="i,j" name="expire_date" />
          <integer name="tmp_is_payed" />
          <set dst="is_payed" dst_index="i,j" src="tmp_is_payed" />
          <integer array_index="i,j" name="is_printed" />
          <integer array_index="i,j" name="account_id" />
          <integer array_index="i,j" name="odt_status" />
          <integer array_index="i,j" name="odt_created" />
          <integer array_index="i,j" name="odt_modified" />
          <integer array_index="i,j" name="is_odt_mailed" />
          <integer array_index="i,j" name="pdf_status" />
          <integer array_index="i,j" name="pdf_created" />
          <integer array_index="i,j" name="pdf_modified" />
          <integer array_index="i,j" name="is_pdf_mailed" />
          <string array_index="i,j" name="full_name" />
          <if condition="ne" value="0" variable="tmp_is_payed">
            <integer array_index="i,j" name="accInvcInfo_id" />
            <integer array_index="i,j" name="accInvcInfo_date" />
            <integer array_index="i,j" name="accInvcInfo_payed_date" />
            <string array_index="i,j" name="payment_ext_number" />
            <integer array_index="i,j" name="is_printed" />
          </if>
          <if condition="eq" value="0" variable="tmp_is_payed">
            <set dst="accInvcInfo_id" dst_index="i,j" value="0" />
            <set dst="accInvcInfo_date" dst_index="i,j" value="0" />
            <set dst="accInvcInfo_payed_date" dst_index="i,j" value="0" />
            <set dst="payment_ext_number" dst_index="i,j" value="" />
            <set dst="is_printed" dst_index="i,j" value="0" />
          </if>
          <integer name="entry_size" />
          <set dst="entry_size_array" dst_index="i,j" src="entry_size" />
          <for count="entry_size" from="0" name="x">
            <string array_index="i,j,x" name="name" />
            <integer array_index="i,j,x" name="invoice_id" />
            <integer array_index="i,j,x" name="slink_id" />
            <integer array_index="i,j,x" name="service_type" />
            <integer array_index="i,j,x" name="discount_period_id" />
            <integer array_index="i,j,x" name="date" />
            <double array_index="i,j,x" name="qnt" />
            <double array_index="i,j,x" name="base" />
            <double array_index="i,j,x" name="sum" />
            <double array_index="i,j,x" name="tax" />
          </for>
          <if condition="ne" value="0" variable="entry_size">
            <double array_index="i,j" name="total_sum" />
            <double array_index="i,j" name="total_tax" />
            <double array_index="i,j" name="total_sum_plus_total_tax" />
          </if>
        </for>
      </for>
    </output>
  </function>


*/
