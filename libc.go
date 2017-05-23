package main

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
// 	void **tables,
// 	unsigned int table_length,
// 	void **applications,
// 	unsigned int app_length,
// 	void **risk_management_data,
// 	unsigned int risk_management_length,
// 	void **acquirers,
// 	unsigned int acq_length,
// 	void *user_data);
//
// tms_error_t TMS_EXPORT tms_get_tables(unsigned char *payload, unsigned int length, tms_store_write_cb write_cb, void *user_data) {
//   return gambiarraDaPorra();
// }
import "C"

func main() {}
