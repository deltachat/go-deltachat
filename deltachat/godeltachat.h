uintptr_t godeltachat_eventhandler(dc_context_t* context, int event,
                                 uintptr_t data1, uintptr_t data2);

dc_context_t* godeltachat_create_context(char *dbLocation);

void godeltachat_do_imap_routine(dc_context_t* context);

void godeltachat_do_smtp_routine(dc_context_t* context);
