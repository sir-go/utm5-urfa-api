package fn

// rpcf_dealer_add_payment_for_account_notify
func x70000017(c conn, p Dict) Dict {
	c.Call(0x70000017)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000017" name="rpcf_dealer_add_payment_for_account_notify">
      <input>
         <integer name="account_id" />
	 <double name="payment_incurrency" />
	 <integer name="currency_id" />
	 <integer default="now()" name="actual_date" />
	 <integer name="burn_date" />
	 <integer name="method" />
	 <string default="" name="admin_comment" />
	 <string default="" name="comment" />
	 <string name="payment_ext_number" />
	 <integer name="payment_to_invoice" />
	 <integer name="turn_on_inet" />
	 <integer name="notify" />
	 <string name="hash" />
      </input>
      <output>
          <integer name="payment_transaction_id" />
          <if condition="eq" value="0" variable="payment_transaction_id">
               <error code="13" comment="payment failed" />
          </if>
      </output>
    </function>


*/
