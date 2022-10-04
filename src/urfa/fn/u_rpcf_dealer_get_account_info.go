package fn

// rpcf_dealer_get_account_info
func x70000013(c conn, p Dict) Dict {
	c.Call(0x70000013)
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000013" name="rpcf_dealer_get_account_info">
      <input>
        <integer name="account_id" />
      </input>
      <output>
        <integer name="account_id" />
        <if condition="eq" value="0" variable="account_id">
          <error code="11" comment=" account not found or you dont have enough privileges" />
        </if>
        <integer name="discount_period_id" />
        <integer name="is_blocked" />
        <double name="vat_rate" />
        <double name="sale_tax_rate" />
        <double name="comission_coefficient" />
        <double name="default_comission_value" />
        <double name="credit" />
        <double name="balance" />
        <integer name="int_status" />
        <integer name="block_recalc_abon" />
        <integer name="block_recalc_prepaid" />
        <integer name="unlimited" />
      </output>
    </function>


*/
