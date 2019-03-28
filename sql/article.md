# article
```
SHOW COLUMNS FROM article;
```
# res
```
"Field";"Type";"Null";"Key";"Default";"Extra"
"id";"bigint(20)";"NO";"PRI";;"auto_increment"
"ticket_id";"bigint(20)";"NO";"MUL";;""
"article_type_id";"smallint(6)";"NO";"MUL";;""
"article_sender_type_id";"smallint(6)";"NO";"MUL";;""
"a_from";"text";"YES";"";;""
"a_reply_to";"text";"YES";"";;""
"a_to";"text";"YES";"";;""
"a_cc";"text";"YES";"";;""
"a_subject";"text";"YES";"";;""
"a_message_id";"text";"YES";"";;""
"a_message_id_md5";"varchar(32)";"YES";"MUL";;""
"a_in_reply_to";"text";"YES";"";;""
"a_references";"text";"YES";"";;""
"a_content_type";"varchar(250)";"YES";"";;""
"a_body";"mediumtext";"NO";"";;""
"incoming_time";"int(11)";"NO";"";;""
"content_path";"varchar(250)";"YES";"";;""
"valid_id";"smallint(6)";"NO";"MUL";;""
"create_time";"datetime";"NO";"";;""
"create_by";"int(11)";"NO";"MUL";;""
"change_time";"datetime";"NO";"";;""
"change_by";"int(11)";"NO";"MUL";;""
```
