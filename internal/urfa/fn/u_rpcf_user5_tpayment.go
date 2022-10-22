package fn

// rpcf_user5_tpayment
func xU4029(c conn, _ Dict) Dict {
	c.Call(-0x4029)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4029" name="rpcf_user5_tpayment">
    <input />
    <output>
      <integer name="is_exist" />
      <if condition="ne" value="0" variable="is_exist">
        <integer name="first_payment_time" />
        <integer name="last_payment_date" />
        <integer name="time2burn" />
        <double name="payment_value" />
        <double name="already_discounted" />
      </if>
    </output>
  </function>


*/
