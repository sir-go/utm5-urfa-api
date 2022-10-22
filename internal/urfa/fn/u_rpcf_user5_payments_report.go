package fn

// rpcf_user5_payments_report
func xU4019(c conn, p Dict) Dict {
	c.Call(-0x4019)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4019" name="rpcf_user5_payments_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="accounts_count" />
      <for count="accounts_count" from="0" name="j">
      	<integer name="atr_size" />
      	<for count="atr_size" from="0" name="i">
       	 <integer array_index="j,i" name="account_id" />
       	 <integer array_index="j,i" name="actual_date" />
       	 <integer array_index="j,i" name="payment_enter_date" />
       	 <double array_index="j,i" name="payment" />
       	 <double array_index="j,i" name="payment_incurrency" />
       	 <integer array_index="j,i" name="currency_id" />
       	 <integer array_index="j,i" name="payment_method_id" />
       	 <string array_index="j,i" name="payment_method" />
       	 <string array_index="j,i" name="comment" />
      	</for>
      </for>
    </output>
  </function>


*/
