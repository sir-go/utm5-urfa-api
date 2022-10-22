package fn

// rpcf_dealer_get_invoices_list
func x7000003c(c conn, p Dict) Dict {
	c.Call(0x7000003c)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000003c" name="rpcf_dealer_get_invoices_list">
      <input>
           <integer name="account_id" />
	   <integer name="time_start" />
	   <integer default="now()" name="time_end" />
	   <integer name="group_id" />
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
	     <if condition="eq" value="0" variable="count_of_invoice">
	       <set dst="currency_id" dst_index="i" value="0" />
	       <set dst="currency_name" dst_index="i" value="" />
	       <set dst="payment_rule" dst_index="i" value="" />
	     </if>
             <for count="count_of_invoice" from="0" name="j">
	       <integer array_index="i,j" name="id" />
	       <string array_index="i,j" name="ext_num" />
	       <integer array_index="i,j" name="invoice_date" />
	       <integer array_index="i,j" name="uid" />
	       <integer array_index="i,j" name="payment_transaction_id" />
	       <integer array_index="i,j" name="expire_date" />
	       <integer name="is_payed" />
	       <set dst="is_payed_array" dst_index="i,j" src="is_payed" />
	       <integer array_index="i,j" name="is_printed" />
	       <integer array_index="i,j" name="account_id" />
	       <string array_index="i,j" name="full_name" />
	       <if condition="ne" value="0" variable="is_payed">
	         <integer array_index="i,j" name="accInvcInfo_id" />
		 <integer array_index="i,j" name="accInvcInfo_date" />
		 <integer array_index="i,j" name="accInvcInfo_payed_date" />
		 <string array_index="i,j" name="payment_ext_number" />
		 <integer array_index="i,j" name="is_printed" />
	       </if>
	       <if condition="eq" value="0" variable="is_payed">
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
	       <if condition="eq" value="0" variable="entry_size">
	         <set dst="total_sum" dst_index="i,j" value="0" />
		 <set dst="total_tax" dst_index="i,j" value="0" />
		 <set dst="total_sum_plus_total_tax" dst_index="i,j" value="0" />
	       </if>
	     </for>
	   </for>
      </output>
    </function>


*/
