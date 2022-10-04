package fn

// rpcf_dealer_save_account
func x70000015(c conn, p Dict) Dict {
	c.Call(0x70000015)

	// fixme: input has a complex param
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000015" name="rpcf_dealer_save_account">
    <input>
      <integer name="account_id" />
      <integer name="discount_period_id" />
      <double name="credit" />
      <integer name="is_blocked" />
      <if condition="ne" value="0" variable="is_blocked">
        <integer default="now()" name="block_start_date" />
        <integer default="max_time()" name="block_end_date" />
      </if>
      <double name="vat_rate" />
      <double name="sale_tax_rate" />
      <integer name="int_status" />
      <integer name="block_recalc_abon" />
      <integer name="block_recalc_prepaid" />
      <integer name="unlimited" />
    </input>
    <output>
         <integer name="aid" />
         <if condition="eq" value="0" variable="aid">
           <error code="11" comment="account not found or you dont have enough privileges" />
         </if>
    </output>
    </function>




*/
