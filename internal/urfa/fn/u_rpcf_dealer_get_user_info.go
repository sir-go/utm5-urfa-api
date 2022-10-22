package fn

// rpcf_dealer_get_user_info
func x70000005(c conn, p Dict) Dict {
	c.Call(0x70000005)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000005" name="rpcf_dealer_get_user_info">
      <input>
        <integer name="user_id" />
      </input>
      <output>
        <integer name="user_id" />
        <if condition="eq" value="0" variable="user_id">
          <error code="10" comment="user not found" />
        </if>
        <integer name="accounts_count" />
        <for count="accounts_count" from="0" name="i">
          <integer array_index="i" name="account_id_array" />
          <string array_index="i" name="account_name_array" />
        </for>
        <string name="login" />
        <string name="password" />
        <integer name="basic_account" />
        <string name="full_name" />
        <integer name="create_date" />
        <integer name="last_change_date" />
        <integer name="who_create" />
        <integer name="who_change" />
        <integer name="is_juridical" />
        <string name="jur_address" />
        <string name="act_address" />
        <string name="work_tel" />
        <string name="home_tel" />
        <string name="mob_tel" />
        <string name="web_page" />
        <string name="icq_number" />
        <string name="tax_number" />
        <string name="kpp_number" />
        <integer name="bank_id" />
        <string name="bank_account" />
        <string name="comments" />
        <string name="personal_manager" />
        <integer name="connect_date" />
        <string name="email" />
        <integer name="is_send_invoice" />
        <integer name="advance_payment" />
        <integer name="house_id" />
        <string name="flat_number" />
        <string name="entrance" />
        <string name="floor" />
        <string name="district" />
        <string name="building" />
        <string name="passport" />
        <integer name="parameters_size" />
        <for count="parameters_size" from="0" name="i">
          <integer array_index="i" name="parameted_id" />
          <string array_index="i" name="parameter_value" />
        </for>
      </output>
    </function>


*/
