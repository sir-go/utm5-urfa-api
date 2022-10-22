package fn

// rpcf_dealer_add_payment_for_account
func x70000016(c conn, p Dict) Dict {
	c.Call(0x70000016)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000016" name="rpcf_dealer_add_payment_for_account">
      <input>
        <integer name="account_id" />
        <double name="payment" />
        <integer default="810" name="currency_id" />
        <integer default="now()" name="payment_date" />
        <integer default="0" name="burn_date" />
        <integer default="1" name="payment_method" />
        <string default="" name="admin_comment" />
        <string default="" name="comment" />
        <string default="" name="payment_ext_number" />
        <integer default="0" name="payment_to_invoice" />
        <integer default="1" name="turn_on_inet" />
      </input>
      <output>
          <integer name="payment_transaction_id" />
          <if condition="eq" value="0" variable="payment_transaction_id">
               <error code="13" comment="payment failed" />
          </if>
      </output>
    </function>


*/
