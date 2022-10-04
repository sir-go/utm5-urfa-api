package fn

// rpcf_edit_dealer
func x13002(c conn, p Dict) Dict {
	c.Call(0x13002)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Modify dealer staff" id="0x13002" name="rpcf_edit_dealer">
      <input>
        <integer name="dealer_id" />
        <string name="login" />
        <string name="password" />
        <ip_address default="0.0.0.0" name="ip4_address" />
        <integer default="32" name="mask4" />
        <ip_address default="::0" name="ip6_address" />
        <integer default="128" name="mask6" />
        <string default="" name="full_name" />
        <string default="" name="act_address" />
        <string default="" name="passport" />
        <string default="" name="work_tel" />
        <string default="" name="home_tel" />
        <string default="" name="mob_tel" />
        <string default="" name="web_page" />
        <string default="" name="icq_number" />
        <string default="" name="email" />
        <string default="" name="comments" />
        <integer default="size(group_id)" name="groups_count" />
        <for count="size(group_id)" from="0" name="i">
	    <integer array_index="i" name="group_id" />
        </for>
      </input>
      <output>
        <integer name="result" />
         <if condition="eq" value="0" variable="result">
            <error code="10" comment="error adding dealer. look debug.log for details" />
        </if>
      </output>
    </function>


*/
