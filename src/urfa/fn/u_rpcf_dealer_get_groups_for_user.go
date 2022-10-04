package fn

// rpcf_dealer_get_groups_for_user
func x70000006(c conn, p Dict) Dict {
	c.Call(0x70000006)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000006" name="rpcf_dealer_get_groups_for_user">
      <input>
        <integer name="user_id" />
      </input>
      <output>
         <integer name="groups_size" />
         <if condition="eq" value="0" variable="groups_size">
            <error code="10" comment="user has no groups or you dont have enough privileges" />
         </if>
         <for count="groups_size" from="0" name="i">
	    <integer array_index="i" name="group_id" />
	    <string array_index="i" name="group_name" />
	 </for>
      </output>
    </function>


*/
