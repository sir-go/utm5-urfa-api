package fn

// rpcf_dealer_add_user
func x70000001(c conn, p Dict) Dict {
	c.Call(0x70000001)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000001" name="rpcf_dealer_add_user">
      <input>
        <string name="login" />
        <string name="password" />
        <string default="" name="full_name" />
        <integer default="0" name="is_juridical" />
        <string default="" name="jur_address" />
        <string default="" name="act_address" />
        <string default="" name="flat_number" />
        <string default="" name="entrance" />
        <string default="" name="floor" />
        <string default="" name="district" />
        <string default="" name="building" />
        <string default="" name="passport" />
        <integer default="0" name="house_id" />
        <string default="" name="work_tel" />
        <string default="" name="home_tel" />
        <string default="" name="mob_tel" />
        <string default="" name="web_page" />
        <string default="" name="icq_number" />
        <string default="" name="tax_number" />
        <string default="" name="kpp_number" />
        <string default="" name="email" />
        <integer default="0" name="bank_id" />
        <string default="" name="bank_account" />
        <string default="" name="comments" />
        <string default="" name="personal_manager" />
        <integer default="0" name="connect_date" />
        <integer default="0" name="is_send_invoice" />
        <integer default="0" name="advance_payment" />
        <integer default="size(parameter_value)" name="parameters_count" />
        <for count="size(parameter_value)" from="0" name="i">
          <integer array_index="i" name="parameter_id" />
          <string array_index="i" name="parameter_value" />
        </for>
      </input>
      <output>
        <integer name="user_id" />
	<if condition="ne" value="0" variable="user_id">
            <integer name="account_id" />
	</if>
      </output>
    </function>


*/
