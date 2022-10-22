package fn

// rpcf_dealer_get_periodic_link_stats
func x70000068(c conn, p Dict) Dict {
	c.Call(0x70000068)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000068" name="rpcf_dealer_get_periodic_link_stats">
        <input>
        <integer name="slink_id" />
        </input>
        <output>
        <integer name="result" />
        <if condition="ne" value="-1" variable="result">
            <integer name="is_active" />
            <integer name="fee_recalc_start" />
            <integer name="fee_recalc_duration" />
            <integer name="ip_recalc_start" />
            <integer name="ip_recalc_durtation" />
            <integer name="tel_recalc_start" />
            <integer name="tel_recalc_duration" />
            <double name="charged" />
            <double name="repaid" />
        </if>
        </output>
    </function>


*/
