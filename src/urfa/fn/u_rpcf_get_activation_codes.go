package fn

// rpcf_get_activation_codes
func x450b(c conn, p Dict) Dict {
	c.Call(0x450b)
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x450b" name="rpcf_get_activation_codes">
      <input>
        <integer name="account_id" />
      </input>
      <output>
        <integer name="ac_count" />
        <for count="ac_count" from="0" name="i">
            <integer array_index="i" name="access_card_number_array" />
            <integer array_index="i" name="part1_array" />
            <integer array_index="i" name="part2_array" />
            <integer array_index="i" name="part3_array" />
            <integer array_index="i" name="part4_array" />
            <integer array_index="i" name="part5_array" />
            <integer array_index="i" name="part6_array" />
            <integer array_index="i" name="status_array" />
            <integer array_index="i" name="create_date_array" />
            <integer array_index="i" name="modify_date_array" />
            <integer array_index="i" name="delete_date_array" />
        </for>
      </output>
    </function>


*/
