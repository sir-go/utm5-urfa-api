package fn

// rpcf_invoices_list_for_dealer
func x13026(c conn, p Dict) Dict {
	c.Call(0x13026)
	putI(c, p, "dealer_id")
	putI(c, p, "start_date")
	putI(c, p, "end_date")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x13026" name="rpcf_invoices_list_for_dealer">
      <input>
	   <integer name="dealer_id" />
	   <integer name="start_date" />
	   <integer name="end_date" />
      </input>
      <output>
           <integer name="count" />
	   <for count="count" from="0" name="i">
	       <integer array_index="i" name="invoice_id" />
	       <string array_index="i" name="ext_num" />
	       <integer array_index="i" name="invoice_date" />
	       <integer array_index="i" name="user_id" />
	       <integer array_index="i" name="payment_transaction_id" />
	       <integer array_index="i" name="expire_date" />
	       <integer array_index="i" name="is_payed" />
	       <integer array_index="i" name="is_printed" />
	       <integer array_index="i" name="account_id" />
	       <string array_index="i" name="full_name" />
	       <double array_index="i" name="sum" />
	       <double array_index="i" name="tax" />
	       <double array_index="i" name="sum_with_tax" />
	   </for>
      </output>
    </function>






*/
