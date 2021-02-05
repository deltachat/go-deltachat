#include <deltachat.h>
#include <stdio.h>

// Context creation because passing a C function as callback value from go does not seeem
// to work
dc_context_t* godeltachat_create_context(char *dbLocation)
{
  return dc_context_new(NULL, dbLocation, NULL);
}
