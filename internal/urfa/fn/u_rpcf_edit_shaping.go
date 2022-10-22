package fn

// rpcf_edit_shaping
func x12010(c conn, p Dict) Dict {
	c.Call(0x12010)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x12010" name="rpcf_edit_shaping">
        <input>
            <integer name="service_id" />
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

        </input>
        <output>
            <integer name="result" />
        </output>
    </function>


*/
