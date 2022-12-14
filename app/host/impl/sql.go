package impl

const (
	InsertResourceSql = `
	INSERT INTO resource (
		id,
		vendor,
		region,
		create_at,
		expire_at,
		type,
		name,
		description,
		status,
		update_at,
		sync_at,
		accout,
		public_ip,
		private_ip
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?);
`
	InsertDescribeSql = `
INSERT INTO host ( resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number )
	VALUES
		( ?,?,?,?,?,?,?,? );
`

	queryHostSQL = `
SELECT *
FROM resource as r
         LEFT JOIN host h ON r.id = h.resource_id
`

	DeleteResourceSql = `delete from resource where id=?`

	DeleteDescribeSql = `delete from host where resource_id=?`

	updateResourceSQL = `UPDATE resource SET vendor=?,region=?,expire_at=?,name=?,description=? WHERE id = ?`
	updateHostSQL     = `UPDATE host SET cpu=?,memory=? WHERE resource_id = ?`
)
