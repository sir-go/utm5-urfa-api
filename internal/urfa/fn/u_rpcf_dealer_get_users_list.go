package fn

// rpcf_dealer_get_users_list
func x70000056(c conn, p Dict) Dict {
	c.Call(0x70000056)
	putI(c, p, "from", 0)
	putI(c, p, "to", -1)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000056" name="rpcf_dealer_get_users_list">
      <input>
         <integer default="0" name="from" />
	 <integer default="-1" name="to" />
      </input>
      <output>
        <integer name="cnt" />
        <for count="cnt" from="0" name="i">
           <integer array_index="i" name="user_id_array" />
           <string array_index="i" name="login_array" />
           <integer array_index="i" name="basic_account" />
           <string array_index="i" name="full_name" />
           <integer array_index="i" name="is_blocked" />
           <double array_index="i" name="balance" />
           <integer name="ip_adr_size" />
	   <set dst="ip_adr_size_array" dst_index="i" src="ip_adr_size" />
           <for count="ip_adr_size" from="0" name="j">
	        <integer name="group_size" />
		<set dst="group_size_array" dst_index="i,j" src="group_size" />
	        <for count="group_size" from="0" name="x">
	           <ip_address array_index="i,j,x" name="ip_address" />
                   <integer array_index="i,j,x" name="mask" />
    	           <integer array_index="i,j,x" name="group_type" />
	        </for>
           </for>
           <integer array_index="i" name="user_int_status" />
        </for>
      </output>
    </function>


*/
