package fn

// rpcf_get_dealer_info
func x13004(c conn, p Dict) Dict {
	c.Call(0x13004)
	putI(c, p, "dealer_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Get dealer info" id="0x13004" name="rpcf_get_dealer_info">
      <input>
        <integer name="dealer_id" />
      </input>
      <output>
          <integer name="result" />
          <if condition="eq" value="0" variable="result">
            <error code="10" comment="error getting dealer info. look debug.log for details" />
          </if>
          <string name="login" />
          <string name="password" />
          <ip_address name="ip4_address" />
          <integer name="mask4" />
          <ip_address name="ip6_address" />
          <integer name="mask6" />
          <string name="full_name" />
          <string name="act_address" />
          <string name="passport" />
          <string name="work_tel" />
          <string name="home_tel" />
          <string name="mob_tel" />
          <string name="web_page" />
          <string name="icq_number" />
          <string name="email" />
          <string name="comments" />
	  <integer name="who_create" />
	  <integer name="who_change" />
	  <integer name="create_date" />
	  <integer name="change_date" />
          <integer name="groups_count" />
          <for count="groups_count" from="0" name="i">
	    <integer array_index="i" name="group_id" />
            <string array_index="i" name="group_name" />
         </for>
      </output>
    </function>


*/
