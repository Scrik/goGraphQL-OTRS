# article_attachment
```
DESC article_attachment
```

# res
```
"Field";"Type";"Null";"Key";"Default";"Extra"
"id";"bigint(20)";"NO";"PRI";;"auto_increment"
"tn";"varchar(50)";"NO";"UNI";;""
"title";"varchar(255)";"YES";"MUL";;""
"queue_id";"int(11)";"NO";"MUL";;""
"ticket_lock_id";"smallint(6)";"NO";"MUL";;""
"type_id";"smallint(6)";"YES";"MUL";;""
"service_id";"int(11)";"YES";"MUL";;""
"sla_id";"int(11)";"YES";"MUL";;""
"user_id";"int(11)";"NO";"MUL";;""
"responsible_user_id";"int(11)";"NO";"MUL";;""
"ticket_priority_id";"smallint(6)";"NO";"MUL";;""
"ticket_state_id";"smallint(6)";"NO";"MUL";;""
"customer_id";"varchar(150)";"YES";"MUL";;""
"customer_user_id";"varchar(250)";"YES";"MUL";;""
"timeout";"int(11)";"NO";"MUL";;""
"until_time";"int(11)";"NO";"MUL";;""
"escalation_time";"int(11)";"NO";"MUL";;""
"escalation_update_time";"int(11)";"NO";"MUL";;""
"escalation_response_time";"int(11)";"NO";"MUL";;""
"escalation_solution_time";"int(11)";"NO";"MUL";;""
"archive_flag";"smallint(6)";"NO";"MUL";"0";""
"create_time_unix";"bigint(20)";"NO";"MUL";;""
"create_time";"datetime";"NO";"MUL";;""
"create_by";"int(11)";"NO";"MUL";;""
"change_time";"datetime";"NO";"";;""
"change_by";"int(11)";"NO";"MUL";;""
```
