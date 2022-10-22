package fn

// rpcf_dealer_blocks_report
func x70000035(c conn, p Dict) Dict {
	c.Call(0x70000035)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000035" name="rpcf_dealer_blocks_report">
       <input>
        <integer default="0" name="user_id" />
        <integer default="0" name="account_id" />
        <integer default="0" name="group_id" />
        <integer default="0" name="apid" />
        <integer name="time_start" />
        <integer default="now()" name="time_end" />
        <integer default="1" name="show_all" />
       </input>
       <output>
        <integer name="accounts_count" />
        <for count="accounts_count" from="0" name="i">
         <integer name="atr_size" />
            <set dst="atr_size_array" dst_index="i" src="atr_size" />
         <for count="atr_size" from="0" name="j">
           <integer array_index="i,j" name="account_id" />
           <string array_index="i,j" name="login" />
           <integer array_index="i,j" name="start_date" />
           <integer array_index="i,j" name="expire_date" />
           <integer array_index="i,j" name="what_blocked" />
           <integer array_index="i,j" name="block_type" />
           <string array_index="i,j" name="comment" />
         </for>
        </for>
       </output>
     </function>


*/
