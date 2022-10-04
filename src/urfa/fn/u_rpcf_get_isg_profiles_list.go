package fn

// rpcf_get_isg_profiles_list
func x1403(c conn, _ Dict) Dict {
	c.Call(0x1403)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1403" name="rpcf_get_isg_profiles_list">
     <output>
       <integer name="size" />
       <for count="size" from="0" name="i">
         <integer array_index="i" name="id" />
         <string array_index="i" name="name" />
         <integer array_index="i" name="login_type" />
         <integer array_index="i" name="password_type" />
         <integer array_index="i" name="blocked_account_code" />
         <integer array_index="i" name="unblocked_account_code" />
         <string array_index="i" name="static_password" />
         <integer array_index="i" name="send_coa" />
         <integer array_index="i" name="session_time" />
       </for>
     </output>
   </function>


*/
