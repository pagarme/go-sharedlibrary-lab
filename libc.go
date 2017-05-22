package main

// #cgo CFLAGS: -I/Users/jonathanlima/src/github.com/pagarme/libabecs/include
//
// #include <libabecs/packet.h>
// #include <libabecs/abecs.h>
// #include <libmpos/mpos.h>
//
// #ifdef _WIN32
// #define TMS_EXPORT __declspec(dllexport)
// #else
// #define TMS_EXPORT
// #endif
//
// typedef int tms_error_t;
//
// typedef tms_error_t (*tms_store_write_cb)(
// 	const char *version,
// 	abecs_table_load_record_cmd_t **tables,
// 	unsigned int table_length,
// 	mpos_application_data **applications,
// 	unsigned int app_length,
// 	mpos_risk_management_data **risk_management_data,
// 	unsigned int risk_management_length,
// 	mpos_acquirer_data **acquirers,
// 	unsigned int acq_length,
// 	void *user_data);
//
// tms_error_t TMS_EXPORT tms_get_tables(unsigned char *payload, unsigned int length, tms_store_write_cb write_cb, void *user_data) {
//   return gambiarraDaPorra();
// }
import "C"

func main() {}
