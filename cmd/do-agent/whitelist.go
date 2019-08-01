package main

var k8sWhitelist = map[string]bool{
	"kube_deployment_spec_replicas":               true,
	"kube_deployment_status_replicas_available":   true,
	"kube_deployment_status_replicas_unavailable": true,

	"kube_daemonset_status_desired_number_scheduled": true,
	"kube_daemonset_status_number_available":         true,
	"kube_daemonset_status_number_unavailable":       true,

	"kube_statefulset_replicas":              true,
	"kube_statefulset_status_replicas_ready": true,

	"kube_node_status_allocatable": true,
	"kube_node_status_capacity":    true,
}

var dbaasWhitelist = map[string]bool{

	"postgresql_pg_stat_activity_conn_count":       true,
	"postgresql_pg_stat_database_blks_hit":         true,
	"postgresql_pg_stat_database_blks_read":        true,
	"postgresql_pg_stat_database_deadlocks":        true,
	"postgresql_pg_stat_replication_bytes_diff":    true,
	"postgresql_pg_stat_user_tables_idx_scan":      true,
	"postgresql_pg_stat_user_tables_n_tup_ins":     true,
	"postgresql_pg_stat_user_tables_n_tup_upd":     true,
	"postgresql_pg_stat_user_tables_seq_scan":      true,
	"postgresql_pg_stat_user_tables_n_tup_del":     true,
	"postgresql_pg_stat_user_tables_idx_tup_fetch": true,
	"postgresql_pg_stat_user_tables_seq_tup_read":  true,

	"mysql_threads_created":   true,
	"mysql_threads_connected": true,
	"mysql_threads_running":   true,

	"mysql_handler_read_key":      true,
	"mysql_handler_read_first":    true,
	"mysql_handler_read_next":     true,
	"mysql_handler_read_prev":     true,
	"mysql_handler_read_last":     true,
	"mysql_handler_read_rnd":      true,
	"mysql_handler_read_rnd_next": true,

	"mysql_com_select": true,
	"mysql_com_insert": true,
	"mysql_com_update": true,
	"mysql_com_delete": true,

	"mysql_perf_schema_table_io_waits_total_fetch":  true,
	"mysql_perf_schema_table_io_waits_total_insert": true,
	"mysql_perf_schema_table_io_waits_total_update": true,
	"mysql_perf_schema_table_io_waits_total_delete": true,

	"mysql_perf_schema_table_io_waits_seconds_total_fetch":  true,
	"mysql_perf_schema_table_io_waits_seconds_total_insert": true,
	"mysql_perf_schema_table_io_waits_seconds_total_update": true,
	"mysql_perf_schema_table_io_waits_seconds_total_delete": true,

	"redis_total_connections_received": true,
	"redis_rejected_connections":       true,
	"redis_evicted_keys":               true,
	"redis_keyspace_hits":              true,
	"redis_keyspace_misses":            true,
	"redis_instantaneous_ops_per_sec":  true,
	"redis_used_memory_rss":            true,
	"redis_used_memory":                true,
	"redis_connected_slaves":           true,
	"redis_clients":                    true,
}

var cadvisorWhitelist = map[string]bool{

	"container_cpu_load_average_10s":         true,
	"container_cpu_user_seconds_total":       true,
	"container_fs_writes_bytes_total":        true,
	"container_fs_reads_bytes_total":         true,
	"container_memory_usage_bytes":           true,
	"container_network_receive_bytes_total":  true,
	"container_network_transmit_bytes_total": true,
}
