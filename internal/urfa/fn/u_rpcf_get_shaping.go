package fn

// rpcf_get_shaping
func x1200e(c conn, p Dict) Dict {
	c.Call(0x1200e)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1200e" name="rpcf_get_shaping">
        <input>
            <integer name="service_id" />
        </input>
        <output>
            <integer name="result" />
            <if condition="eq" value="0" variable="result">
                <integer name="flags" />


		    <integer name="in_tclass_count" />
		    <for count="in_tclass_count" from="0" name="i">
			<integer array_index="i" name="in_tclass_id_array" />
		    </for>
		    <integer name="in_limits_count" />
		    <for count="in_limits_count" from="0" name="l">
			<integer array_index="l" name="in_acc_states_array" />
			<long array_index="l" name="in_borders_array" />
			<integer array_index="l" name="in_timeranges_array" />
			<integer array_index="l" name="in_limits_array" />
		    </for>


		    <integer name="out_tclass_count" />
		    <for count="out_tclass_count" from="0" name="i">
			<integer array_index="i" name="out_tclass_id_array" />
		    </for>
		    <integer name="out_limits_count" />
		    <for count="out_limits_count" from="0" name="l">
			<integer array_index="l" name="out_acc_states_array" />
			<long array_index="l" name="out_borders_array" />
			<integer array_index="l" name="out_timeranges_array" />
			<integer array_index="l" name="out_limits_array" />
		    </for>

                <integer name="rad_count" />
                <for count="rad_count" from="0" name="i">
                    <integer array_index="i" name="rad_vendor" />
                    <integer array_index="i" name="rad_attr" />
                    <integer array_index="i" name="rad_type" />
                    <string array_index="i" name="rad_value" />
                </for>

                <integer name="settings_cnt" />
                <for count="settings_cnt" from="0" name="i">
                    <integer array_index="i" name="turbo_mode_service_id" />
                    <integer array_index="i" name="turbo_mode_incoming_rate" />
                    <integer array_index="i" name="turbo_mode_outgoing_rate" />
                    <long array_index="i" name="turbo_mode_incoming_limit" />
                    <long array_index="i" name="turbo_mode_outgoing_limit" />
                    <integer array_index="i" name="name" />
                    <integer array_index="i" name="turbo_mode_duration" />
                    <integer array_index="i" name="turbo_mode_acct_period_flag" />
                </for>
            </if>

        </output>
    </function>


*/
