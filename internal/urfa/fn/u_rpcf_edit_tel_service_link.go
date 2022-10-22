package fn

// rpcf_edit_tel_service_link
func x1033b(c conn, p Dict) Dict {
	c.Call(0x1033b)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"slinkId": c.GetI(),
	}
}

/* xml:
<function id="0x1033b" name="rpcf_edit_tel_service_link">
		<input>
			<integer name="slinkId" />
			<integer name="start_date" />
			<integer name="expire_date" />
            <integer name="policy_id" />
            <double name="cost_coef" />
            <integer default="0" name="house_id" />
            <string default="" name="comment" />
			<integer name="loginCount" />
			<for count="loginCount" from="0" name="i">
				<string array_index="i" name="login" />
				<string array_index="i" name="number" />
				<string array_index="i" name="incoming_trunk" />
				<string array_index="i" name="outgoing_trunk" />
				<string array_index="i" name="pbx_id" />
				<string array_index="i" name="password" />
				<string array_index="i" name="allowed_cid" />
			</for>
		</input>
		<output>
			<integer name="slinkId" />
		</output>
  </function>


*/
