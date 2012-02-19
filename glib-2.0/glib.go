package glib

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

extern void _GLib_go_callback_cleanup(void *gofunc);
static void _c_callback_cleanup(void *userdata)
{
	_GLib_go_callback_cleanup(userdata);
}
typedef struct _GAllocator GAllocator;
typedef struct _GArray GArray;
typedef uint32_t GAsciiType;
typedef struct _GAsyncQueue GAsyncQueue;
typedef struct _GBookmarkFile GBookmarkFile;
typedef uint32_t GBookmarkFileError;
typedef struct _GByteArray GByteArray;
typedef struct _GCache GCache;
typedef void* GCacheDestroyFunc;
extern void _GCacheDestroyFunc_c_wrapper();
extern void _GCacheDestroyFunc_c_wrapper_once();
typedef struct _GChecksum GChecksum;
typedef uint32_t GChecksumType;
typedef void* GChildWatchFunc;
extern void _GChildWatchFunc_c_wrapper();
extern void _GChildWatchFunc_c_wrapper_once();
typedef void* GCompareDataFunc;
extern void _GCompareDataFunc_c_wrapper();
extern void _GCompareDataFunc_c_wrapper_once();
typedef void* GCompareFunc;
extern void _GCompareFunc_c_wrapper();
extern void _GCompareFunc_c_wrapper_once();
typedef struct _GCompletion GCompletion;
typedef void* GCompletionFunc;
extern void _GCompletionFunc_c_wrapper();
extern void _GCompletionFunc_c_wrapper_once();
typedef void* GCompletionStrncmpFunc;
extern void _GCompletionStrncmpFunc_c_wrapper();
extern void _GCompletionStrncmpFunc_c_wrapper_once();
typedef struct _GCond GCond;
typedef uint32_t GConvertError;
typedef struct _GData GData;
typedef void* GDataForeachFunc;
extern void _GDataForeachFunc_c_wrapper();
extern void _GDataForeachFunc_c_wrapper_once();
typedef struct _GDate GDate;
typedef uint32_t GDateDMY;
typedef uint32_t GDateMonth;
typedef struct _GDateTime GDateTime;
typedef uint32_t GDateWeekday;
typedef struct _GDebugKey GDebugKey;
typedef void* GDestroyNotify;
extern void _GDestroyNotify_c_wrapper();
extern void _GDestroyNotify_c_wrapper_once();
typedef struct _GDir GDir;
typedef struct _GDoubleIEEE754 GDoubleIEEE754;
typedef void* GEqualFunc;
extern void _GEqualFunc_c_wrapper();
extern void _GEqualFunc_c_wrapper_once();
typedef struct _GError GError;
typedef uint32_t GErrorType;
typedef uint32_t GFileError;
typedef uint32_t GFileTest;
typedef struct _GFloatIEEE754 GFloatIEEE754;
typedef uint32_t GFormatSizeFlags;
typedef void* GFreeFunc;
extern void _GFreeFunc_c_wrapper();
extern void _GFreeFunc_c_wrapper_once();
typedef void* GFunc;
extern void _GFunc_c_wrapper();
extern void _GFunc_c_wrapper_once();
typedef void* GHFunc;
extern void _GHFunc_c_wrapper();
extern void _GHFunc_c_wrapper_once();
typedef void* GHRFunc;
extern void _GHRFunc_c_wrapper();
extern void _GHRFunc_c_wrapper_once();
typedef void* GHashFunc;
extern void _GHashFunc_c_wrapper();
extern void _GHashFunc_c_wrapper_once();
typedef struct _GHashTable GHashTable;
typedef struct _GHashTableIter GHashTableIter;
typedef struct _GHmac GHmac;
typedef struct _GHook GHook;
typedef void* GHookCheckFunc;
extern void _GHookCheckFunc_c_wrapper();
extern void _GHookCheckFunc_c_wrapper_once();
typedef void* GHookCheckMarshaller;
extern void _GHookCheckMarshaller_c_wrapper();
extern void _GHookCheckMarshaller_c_wrapper_once();
typedef void* GHookCompareFunc;
extern void _GHookCompareFunc_c_wrapper();
extern void _GHookCompareFunc_c_wrapper_once();
typedef void* GHookFinalizeFunc;
extern void _GHookFinalizeFunc_c_wrapper();
extern void _GHookFinalizeFunc_c_wrapper_once();
typedef void* GHookFindFunc;
extern void _GHookFindFunc_c_wrapper();
extern void _GHookFindFunc_c_wrapper_once();
typedef uint32_t GHookFlagMask;
typedef void* GHookFunc;
extern void _GHookFunc_c_wrapper();
extern void _GHookFunc_c_wrapper_once();
typedef struct _GHookList GHookList;
typedef void* GHookMarshaller;
extern void _GHookMarshaller_c_wrapper();
extern void _GHookMarshaller_c_wrapper_once();
typedef struct _GIConv GIConv;
typedef struct _GIOChannel GIOChannel;
typedef uint32_t GIOChannelError;
typedef uint32_t GIOCondition;
typedef uint32_t GIOError;
typedef uint32_t GIOFlags;
typedef void* GIOFunc;
extern void _GIOFunc_c_wrapper();
extern void _GIOFunc_c_wrapper_once();
typedef struct _GIOFuncs GIOFuncs;
typedef uint32_t GIOStatus;
typedef struct _GKeyFile GKeyFile;
typedef uint32_t GKeyFileError;
typedef uint32_t GKeyFileFlags;
typedef struct _GList GList;
typedef void* GLogFunc;
extern void _GLogFunc_c_wrapper();
extern void _GLogFunc_c_wrapper_once();
typedef int32_t GLogLevelFlags;
typedef struct _GMainContext GMainContext;
typedef struct _GMainLoop GMainLoop;
typedef struct _GMappedFile GMappedFile;
typedef uint32_t GMarkupCollectType;
typedef uint32_t GMarkupError;
typedef struct _GMarkupParseContext GMarkupParseContext;
typedef uint32_t GMarkupParseFlags;
typedef struct _GMarkupParser GMarkupParser;
typedef struct _GMatchInfo GMatchInfo;
typedef struct _GMemChunk GMemChunk;
typedef struct _GMemVTable GMemVTable;
typedef struct _GMutex GMutex;
typedef struct _GNode GNode;
typedef void* GNodeForeachFunc;
extern void _GNodeForeachFunc_c_wrapper();
extern void _GNodeForeachFunc_c_wrapper_once();
typedef void* GNodeTraverseFunc;
extern void _GNodeTraverseFunc_c_wrapper();
extern void _GNodeTraverseFunc_c_wrapper_once();
typedef uint32_t GNormalizeMode;
typedef struct _GOnce GOnce;
typedef uint32_t GOnceStatus;
typedef uint32_t GOptionArg;
typedef void* GOptionArgFunc;
extern void _GOptionArgFunc_c_wrapper();
extern void _GOptionArgFunc_c_wrapper_once();
typedef struct _GOptionContext GOptionContext;
typedef struct _GOptionEntry GOptionEntry;
typedef uint32_t GOptionError;
typedef void* GOptionErrorFunc;
extern void _GOptionErrorFunc_c_wrapper();
extern void _GOptionErrorFunc_c_wrapper_once();
typedef uint32_t GOptionFlags;
typedef struct _GOptionGroup GOptionGroup;
typedef void* GOptionParseFunc;
extern void _GOptionParseFunc_c_wrapper();
extern void _GOptionParseFunc_c_wrapper_once();
typedef struct _GPatternSpec GPatternSpec;
typedef struct _GPollFD GPollFD;
typedef void* GPollFunc;
extern void _GPollFunc_c_wrapper();
extern void _GPollFunc_c_wrapper_once();
typedef void* GPrintFunc;
extern void _GPrintFunc_c_wrapper();
extern void _GPrintFunc_c_wrapper_once();
typedef struct _GPrivate GPrivate;
typedef struct _GPtrArray GPtrArray;
typedef struct _GQueue GQueue;
typedef struct _GRand GRand;
typedef struct _GRegex GRegex;
typedef uint32_t GRegexCompileFlags;
typedef uint32_t GRegexError;
typedef void* GRegexEvalCallback;
extern void _GRegexEvalCallback_c_wrapper();
extern void _GRegexEvalCallback_c_wrapper_once();
typedef uint32_t GRegexMatchFlags;
typedef struct _GRelation GRelation;
typedef struct _GSList GSList;
typedef struct _GScanner GScanner;
typedef struct _GScannerConfig GScannerConfig;
typedef void* GScannerMsgFunc;
extern void _GScannerMsgFunc_c_wrapper();
extern void _GScannerMsgFunc_c_wrapper_once();
typedef uint32_t GSeekType;
typedef struct _GSequence GSequence;
typedef struct _GSequenceIter GSequenceIter;
typedef void* GSequenceIterCompareFunc;
extern void _GSequenceIterCompareFunc_c_wrapper();
extern void _GSequenceIterCompareFunc_c_wrapper_once();
typedef uint32_t GShellError;
typedef uint32_t GSliceConfig;
typedef struct _GSource GSource;
typedef struct _GSourceCallbackFuncs GSourceCallbackFuncs;
typedef void* GSourceDummyMarshal;
extern void _GSourceDummyMarshal_c_wrapper();
extern void _GSourceDummyMarshal_c_wrapper_once();
typedef void* GSourceFunc;
extern void _GSourceFunc_c_wrapper();
extern void _GSourceFunc_c_wrapper_once();
typedef struct _GSourceFuncs GSourceFuncs;
typedef struct _GSourcePrivate GSourcePrivate;
typedef void* GSpawnChildSetupFunc;
extern void _GSpawnChildSetupFunc_c_wrapper();
extern void _GSpawnChildSetupFunc_c_wrapper_once();
typedef uint32_t GSpawnError;
typedef uint32_t GSpawnFlags;
typedef struct _GStatBuf GStatBuf;
typedef struct _GStaticMutex GStaticMutex;
typedef struct _GStaticPrivate GStaticPrivate;
typedef struct _GStaticRWLock GStaticRWLock;
typedef struct _GStaticRecMutex GStaticRecMutex;
typedef struct _GString GString;
typedef struct _GStringChunk GStringChunk;
typedef struct _GSystemThread GSystemThread;
typedef struct _GTestCase GTestCase;
typedef struct _GTestConfig GTestConfig;
typedef void* GTestDataFunc;
extern void _GTestDataFunc_c_wrapper();
extern void _GTestDataFunc_c_wrapper_once();
typedef void* GTestFixtureFunc;
extern void _GTestFixtureFunc_c_wrapper();
extern void _GTestFixtureFunc_c_wrapper_once();
typedef void* GTestFunc;
extern void _GTestFunc_c_wrapper();
extern void _GTestFunc_c_wrapper_once();
typedef struct _GTestLogBuffer GTestLogBuffer;
typedef void* GTestLogFatalFunc;
extern void _GTestLogFatalFunc_c_wrapper();
extern void _GTestLogFatalFunc_c_wrapper_once();
typedef struct _GTestLogMsg GTestLogMsg;
typedef uint32_t GTestLogType;
typedef struct _GTestSuite GTestSuite;
typedef uint32_t GTestTrapFlags;
typedef struct _GThread GThread;
typedef uint32_t GThreadError;
typedef struct _GThreadFunctions GThreadFunctions;
typedef struct _GThreadPool GThreadPool;
typedef uint32_t GThreadPriority;
typedef uint32_t GTimeType;
typedef struct _GTimeVal GTimeVal;
typedef struct _GTimeZone GTimeZone;
typedef struct _GTimer GTimer;
typedef uint32_t GTokenType;
typedef struct _GTokenValue GTokenValue;
typedef void* GTranslateFunc;
extern void _GTranslateFunc_c_wrapper();
extern void _GTranslateFunc_c_wrapper_once();
typedef struct _GTrashStack GTrashStack;
typedef uint32_t GTraverseFlags;
typedef void* GTraverseFunc;
extern void _GTraverseFunc_c_wrapper();
extern void _GTraverseFunc_c_wrapper_once();
typedef uint32_t GTraverseType;
typedef struct _GTree GTree;
typedef struct _GTuples GTuples;
typedef uint32_t GUnicodeBreakType;
typedef int32_t GUnicodeScript;
typedef uint32_t GUnicodeType;
typedef uint32_t GUserDirectory;
typedef struct _GVariant GVariant;
typedef struct _GVariantBuilder GVariantBuilder;
typedef uint32_t GVariantClass;
typedef uint32_t GVariantParseError;
typedef struct _GVariantType GVariantType;
typedef void* GVoidFunc;
extern void _GVoidFunc_c_wrapper();
extern void _GVoidFunc_c_wrapper_once();
typedef struct _G_StaticAssert_347 G_StaticAssert_347;
extern void g_allocator_free(GAllocator*);
extern char* g_array_free(GArray*, int);
extern uint32_t g_array_get_element_size(GArray*);
extern void g_array_unref(GArray*);
extern int32_t g_async_queue_length(GAsyncQueue*);
extern int32_t g_async_queue_length_unlocked(GAsyncQueue*);
extern void g_async_queue_lock(GAsyncQueue*);
extern void g_async_queue_push(GAsyncQueue*, void*);
extern void g_async_queue_push_unlocked(GAsyncQueue*, void*);
extern void g_async_queue_ref_unlocked(GAsyncQueue*);
extern void g_async_queue_unlock(GAsyncQueue*);
extern void g_async_queue_unref(GAsyncQueue*);
extern void g_async_queue_unref_and_unlock(GAsyncQueue*);
extern void g_bookmark_file_add_application(GBookmarkFile*, char*, char*, char*);
extern void g_bookmark_file_add_group(GBookmarkFile*, char*, char*);
extern void g_bookmark_file_free(GBookmarkFile*);
extern int64_t g_bookmark_file_get_added(GBookmarkFile*, char*, GError**);
extern int g_bookmark_file_get_app_info(GBookmarkFile*, char*, char*, char*, uint32_t*, int64_t*, GError**);
extern char* g_bookmark_file_get_description(GBookmarkFile*, char*, GError**);
extern int g_bookmark_file_get_icon(GBookmarkFile*, char*, char*, char*, GError**);
extern int g_bookmark_file_get_is_private(GBookmarkFile*, char*, GError**);
extern char* g_bookmark_file_get_mime_type(GBookmarkFile*, char*, GError**);
extern int64_t g_bookmark_file_get_modified(GBookmarkFile*, char*, GError**);
extern int32_t g_bookmark_file_get_size(GBookmarkFile*);
extern char* g_bookmark_file_get_title(GBookmarkFile*, char*, GError**);
extern int64_t g_bookmark_file_get_visited(GBookmarkFile*, char*, GError**);
extern int g_bookmark_file_has_application(GBookmarkFile*, char*, char*, GError**);
extern int g_bookmark_file_has_group(GBookmarkFile*, char*, char*, GError**);
extern int g_bookmark_file_has_item(GBookmarkFile*, char*);
extern int g_bookmark_file_load_from_data(GBookmarkFile*, char*, uint64_t, GError**);
extern int g_bookmark_file_load_from_data_dirs(GBookmarkFile*, char*, char*, GError**);
extern int g_bookmark_file_load_from_file(GBookmarkFile*, char*, GError**);
extern int g_bookmark_file_move_item(GBookmarkFile*, char*, char*, GError**);
extern int g_bookmark_file_remove_application(GBookmarkFile*, char*, char*, GError**);
extern int g_bookmark_file_remove_group(GBookmarkFile*, char*, char*, GError**);
extern int g_bookmark_file_remove_item(GBookmarkFile*, char*, GError**);
extern void g_bookmark_file_set_added(GBookmarkFile*, char*, int64_t);
extern int g_bookmark_file_set_app_info(GBookmarkFile*, char*, char*, char*, int32_t, int64_t, GError**);
extern void g_bookmark_file_set_description(GBookmarkFile*, char*, char*);
extern void g_bookmark_file_set_groups(GBookmarkFile*, char*, char*, uint64_t);
extern void g_bookmark_file_set_icon(GBookmarkFile*, char*, char*, char*);
extern void g_bookmark_file_set_is_private(GBookmarkFile*, char*, int);
extern void g_bookmark_file_set_mime_type(GBookmarkFile*, char*, char*);
extern void g_bookmark_file_set_modified(GBookmarkFile*, char*, int64_t);
extern void g_bookmark_file_set_title(GBookmarkFile*, char*, char*);
extern void g_bookmark_file_set_visited(GBookmarkFile*, char*, int64_t);
extern char* g_bookmark_file_to_data(GBookmarkFile*, uint64_t*, GError**);
extern int g_bookmark_file_to_file(GBookmarkFile*, char*, GError**);
extern uint32_t g_bookmark_file_error_quark();
extern uint8_t* g_byte_array_free(GByteArray*, int);
extern void g_byte_array_unref(GByteArray*);
extern void g_cache_destroy(GCache*);
extern void g_cache_remove(GCache*, void*);
extern void g_checksum_free(GChecksum*);
extern void g_checksum_get_digest(GChecksum*, uint8_t*, uint64_t*);
extern char* g_checksum_get_string(GChecksum*);
extern void g_checksum_reset(GChecksum*);
extern void g_checksum_update(GChecksum*, uint8_t*, int64_t);
extern int64_t g_checksum_type_get_length(GChecksumType);
extern void g_completion_clear_items(GCompletion*);
extern GList* g_completion_complete_utf8(GCompletion*, char*, char*);
extern void g_completion_free(GCompletion*);
extern GDate* g_date_new();
extern GDate* g_date_new_dmy(uint8_t, GDateMonth, uint16_t);
extern GDate* g_date_new_julian(uint32_t);
extern void g_date_add_days(GDate*, uint32_t);
extern void g_date_add_months(GDate*, uint32_t);
extern void g_date_add_years(GDate*, uint32_t);
extern void g_date_clamp(GDate*, GDate*, GDate*);
extern void g_date_clear(GDate*, uint32_t);
extern int32_t g_date_compare(GDate*, GDate*);
extern int32_t g_date_days_between(GDate*, GDate*);
extern void g_date_free(GDate*);
extern uint8_t g_date_get_day(GDate*);
extern uint32_t g_date_get_day_of_year(GDate*);
extern uint32_t g_date_get_iso8601_week_of_year(GDate*);
extern uint32_t g_date_get_julian(GDate*);
extern uint32_t g_date_get_monday_week_of_year(GDate*);
extern GDateMonth g_date_get_month(GDate*);
extern uint32_t g_date_get_sunday_week_of_year(GDate*);
extern GDateWeekday g_date_get_weekday(GDate*);
extern uint16_t g_date_get_year(GDate*);
extern int g_date_is_first_of_month(GDate*);
extern int g_date_is_last_of_month(GDate*);
extern void g_date_order(GDate*, GDate*);
extern void g_date_set_day(GDate*, uint8_t);
extern void g_date_set_dmy(GDate*, uint8_t, GDateMonth, uint16_t);
extern void g_date_set_julian(GDate*, uint32_t);
extern void g_date_set_month(GDate*, GDateMonth);
extern void g_date_set_parse(GDate*, char*);
extern void g_date_set_time(GDate*, int32_t);
extern void g_date_set_time_t(GDate*, int64_t);
extern void g_date_set_time_val(GDate*, GTimeVal*);
extern void g_date_set_year(GDate*, uint16_t);
extern void g_date_subtract_days(GDate*, uint32_t);
extern void g_date_subtract_months(GDate*, uint32_t);
extern void g_date_subtract_years(GDate*, uint32_t);
extern void g_date_to_struct_tm(GDate*, void*);
extern int g_date_valid(GDate*);
extern uint8_t g_date_get_days_in_month(GDateMonth, uint16_t);
extern uint8_t g_date_get_monday_weeks_in_year(uint16_t);
extern uint8_t g_date_get_sunday_weeks_in_year(uint16_t);
extern int g_date_is_leap_year(uint16_t);
extern uint64_t g_date_strftime(char*, uint64_t, char*, GDate*);
extern int g_date_valid_day(uint8_t);
extern int g_date_valid_dmy(uint8_t, GDateMonth, uint16_t);
extern int g_date_valid_julian(uint32_t);
extern int g_date_valid_month(GDateMonth);
extern int g_date_valid_weekday(GDateWeekday);
extern int g_date_valid_year(uint16_t);
extern GDateTime* g_date_time_new(GTimeZone*, int32_t, int32_t, int32_t, int32_t, int32_t, double);
extern GDateTime* g_date_time_new_from_timeval_local(GTimeVal*);
extern GDateTime* g_date_time_new_from_timeval_utc(GTimeVal*);
extern GDateTime* g_date_time_new_from_unix_local(int64_t);
extern GDateTime* g_date_time_new_from_unix_utc(int64_t);
extern GDateTime* g_date_time_new_local(int32_t, int32_t, int32_t, int32_t, int32_t, double);
extern GDateTime* g_date_time_new_now(GTimeZone*);
extern GDateTime* g_date_time_new_now_local();
extern GDateTime* g_date_time_new_now_utc();
extern GDateTime* g_date_time_new_utc(int32_t, int32_t, int32_t, int32_t, int32_t, double);
extern GDateTime* g_date_time_add(GDateTime*, int64_t);
extern GDateTime* g_date_time_add_days(GDateTime*, int32_t);
extern GDateTime* g_date_time_add_full(GDateTime*, int32_t, int32_t, int32_t, int32_t, int32_t, double);
extern GDateTime* g_date_time_add_hours(GDateTime*, int32_t);
extern GDateTime* g_date_time_add_minutes(GDateTime*, int32_t);
extern GDateTime* g_date_time_add_months(GDateTime*, int32_t);
extern GDateTime* g_date_time_add_seconds(GDateTime*, double);
extern GDateTime* g_date_time_add_weeks(GDateTime*, int32_t);
extern GDateTime* g_date_time_add_years(GDateTime*, int32_t);
extern int64_t g_date_time_difference(GDateTime*, GDateTime*);
extern char* g_date_time_format(GDateTime*, char*);
extern int32_t g_date_time_get_day_of_month(GDateTime*);
extern int32_t g_date_time_get_day_of_week(GDateTime*);
extern int32_t g_date_time_get_day_of_year(GDateTime*);
extern int32_t g_date_time_get_hour(GDateTime*);
extern int32_t g_date_time_get_microsecond(GDateTime*);
extern int32_t g_date_time_get_minute(GDateTime*);
extern int32_t g_date_time_get_month(GDateTime*);
extern int32_t g_date_time_get_second(GDateTime*);
extern double g_date_time_get_seconds(GDateTime*);
extern char* g_date_time_get_timezone_abbreviation(GDateTime*);
extern int64_t g_date_time_get_utc_offset(GDateTime*);
extern int32_t g_date_time_get_week_numbering_year(GDateTime*);
extern int32_t g_date_time_get_week_of_year(GDateTime*);
extern int32_t g_date_time_get_year(GDateTime*);
extern void g_date_time_get_ymd(GDateTime*, int32_t*, int32_t*, int32_t*);
extern int g_date_time_is_daylight_savings(GDateTime*);
extern GDateTime* g_date_time_ref(GDateTime*);
extern GDateTime* g_date_time_to_local(GDateTime*);
extern int g_date_time_to_timeval(GDateTime*, GTimeVal*);
extern GDateTime* g_date_time_to_timezone(GDateTime*, GTimeZone*);
extern int64_t g_date_time_to_unix(GDateTime*);
extern GDateTime* g_date_time_to_utc(GDateTime*);
extern void g_date_time_unref(GDateTime*);
extern int32_t g_date_time_compare(void*, void*);
extern int g_date_time_equal(void*, void*);
extern uint32_t g_date_time_hash(void*);
extern void g_dir_close(GDir*);
extern char* g_dir_read_name(GDir*);
extern void g_dir_rewind(GDir*);
extern char* g_dir_make_tmp(char*, GError**);
extern GError* g_error_new_literal(uint32_t, int32_t, char*);
extern GError* g_error_copy(GError*);
extern void g_error_free(GError*);
extern int g_error_matches(GError*, uint32_t, int32_t);
extern void g_hash_table_destroy(GHashTable*);
extern void g_hash_table_insert(GHashTable*, void*, void*);
extern int g_hash_table_lookup_extended(GHashTable*, void*, void*, void*);
extern int g_hash_table_remove(GHashTable*, void*);
extern void g_hash_table_remove_all(GHashTable*);
extern void g_hash_table_replace(GHashTable*, void*, void*);
extern uint32_t g_hash_table_size(GHashTable*);
extern int g_hash_table_steal(GHashTable*, void*);
extern void g_hash_table_steal_all(GHashTable*);
extern void g_hash_table_unref(GHashTable*);
extern void g_hash_table_iter_init(GHashTableIter*, GHashTable*);
extern int g_hash_table_iter_next(GHashTableIter*, void*, void*);
extern void g_hash_table_iter_remove(GHashTableIter*);
extern void g_hash_table_iter_replace(GHashTableIter*, void*);
extern void g_hash_table_iter_steal(GHashTableIter*);
extern void g_hmac_get_digest(GHmac*, uint8_t*, uint64_t*);
extern char* g_hmac_get_string(GHmac*);
extern void g_hmac_unref(GHmac*);
extern void g_hmac_update(GHmac*, uint8_t*, int64_t);
extern int32_t g_hook_compare_ids(GHook*, GHook*);
extern int g_hook_destroy(GHookList*, uint64_t);
extern void g_hook_destroy_link(GHookList*, GHook*);
extern void g_hook_free(GHookList*, GHook*);
extern void g_hook_insert_before(GHookList*, GHook*, GHook*);
extern void g_hook_prepend(GHookList*, GHook*);
extern void g_hook_unref(GHookList*, GHook*);
extern void g_hook_list_clear(GHookList*);
extern void g_hook_list_init(GHookList*, uint32_t);
extern void g_hook_list_invoke(GHookList*, int);
extern void g_hook_list_invoke_check(GHookList*, int);
extern uint64_t g_iconv(GIConv*, char*, uint64_t*, char*, uint64_t*);
extern int32_t g_iconv_close(GIConv*);
extern GIOChannel* g_io_channel_new_file(char*, char*, GError**);
extern GIOChannel* g_io_channel_unix_new(int32_t);
extern void g_io_channel_close(GIOChannel*);
extern GIOStatus g_io_channel_flush(GIOChannel*, GError**);
extern GIOCondition g_io_channel_get_buffer_condition(GIOChannel*);
extern uint64_t g_io_channel_get_buffer_size(GIOChannel*);
extern int g_io_channel_get_buffered(GIOChannel*);
extern int g_io_channel_get_close_on_unref(GIOChannel*);
extern char* g_io_channel_get_encoding(GIOChannel*);
extern GIOFlags g_io_channel_get_flags(GIOChannel*);
extern char* g_io_channel_get_line_term(GIOChannel*, int32_t*);
extern void g_io_channel_init(GIOChannel*);
extern GIOError g_io_channel_read(GIOChannel*, char*, uint64_t, uint64_t*);
extern GIOStatus g_io_channel_read_chars(GIOChannel*, char*, uint64_t, uint64_t*, GError**);
extern GIOStatus g_io_channel_read_line(GIOChannel*, char*, uint64_t*, uint64_t*, GError**);
extern GIOStatus g_io_channel_read_line_string(GIOChannel*, GString*, uint64_t*, GError**);
extern GIOStatus g_io_channel_read_to_end(GIOChannel*, char*, uint64_t*, GError**);
extern GIOStatus g_io_channel_read_unichar(GIOChannel*, uint32_t*, GError**);
extern GIOChannel* g_io_channel_ref(GIOChannel*);
extern GIOError g_io_channel_seek(GIOChannel*, int64_t, GSeekType);
extern GIOStatus g_io_channel_seek_position(GIOChannel*, int64_t, GSeekType, GError**);
extern void g_io_channel_set_buffer_size(GIOChannel*, uint64_t);
extern void g_io_channel_set_buffered(GIOChannel*, int);
extern void g_io_channel_set_close_on_unref(GIOChannel*, int);
extern GIOStatus g_io_channel_set_encoding(GIOChannel*, char*, GError**);
extern GIOStatus g_io_channel_set_flags(GIOChannel*, GIOFlags, GError**);
extern void g_io_channel_set_line_term(GIOChannel*, char*, int32_t);
extern GIOStatus g_io_channel_shutdown(GIOChannel*, int, GError**);
extern int32_t g_io_channel_unix_get_fd(GIOChannel*);
extern void g_io_channel_unref(GIOChannel*);
extern GIOError g_io_channel_write(GIOChannel*, char*, uint64_t, uint64_t*);
extern GIOStatus g_io_channel_write_chars(GIOChannel*, char*, int64_t, uint64_t*, GError**);
extern GIOStatus g_io_channel_write_unichar(GIOChannel*, uint32_t, GError**);
extern GIOChannelError g_io_channel_error_from_errno(int32_t);
extern uint32_t g_io_channel_error_quark();
extern void g_key_file_free(GKeyFile*);
extern int g_key_file_get_boolean(GKeyFile*, char*, char*, GError**);
extern int* g_key_file_get_boolean_list(GKeyFile*, char*, char*, uint64_t*, GError**);
extern char* g_key_file_get_comment(GKeyFile*, char*, char*, GError**);
extern double g_key_file_get_double(GKeyFile*, char*, char*, GError**);
extern double* g_key_file_get_double_list(GKeyFile*, char*, char*, uint64_t*, GError**);
extern int64_t g_key_file_get_int64(GKeyFile*, char*, char*, GError**);
extern int32_t g_key_file_get_integer(GKeyFile*, char*, char*, GError**);
extern int32_t* g_key_file_get_integer_list(GKeyFile*, char*, char*, uint64_t*, GError**);
extern char* g_key_file_get_locale_string(GKeyFile*, char*, char*, char*, GError**);
extern char** g_key_file_get_locale_string_list(GKeyFile*, char*, char*, char*, uint64_t*, GError**);
extern char* g_key_file_get_start_group(GKeyFile*);
extern char* g_key_file_get_string(GKeyFile*, char*, char*, GError**);
extern char** g_key_file_get_string_list(GKeyFile*, char*, char*, uint64_t*, GError**);
extern uint64_t g_key_file_get_uint64(GKeyFile*, char*, char*, GError**);
extern char* g_key_file_get_value(GKeyFile*, char*, char*, GError**);
extern int g_key_file_has_group(GKeyFile*, char*);
extern int g_key_file_load_from_data(GKeyFile*, char*, uint64_t, GKeyFileFlags, GError**);
extern int g_key_file_load_from_data_dirs(GKeyFile*, char*, char*, GKeyFileFlags, GError**);
extern int g_key_file_load_from_dirs(GKeyFile*, char*, char*, char*, GKeyFileFlags, GError**);
extern int g_key_file_load_from_file(GKeyFile*, char*, GKeyFileFlags, GError**);
extern int g_key_file_remove_comment(GKeyFile*, char*, char*, GError**);
extern int g_key_file_remove_group(GKeyFile*, char*, GError**);
extern int g_key_file_remove_key(GKeyFile*, char*, char*, GError**);
extern void g_key_file_set_boolean(GKeyFile*, char*, char*, int);
extern void g_key_file_set_boolean_list(GKeyFile*, char*, char*, int, uint64_t);
extern int g_key_file_set_comment(GKeyFile*, char*, char*, char*, GError**);
extern void g_key_file_set_double(GKeyFile*, char*, char*, double);
extern void g_key_file_set_double_list(GKeyFile*, char*, char*, double, uint64_t);
extern void g_key_file_set_int64(GKeyFile*, char*, char*, int64_t);
extern void g_key_file_set_integer(GKeyFile*, char*, char*, int32_t);
extern void g_key_file_set_integer_list(GKeyFile*, char*, char*, int32_t, uint64_t);
extern void g_key_file_set_list_separator(GKeyFile*, uint8_t);
extern void g_key_file_set_locale_string(GKeyFile*, char*, char*, char*, char*);
extern void g_key_file_set_locale_string_list(GKeyFile*, char*, char*, char*, char*, uint64_t);
extern void g_key_file_set_string(GKeyFile*, char*, char*, char*);
extern void g_key_file_set_string_list(GKeyFile*, char*, char*, char**, uint64_t);
extern void g_key_file_set_uint64(GKeyFile*, char*, char*, uint64_t);
extern void g_key_file_set_value(GKeyFile*, char*, char*, char*);
extern char* g_key_file_to_data(GKeyFile*, uint64_t*, GError**);
extern uint32_t g_key_file_error_quark();
extern void g_list_pop_allocator();
extern void g_list_push_allocator(void*);
extern GMainContext* g_main_context_new();
extern int g_main_context_acquire(GMainContext*);
extern void g_main_context_add_poll(GMainContext*, GPollFD*, int32_t);
extern int32_t g_main_context_check(GMainContext*, int32_t, GPollFD*, int32_t);
extern void g_main_context_dispatch(GMainContext*);
extern GSource* g_main_context_find_source_by_funcs_user_data(GMainContext*, GSourceFuncs*, void*);
extern GSource* g_main_context_find_source_by_id(GMainContext*, uint32_t);
extern GSource* g_main_context_find_source_by_user_data(GMainContext*, void*);
extern void g_main_context_invoke_full(GMainContext*, int32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_main_context_invoke_full(GMainContext* this, int32_t arg0, void* gofunc) {
	if (gofunc) {
		g_main_context_invoke_full(this, arg0, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		g_main_context_invoke_full(this, arg0, 0, 0, 0);
	}
}
extern int g_main_context_is_owner(GMainContext*);
extern int g_main_context_iteration(GMainContext*, int);
extern int g_main_context_pending(GMainContext*);
extern void g_main_context_pop_thread_default(GMainContext*);
extern int g_main_context_prepare(GMainContext*, int32_t*);
extern void g_main_context_push_thread_default(GMainContext*);
extern int32_t g_main_context_query(GMainContext*, int32_t, int32_t*, GPollFD**, int32_t*);
extern GMainContext* g_main_context_ref(GMainContext*);
extern void g_main_context_release(GMainContext*);
extern void g_main_context_remove_poll(GMainContext*, GPollFD*);
extern void g_main_context_unref(GMainContext*);
extern int g_main_context_wait(GMainContext*, GCond*, GMutex*);
extern void g_main_context_wakeup(GMainContext*);
extern GMainContext* g_main_context_default();
extern GMainContext* g_main_context_get_thread_default();
extern GMainLoop* g_main_loop_new(GMainContext*, int);
extern GMainContext* g_main_loop_get_context(GMainLoop*);
extern int g_main_loop_is_running(GMainLoop*);
extern void g_main_loop_quit(GMainLoop*);
extern GMainLoop* g_main_loop_ref(GMainLoop*);
extern void g_main_loop_run(GMainLoop*);
extern void g_main_loop_unref(GMainLoop*);
extern void g_mapped_file_free(GMappedFile*);
extern char* g_mapped_file_get_contents(GMappedFile*);
extern uint64_t g_mapped_file_get_length(GMappedFile*);
extern void g_mapped_file_unref(GMappedFile*);
extern int g_markup_parse_context_end_parse(GMarkupParseContext*, GError**);
extern void g_markup_parse_context_free(GMarkupParseContext*);
extern char* g_markup_parse_context_get_element(GMarkupParseContext*);
extern void g_markup_parse_context_get_position(GMarkupParseContext*, int32_t*, int32_t*);
extern int g_markup_parse_context_parse(GMarkupParseContext*, char*, int64_t, GError**);
extern void g_markup_parse_context_push(GMarkupParseContext*, GMarkupParser*, void*);
extern char* g_match_info_expand_references(GMatchInfo*, char*, GError**);
extern char* g_match_info_fetch(GMatchInfo*, int32_t);
extern char* g_match_info_fetch_named(GMatchInfo*, char*);
extern int g_match_info_fetch_named_pos(GMatchInfo*, char*, int32_t*, int32_t*);
extern int g_match_info_fetch_pos(GMatchInfo*, int32_t, int32_t*, int32_t*);
extern void g_match_info_free(GMatchInfo*);
extern int32_t g_match_info_get_match_count(GMatchInfo*);
extern GRegex* g_match_info_get_regex(GMatchInfo*);
extern char* g_match_info_get_string(GMatchInfo*);
extern int g_match_info_is_partial_match(GMatchInfo*);
extern int g_match_info_matches(GMatchInfo*);
extern int g_match_info_next(GMatchInfo*, GError**);
extern GMatchInfo* g_match_info_ref(GMatchInfo*);
extern void g_match_info_unref(GMatchInfo*);
extern void g_mem_chunk_clean(GMemChunk*);
extern void g_mem_chunk_destroy(GMemChunk*);
extern void g_mem_chunk_free(GMemChunk*, void*);
extern void g_mem_chunk_print(GMemChunk*);
extern void g_mem_chunk_reset(GMemChunk*);
extern void g_mem_chunk_info();
extern int32_t g_node_child_index(GNode*, void*);
extern int32_t g_node_child_position(GNode*, GNode*);
extern uint32_t g_node_depth(GNode*);
extern void g_node_destroy(GNode*);
extern int g_node_is_ancestor(GNode*, GNode*);
extern uint32_t g_node_max_height(GNode*);
extern uint32_t g_node_n_children(GNode*);
extern uint32_t g_node_n_nodes(GNode*, GTraverseFlags);
extern void g_node_reverse_children(GNode*);
extern void g_node_unlink(GNode*);
extern void g_node_pop_allocator();
extern void g_node_push_allocator(void*);
extern int g_once_init_enter(uint64_t*);
extern int g_once_init_enter_impl(uint64_t*);
extern void g_once_init_leave(uint64_t*, uint64_t);
extern void g_option_context_add_group(GOptionContext*, GOptionGroup*);
extern void g_option_context_add_main_entries(GOptionContext*, GOptionEntry*, char*);
extern void g_option_context_free(GOptionContext*);
extern char* g_option_context_get_description(GOptionContext*);
extern char* g_option_context_get_help(GOptionContext*, int, GOptionGroup*);
extern int g_option_context_get_help_enabled(GOptionContext*);
extern int g_option_context_get_ignore_unknown_options(GOptionContext*);
extern char* g_option_context_get_summary(GOptionContext*);
extern int g_option_context_parse(GOptionContext*, int32_t*, char***, GError**);
extern void g_option_context_set_description(GOptionContext*, char*);
extern void g_option_context_set_help_enabled(GOptionContext*, int);
extern void g_option_context_set_ignore_unknown_options(GOptionContext*, int);
extern void g_option_context_set_main_group(GOptionContext*, GOptionGroup*);
extern void g_option_context_set_summary(GOptionContext*, char*);
extern void g_option_context_set_translate_func(GOptionContext*, GTranslateFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_option_context_set_translate_func(GOptionContext* this, void* gofunc) {
	if (gofunc) {
		g_option_context_set_translate_func(this, _GTranslateFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		g_option_context_set_translate_func(this, 0, 0, 0);
	}
}
extern void g_option_context_set_translation_domain(GOptionContext*, char*);
extern void g_option_group_add_entries(GOptionGroup*, GOptionEntry*);
extern void g_option_group_free(GOptionGroup*);
extern void g_option_group_set_translate_func(GOptionGroup*, GTranslateFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_option_group_set_translate_func(GOptionGroup* this, void* gofunc) {
	if (gofunc) {
		g_option_group_set_translate_func(this, _GTranslateFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		g_option_group_set_translate_func(this, 0, 0, 0);
	}
}
extern void g_option_group_set_translation_domain(GOptionGroup*, char*);
extern int g_pattern_spec_equal(GPatternSpec*, GPatternSpec*);
extern void g_pattern_spec_free(GPatternSpec*);
extern void g_ptr_array_add(GPtrArray*, void*);
extern int g_ptr_array_remove(GPtrArray*, void*);
extern int g_ptr_array_remove_fast(GPtrArray*, void*);
extern void g_ptr_array_remove_range(GPtrArray*, uint32_t, uint32_t);
extern void g_ptr_array_set_free_func(GPtrArray*, GDestroyNotify);
extern void g_ptr_array_set_size(GPtrArray*, int32_t);
extern void g_ptr_array_unref(GPtrArray*);
extern void g_queue_clear(GQueue*);
extern void g_queue_free(GQueue*);
extern uint32_t g_queue_get_length(GQueue*);
extern int32_t g_queue_index(GQueue*, void*);
extern void g_queue_init(GQueue*);
extern int g_queue_is_empty(GQueue*);
extern void g_queue_push_head(GQueue*, void*);
extern void g_queue_push_nth(GQueue*, void*, int32_t);
extern void g_queue_push_tail(GQueue*, void*);
extern int g_queue_remove(GQueue*, void*);
extern uint32_t g_queue_remove_all(GQueue*, void*);
extern void g_queue_reverse(GQueue*);
extern double g_rand_double(GRand*);
extern double g_rand_double_range(GRand*, double, double);
extern void g_rand_free(GRand*);
extern uint32_t g_rand_int(GRand*);
extern int32_t g_rand_int_range(GRand*, int32_t, int32_t);
extern void g_rand_set_seed(GRand*, uint32_t);
extern void g_rand_set_seed_array(GRand*, uint32_t*, uint32_t);
extern GRegex* g_regex_new(char*, GRegexCompileFlags, GRegexMatchFlags, GError**);
extern int32_t g_regex_get_capture_count(GRegex*);
extern GRegexCompileFlags g_regex_get_compile_flags(GRegex*);
extern GRegexMatchFlags g_regex_get_match_flags(GRegex*);
extern int32_t g_regex_get_max_backref(GRegex*);
extern char* g_regex_get_pattern(GRegex*);
extern int32_t g_regex_get_string_number(GRegex*, char*);
extern int g_regex_match(GRegex*, char*, GRegexMatchFlags, GMatchInfo**);
extern int g_regex_match_all(GRegex*, char*, GRegexMatchFlags, GMatchInfo**);
extern int g_regex_match_all_full(GRegex*, char**, int64_t, int32_t, GRegexMatchFlags, GMatchInfo**, GError**);
extern int g_regex_match_full(GRegex*, char**, int64_t, int32_t, GRegexMatchFlags, GMatchInfo**, GError**);
extern GRegex* g_regex_ref(GRegex*);
extern char* g_regex_replace(GRegex*, char**, int64_t, int32_t, char*, GRegexMatchFlags, GError**);
extern char* g_regex_replace_literal(GRegex*, char**, int64_t, int32_t, char*, GRegexMatchFlags, GError**);
extern void g_regex_unref(GRegex*);
extern int g_regex_check_replacement(char*, int*, GError**);
extern uint32_t g_regex_error_quark();
extern char* g_regex_escape_nul(char*, int32_t);
extern char* g_regex_escape_string(char**, int32_t);
extern int g_regex_match_simple(char*, char*, GRegexCompileFlags, GRegexMatchFlags);
extern int32_t g_relation_count(GRelation*, void*, int32_t);
extern int32_t g_relation_delete(GRelation*, void*, int32_t);
extern void g_relation_destroy(GRelation*);
extern void g_relation_print(GRelation*);
extern void g_slist_pop_allocator();
extern void g_slist_push_allocator(void*);
extern uint32_t g_scanner_cur_line(GScanner*);
extern uint32_t g_scanner_cur_position(GScanner*);
extern GTokenType g_scanner_cur_token(GScanner*);
extern void g_scanner_destroy(GScanner*);
extern int g_scanner_eof(GScanner*);
extern GTokenType g_scanner_get_next_token(GScanner*);
extern void g_scanner_input_file(GScanner*, int32_t);
extern void g_scanner_input_text(GScanner*, char*, uint32_t);
extern GTokenType g_scanner_peek_next_token(GScanner*);
extern void g_scanner_scope_add_symbol(GScanner*, uint32_t, char*, void*);
extern void g_scanner_scope_remove_symbol(GScanner*, uint32_t, char*);
extern uint32_t g_scanner_set_scope(GScanner*, uint32_t);
extern void g_scanner_sync_file_offset(GScanner*);
extern void g_scanner_unexp_token(GScanner*, GTokenType, char*, char*, char*, char*, int32_t);
extern void g_sequence_free(GSequence*);
extern int32_t g_sequence_get_length(GSequence*);
extern void g_sequence_move(GSequenceIter*, GSequenceIter*);
extern void g_sequence_move_range(GSequenceIter*, GSequenceIter*, GSequenceIter*);
extern void g_sequence_remove(GSequenceIter*);
extern void g_sequence_remove_range(GSequenceIter*, GSequenceIter*);
extern void g_sequence_set(GSequenceIter*, void*);
extern void g_sequence_swap(GSequenceIter*, GSequenceIter*);
extern int32_t g_sequence_iter_compare(GSequenceIter*, GSequenceIter*);
extern int32_t g_sequence_iter_get_position(GSequenceIter*);
extern int g_sequence_iter_is_begin(GSequenceIter*);
extern int g_sequence_iter_is_end(GSequenceIter*);
extern GSource* g_source_new(GSourceFuncs*, uint32_t);
extern void g_source_add_child_source(GSource*, GSource*);
extern void g_source_add_poll(GSource*, GPollFD*);
extern uint32_t g_source_attach(GSource*, GMainContext*);
extern void g_source_destroy(GSource*);
extern int g_source_get_can_recurse(GSource*);
extern GMainContext* g_source_get_context(GSource*);
extern void g_source_get_current_time(GSource*, GTimeVal*);
extern uint32_t g_source_get_id(GSource*);
extern char* g_source_get_name(GSource*);
extern int32_t g_source_get_priority(GSource*);
extern int64_t g_source_get_time(GSource*);
extern int g_source_is_destroyed(GSource*);
extern GSource* g_source_ref(GSource*);
extern void g_source_remove_child_source(GSource*, GSource*);
extern void g_source_remove_poll(GSource*, GPollFD*);
extern void g_source_set_callback(GSource*, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_source_set_callback(GSource* this, void* gofunc) {
	if (gofunc) {
		g_source_set_callback(this, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		g_source_set_callback(this, 0, 0, 0);
	}
}
extern void g_source_set_callback_indirect(GSource*, void*, GSourceCallbackFuncs*);
extern void g_source_set_can_recurse(GSource*, int);
extern void g_source_set_funcs(GSource*, GSourceFuncs*);
extern void g_source_set_name(GSource*, char*);
extern void g_source_set_priority(GSource*, int32_t);
extern void g_source_unref(GSource*);
extern int g_source_remove(uint32_t);
extern int g_source_remove_by_funcs_user_data(GSourceFuncs*, void*);
extern int g_source_remove_by_user_data(void*);
extern void g_source_set_name_by_id(uint32_t, char*);
extern void g_static_mutex_free(GStaticMutex*);
extern void g_static_mutex_init(GStaticMutex*);
extern void g_static_private_free(GStaticPrivate*);
extern void g_static_private_init(GStaticPrivate*);
extern void g_static_private_set(GStaticPrivate*, void*, GDestroyNotify);
extern void g_static_rw_lock_free(GStaticRWLock*);
extern void g_static_rw_lock_init(GStaticRWLock*);
extern void g_static_rw_lock_reader_lock(GStaticRWLock*);
extern int g_static_rw_lock_reader_trylock(GStaticRWLock*);
extern void g_static_rw_lock_reader_unlock(GStaticRWLock*);
extern void g_static_rw_lock_writer_lock(GStaticRWLock*);
extern int g_static_rw_lock_writer_trylock(GStaticRWLock*);
extern void g_static_rw_lock_writer_unlock(GStaticRWLock*);
extern void g_static_rec_mutex_free(GStaticRecMutex*);
extern void g_static_rec_mutex_init(GStaticRecMutex*);
extern void g_static_rec_mutex_lock(GStaticRecMutex*);
extern void g_static_rec_mutex_lock_full(GStaticRecMutex*, uint32_t);
extern int g_static_rec_mutex_trylock(GStaticRecMutex*);
extern void g_static_rec_mutex_unlock(GStaticRecMutex*);
extern uint32_t g_static_rec_mutex_unlock_full(GStaticRecMutex*);
extern GString* g_string_append(GString*, char*);
extern GString* g_string_append_c(GString*, uint8_t);
extern GString* g_string_append_len(GString*, char*, int64_t);
extern GString* g_string_append_unichar(GString*, uint32_t);
extern GString* g_string_append_uri_escaped(GString*, char*, char*, int);
extern GString* g_string_ascii_down(GString*);
extern GString* g_string_ascii_up(GString*);
extern GString* g_string_assign(GString*, char*);
extern GString* g_string_down(GString*);
extern int g_string_equal(GString*, GString*);
extern GString* g_string_erase(GString*, int64_t, int64_t);
extern char* g_string_free(GString*, int);
extern uint32_t g_string_hash(GString*);
extern GString* g_string_insert(GString*, int64_t, char*);
extern GString* g_string_insert_c(GString*, int64_t, uint8_t);
extern GString* g_string_insert_len(GString*, int64_t, char*, int64_t);
extern GString* g_string_insert_unichar(GString*, int64_t, uint32_t);
extern GString* g_string_overwrite(GString*, uint64_t, char*);
extern GString* g_string_overwrite_len(GString*, uint64_t, char*, int64_t);
extern GString* g_string_prepend(GString*, char*);
extern GString* g_string_prepend_c(GString*, uint8_t);
extern GString* g_string_prepend_len(GString*, char*, int64_t);
extern GString* g_string_prepend_unichar(GString*, uint32_t);
extern GString* g_string_set_size(GString*, uint64_t);
extern GString* g_string_truncate(GString*, uint64_t);
extern GString* g_string_up(GString*);
extern void g_string_chunk_clear(GStringChunk*);
extern void g_string_chunk_free(GStringChunk*);
extern char* g_string_chunk_insert(GStringChunk*, char*);
extern char* g_string_chunk_insert_const(GStringChunk*, char*);
extern char* g_string_chunk_insert_len(GStringChunk*, char*, int64_t);
extern void g_test_log_buffer_free(GTestLogBuffer*);
extern void g_test_log_buffer_push(GTestLogBuffer*, uint32_t, uint8_t*);
extern void g_test_log_msg_free(GTestLogMsg*);
extern void g_test_suite_add(GTestSuite*, GTestCase*);
extern void g_test_suite_add_suite(GTestSuite*, GTestSuite*);
extern void g_thread_set_priority(GThread*, GThreadPriority);
extern uint32_t g_thread_error_quark();
extern void g_thread_exit(void*);
extern int g_thread_get_initialized();
extern void g_thread_init(GThreadFunctions*);
extern void g_thread_init_with_errorcheck_mutexes(GThreadFunctions*);
extern void g_thread_pool_free(GThreadPool*, int, int);
extern int32_t g_thread_pool_get_max_threads(GThreadPool*);
extern uint32_t g_thread_pool_get_num_threads(GThreadPool*);
extern void g_thread_pool_push(GThreadPool*, void*, GError**);
extern void g_thread_pool_set_max_threads(GThreadPool*, int32_t, GError**);
extern uint32_t g_thread_pool_unprocessed(GThreadPool*);
extern uint32_t g_thread_pool_get_max_idle_time();
extern int32_t g_thread_pool_get_max_unused_threads();
extern uint32_t g_thread_pool_get_num_unused_threads();
extern void g_thread_pool_set_max_idle_time(uint32_t);
extern void g_thread_pool_set_max_unused_threads(int32_t);
extern void g_thread_pool_stop_unused_threads();
extern void g_time_val_add(GTimeVal*, int64_t);
extern char* g_time_val_to_iso8601(GTimeVal*);
extern int g_time_val_from_iso8601(char*, GTimeVal*);
extern int32_t g_time_zone_adjust_time(GTimeZone*, GTimeType, int64_t*);
extern int32_t g_time_zone_find_interval(GTimeZone*, GTimeType, int64_t);
extern char* g_time_zone_get_abbreviation(GTimeZone*, int32_t);
extern int32_t g_time_zone_get_offset(GTimeZone*, int32_t);
extern int g_time_zone_is_dst(GTimeZone*, int32_t);
extern void g_time_zone_unref(GTimeZone*);
extern void g_timer_continue(GTimer*);
extern void g_timer_destroy(GTimer*);
extern double g_timer_elapsed(GTimer*, uint64_t*);
extern void g_timer_reset(GTimer*);
extern void g_timer_start(GTimer*);
extern void g_timer_stop(GTimer*);
extern uint32_t g_trash_stack_height(GTrashStack*);
extern void g_trash_stack_push(GTrashStack*, void*);
extern void g_tree_destroy(GTree*);
extern int32_t g_tree_height(GTree*);
extern void g_tree_insert(GTree*, void*, void*);
extern int g_tree_lookup_extended(GTree*, void*, void*, void*);
extern int32_t g_tree_nnodes(GTree*);
extern int g_tree_remove(GTree*, void*);
extern void g_tree_replace(GTree*, void*, void*);
extern int g_tree_steal(GTree*, void*);
extern void g_tree_unref(GTree*);
extern void g_tuples_destroy(GTuples*);
extern GVariant* g_variant_new_array(GVariantType*, GVariant**, uint64_t);
extern GVariant* g_variant_new_boolean(int);
extern GVariant* g_variant_new_byte(uint8_t);
extern GVariant* g_variant_new_bytestring(uint8_t*);
extern GVariant* g_variant_new_bytestring_array(char**, int64_t);
extern GVariant* g_variant_new_dict_entry(GVariant*, GVariant*);
extern GVariant* g_variant_new_double(double);
extern GVariant* g_variant_new_from_data(GVariantType*, uint8_t*, uint64_t, int, GDestroyNotify, void*);
extern GVariant* g_variant_new_handle(int32_t);
extern GVariant* g_variant_new_int16(int16_t);
extern GVariant* g_variant_new_int32(int32_t);
extern GVariant* g_variant_new_int64(int64_t);
extern GVariant* g_variant_new_maybe(GVariantType*, GVariant*);
extern GVariant* g_variant_new_object_path(char*);
extern GVariant* g_variant_new_objv(char**, int64_t);
extern GVariant* g_variant_new_signature(char*);
extern GVariant* g_variant_new_string(char*);
extern GVariant* g_variant_new_strv(char**, int64_t);
extern GVariant* g_variant_new_tuple(GVariant**, uint64_t);
extern GVariant* g_variant_new_uint16(uint16_t);
extern GVariant* g_variant_new_uint32(uint32_t);
extern GVariant* g_variant_new_uint64(uint64_t);
extern GVariant* g_variant_new_variant(GVariant*);
extern GVariant* g_variant_byteswap(GVariant*);
extern GVariantClass g_variant_classify(GVariant*);
extern int32_t g_variant_compare(GVariant*, GVariant*);
extern uint8_t* g_variant_dup_bytestring(GVariant*, uint64_t*);
extern char** g_variant_dup_bytestring_array(GVariant*, uint64_t*);
extern char** g_variant_dup_objv(GVariant*, uint64_t*);
extern char* g_variant_dup_string(GVariant*, uint64_t*);
extern char** g_variant_dup_strv(GVariant*, uint64_t*);
extern int g_variant_equal(GVariant*, GVariant*);
extern int g_variant_get_boolean(GVariant*);
extern uint8_t g_variant_get_byte(GVariant*);
extern uint8_t* g_variant_get_bytestring(GVariant*);
extern char** g_variant_get_bytestring_array(GVariant*, uint64_t*);
extern GVariant* g_variant_get_child_value(GVariant*, uint64_t);
extern void* g_variant_get_data(GVariant*);
extern double g_variant_get_double(GVariant*);
extern void** g_variant_get_fixed_array(GVariant*, uint64_t*, uint64_t);
extern int32_t g_variant_get_handle(GVariant*);
extern int16_t g_variant_get_int16(GVariant*);
extern int32_t g_variant_get_int32(GVariant*);
extern int64_t g_variant_get_int64(GVariant*);
extern GVariant* g_variant_get_maybe(GVariant*);
extern GVariant* g_variant_get_normal_form(GVariant*);
extern char** g_variant_get_objv(GVariant*, uint64_t*);
extern uint64_t g_variant_get_size(GVariant*);
extern char* g_variant_get_string(GVariant*, uint64_t*);
extern char** g_variant_get_strv(GVariant*, uint64_t*);
extern char* g_variant_get_type_string(GVariant*);
extern uint16_t g_variant_get_uint16(GVariant*);
extern uint32_t g_variant_get_uint32(GVariant*);
extern uint64_t g_variant_get_uint64(GVariant*);
extern GVariant* g_variant_get_variant(GVariant*);
extern uint32_t g_variant_hash(GVariant*);
extern int g_variant_is_container(GVariant*);
extern int g_variant_is_floating(GVariant*);
extern int g_variant_is_normal_form(GVariant*);
extern int g_variant_is_of_type(GVariant*, GVariantType*);
extern GVariant* g_variant_lookup_value(GVariant*, char*, GVariantType*);
extern uint64_t g_variant_n_children(GVariant*);
extern char* g_variant_print(GVariant*, int);
extern GVariant* g_variant_ref(GVariant*);
extern GVariant* g_variant_ref_sink(GVariant*);
extern void g_variant_store(GVariant*, void*);
extern GVariant* g_variant_take_ref(GVariant*);
extern void g_variant_unref(GVariant*);
extern int g_variant_is_object_path(char*);
extern int g_variant_is_signature(char*);
extern GVariant* g_variant_parse(GVariantType*, char*, char*, char*, GError**);
extern uint32_t g_variant_parser_get_error_quark();
extern GVariantBuilder* g_variant_builder_new(GVariantType*);
extern void g_variant_builder_add_value(GVariantBuilder*, GVariant*);
extern void g_variant_builder_close(GVariantBuilder*);
extern GVariant* g_variant_builder_end(GVariantBuilder*);
extern void g_variant_builder_open(GVariantBuilder*, GVariantType*);
extern GVariantBuilder* g_variant_builder_ref(GVariantBuilder*);
extern void g_variant_builder_unref(GVariantBuilder*);
extern GVariantType* g_variant_type_new(char*);
extern GVariantType* g_variant_type_new_tuple(GVariantType**, int32_t);
extern GVariantType* g_variant_type_copy(GVariantType*);
extern char* g_variant_type_dup_string(GVariantType*);
extern GVariantType* g_variant_type_element(GVariantType*);
extern int g_variant_type_equal(GVariantType*, GVariantType*);
extern GVariantType* g_variant_type_first(GVariantType*);
extern void g_variant_type_free(GVariantType*);
extern uint64_t g_variant_type_get_string_length(GVariantType*);
extern uint32_t g_variant_type_hash(GVariantType*);
extern int g_variant_type_is_array(GVariantType*);
extern int g_variant_type_is_basic(GVariantType*);
extern int g_variant_type_is_container(GVariantType*);
extern int g_variant_type_is_definite(GVariantType*);
extern int g_variant_type_is_dict_entry(GVariantType*);
extern int g_variant_type_is_maybe(GVariantType*);
extern int g_variant_type_is_subtype_of(GVariantType*, GVariantType*);
extern int g_variant_type_is_tuple(GVariantType*);
extern int g_variant_type_is_variant(GVariantType*);
extern GVariantType* g_variant_type_key(GVariantType*);
extern uint64_t g_variant_type_n_items(GVariantType*);
extern GVariantType* g_variant_type_new_array(GVariantType*);
extern GVariantType* g_variant_type_new_dict_entry(GVariantType*, GVariantType*);
extern GVariantType* g_variant_type_new_maybe(GVariantType*);
extern GVariantType* g_variant_type_next(GVariantType*);
extern GVariantType* g_variant_type_value(GVariantType*);
extern GVariantType* g_variant_type_checked_(char*);
extern int g_variant_type_string_is_valid(char*);
extern int g_variant_type_string_scan(char*, char*, char**);
extern int32_t g_access(char*, int32_t);
extern int32_t g_ascii_digit_value(uint8_t);
extern char* g_ascii_dtostr(char*, int32_t, double);
extern char* g_ascii_formatd(char*, int32_t, char*, double);
extern int32_t g_ascii_strcasecmp(char*, char*);
extern char* g_ascii_strdown(char*, int64_t);
extern int32_t g_ascii_strncasecmp(char*, char*, uint64_t);
extern double g_ascii_strtod(char*, char*);
extern int64_t g_ascii_strtoll(char*, char*, uint32_t);
extern uint64_t g_ascii_strtoull(char*, char*, uint32_t);
extern char* g_ascii_strup(char*, int64_t);
extern uint8_t g_ascii_tolower(uint8_t);
extern uint8_t g_ascii_toupper(uint8_t);
extern int32_t g_ascii_xdigit_value(uint8_t);
extern void g_assert_warning(char*, char*, int32_t, char*, char*);
extern void g_assertion_message(char*, char*, int32_t, char*, char*);
extern void g_assertion_message_cmpstr(char*, char*, int32_t, char*, char*, char*, char*, char*);
extern void g_assertion_message_error(char*, char*, int32_t, char*, char*, GError*, uint32_t, int32_t);
extern void g_assertion_message_expr(char*, char*, int32_t, char*, char*);
extern void g_atexit(GVoidFunc);
extern int32_t g_atomic_int_add(int32_t*, int32_t);
extern uint32_t g_atomic_int_and(uint32_t*, uint32_t);
extern int g_atomic_int_compare_and_exchange(int32_t*, int32_t, int32_t);
extern int g_atomic_int_dec_and_test(int32_t*);
extern int32_t g_atomic_int_exchange_and_add(int32_t*, int32_t);
extern int32_t g_atomic_int_get(int32_t*);
extern void g_atomic_int_inc(int32_t*);
extern uint32_t g_atomic_int_or(uint32_t*, uint32_t);
extern void g_atomic_int_set(int32_t*, int32_t);
extern uint32_t g_atomic_int_xor(uint32_t*, uint32_t);
extern int64_t g_atomic_pointer_add(void*, int64_t);
extern uint64_t g_atomic_pointer_and(void*, uint64_t);
extern int g_atomic_pointer_compare_and_exchange(void*, void*, void*);
extern uint64_t g_atomic_pointer_or(void*, uint64_t);
extern void g_atomic_pointer_set(void*, void*);
extern uint64_t g_atomic_pointer_xor(void*, uint64_t);
extern uint8_t* g_base64_decode(char*, uint64_t*);
extern uint8_t* g_base64_decode_inplace(uint8_t**, uint64_t*);
extern uint64_t g_base64_decode_step(uint8_t*, uint64_t, uint8_t**, int32_t*, uint32_t*);
extern char* g_base64_encode(uint8_t*, uint64_t);
extern uint64_t g_base64_encode_close(int, uint8_t**, int32_t*, int32_t*);
extern uint64_t g_base64_encode_step(uint8_t*, uint64_t, int, uint8_t**, int32_t*, int32_t*);
extern char* g_basename(char*);
extern void g_bit_lock(int32_t*, int32_t);
extern int32_t g_bit_nth_lsf(uint64_t, int32_t);
extern int32_t g_bit_nth_msf(uint64_t, int32_t);
extern uint32_t g_bit_storage(uint64_t);
extern int g_bit_trylock(int32_t*, int32_t);
extern void g_bit_unlock(int32_t*, int32_t);
extern void g_blow_chunks();
extern char* g_build_filenamev(char**);
extern char* g_build_pathv(char*, char**);
extern int32_t g_chdir(char*);
extern char* glib_check_version(uint32_t, uint32_t, uint32_t);
extern uint32_t g_child_watch_add_full(int32_t, int32_t, GChildWatchFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_child_watch_add_full(int32_t arg0, int32_t arg1, void* gofunc) {
	if (gofunc) {
		return g_child_watch_add_full(arg0, arg1, _GChildWatchFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_child_watch_add_full(arg0, arg1, 0, 0, 0);
	}
}
extern GSource* g_child_watch_source_new(int32_t);
extern void g_clear_error(GError**);
extern char* g_compute_checksum_for_data(GChecksumType, uint8_t*, uint64_t);
extern char* g_compute_checksum_for_string(GChecksumType, char*, int64_t);
extern char* g_compute_hmac_for_data(GChecksumType, uint8_t*, uint64_t, uint8_t*, uint64_t);
extern char* g_compute_hmac_for_string(GChecksumType, uint8_t*, uint64_t, char*, int64_t);
extern char* g_convert(char*, int64_t, char*, char*, uint64_t*, uint64_t*, GError**);
extern uint32_t g_convert_error_quark();
extern char* g_convert_with_fallback(char*, int64_t, char*, char*, char*, uint64_t*, uint64_t*, GError**);
extern char* g_convert_with_iconv(char*, int64_t, GIConv*, uint64_t*, uint64_t*, GError**);
extern void g_datalist_clear(GData*);
extern uint32_t g_datalist_get_flags(GData*);
extern void g_datalist_id_set_data_full(GData*, uint32_t, void*, GDestroyNotify);
extern void g_datalist_init(GData*);
extern void g_datalist_set_flags(GData*, uint32_t);
extern void g_datalist_unset_flags(GData*, uint32_t);
extern void g_dataset_destroy(void*);
extern void g_dataset_id_set_data_full(void*, uint32_t, void*, GDestroyNotify);
extern char* g_dcgettext(char*, char*, int32_t);
extern char* g_dgettext(char*, char*);
extern int g_direct_equal(void*, void*);
extern uint32_t g_direct_hash(void*);
extern char* g_dngettext(char*, char*, char*, uint64_t);
extern int g_double_equal(void*, void*);
extern uint32_t g_double_hash(void*);
extern char* g_dpgettext(char*, char*, uint64_t);
extern char* g_dpgettext2(char*, char*, char*);
extern void glib_dummy_decl();
extern GFileError g_file_error_from_errno(int32_t);
extern uint32_t g_file_error_quark();
extern int g_file_get_contents(char*, uint8_t**, uint64_t*, GError**);
extern int32_t g_file_open_tmp(char*, char**, GError**);
extern char* g_file_read_link(char*, GError**);
extern int g_file_set_contents(char*, uint8_t*, int64_t, GError**);
extern int g_file_test(char*, GFileTest);
extern char* g_filename_display_basename(char*);
extern char* g_filename_display_name(char*);
extern char* g_filename_from_uri(char*, char*, GError**);
extern char* g_filename_from_utf8(char*, int64_t, uint64_t*, uint64_t*, GError**);
extern char* g_filename_to_uri(char*, char*, GError**);
extern char* g_filename_to_utf8(char*, int64_t, uint64_t*, uint64_t*, GError**);
extern char* g_find_program_in_path(char*);
extern char* g_format_size(uint64_t);
extern char* g_format_size_for_display(int64_t);
extern char* g_format_size_full(uint64_t, GFormatSizeFlags);
extern void g_free(void*);
extern char* g_get_application_name();
extern int g_get_charset(char*);
extern char* g_get_current_dir();
extern void g_get_current_time(GTimeVal*);
extern char** g_get_environ();
extern int g_get_filename_charsets(char*);
extern char* g_get_home_dir();
extern char* g_get_host_name();
extern char** g_get_language_names();
extern char** g_get_locale_variants(char*);
extern int64_t g_get_monotonic_time();
extern char* g_get_prgname();
extern char* g_get_real_name();
extern int64_t g_get_real_time();
extern char** g_get_system_config_dirs();
extern char** g_get_system_data_dirs();
extern char* g_get_tmp_dir();
extern char* g_get_user_cache_dir();
extern char* g_get_user_config_dir();
extern char* g_get_user_data_dir();
extern char* g_get_user_name();
extern char* g_get_user_runtime_dir();
extern char* g_get_user_special_dir(GUserDirectory);
extern char* g_getenv(char*);
extern int g_hostname_is_ascii_encoded(char*);
extern int g_hostname_is_ip_address(char*);
extern int g_hostname_is_non_ascii(char*);
extern char* g_hostname_to_ascii(char*);
extern char* g_hostname_to_unicode(char*);
extern uint32_t g_idle_add_full(int32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_idle_add_full(int32_t arg0, void* gofunc) {
	if (gofunc) {
		return g_idle_add_full(arg0, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_idle_add_full(arg0, 0, 0, 0);
	}
}
extern int g_idle_remove_by_data(void*);
extern GSource* g_idle_source_new();
extern int g_int64_equal(void*, void*);
extern uint32_t g_int64_hash(void*);
extern int g_int_equal(void*, void*);
extern uint32_t g_int_hash(void*);
extern char* g_intern_static_string(char*);
extern char* g_intern_string(char*);
extern uint32_t g_io_add_watch_full(GIOChannel*, int32_t, GIOCondition, GIOFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_io_add_watch_full(GIOChannel* arg0, int32_t arg1, GIOCondition arg2, void* gofunc) {
	if (gofunc) {
		return g_io_add_watch_full(arg0, arg1, arg2, _GIOFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_io_add_watch_full(arg0, arg1, arg2, 0, 0, 0);
	}
}
extern GSource* g_io_create_watch(GIOChannel*, GIOCondition);
extern char** g_listenv();
extern char* g_locale_from_utf8(char*, int64_t, uint64_t*, uint64_t*, GError**);
extern char* g_locale_to_utf8(char*, int64_t, uint64_t*, uint64_t*, GError**);
extern void g_log_default_handler(char*, GLogLevelFlags, char*, void*);
extern void g_log_remove_handler(char*, uint32_t);
extern GLogLevelFlags g_log_set_always_fatal(GLogLevelFlags);
extern GLogLevelFlags g_log_set_fatal_mask(char*, GLogLevelFlags);
extern GSource* g_main_current_source();
extern int32_t g_main_depth();
extern uint32_t g_markup_error_quark();
extern char* g_markup_escape_text(char*, int64_t);
extern int g_mem_is_system_malloc();
extern void g_mem_profile();
extern void g_mem_set_vtable(GMemVTable*);
extern int32_t g_mkdir_with_parents(char*, int32_t);
extern char* g_mkdtemp(char*);
extern char* g_mkdtemp_full(char*, int32_t);
extern int32_t g_mkstemp(char*);
extern int32_t g_mkstemp_full(char*, int32_t, int32_t);
extern void g_nullify_pointer(void*);
extern void g_on_error_query(char*);
extern void g_on_error_stack_trace(char*);
extern uint32_t g_option_error_quark();
extern uint32_t g_parse_debug_string(char*, GDebugKey*, uint32_t);
extern char* g_path_get_basename(char*);
extern char* g_path_get_dirname(char*);
extern int g_path_is_absolute(char*);
extern char* g_path_skip_root(char*);
extern int g_pattern_match(GPatternSpec*, uint32_t, char*, char*);
extern int g_pattern_match_simple(char*, char*);
extern int g_pattern_match_string(GPatternSpec*, char*);
extern void g_pointer_bit_lock(void*, int32_t);
extern int g_pointer_bit_trylock(void*, int32_t);
extern void g_pointer_bit_unlock(void*, int32_t);
extern int32_t g_poll(GPollFD*, uint32_t, int32_t);
extern void g_propagate_error(GError*, GError*);
extern uint32_t g_quark_from_static_string(char*);
extern uint32_t g_quark_from_string(char*);
extern char* g_quark_to_string(uint32_t);
extern uint32_t g_quark_try_string(char*);
extern double g_random_double();
extern double g_random_double_range(double, double);
extern uint32_t g_random_int();
extern int32_t g_random_int_range(int32_t, int32_t);
extern void g_random_set_seed(uint32_t);
extern void g_reload_user_special_dirs_cache();
extern void g_return_if_fail_warning(char*, char*, char*);
extern int32_t g_rmdir(char*);
extern void g_set_application_name(char*);
extern void g_set_error_literal(GError*, uint32_t, int32_t, char*);
extern void g_set_prgname(char*);
extern int g_setenv(char*, char*, int);
extern uint32_t g_shell_error_quark();
extern int g_shell_parse_argv(char*, int32_t*, char***, GError**);
extern char* g_shell_quote(char*);
extern char* g_shell_unquote(char*, GError**);
extern void g_slice_free1(uint64_t, void*);
extern void g_slice_free_chain_with_offset(uint64_t, void*, uint64_t);
extern int64_t g_slice_get_config(GSliceConfig);
extern int64_t* g_slice_get_config_state(GSliceConfig, int64_t, uint32_t*);
extern void g_slice_set_config(GSliceConfig, int64_t);
extern uint32_t g_spaced_primes_closest(uint32_t);
extern int g_spawn_async(char*, char**, char**, GSpawnFlags, GSpawnChildSetupFunc, void*, int32_t*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_spawn_async(char* arg0, char** arg1, char** arg2, GSpawnFlags arg3, void* gofunc, int32_t* arg6, GError** arg7) {
	if (gofunc) {
		return g_spawn_async(arg0, arg1, arg2, arg3, _GSpawnChildSetupFunc_c_wrapper_once, gofunc, arg6, arg7);
	} else {
		return g_spawn_async(arg0, arg1, arg2, arg3, 0, 0, arg6, arg7);
	}
}
extern int g_spawn_async_with_pipes(char*, char**, char**, GSpawnFlags, GSpawnChildSetupFunc, void*, int32_t*, int32_t*, int32_t*, int32_t*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_spawn_async_with_pipes(char* arg0, char** arg1, char** arg2, GSpawnFlags arg3, void* gofunc, int32_t* arg6, int32_t* arg7, int32_t* arg8, int32_t* arg9, GError** arg10) {
	if (gofunc) {
		return g_spawn_async_with_pipes(arg0, arg1, arg2, arg3, _GSpawnChildSetupFunc_c_wrapper_once, gofunc, arg6, arg7, arg8, arg9, arg10);
	} else {
		return g_spawn_async_with_pipes(arg0, arg1, arg2, arg3, 0, 0, arg6, arg7, arg8, arg9, arg10);
	}
}
extern void g_spawn_close_pid(int32_t);
extern int g_spawn_command_line_async(char*, GError**);
extern int g_spawn_command_line_sync(char*, uint8_t**, uint8_t**, int32_t*, GError**);
extern uint32_t g_spawn_error_quark();
extern int g_spawn_sync(char*, char**, char**, GSpawnFlags, GSpawnChildSetupFunc, void*, uint8_t**, uint8_t**, int32_t*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_spawn_sync(char* arg0, char** arg1, char** arg2, GSpawnFlags arg3, void* gofunc, uint8_t** arg6, uint8_t** arg7, int32_t* arg8, GError** arg9) {
	if (gofunc) {
		return g_spawn_sync(arg0, arg1, arg2, arg3, _GSpawnChildSetupFunc_c_wrapper_once, gofunc, arg6, arg7, arg8, arg9);
	} else {
		return g_spawn_sync(arg0, arg1, arg2, arg3, 0, 0, arg6, arg7, arg8, arg9);
	}
}
extern char* g_stpcpy(char*, char*);
extern int g_str_equal(void*, void*);
extern int g_str_has_prefix(char*, char*);
extern int g_str_has_suffix(char*, char*);
extern uint32_t g_str_hash(void*);
extern char* g_strcanon(char*, char*, uint8_t);
extern int32_t g_strcasecmp(char*, char*);
extern char* g_strchomp(char*);
extern char* g_strchug(char*);
extern int32_t g_strcmp0(char*, char*);
extern char* g_strcompress(char*);
extern char* g_strdelimit(char*, char*, uint8_t);
extern char* g_strdown(char*);
extern char* g_strdup(char*);
extern char* g_strerror(int32_t);
extern char* g_strescape(char*, char*);
extern void g_strfreev(char*);
extern GString* g_string_new(char*);
extern GString* g_string_new_len(char*, int64_t);
extern GString* g_string_sized_new(uint64_t);
extern char* g_strip_context(char*, char*);
extern char* g_strjoinv(char*, char*);
extern uint64_t g_strlcat(char*, char*, uint64_t);
extern uint64_t g_strlcpy(char*, char*, uint64_t);
extern int32_t g_strncasecmp(char*, char*, uint32_t);
extern char* g_strndup(char*, uint64_t);
extern char* g_strnfill(uint64_t, uint8_t);
extern char* g_strreverse(char*);
extern char* g_strrstr(char*, char*);
extern char* g_strrstr_len(char*, int64_t, char*);
extern char* g_strsignal(int32_t);
extern char* g_strstr_len(char*, int64_t, char*);
extern double g_strtod(char*, char*);
extern char* g_strup(char*);
extern GType g_strv_get_type();
extern uint32_t g_strv_length(char*);
extern void g_test_bug(char*);
extern void g_test_bug_base(char*);
extern void g_test_fail();
extern char* g_test_log_type_name(GTestLogType);
extern void g_test_queue_destroy(GDestroyNotify, void*);
extern void g_test_queue_free(void*);
extern double g_test_rand_double();
extern double g_test_rand_double_range(double, double);
extern int32_t g_test_rand_int();
extern int32_t g_test_rand_int_range(int32_t, int32_t);
extern int32_t g_test_run();
extern int32_t g_test_run_suite(GTestSuite*);
extern double g_test_timer_elapsed();
extern double g_test_timer_last();
extern void g_test_timer_start();
extern void g_test_trap_assertions(char*, char*, int32_t, char*, uint64_t, char*);
extern int g_test_trap_fork(uint64_t, GTestTrapFlags);
extern int g_test_trap_has_passed();
extern int g_test_trap_reached_timeout();
extern uint32_t g_timeout_add_full(int32_t, uint32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_timeout_add_full(int32_t arg0, uint32_t arg1, void* gofunc) {
	if (gofunc) {
		return g_timeout_add_full(arg0, arg1, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_timeout_add_full(arg0, arg1, 0, 0, 0);
	}
}
extern uint32_t g_timeout_add_seconds_full(int32_t, uint32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_timeout_add_seconds_full(int32_t arg0, uint32_t arg1, void* gofunc) {
	if (gofunc) {
		return g_timeout_add_seconds_full(arg0, arg1, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_timeout_add_seconds_full(arg0, arg1, 0, 0, 0);
	}
}
extern GSource* g_timeout_source_new(uint32_t);
extern GSource* g_timeout_source_new_seconds(uint32_t);
extern uint16_t* g_ucs4_to_utf16(uint32_t*, int64_t, int64_t*, int64_t*, GError**);
extern char* g_ucs4_to_utf8(uint32_t*, int64_t, int64_t*, int64_t*, GError**);
extern GUnicodeBreakType g_unichar_break_type(uint32_t);
extern int32_t g_unichar_combining_class(uint32_t);
extern int g_unichar_compose(uint32_t, uint32_t, uint32_t*);
extern int g_unichar_decompose(uint32_t, uint32_t*, uint32_t*);
extern int32_t g_unichar_digit_value(uint32_t);
extern uint64_t g_unichar_fully_decompose(uint32_t, int, uint32_t*, uint64_t);
extern int g_unichar_get_mirror_char(uint32_t, uint32_t*);
extern GUnicodeScript g_unichar_get_script(uint32_t);
extern int g_unichar_isalnum(uint32_t);
extern int g_unichar_isalpha(uint32_t);
extern int g_unichar_iscntrl(uint32_t);
extern int g_unichar_isdefined(uint32_t);
extern int g_unichar_isdigit(uint32_t);
extern int g_unichar_isgraph(uint32_t);
extern int g_unichar_islower(uint32_t);
extern int g_unichar_ismark(uint32_t);
extern int g_unichar_isprint(uint32_t);
extern int g_unichar_ispunct(uint32_t);
extern int g_unichar_isspace(uint32_t);
extern int g_unichar_istitle(uint32_t);
extern int g_unichar_isupper(uint32_t);
extern int g_unichar_iswide(uint32_t);
extern int g_unichar_iswide_cjk(uint32_t);
extern int g_unichar_isxdigit(uint32_t);
extern int g_unichar_iszerowidth(uint32_t);
extern int32_t g_unichar_to_utf8(uint32_t, char*);
extern uint32_t g_unichar_tolower(uint32_t);
extern uint32_t g_unichar_totitle(uint32_t);
extern uint32_t g_unichar_toupper(uint32_t);
extern GUnicodeType g_unichar_type(uint32_t);
extern int g_unichar_validate(uint32_t);
extern int32_t g_unichar_xdigit_value(uint32_t);
extern uint32_t* g_unicode_canonical_decomposition(uint32_t, uint64_t*);
extern void g_unicode_canonical_ordering(uint32_t*, uint64_t);
extern GUnicodeScript g_unicode_script_from_iso15924(uint32_t);
extern uint32_t g_unicode_script_to_iso15924(GUnicodeScript);
extern int32_t g_unlink(char*);
extern void g_unsetenv(char*);
extern char* g_uri_escape_string(char*, char*, int);
extern char* g_uri_parse_scheme(char*);
extern char* g_uri_unescape_segment(char*, char*, char*);
extern char* g_uri_unescape_string(char*, char*);
extern void g_usleep(uint64_t);
extern uint32_t* g_utf16_to_ucs4(uint16_t*, int64_t, int64_t*, int64_t*, GError**);
extern char* g_utf16_to_utf8(uint16_t*, int64_t, int64_t*, int64_t*, GError**);
extern char* g_utf8_casefold(char*, int64_t);
extern int32_t g_utf8_collate(char*, char*);
extern char* g_utf8_collate_key(char*, int64_t);
extern char* g_utf8_collate_key_for_filename(char*, int64_t);
extern char* g_utf8_find_next_char(char*, char*);
extern char* g_utf8_find_prev_char(char*, char*);
extern uint32_t g_utf8_get_char(char*);
extern uint32_t g_utf8_get_char_validated(char*, int64_t);
extern char* g_utf8_normalize(char*, int64_t, GNormalizeMode);
extern char* g_utf8_offset_to_pointer(char*, int64_t);
extern int64_t g_utf8_pointer_to_offset(char*, char*);
extern char* g_utf8_prev_char(char*);
extern char* g_utf8_strchr(char*, int64_t, uint32_t);
extern char* g_utf8_strdown(char*, int64_t);
extern int64_t g_utf8_strlen(char*, int64_t);
extern char* g_utf8_strncpy(char*, char*, uint64_t);
extern char* g_utf8_strrchr(char*, int64_t, uint32_t);
extern char* g_utf8_strreverse(char*, int64_t);
extern char* g_utf8_strup(char*, int64_t);
extern char* g_utf8_substring(char*, int64_t, int64_t);
extern uint32_t* g_utf8_to_ucs4(char*, int64_t, int64_t*, int64_t*, GError**);
extern uint32_t* g_utf8_to_ucs4_fast(char*, int64_t, int64_t*);
extern uint16_t* g_utf8_to_utf16(char*, int64_t, int64_t*, int64_t*, GError**);
extern int g_utf8_validate(char*, int64_t, char**);
extern GType g_variant_get_gtype();
extern GVariantType* g_variant_get_type(GVariant*);
extern void g_warn_message(char*, char*, int32_t, char*, char*);
struct _GAllocator {};
struct _GArray { uint8_t _data[16]; };
struct _GAsyncQueue {};
struct _GBookmarkFile {};
struct _GByteArray { uint8_t _data[16]; };
struct _GCache {};
struct _GChecksum {};
struct _GCompletion { uint8_t _data[40]; };
struct _GCond {};
struct _GData {};
struct _GDate { uint8_t _data[24]; };
struct _GDateTime {};
struct _GDebugKey { uint8_t _data[16]; };
struct _GDir {};
struct _GDoubleIEEE754 { uint8_t _data[8]; };
struct _GError { uint8_t _data[16]; };
struct _GFloatIEEE754 { uint8_t _data[4]; };
struct _GHashTable {};
struct _GHashTableIter { uint8_t _data[40]; };
struct _GHmac {};
struct _GHook { uint8_t _data[64]; };
struct _GHookList { uint8_t _data[56]; };
struct _GIConv {};
struct _GIOChannel { uint8_t _data[136]; };
struct _GIOFuncs { uint8_t _data[64]; };
struct _GKeyFile {};
struct _GList { uint8_t _data[24]; };
struct _GMainContext {};
struct _GMainLoop {};
struct _GMappedFile {};
struct _GMarkupParseContext {};
struct _GMarkupParser { uint8_t _data[40]; };
struct _GMatchInfo {};
struct _GMemChunk {};
struct _GMemVTable { uint8_t _data[48]; };
struct _GMutex {};
struct _GNode { uint8_t _data[40]; };
struct _GOnce { uint8_t _data[16]; };
struct _GOptionContext {};
struct _GOptionEntry { uint8_t _data[48]; };
struct _GOptionGroup {};
struct _GPatternSpec {};
struct _GPollFD { uint8_t _data[8]; };
struct _GPrivate {};
struct _GPtrArray { uint8_t _data[16]; };
struct _GQueue { uint8_t _data[24]; };
struct _GRand {};
struct _GRegex {};
struct _GRelation {};
struct _GSList { uint8_t _data[16]; };
struct _GScanner { uint8_t _data[144]; };
struct _GScannerConfig { uint8_t _data[128]; };
struct _GSequence {};
struct _GSequenceIter {};
struct _GSource { uint8_t _data[96]; };
struct _GSourceCallbackFuncs { uint8_t _data[24]; };
struct _GSourceFuncs { uint8_t _data[48]; };
struct _GSourcePrivate {};
struct _GStatBuf {};
struct _GStaticMutex { uint8_t _data[8]; };
struct _GStaticPrivate { uint8_t _data[4]; };
struct _GStaticRWLock { uint8_t _data[40]; };
struct _GStaticRecMutex { uint8_t _data[24]; };
struct _GString { uint8_t _data[24]; };
struct _GStringChunk {};
struct _GSystemThread { uint8_t _data[8]; };
struct _GTestCase {};
struct _GTestConfig { uint8_t _data[20]; };
struct _GTestLogBuffer { uint8_t _data[16]; };
struct _GTestLogMsg { uint8_t _data[32]; };
struct _GTestSuite {};
struct _GThread { uint8_t _data[24]; };
struct _GThreadFunctions { uint8_t _data[168]; };
struct _GThreadPool { uint8_t _data[24]; };
struct _GTimeVal { uint8_t _data[16]; };
struct _GTimeZone {};
struct _GTimer {};
struct _GTokenValue { uint8_t _data[8]; };
struct _GTrashStack { uint8_t _data[8]; };
struct _GTree {};
struct _GTuples { uint8_t _data[4]; };
struct _GVariant {};
struct _GVariantBuilder { uint8_t _data[128]; };
struct _GVariantType {};
struct _G_StaticAssert_347 { uint8_t _data[1]; };


#cgo pkg-config: glib-2.0
*/
import "C"
import "unsafe"
import "errors"

const alot = 999999

type _GSList struct {
	data unsafe.Pointer
	next *_GSList
}

type _GList struct {
	data unsafe.Pointer
	next *_GList
	prev *_GList
}

type _GError struct {
	domain uint32
	code int32
	message *C.char
}

func _GoStringToGString(x string) *C.char {
	if x == "\x00" {
		return nil
	}
	return C.CString(x)
}

func _GoBoolToCBool(x bool) C.int {
	if x { return 1 }
	return 0
}

func _CInterfaceToGoInterface(iface [2]unsafe.Pointer) interface{} {
	return *(*interface{})(unsafe.Pointer(&iface))
}

func _GoInterfaceToCInterface(iface interface{}) *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(&iface))
}


const AllocatorList = 1
const AllocatorNode = 3
const AllocatorSlist = 2
const AllocAndFree = 2
const AllocOnly = 1
const AsciiDtostrBufSize = 39
const AtomicOpUseGccBuiltins = 1
// blacklisted: Allocator (struct)
// blacklisted: Array (struct)
type AsciiType C.uint32_t
const (
	AsciiTypeAlnum AsciiType = 1
	AsciiTypeAlpha AsciiType = 2
	AsciiTypeCntrl AsciiType = 4
	AsciiTypeDigit AsciiType = 8
	AsciiTypeGraph AsciiType = 16
	AsciiTypeLower AsciiType = 32
	AsciiTypePrint AsciiType = 64
	AsciiTypePunct AsciiType = 128
	AsciiTypeSpace AsciiType = 256
	AsciiTypeUpper AsciiType = 512
	AsciiTypeXdigit AsciiType = 1024
)
// blacklisted: AsyncQueue (struct)
const BigEndian = 4321
// blacklisted: BookmarkFile (struct)
type BookmarkFileError C.uint32_t
const (
	BookmarkFileErrorInvalidURI BookmarkFileError = 0
	BookmarkFileErrorInvalidValue BookmarkFileError = 1
	BookmarkFileErrorAppNotRegistered BookmarkFileError = 2
	BookmarkFileErrorURINotFound BookmarkFileError = 3
	BookmarkFileErrorRead BookmarkFileError = 4
	BookmarkFileErrorUnknownEncoding BookmarkFileError = 5
	BookmarkFileErrorWrite BookmarkFileError = 6
	BookmarkFileErrorFileNotFound BookmarkFileError = 7
)
// blacklisted: ByteArray (struct)
const CanInline = 1
// blacklisted: CSET_A_2_Z (constant)
const CsetDigits = "0123456789"
// blacklisted: CSET_a_2_z (constant)
// blacklisted: Cache (struct)
// blacklisted: CacheDestroyFunc (callback)
// blacklisted: Checksum (struct)
type ChecksumType C.uint32_t
const (
	ChecksumTypeMd5 ChecksumType = 0
	ChecksumTypeSha1 ChecksumType = 1
	ChecksumTypeSha256 ChecksumType = 2
)
// blacklisted: ChildWatchFunc (callback)
// blacklisted: CompareDataFunc (callback)
// blacklisted: CompareFunc (callback)
// blacklisted: Completion (struct)
// blacklisted: CompletionFunc (callback)
// blacklisted: CompletionStrncmpFunc (callback)
// blacklisted: Cond (struct)
type ConvertError C.uint32_t
const (
	ConvertErrorNoConversion ConvertError = 0
	ConvertErrorIllegalSequence ConvertError = 1
	ConvertErrorFailed ConvertError = 2
	ConvertErrorPartialInput ConvertError = 3
	ConvertErrorBadURI ConvertError = 4
	ConvertErrorNotAbsolutePath ConvertError = 5
)
const DatalistFlagsMask = 3
const DateBadDay = 0
const DateBadJulian = 0
const DateBadYear = 0
const DirSeparator = 92
const DirSeparatorS = "\\"
// blacklisted: Data (struct)
// blacklisted: DataForeachFunc (callback)
// blacklisted: Date (struct)
type DateDMY C.uint32_t
const (
	DateDMYDay DateDMY = 0
	DateDMYMonth DateDMY = 1
	DateDMYYear DateDMY = 2
)
type DateMonth C.uint32_t
const (
	DateMonthBadMonth DateMonth = 0
	DateMonthJanuary DateMonth = 1
	DateMonthFebruary DateMonth = 2
	DateMonthMarch DateMonth = 3
	DateMonthApril DateMonth = 4
	DateMonthMay DateMonth = 5
	DateMonthJune DateMonth = 6
	DateMonthJuly DateMonth = 7
	DateMonthAugust DateMonth = 8
	DateMonthSeptember DateMonth = 9
	DateMonthOctober DateMonth = 10
	DateMonthNovember DateMonth = 11
	DateMonthDecember DateMonth = 12
)
// blacklisted: DateTime (struct)
type DateWeekday C.uint32_t
const (
	DateWeekdayBadWeekday DateWeekday = 0
	DateWeekdayMonday DateWeekday = 1
	DateWeekdayTuesday DateWeekday = 2
	DateWeekdayWednesday DateWeekday = 3
	DateWeekdayThursday DateWeekday = 4
	DateWeekdayFriday DateWeekday = 5
	DateWeekdaySaturday DateWeekday = 6
	DateWeekdaySunday DateWeekday = 7
)
// blacklisted: DebugKey (struct)
// blacklisted: DestroyNotify (callback)
// blacklisted: Dir (struct)
type DoubleIEEE754 struct {
	_data [8]byte
}
const E = 2.718282
// blacklisted: EqualFunc (callback)
// blacklisted: Error (struct)
type ErrorType C.uint32_t
const (
	ErrorTypeUnknown ErrorType = 0
	ErrorTypeUnexpEof ErrorType = 1
	ErrorTypeUnexpEofInString ErrorType = 2
	ErrorTypeUnexpEofInComment ErrorType = 3
	ErrorTypeNonDigitInConst ErrorType = 4
	ErrorTypeDigitRadix ErrorType = 5
	ErrorTypeFloatRadix ErrorType = 6
	ErrorTypeFloatMalformed ErrorType = 7
)
type FileError C.uint32_t
const (
	FileErrorExist FileError = 0
	FileErrorIsdir FileError = 1
	FileErrorAcces FileError = 2
	FileErrorNametoolong FileError = 3
	FileErrorNoent FileError = 4
	FileErrorNotdir FileError = 5
	FileErrorNxio FileError = 6
	FileErrorNodev FileError = 7
	FileErrorRofs FileError = 8
	FileErrorTxtbsy FileError = 9
	FileErrorFault FileError = 10
	FileErrorLoop FileError = 11
	FileErrorNospc FileError = 12
	FileErrorNomem FileError = 13
	FileErrorMfile FileError = 14
	FileErrorNfile FileError = 15
	FileErrorBadf FileError = 16
	FileErrorInval FileError = 17
	FileErrorPipe FileError = 18
	FileErrorAgain FileError = 19
	FileErrorIntr FileError = 20
	FileErrorIo FileError = 21
	FileErrorPerm FileError = 22
	FileErrorNosys FileError = 23
	FileErrorFailed FileError = 24
)
type FileTest C.uint32_t
const (
	FileTestIsRegular FileTest = 1
	FileTestIsSymlink FileTest = 2
	FileTestIsDir FileTest = 4
	FileTestIsExecutable FileTest = 8
	FileTestExists FileTest = 16
)
type FloatIEEE754 struct {
	_data [4]byte
}
type FormatSizeFlags C.uint32_t
const (
	FormatSizeFlagsDefault FormatSizeFlags = 0
	FormatSizeFlagsLongFormat FormatSizeFlags = 1
	FormatSizeFlagsIecUnits FormatSizeFlags = 2
)
// blacklisted: FreeFunc (callback)
// blacklisted: Func (callback)
const Gint16Format = "hi"
const Gint16Modifier = "h"
const Gint32Format = "i"
const Gint32Modifier = ""
const Gint64Format = "li"
const Gint64Modifier = "l"
const GintptrFormat = "li"
const GintptrModifier = "l"
const GnucFunction = ""
const GnucPrettyFunction = ""
const GsizeFormat = "lu"
const GsizeModifier = "l"
const GssizeFormat = "li"
const Guint16Format = "hu"
const Guint32Format = "u"
const Guint64Format = "lu"
const GuintptrFormat = "lu"
const HaveGint64 = 1
const HaveGnucVarargs = 1
const HaveGnucVisibility = 1
const HaveGrowingStack = 1
const HaveInline = 1
const HaveIsoVarargs = 1
const Have__Inline = 1
const Have__Inline__ = 1
// blacklisted: HFunc (callback)
const HookFlagUserShift = 4
// blacklisted: HRFunc (callback)
// blacklisted: HashFunc (callback)
// blacklisted: HashTable (struct)
// blacklisted: HashTableIter (struct)
// blacklisted: Hmac (struct)
// blacklisted: Hook (struct)
// blacklisted: HookCheckFunc (callback)
// blacklisted: HookCheckMarshaller (callback)
// blacklisted: HookCompareFunc (callback)
// blacklisted: HookFinalizeFunc (callback)
// blacklisted: HookFindFunc (callback)
type HookFlagMask C.uint32_t
const (
	HookFlagMaskActive HookFlagMask = 1
	HookFlagMaskInCall HookFlagMask = 2
	HookFlagMaskMask HookFlagMask = 15
)
// blacklisted: HookFunc (callback)
// blacklisted: HookList (struct)
// blacklisted: HookMarshaller (callback)
// blacklisted: IConv (struct)
const Ieee754DoubleBias = 1023
const Ieee754FloatBias = 127
// blacklisted: IOChannel (struct)
type IOChannelError C.uint32_t
const (
	IOChannelErrorFbig IOChannelError = 0
	IOChannelErrorInval IOChannelError = 1
	IOChannelErrorIo IOChannelError = 2
	IOChannelErrorIsdir IOChannelError = 3
	IOChannelErrorNospc IOChannelError = 4
	IOChannelErrorNxio IOChannelError = 5
	IOChannelErrorOverflow IOChannelError = 6
	IOChannelErrorPipe IOChannelError = 7
	IOChannelErrorFailed IOChannelError = 8
)
type IOCondition C.uint32_t
const (
	IOConditionIn IOCondition = 1
	IOConditionOut IOCondition = 4
	IOConditionPri IOCondition = 2
	IOConditionErr IOCondition = 8
	IOConditionHup IOCondition = 16
	IOConditionNval IOCondition = 32
)
type IOError C.uint32_t
const (
	IOErrorNone IOError = 0
	IOErrorAgain IOError = 1
	IOErrorInval IOError = 2
	IOErrorUnknown IOError = 3
)
type IOFlags C.uint32_t
const (
	IOFlagsAppend IOFlags = 1
	IOFlagsNonblock IOFlags = 2
	IOFlagsIsReadable IOFlags = 4
	IOFlagsIsWriteable IOFlags = 8
	IOFlagsIsSeekable IOFlags = 16
	IOFlagsMask IOFlags = 31
	IOFlagsGetMask IOFlags = 31
	IOFlagsSetMask IOFlags = 3
)
// blacklisted: IOFunc (callback)
// blacklisted: IOFuncs (struct)
type IOStatus C.uint32_t
const (
	IOStatusError IOStatus = 0
	IOStatusNormal IOStatus = 1
	IOStatusEof IOStatus = 2
	IOStatusAgain IOStatus = 3
)
const KeyFileDesktopGroup = "Desktop Entry"
const KeyFileDesktopKeyCategories = "Categories"
const KeyFileDesktopKeyComment = "Comment"
const KeyFileDesktopKeyExec = "Exec"
const KeyFileDesktopKeyGenericName = "GenericName"
const KeyFileDesktopKeyHidden = "Hidden"
const KeyFileDesktopKeyIcon = "Icon"
const KeyFileDesktopKeyMIMEType = "MimeType"
const KeyFileDesktopKeyName = "Name"
const KeyFileDesktopKeyNotShowIn = "NotShowIn"
const KeyFileDesktopKeyNoDisplay = "NoDisplay"
const KeyFileDesktopKeyOnlyShowIn = "OnlyShowIn"
const KeyFileDesktopKeyPath = "Path"
const KeyFileDesktopKeyStartupNotify = "StartupNotify"
const KeyFileDesktopKeyStartupWmClass = "StartupWMClass"
const KeyFileDesktopKeyTerminal = "Terminal"
const KeyFileDesktopKeyTryExec = "TryExec"
const KeyFileDesktopKeyType = "Type"
const KeyFileDesktopKeyUrl = "URL"
const KeyFileDesktopKeyVersion = "Version"
const KeyFileDesktopTypeApplication = "Application"
const KeyFileDesktopTypeDirectory = "Directory"
const KeyFileDesktopTypeLink = "Link"
type KeyFile struct {}
func (this0 *KeyFile) Free() {
	var this1 *C.GKeyFile
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	C.g_key_file_free(this1)
}
func (this0 *KeyFile) GetBoolean(group_name0 string, key0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_boolean(this1, group_name1, key1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetBooleanList(group_name0 string, key0 string) (uint64, []bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_boolean_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []bool
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]bool, length1)
	for i := range ret2 {
		ret2[i] = (*(*[999999]C.int)(unsafe.Pointer(ret1)))[i] != 0
	}
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetComment(group_name0 string, key0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_comment(this1, group_name1, key1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetDouble(group_name0 string, key0 string) (float64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_double(this1, group_name1, key1, &err1)
	var ret2 float64
	var err2 error
	ret2 = float64(ret1)
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetDoubleList(group_name0 string, key0 string) (uint64, []float64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_double_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []float64
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]float64, length1)
	for i := range ret2 {
		ret2[i] = float64((*(*[999999]C.double)(unsafe.Pointer(ret1)))[i])
	}
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetInt64(group_name0 string, key0 string) (int64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_int64(this1, group_name1, key1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetInteger(group_name0 string, key0 string) (int, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_integer(this1, group_name1, key1, &err1)
	var ret2 int
	var err2 error
	ret2 = int(ret1)
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetIntegerList(group_name0 string, key0 string) (uint64, []int, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_integer_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []int
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]int, length1)
	for i := range ret2 {
		ret2[i] = int((*(*[999999]C.int32_t)(unsafe.Pointer(ret1)))[i])
	}
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetLocaleString(group_name0 string, key0 string, locale0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	ret1 := C.g_key_file_get_locale_string(this1, group_name1, key1, locale1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetLocaleStringList(group_name0 string, key0 string, locale0 string) (uint64, []string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	ret1 := C.g_key_file_get_locale_string_list(this1, group_name1, key1, locale1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetStartGroup() string {
	var this1 *C.GKeyFile
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	ret1 := C.g_key_file_get_start_group(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *KeyFile) GetString(group_name0 string, key0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_string(this1, group_name1, key1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetStringList(group_name0 string, key0 string) (uint64, []string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_string_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetUint64(group_name0 string, key0 string) (uint64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_uint64(this1, group_name1, key1, &err1)
	var ret2 uint64
	var err2 error
	ret2 = uint64(ret1)
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetValue(group_name0 string, key0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_value(this1, group_name1, key1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) HasGroup(group_name0 string) bool {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	ret1 := C.g_key_file_has_group(this1, group_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *KeyFile) LoadFromData(data0 string, length0 uint64, flags0 KeyFileFlags) (bool, error) {
	var this1 *C.GKeyFile
	var data1 *C.char
	var length1 C.uint64_t
	var flags1 C.GKeyFileFlags
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	data1 = _GoStringToGString(data0)
	defer C.free(unsafe.Pointer(data1))
	length1 = C.uint64_t(length0)
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_data(this1, data1, length1, flags1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) LoadFromDataDirs(file0 string, full_path0 string, flags0 KeyFileFlags) (bool, error) {
	var this1 *C.GKeyFile
	var file1 *C.char
	var full_path1 *C.char
	var flags1 C.GKeyFileFlags
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	full_path1 = _GoStringToGString(full_path0)
	defer C.free(unsafe.Pointer(full_path1))
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_data_dirs(this1, file1, full_path1, flags1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) LoadFromDirs(file0 string, search_dirs0 string, full_path0 string, flags0 KeyFileFlags) (bool, error) {
	var this1 *C.GKeyFile
	var file1 *C.char
	var search_dirs1 *C.char
	var full_path1 *C.char
	var flags1 C.GKeyFileFlags
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	search_dirs1 = _GoStringToGString(search_dirs0)
	defer C.free(unsafe.Pointer(search_dirs1))
	full_path1 = _GoStringToGString(full_path0)
	defer C.free(unsafe.Pointer(full_path1))
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_dirs(this1, file1, search_dirs1, full_path1, flags1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) LoadFromFile(file0 string, flags0 KeyFileFlags) (bool, error) {
	var this1 *C.GKeyFile
	var file1 *C.char
	var flags1 C.GKeyFileFlags
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_file(this1, file1, flags1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) RemoveComment(group_name0 string, key0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_remove_comment(this1, group_name1, key1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) RemoveGroup(group_name0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	ret1 := C.g_key_file_remove_group(this1, group_name1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) RemoveKey(group_name0 string, key0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_remove_key(this1, group_name1, key1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) SetBoolean(group_name0 string, key0 string, value0 bool) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.int
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = _GoBoolToCBool(value0)
	C.g_key_file_set_boolean(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetBooleanList(group_name0 string, key0 string, list0 bool, length0 uint64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 C.int
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = _GoBoolToCBool(list0)
	length1 = C.uint64_t(length0)
	C.g_key_file_set_boolean_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetComment(group_name0 string, key0 string, comment0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var comment1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	comment1 = _GoStringToGString(comment0)
	defer C.free(unsafe.Pointer(comment1))
	ret1 := C.g_key_file_set_comment(this1, group_name1, key1, comment1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) SetDouble(group_name0 string, key0 string, value0 float64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.double
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.double(value0)
	C.g_key_file_set_double(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetDoubleList(group_name0 string, key0 string, list0 float64, length0 uint64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 C.double
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = C.double(list0)
	length1 = C.uint64_t(length0)
	C.g_key_file_set_double_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetInt64(group_name0 string, key0 string, value0 int64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.int64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int64_t(value0)
	C.g_key_file_set_int64(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetInteger(group_name0 string, key0 string, value0 int) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.int32_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int32_t(value0)
	C.g_key_file_set_integer(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetIntegerList(group_name0 string, key0 string, list0 int, length0 uint64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 C.int32_t
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = C.int32_t(list0)
	length1 = C.uint64_t(length0)
	C.g_key_file_set_integer_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetListSeparator(separator0 int) {
	var this1 *C.GKeyFile
	var separator1 C.uint8_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	separator1 = C.uint8_t(separator0)
	C.g_key_file_set_list_separator(this1, separator1)
}
func (this0 *KeyFile) SetLocaleString(group_name0 string, key0 string, locale0 string, string0 string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var string1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	C.g_key_file_set_locale_string(this1, group_name1, key1, locale1, string1)
}
func (this0 *KeyFile) SetLocaleStringList(group_name0 string, key0 string, locale0 string, list0 string, length0 uint64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var list1 *C.char
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	list1 = _GoStringToGString(list0)
	defer C.free(unsafe.Pointer(list1))
	length1 = C.uint64_t(length0)
	C.g_key_file_set_locale_string_list(this1, group_name1, key1, locale1, list1, length1)
}
func (this0 *KeyFile) SetString(group_name0 string, key0 string, string0 string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var string1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	C.g_key_file_set_string(this1, group_name1, key1, string1)
}
func (this0 *KeyFile) SetStringList(group_name0 string, key0 string, list0 []string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 **C.char
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*list1)) * (len(list0) + 1))))
	defer C.free(unsafe.Pointer(list1))
	for i, e := range list0 {
		(*(*[999999]*C.char)(unsafe.Pointer(list1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(list1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(list1)))[len(list0)] = nil
	length1 = C.uint64_t(len(list0))
	C.g_key_file_set_string_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetUint64(group_name0 string, key0 string, value0 uint64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.uint64_t(value0)
	C.g_key_file_set_uint64(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetValue(group_name0 string, key0 string, value0 string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	C.g_key_file_set_value(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) ToData(length0 *uint64) (string, error) {
	var this1 *C.GKeyFile
	var length1 *C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	length1 = (*C.uint64_t)(unsafe.Pointer(length0))
	ret1 := C.g_key_file_to_data(this1, length1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func KeyFileErrorQuark() int {
	ret1 := C.g_key_file_error_quark()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
type KeyFileError C.uint32_t
const (
	KeyFileErrorUnknownEncoding KeyFileError = 0
	KeyFileErrorParse KeyFileError = 1
	KeyFileErrorNotFound KeyFileError = 2
	KeyFileErrorKeyNotFound KeyFileError = 3
	KeyFileErrorGroupNotFound KeyFileError = 4
	KeyFileErrorInvalidValue KeyFileError = 5
)
type KeyFileFlags C.uint32_t
const (
	KeyFileFlagsNone KeyFileFlags = 0
	KeyFileFlagsKeepComments KeyFileFlags = 1
	KeyFileFlagsKeepTranslations KeyFileFlags = 2
)
const LittleEndian = 1234
const Ln10 = 2.302585
const Ln2 = 0.693147
const Log2Base10 = 0.30103
const LogFatalMask = 0
const LogLevelUserShift = 8
// blacklisted: List (struct)
// blacklisted: LogFunc (callback)
type LogLevelFlags C.int32_t
const (
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	LogLevelFlagsFlagFatal LogLevelFlags = 2
	LogLevelFlagsLevelError LogLevelFlags = 4
	LogLevelFlagsLevelCritical LogLevelFlags = 8
	LogLevelFlagsLevelWarning LogLevelFlags = 16
	LogLevelFlagsLevelMessage LogLevelFlags = 32
	LogLevelFlagsLevelInfo LogLevelFlags = 64
	LogLevelFlagsLevelDebug LogLevelFlags = 128
	LogLevelFlagsLevelMask LogLevelFlags = -4
)
const MajorVersion = 2
const MicroVersion = 92
const MinorVersion = 29
const ModuleSuffix = "so"
const MutexDebugMagic = -119436585
// blacklisted: MainContext (struct)
// blacklisted: MainLoop (struct)
// blacklisted: MappedFile (struct)
type MarkupCollectType C.uint32_t
const (
	MarkupCollectTypeInvalid MarkupCollectType = 0
	MarkupCollectTypeString MarkupCollectType = 1
	MarkupCollectTypeStrdup MarkupCollectType = 2
	MarkupCollectTypeBoolean MarkupCollectType = 3
	MarkupCollectTypeTristate MarkupCollectType = 4
	MarkupCollectTypeOptional MarkupCollectType = 65536
)
type MarkupError C.uint32_t
const (
	MarkupErrorBadUtf8 MarkupError = 0
	MarkupErrorEmpty MarkupError = 1
	MarkupErrorParse MarkupError = 2
	MarkupErrorUnknownElement MarkupError = 3
	MarkupErrorUnknownAttribute MarkupError = 4
	MarkupErrorInvalidContent MarkupError = 5
	MarkupErrorMissingAttribute MarkupError = 6
)
// blacklisted: MarkupParseContext (struct)
type MarkupParseFlags C.uint32_t
const (
	MarkupParseFlagsDoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	MarkupParseFlagsTreatCdataAsText MarkupParseFlags = 2
	MarkupParseFlagsPrefixErrorPosition MarkupParseFlags = 4
)
// blacklisted: MarkupParser (struct)
// blacklisted: MatchInfo (struct)
// blacklisted: MemChunk (struct)
// blacklisted: MemVTable (struct)
// blacklisted: Mutex (struct)
// blacklisted: Node (struct)
// blacklisted: NodeForeachFunc (callback)
// blacklisted: NodeTraverseFunc (callback)
type NormalizeMode C.uint32_t
const (
	NormalizeModeDefault NormalizeMode = 0
	NormalizeModeNfd NormalizeMode = 0
	NormalizeModeDefaultCompose NormalizeMode = 1
	NormalizeModeNfc NormalizeMode = 1
	NormalizeModeAll NormalizeMode = 2
	NormalizeModeNfkd NormalizeMode = 2
	NormalizeModeAllCompose NormalizeMode = 3
	NormalizeModeNfkc NormalizeMode = 3
)
const OptionRemaining = ""
// blacklisted: Once (struct)
type OnceStatus C.uint32_t
const (
	OnceStatusNotcalled OnceStatus = 0
	OnceStatusProgress OnceStatus = 1
	OnceStatusReady OnceStatus = 2
)
type OptionArg C.uint32_t
const (
	OptionArgNone OptionArg = 0
	OptionArgString OptionArg = 1
	OptionArgInt OptionArg = 2
	OptionArgCallback OptionArg = 3
	OptionArgFilename OptionArg = 4
	OptionArgStringArray OptionArg = 5
	OptionArgFilenameArray OptionArg = 6
	OptionArgDouble OptionArg = 7
	OptionArgInt64 OptionArg = 8
)
// blacklisted: OptionArgFunc (callback)
// blacklisted: OptionContext (struct)
// blacklisted: OptionEntry (struct)
type OptionError C.uint32_t
const (
	OptionErrorUnknownOption OptionError = 0
	OptionErrorBadValue OptionError = 1
	OptionErrorFailed OptionError = 2
)
// blacklisted: OptionErrorFunc (callback)
type OptionFlags C.uint32_t
const (
	OptionFlagsHidden OptionFlags = 1
	OptionFlagsInMain OptionFlags = 2
	OptionFlagsReverse OptionFlags = 4
	OptionFlagsNoArg OptionFlags = 8
	OptionFlagsFilename OptionFlags = 16
	OptionFlagsOptionalArg OptionFlags = 32
	OptionFlagsNoalias OptionFlags = 64
)
// blacklisted: OptionGroup (struct)
// blacklisted: OptionParseFunc (callback)
const PdpEndian = 3412
const Pi = 3.141593
const Pi2 = 1.570796
const Pi4 = 0.785398
const PollfdFormat = "%#I64x"
const PriorityDefault = 0
const PriorityDefaultIdle = 200
const PriorityHigh = -100
const PriorityHighIdle = 100
const PriorityLow = 300
// blacklisted: PatternSpec (struct)
type PollFD struct {
	FD int32
	Events uint16
	Revents uint16
}
// blacklisted: PollFunc (callback)
// blacklisted: PrintFunc (callback)
// blacklisted: Private (struct)
// blacklisted: PtrArray (struct)
// blacklisted: Queue (struct)
// blacklisted: Rand (struct)
// blacklisted: Regex (struct)
type RegexCompileFlags C.uint32_t
const (
	RegexCompileFlagsCaseless RegexCompileFlags = 1
	RegexCompileFlagsMultiline RegexCompileFlags = 2
	RegexCompileFlagsDotall RegexCompileFlags = 4
	RegexCompileFlagsExtended RegexCompileFlags = 8
	RegexCompileFlagsAnchored RegexCompileFlags = 16
	RegexCompileFlagsDollarEndonly RegexCompileFlags = 32
	RegexCompileFlagsUngreedy RegexCompileFlags = 512
	RegexCompileFlagsRaw RegexCompileFlags = 2048
	RegexCompileFlagsNoAutoCapture RegexCompileFlags = 4096
	RegexCompileFlagsOptimize RegexCompileFlags = 8192
	RegexCompileFlagsDupnames RegexCompileFlags = 524288
	RegexCompileFlagsNewlineCr RegexCompileFlags = 1048576
	RegexCompileFlagsNewlineLf RegexCompileFlags = 2097152
	RegexCompileFlagsNewlineCrlf RegexCompileFlags = 3145728
)
type RegexError C.uint32_t
const (
	RegexErrorCompile RegexError = 0
	RegexErrorOptimize RegexError = 1
	RegexErrorReplace RegexError = 2
	RegexErrorMatch RegexError = 3
	RegexErrorInternal RegexError = 4
	RegexErrorStrayBackslash RegexError = 101
	RegexErrorMissingControlChar RegexError = 102
	RegexErrorUnrecognizedEscape RegexError = 103
	RegexErrorQuantifiersOutOfOrder RegexError = 104
	RegexErrorQuantifierTooBig RegexError = 105
	RegexErrorUnterminatedCharacterClass RegexError = 106
	RegexErrorInvalidEscapeInCharacterClass RegexError = 107
	RegexErrorRangeOutOfOrder RegexError = 108
	RegexErrorNothingToRepeat RegexError = 109
	RegexErrorUnrecognizedCharacter RegexError = 112
	RegexErrorPosixNamedClassOutsideClass RegexError = 113
	RegexErrorUnmatchedParenthesis RegexError = 114
	RegexErrorInexistentSubpatternReference RegexError = 115
	RegexErrorUnterminatedComment RegexError = 118
	RegexErrorExpressionTooLarge RegexError = 120
	RegexErrorMemoryError RegexError = 121
	RegexErrorVariableLengthLookbehind RegexError = 125
	RegexErrorMalformedCondition RegexError = 126
	RegexErrorTooManyConditionalBranches RegexError = 127
	RegexErrorAssertionExpected RegexError = 128
	RegexErrorUnknownPosixClassName RegexError = 130
	RegexErrorPosixCollatingElementsNotSupported RegexError = 131
	RegexErrorHexCodeTooLarge RegexError = 134
	RegexErrorInvalidCondition RegexError = 135
	RegexErrorSingleByteMatchInLookbehind RegexError = 136
	RegexErrorInfiniteLoop RegexError = 140
	RegexErrorMissingSubpatternNameTerminator RegexError = 142
	RegexErrorDuplicateSubpatternName RegexError = 143
	RegexErrorMalformedProperty RegexError = 146
	RegexErrorUnknownProperty RegexError = 147
	RegexErrorSubpatternNameTooLong RegexError = 148
	RegexErrorTooManySubpatterns RegexError = 149
	RegexErrorInvalidOctalValue RegexError = 151
	RegexErrorTooManyBranchesInDefine RegexError = 154
	RegexErrorDefineRepetion RegexError = 155
	RegexErrorInconsistentNewlineOptions RegexError = 156
	RegexErrorMissingBackReference RegexError = 157
)
// blacklisted: RegexEvalCallback (callback)
type RegexMatchFlags C.uint32_t
const (
	RegexMatchFlagsAnchored RegexMatchFlags = 16
	RegexMatchFlagsNotbol RegexMatchFlags = 128
	RegexMatchFlagsNoteol RegexMatchFlags = 256
	RegexMatchFlagsNotempty RegexMatchFlags = 1024
	RegexMatchFlagsPartial RegexMatchFlags = 32768
	RegexMatchFlagsNewlineCr RegexMatchFlags = 1048576
	RegexMatchFlagsNewlineLf RegexMatchFlags = 2097152
	RegexMatchFlagsNewlineCrlf RegexMatchFlags = 3145728
	RegexMatchFlagsNewlineAny RegexMatchFlags = 4194304
)
// blacklisted: Relation (struct)
const SearchpathSeparator = 59
const SearchpathSeparatorS = ";"
const SizeofLong = 8
const SizeofSizeT = 8
const SizeofVoidP = 8
// blacklisted: SList (struct)
const Sqrt2 = 1.414214
const StrDelimiters = "_-|> <."
const SysdefAfInet = 2
const SysdefAfInet6 = 10
const SysdefAfUnix = 1
const SysdefMsgDontroute = 4
const SysdefMsgOob = 1
const SysdefMsgPeek = 2
// blacklisted: Scanner (struct)
// blacklisted: ScannerConfig (struct)
// blacklisted: ScannerMsgFunc (callback)
type SeekType C.uint32_t
const (
	SeekTypeCur SeekType = 0
	SeekTypeSet SeekType = 1
	SeekTypeEnd SeekType = 2
)
// blacklisted: Sequence (struct)
// blacklisted: SequenceIter (struct)
// blacklisted: SequenceIterCompareFunc (callback)
type ShellError C.uint32_t
const (
	ShellErrorBadQuoting ShellError = 0
	ShellErrorEmptyString ShellError = 1
	ShellErrorFailed ShellError = 2
)
type SliceConfig C.uint32_t
const (
	SliceConfigAlwaysMalloc SliceConfig = 1
	SliceConfigBypassMagazines SliceConfig = 2
	SliceConfigWorkingSetMsecs SliceConfig = 3
	SliceConfigColorIncrement SliceConfig = 4
	SliceConfigChunkSizes SliceConfig = 5
	SliceConfigContentionCounter SliceConfig = 6
)
// blacklisted: Source (struct)
// blacklisted: SourceCallbackFuncs (struct)
// blacklisted: SourceDummyMarshal (callback)
// blacklisted: SourceFunc (callback)
// blacklisted: SourceFuncs (struct)
// blacklisted: SourcePrivate (struct)
// blacklisted: SpawnChildSetupFunc (callback)
type SpawnError C.uint32_t
const (
	SpawnErrorFork SpawnError = 0
	SpawnErrorRead SpawnError = 1
	SpawnErrorChdir SpawnError = 2
	SpawnErrorAcces SpawnError = 3
	SpawnErrorPerm SpawnError = 4
	SpawnError2big SpawnError = 5
	SpawnErrorNoexec SpawnError = 6
	SpawnErrorNametoolong SpawnError = 7
	SpawnErrorNoent SpawnError = 8
	SpawnErrorNomem SpawnError = 9
	SpawnErrorNotdir SpawnError = 10
	SpawnErrorLoop SpawnError = 11
	SpawnErrorTxtbusy SpawnError = 12
	SpawnErrorIo SpawnError = 13
	SpawnErrorNfile SpawnError = 14
	SpawnErrorMfile SpawnError = 15
	SpawnErrorInval SpawnError = 16
	SpawnErrorIsdir SpawnError = 17
	SpawnErrorLibbad SpawnError = 18
	SpawnErrorFailed SpawnError = 19
)
type SpawnFlags C.uint32_t
const (
	SpawnFlagsLeaveDescriptorsOpen SpawnFlags = 1
	SpawnFlagsDoNotReapChild SpawnFlags = 2
	SpawnFlagsSearchPath SpawnFlags = 4
	SpawnFlagsStdoutToDevNull SpawnFlags = 8
	SpawnFlagsStderrToDevNull SpawnFlags = 16
	SpawnFlagsChildInheritsStdin SpawnFlags = 32
	SpawnFlagsFileAndArgvZero SpawnFlags = 64
)
// blacklisted: StatBuf (struct)
// blacklisted: StaticMutex (struct)
// blacklisted: StaticPrivate (struct)
// blacklisted: StaticRWLock (struct)
// blacklisted: StaticRecMutex (struct)
// blacklisted: String (struct)
// blacklisted: StringChunk (struct)
type SystemThread struct {
	_data [8]byte
}
// blacklisted: TestCase (struct)
// blacklisted: TestConfig (struct)
// blacklisted: TestDataFunc (callback)
// blacklisted: TestFixtureFunc (callback)
// blacklisted: TestFunc (callback)
// blacklisted: TestLogBuffer (struct)
// blacklisted: TestLogFatalFunc (callback)
// blacklisted: TestLogMsg (struct)
type TestLogType C.uint32_t
const (
	TestLogTypeNone TestLogType = 0
	TestLogTypeError TestLogType = 1
	TestLogTypeStartBinary TestLogType = 2
	TestLogTypeListCase TestLogType = 3
	TestLogTypeSkipCase TestLogType = 4
	TestLogTypeStartCase TestLogType = 5
	TestLogTypeStopCase TestLogType = 6
	TestLogTypeMinResult TestLogType = 7
	TestLogTypeMaxResult TestLogType = 8
	TestLogTypeMessage TestLogType = 9
)
// blacklisted: TestSuite (struct)
type TestTrapFlags C.uint32_t
const (
	TestTrapFlagsSilenceStdout TestTrapFlags = 128
	TestTrapFlagsSilenceStderr TestTrapFlags = 256
	TestTrapFlagsInheritStdin TestTrapFlags = 512
)
// blacklisted: Thread (struct)
type ThreadError C.uint32_t
const (
	ThreadErrorThreadErrorAgain ThreadError = 0
)
// blacklisted: ThreadFunctions (struct)
// blacklisted: ThreadPool (struct)
type ThreadPriority C.uint32_t
const (
	ThreadPriorityLow ThreadPriority = 0
	ThreadPriorityNormal ThreadPriority = 1
	ThreadPriorityHigh ThreadPriority = 2
	ThreadPriorityUrgent ThreadPriority = 3
)
type TimeType C.uint32_t
const (
	TimeTypeStandard TimeType = 0
	TimeTypeDaylight TimeType = 1
	TimeTypeUniversal TimeType = 2
)
// blacklisted: TimeVal (struct)
// blacklisted: TimeZone (struct)
// blacklisted: Timer (struct)
type TokenType C.uint32_t
const (
	TokenTypeEof TokenType = 0
	TokenTypeLeftParen TokenType = 40
	TokenTypeRightParen TokenType = 41
	TokenTypeLeftCurly TokenType = 123
	TokenTypeRightCurly TokenType = 125
	TokenTypeLeftBrace TokenType = 91
	TokenTypeRightBrace TokenType = 93
	TokenTypeEqualSign TokenType = 61
	TokenTypeComma TokenType = 44
	TokenTypeNone TokenType = 256
	TokenTypeError TokenType = 257
	TokenTypeChar TokenType = 258
	TokenTypeBinary TokenType = 259
	TokenTypeOctal TokenType = 260
	TokenTypeInt TokenType = 261
	TokenTypeHex TokenType = 262
	TokenTypeFloat TokenType = 263
	TokenTypeString TokenType = 264
	TokenTypeSymbol TokenType = 265
	TokenTypeIdentifier TokenType = 266
	TokenTypeIdentifierNull TokenType = 267
	TokenTypeCommentSingle TokenType = 268
	TokenTypeCommentMulti TokenType = 269
	TokenTypeLast TokenType = 270
)
type TokenValue struct {
	_data [8]byte
}
// blacklisted: TranslateFunc (callback)
// blacklisted: TrashStack (struct)
type TraverseFlags C.uint32_t
const (
	TraverseFlagsLeaves TraverseFlags = 1
	TraverseFlagsNonLeaves TraverseFlags = 2
	TraverseFlagsAll TraverseFlags = 3
	TraverseFlagsMask TraverseFlags = 3
	TraverseFlagsLeafs TraverseFlags = 1
	TraverseFlagsNonLeafs TraverseFlags = 2
)
// blacklisted: TraverseFunc (callback)
type TraverseType C.uint32_t
const (
	TraverseTypeInOrder TraverseType = 0
	TraverseTypePreOrder TraverseType = 1
	TraverseTypePostOrder TraverseType = 2
	TraverseTypeLevelOrder TraverseType = 3
)
// blacklisted: Tree (struct)
// blacklisted: Tuples (struct)
const URIReservedCharsGenericDelimiters = ":/?#[]@"
const URIReservedCharsSubcomponentDelimiters = "!$&'()*+,;="
const UsecPerSec = 1000000
type UnicodeBreakType C.uint32_t
const (
	UnicodeBreakTypeMandatory UnicodeBreakType = 0
	UnicodeBreakTypeCarriageReturn UnicodeBreakType = 1
	UnicodeBreakTypeLineFeed UnicodeBreakType = 2
	UnicodeBreakTypeCombiningMark UnicodeBreakType = 3
	UnicodeBreakTypeSurrogate UnicodeBreakType = 4
	UnicodeBreakTypeZeroWidthSpace UnicodeBreakType = 5
	UnicodeBreakTypeInseparable UnicodeBreakType = 6
	UnicodeBreakTypeNonBreakingGlue UnicodeBreakType = 7
	UnicodeBreakTypeContingent UnicodeBreakType = 8
	UnicodeBreakTypeSpace UnicodeBreakType = 9
	UnicodeBreakTypeAfter UnicodeBreakType = 10
	UnicodeBreakTypeBefore UnicodeBreakType = 11
	UnicodeBreakTypeBeforeAndAfter UnicodeBreakType = 12
	UnicodeBreakTypeHyphen UnicodeBreakType = 13
	UnicodeBreakTypeNonStarter UnicodeBreakType = 14
	UnicodeBreakTypeOpenPunctuation UnicodeBreakType = 15
	UnicodeBreakTypeClosePunctuation UnicodeBreakType = 16
	UnicodeBreakTypeQuotation UnicodeBreakType = 17
	UnicodeBreakTypeExclamation UnicodeBreakType = 18
	UnicodeBreakTypeIdeographic UnicodeBreakType = 19
	UnicodeBreakTypeNumeric UnicodeBreakType = 20
	UnicodeBreakTypeInfixSeparator UnicodeBreakType = 21
	UnicodeBreakTypeSymbol UnicodeBreakType = 22
	UnicodeBreakTypeAlphabetic UnicodeBreakType = 23
	UnicodeBreakTypePrefix UnicodeBreakType = 24
	UnicodeBreakTypePostfix UnicodeBreakType = 25
	UnicodeBreakTypeComplexContext UnicodeBreakType = 26
	UnicodeBreakTypeAmbiguous UnicodeBreakType = 27
	UnicodeBreakTypeUnknown UnicodeBreakType = 28
	UnicodeBreakTypeNextLine UnicodeBreakType = 29
	UnicodeBreakTypeWordJoiner UnicodeBreakType = 30
	UnicodeBreakTypeHangulLJamo UnicodeBreakType = 31
	UnicodeBreakTypeHangulVJamo UnicodeBreakType = 32
	UnicodeBreakTypeHangulTJamo UnicodeBreakType = 33
	UnicodeBreakTypeHangulLvSyllable UnicodeBreakType = 34
	UnicodeBreakTypeHangulLvtSyllable UnicodeBreakType = 35
	UnicodeBreakTypeCloseParanthesis UnicodeBreakType = 36
)
type UnicodeScript C.int32_t
const (
	UnicodeScriptInvalidCode UnicodeScript = -1
	UnicodeScriptCommon UnicodeScript = 0
	UnicodeScriptInherited UnicodeScript = 1
	UnicodeScriptArabic UnicodeScript = 2
	UnicodeScriptArmenian UnicodeScript = 3
	UnicodeScriptBengali UnicodeScript = 4
	UnicodeScriptBopomofo UnicodeScript = 5
	UnicodeScriptCherokee UnicodeScript = 6
	UnicodeScriptCoptic UnicodeScript = 7
	UnicodeScriptCyrillic UnicodeScript = 8
	UnicodeScriptDeseret UnicodeScript = 9
	UnicodeScriptDevanagari UnicodeScript = 10
	UnicodeScriptEthiopic UnicodeScript = 11
	UnicodeScriptGeorgian UnicodeScript = 12
	UnicodeScriptGothic UnicodeScript = 13
	UnicodeScriptGreek UnicodeScript = 14
	UnicodeScriptGujarati UnicodeScript = 15
	UnicodeScriptGurmukhi UnicodeScript = 16
	UnicodeScriptHan UnicodeScript = 17
	UnicodeScriptHangul UnicodeScript = 18
	UnicodeScriptHebrew UnicodeScript = 19
	UnicodeScriptHiragana UnicodeScript = 20
	UnicodeScriptKannada UnicodeScript = 21
	UnicodeScriptKatakana UnicodeScript = 22
	UnicodeScriptKhmer UnicodeScript = 23
	UnicodeScriptLao UnicodeScript = 24
	UnicodeScriptLatin UnicodeScript = 25
	UnicodeScriptMalayalam UnicodeScript = 26
	UnicodeScriptMongolian UnicodeScript = 27
	UnicodeScriptMyanmar UnicodeScript = 28
	UnicodeScriptOgham UnicodeScript = 29
	UnicodeScriptOldItalic UnicodeScript = 30
	UnicodeScriptOriya UnicodeScript = 31
	UnicodeScriptRunic UnicodeScript = 32
	UnicodeScriptSinhala UnicodeScript = 33
	UnicodeScriptSyriac UnicodeScript = 34
	UnicodeScriptTamil UnicodeScript = 35
	UnicodeScriptTelugu UnicodeScript = 36
	UnicodeScriptThaana UnicodeScript = 37
	UnicodeScriptThai UnicodeScript = 38
	UnicodeScriptTibetan UnicodeScript = 39
	UnicodeScriptCanadianAboriginal UnicodeScript = 40
	UnicodeScriptYi UnicodeScript = 41
	UnicodeScriptTagalog UnicodeScript = 42
	UnicodeScriptHanunoo UnicodeScript = 43
	UnicodeScriptBuhid UnicodeScript = 44
	UnicodeScriptTagbanwa UnicodeScript = 45
	UnicodeScriptBraille UnicodeScript = 46
	UnicodeScriptCypriot UnicodeScript = 47
	UnicodeScriptLimbu UnicodeScript = 48
	UnicodeScriptOsmanya UnicodeScript = 49
	UnicodeScriptShavian UnicodeScript = 50
	UnicodeScriptLinearB UnicodeScript = 51
	UnicodeScriptTaiLe UnicodeScript = 52
	UnicodeScriptUgaritic UnicodeScript = 53
	UnicodeScriptNewTaiLue UnicodeScript = 54
	UnicodeScriptBuginese UnicodeScript = 55
	UnicodeScriptGlagolitic UnicodeScript = 56
	UnicodeScriptTifinagh UnicodeScript = 57
	UnicodeScriptSylotiNagri UnicodeScript = 58
	UnicodeScriptOldPersian UnicodeScript = 59
	UnicodeScriptKharoshthi UnicodeScript = 60
	UnicodeScriptUnknown UnicodeScript = 61
	UnicodeScriptBalinese UnicodeScript = 62
	UnicodeScriptCuneiform UnicodeScript = 63
	UnicodeScriptPhoenician UnicodeScript = 64
	UnicodeScriptPhagsPa UnicodeScript = 65
	UnicodeScriptNko UnicodeScript = 66
	UnicodeScriptKayahLi UnicodeScript = 67
	UnicodeScriptLepcha UnicodeScript = 68
	UnicodeScriptRejang UnicodeScript = 69
	UnicodeScriptSundanese UnicodeScript = 70
	UnicodeScriptSaurashtra UnicodeScript = 71
	UnicodeScriptCham UnicodeScript = 72
	UnicodeScriptOlChiki UnicodeScript = 73
	UnicodeScriptVai UnicodeScript = 74
	UnicodeScriptCarian UnicodeScript = 75
	UnicodeScriptLycian UnicodeScript = 76
	UnicodeScriptLydian UnicodeScript = 77
	UnicodeScriptAvestan UnicodeScript = 78
	UnicodeScriptBamum UnicodeScript = 79
	UnicodeScriptEgyptianHieroglyphs UnicodeScript = 80
	UnicodeScriptImperialAramaic UnicodeScript = 81
	UnicodeScriptInscriptionalPahlavi UnicodeScript = 82
	UnicodeScriptInscriptionalParthian UnicodeScript = 83
	UnicodeScriptJavanese UnicodeScript = 84
	UnicodeScriptKaithi UnicodeScript = 85
	UnicodeScriptLisu UnicodeScript = 86
	UnicodeScriptMeeteiMayek UnicodeScript = 87
	UnicodeScriptOldSouthArabian UnicodeScript = 88
	UnicodeScriptOldTurkic UnicodeScript = 89
	UnicodeScriptSamaritan UnicodeScript = 90
	UnicodeScriptTaiTham UnicodeScript = 91
	UnicodeScriptTaiViet UnicodeScript = 92
	UnicodeScriptBatak UnicodeScript = 93
	UnicodeScriptBrahmi UnicodeScript = 94
	UnicodeScriptMandaic UnicodeScript = 95
)
type UnicodeType C.uint32_t
const (
	UnicodeTypeControl UnicodeType = 0
	UnicodeTypeFormat UnicodeType = 1
	UnicodeTypeUnassigned UnicodeType = 2
	UnicodeTypePrivateUse UnicodeType = 3
	UnicodeTypeSurrogate UnicodeType = 4
	UnicodeTypeLowercaseLetter UnicodeType = 5
	UnicodeTypeModifierLetter UnicodeType = 6
	UnicodeTypeOtherLetter UnicodeType = 7
	UnicodeTypeTitlecaseLetter UnicodeType = 8
	UnicodeTypeUppercaseLetter UnicodeType = 9
	UnicodeTypeSpacingMark UnicodeType = 10
	UnicodeTypeEnclosingMark UnicodeType = 11
	UnicodeTypeNonSpacingMark UnicodeType = 12
	UnicodeTypeDecimalNumber UnicodeType = 13
	UnicodeTypeLetterNumber UnicodeType = 14
	UnicodeTypeOtherNumber UnicodeType = 15
	UnicodeTypeConnectPunctuation UnicodeType = 16
	UnicodeTypeDashPunctuation UnicodeType = 17
	UnicodeTypeClosePunctuation UnicodeType = 18
	UnicodeTypeFinalPunctuation UnicodeType = 19
	UnicodeTypeInitialPunctuation UnicodeType = 20
	UnicodeTypeOtherPunctuation UnicodeType = 21
	UnicodeTypeOpenPunctuation UnicodeType = 22
	UnicodeTypeCurrencySymbol UnicodeType = 23
	UnicodeTypeModifierSymbol UnicodeType = 24
	UnicodeTypeMathSymbol UnicodeType = 25
	UnicodeTypeOtherSymbol UnicodeType = 26
	UnicodeTypeLineSeparator UnicodeType = 27
	UnicodeTypeParagraphSeparator UnicodeType = 28
	UnicodeTypeSpaceSeparator UnicodeType = 29
)
type UserDirectory C.uint32_t
const (
	UserDirectoryDirectoryDesktop UserDirectory = 0
	UserDirectoryDirectoryDocuments UserDirectory = 1
	UserDirectoryDirectoryDownload UserDirectory = 2
	UserDirectoryDirectoryMusic UserDirectory = 3
	UserDirectoryDirectoryPictures UserDirectory = 4
	UserDirectoryDirectoryPublicShare UserDirectory = 5
	UserDirectoryDirectoryTemplates UserDirectory = 6
	UserDirectoryDirectoryVideos UserDirectory = 7
	UserDirectoryNDirectories UserDirectory = 8
)
const VaCopyAsArray = 1
type Variant struct {}
func NewVariantArray(child_type0 *VariantType, children0 []*Variant) *Variant {
	var child_type1 *C.GVariantType
	var children1 **C.GVariant
	var n_children1 C.uint64_t
	child_type1 = (*C.GVariantType)(unsafe.Pointer(child_type0))
	children1 = (**C.GVariant)(C.malloc(C.size_t(int(unsafe.Sizeof(*children1)) * len(children0))))
	defer C.free(unsafe.Pointer(children1))
	for i, e := range children0 {
		(*(*[999999]*C.GVariant)(unsafe.Pointer(children1)))[i] = (*C.GVariant)(unsafe.Pointer(e))
	}
	n_children1 = C.uint64_t(len(children0))
	ret1 := C.g_variant_new_array(child_type1, children1, n_children1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantBoolean(value0 bool) *Variant {
	var value1 C.int
	value1 = _GoBoolToCBool(value0)
	ret1 := C.g_variant_new_boolean(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantByte(value0 int) *Variant {
	var value1 C.uint8_t
	value1 = C.uint8_t(value0)
	ret1 := C.g_variant_new_byte(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantBytestringArray(strv0 []string) *Variant {
	var strv1 **C.char
	var length1 C.int64_t
	strv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*strv1)) * len(strv0))))
	defer C.free(unsafe.Pointer(strv1))
	for i, e := range strv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i]))
	}
	length1 = C.int64_t(len(strv0))
	ret1 := C.g_variant_new_bytestring_array(strv1, length1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantDictEntry(key0 *Variant, value0 *Variant) *Variant {
	var key1 *C.GVariant
	var value1 *C.GVariant
	key1 = (*C.GVariant)(unsafe.Pointer(key0))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_variant_new_dict_entry(key1, value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantDouble(value0 float64) *Variant {
	var value1 C.double
	value1 = C.double(value0)
	ret1 := C.g_variant_new_double(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantHandle(value0 int) *Variant {
	var value1 C.int32_t
	value1 = C.int32_t(value0)
	ret1 := C.g_variant_new_handle(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantInt16(value0 int) *Variant {
	var value1 C.int16_t
	value1 = C.int16_t(value0)
	ret1 := C.g_variant_new_int16(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantInt32(value0 int) *Variant {
	var value1 C.int32_t
	value1 = C.int32_t(value0)
	ret1 := C.g_variant_new_int32(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantInt64(value0 int64) *Variant {
	var value1 C.int64_t
	value1 = C.int64_t(value0)
	ret1 := C.g_variant_new_int64(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantMaybe(child_type0 *VariantType, child0 *Variant) *Variant {
	var child_type1 *C.GVariantType
	var child1 *C.GVariant
	child_type1 = (*C.GVariantType)(unsafe.Pointer(child_type0))
	child1 = (*C.GVariant)(unsafe.Pointer(child0))
	ret1 := C.g_variant_new_maybe(child_type1, child1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantObjectPath(object_path0 string) *Variant {
	var object_path1 *C.char
	object_path1 = _GoStringToGString(object_path0)
	defer C.free(unsafe.Pointer(object_path1))
	ret1 := C.g_variant_new_object_path(object_path1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantObjv(strv0 []string) *Variant {
	var strv1 **C.char
	var length1 C.int64_t
	strv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*strv1)) * len(strv0))))
	defer C.free(unsafe.Pointer(strv1))
	for i, e := range strv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i]))
	}
	length1 = C.int64_t(len(strv0))
	ret1 := C.g_variant_new_objv(strv1, length1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantSignature(signature0 string) *Variant {
	var signature1 *C.char
	signature1 = _GoStringToGString(signature0)
	defer C.free(unsafe.Pointer(signature1))
	ret1 := C.g_variant_new_signature(signature1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantString(string0 string) *Variant {
	var string1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	ret1 := C.g_variant_new_string(string1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantStrv(strv0 []string) *Variant {
	var strv1 **C.char
	var length1 C.int64_t
	strv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*strv1)) * len(strv0))))
	defer C.free(unsafe.Pointer(strv1))
	for i, e := range strv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i]))
	}
	length1 = C.int64_t(len(strv0))
	ret1 := C.g_variant_new_strv(strv1, length1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTuple(children0 []*Variant) *Variant {
	var children1 **C.GVariant
	var n_children1 C.uint64_t
	children1 = (**C.GVariant)(C.malloc(C.size_t(int(unsafe.Sizeof(*children1)) * len(children0))))
	defer C.free(unsafe.Pointer(children1))
	for i, e := range children0 {
		(*(*[999999]*C.GVariant)(unsafe.Pointer(children1)))[i] = (*C.GVariant)(unsafe.Pointer(e))
	}
	n_children1 = C.uint64_t(len(children0))
	ret1 := C.g_variant_new_tuple(children1, n_children1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantUint16(value0 int) *Variant {
	var value1 C.uint16_t
	value1 = C.uint16_t(value0)
	ret1 := C.g_variant_new_uint16(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantUint32(value0 int) *Variant {
	var value1 C.uint32_t
	value1 = C.uint32_t(value0)
	ret1 := C.g_variant_new_uint32(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantUint64(value0 uint64) *Variant {
	var value1 C.uint64_t
	value1 = C.uint64_t(value0)
	ret1 := C.g_variant_new_uint64(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantVariant(value0 *Variant) *Variant {
	var value1 *C.GVariant
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_variant_new_variant(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Byteswap() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_byteswap(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Classify() VariantClass {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_classify(this1)
	var ret2 VariantClass
	ret2 = VariantClass(ret1)
	return ret2
}
func (this0 *Variant) Compare(two0 *Variant) int {
	var this1 *C.GVariant
	var two1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	two1 = (*C.GVariant)(unsafe.Pointer(two0))
	ret1 := C.g_variant_compare(this1, two1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) DupBytestring() (uint64, []int) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_bytestring(this1, &length1)
	var length2 uint64
	var ret2 []int
	length2 = uint64(length1)
	ret2 = make([]int, length1)
	for i := range ret2 {
		ret2[i] = int((*(*[999999]C.uint8_t)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) DupBytestringArray() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_bytestring_array(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *Variant) DupObjv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_objv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *Variant) DupString() (uint64, string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_string(this1, &length1)
	var length2 uint64
	var ret2 string
	length2 = uint64(length1)
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return length2, ret2
}
func (this0 *Variant) DupStrv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_strv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *Variant) Equal(two0 *Variant) bool {
	var this1 *C.GVariant
	var two1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	two1 = (*C.GVariant)(unsafe.Pointer(two0))
	ret1 := C.g_variant_equal(this1, two1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) GetBoolean() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_boolean(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) GetByte() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_byte(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) GetBytestring() []int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_bytestring(this1)
	var ret2 []int
	for i := range ret2 {
		ret2[i] = int((*(*[999999]C.uint8_t)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *Variant) GetBytestringArray() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_bytestring_array(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) GetChildValue(index_0 uint64) *Variant {
	var this1 *C.GVariant
	var index_1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	index_1 = C.uint64_t(index_0)
	ret1 := C.g_variant_get_child_value(this1, index_1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetData() {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	C.g_variant_get_data(this1)
}
func (this0 *Variant) GetDouble() float64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_double(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Variant) GetFixedArray(element_size0 uint64) (uint64, []unsafe.Pointer) {
	var this1 *C.GVariant
	var element_size1 C.uint64_t
	var n_elements1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	element_size1 = C.uint64_t(element_size0)
	ret1 := C.g_variant_get_fixed_array(this1, &n_elements1, element_size1)
	var n_elements2 uint64
	var ret2 []unsafe.Pointer
	n_elements2 = uint64(n_elements1)
	ret2 = make([]unsafe.Pointer, n_elements1)
	for i := range ret2 {
		ret2[i] = (*(*[999999]unsafe.Pointer)(unsafe.Pointer(ret1)))[i]
	}
	return n_elements2, ret2
}
func (this0 *Variant) GetHandle() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_handle(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) GetInt16() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_int16(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) GetInt32() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_int32(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) GetInt64() int64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_int64(this1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *Variant) GetMaybe() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_maybe(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetNormalForm() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_normal_form(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetObjv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_objv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) GetSize() uint64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_size(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Variant) GetString() (uint64, string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_string(this1, &length1)
	var length2 uint64
	var ret2 string
	length2 = uint64(length1)
	ret2 = C.GoString(ret1)
	return length2, ret2
}
func (this0 *Variant) GetStrv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_strv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) GetTypeString() string {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_type_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Variant) GetUint16() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_uint16(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) GetUint32() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_uint32(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) GetUint64() uint64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_uint64(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Variant) GetVariant() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_variant(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Hash() int {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_hash(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Variant) IsContainer() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_is_container(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) IsFloating() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_is_floating(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) IsNormalForm() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_is_normal_form(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) IsOfType(type0 *VariantType) bool {
	var this1 *C.GVariant
	var type1 *C.GVariantType
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	type1 = (*C.GVariantType)(unsafe.Pointer(type0))
	ret1 := C.g_variant_is_of_type(this1, type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) LookupValue(key0 string, expected_type0 *VariantType) *Variant {
	var this1 *C.GVariant
	var key1 *C.char
	var expected_type1 *C.GVariantType
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	expected_type1 = (*C.GVariantType)(unsafe.Pointer(expected_type0))
	ret1 := C.g_variant_lookup_value(this1, key1, expected_type1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) NChildren() uint64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_n_children(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Variant) Print(type_annotate0 bool) string {
	var this1 *C.GVariant
	var type_annotate1 C.int
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	type_annotate1 = _GoBoolToCBool(type_annotate0)
	ret1 := C.g_variant_print(this1, type_annotate1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) RefSink() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_ref_sink(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Store(data0 unsafe.Pointer) {
	var this1 *C.GVariant
	var data1 unsafe.Pointer
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	data1 = unsafe.Pointer(data0)
	C.g_variant_store(this1, data1)
}
func (this0 *Variant) TakeRef() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_take_ref(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func VariantIsObjectPath(string0 string) bool {
	var string1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	ret1 := C.g_variant_is_object_path(string1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func VariantIsSignature(string0 string) bool {
	var string1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	ret1 := C.g_variant_is_signature(string1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func VariantParse(type0 *VariantType, text0 string, limit0 string, endptr0 string) (*Variant, error) {
	var type1 *C.GVariantType
	var text1 *C.char
	var limit1 *C.char
	var endptr1 *C.char
	var err1 *C.GError
	type1 = (*C.GVariantType)(unsafe.Pointer(type0))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	limit1 = _GoStringToGString(limit0)
	defer C.free(unsafe.Pointer(limit1))
	endptr1 = _GoStringToGString(endptr0)
	defer C.free(unsafe.Pointer(endptr1))
	ret1 := C.g_variant_parse(type1, text1, limit1, endptr1, &err1)
	var ret2 *Variant
	var err2 error
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func VariantParserGetErrorQuark() int {
	ret1 := C.g_variant_parser_get_error_quark()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
// blacklisted: VariantBuilder (struct)
type VariantClass C.uint32_t
const (
	VariantClassBoolean VariantClass = 98
	VariantClassByte VariantClass = 121
	VariantClassInt16 VariantClass = 110
	VariantClassUint16 VariantClass = 113
	VariantClassInt32 VariantClass = 105
	VariantClassUint32 VariantClass = 117
	VariantClassInt64 VariantClass = 120
	VariantClassUint64 VariantClass = 116
	VariantClassHandle VariantClass = 104
	VariantClassDouble VariantClass = 100
	VariantClassString VariantClass = 115
	VariantClassObjectPath VariantClass = 111
	VariantClassSignature VariantClass = 103
	VariantClassVariant VariantClass = 118
	VariantClassMaybe VariantClass = 109
	VariantClassArray VariantClass = 97
	VariantClassTuple VariantClass = 40
	VariantClassDictEntry VariantClass = 123
)
type VariantParseError C.uint32_t
const (
	VariantParseErrorFailed VariantParseError = 0
	VariantParseErrorBasicTypeExpected VariantParseError = 1
	VariantParseErrorCannotInferType VariantParseError = 2
	VariantParseErrorDefiniteTypeExpected VariantParseError = 3
	VariantParseErrorInputNotAtEnd VariantParseError = 4
	VariantParseErrorInvalidCharacter VariantParseError = 5
	VariantParseErrorInvalidFormatString VariantParseError = 6
	VariantParseErrorInvalidObjectPath VariantParseError = 7
	VariantParseErrorInvalidSignature VariantParseError = 8
	VariantParseErrorInvalidTypeString VariantParseError = 9
	VariantParseErrorNoCommonType VariantParseError = 10
	VariantParseErrorNumberOutOfRange VariantParseError = 11
	VariantParseErrorNumberTooBig VariantParseError = 12
	VariantParseErrorTypeError VariantParseError = 13
	VariantParseErrorUnexpectedToken VariantParseError = 14
	VariantParseErrorUnknownKeyword VariantParseError = 15
	VariantParseErrorUnterminatedStringConstant VariantParseError = 16
	VariantParseErrorValueExpected VariantParseError = 17
)
type VariantType struct {}
func NewVariantType(type_string0 string) *VariantType {
	var type_string1 *C.char
	type_string1 = _GoStringToGString(type_string0)
	defer C.free(unsafe.Pointer(type_string1))
	ret1 := C.g_variant_type_new(type_string1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTypeTuple(items0 []*VariantType) *VariantType {
	var items1 **C.GVariantType
	var length1 C.int32_t
	items1 = (**C.GVariantType)(C.malloc(C.size_t(int(unsafe.Sizeof(*items1)) * len(items0))))
	defer C.free(unsafe.Pointer(items1))
	for i, e := range items0 {
		(*(*[999999]*C.GVariantType)(unsafe.Pointer(items1)))[i] = (*C.GVariantType)(unsafe.Pointer(e))
	}
	length1 = C.int32_t(len(items0))
	ret1 := C.g_variant_type_new_tuple(items1, length1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Copy() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_copy(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) DupString() string {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_dup_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Element() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_element(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Equal(type20 *VariantType) bool {
	var this1 *C.GVariantType
	var type21 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	type21 = (*C.GVariantType)(unsafe.Pointer(type20))
	ret1 := C.g_variant_type_equal(this1, type21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) First() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_first(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Free() {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	C.g_variant_type_free(this1)
}
func (this0 *VariantType) GetStringLength() uint64 {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_get_string_length(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *VariantType) Hash() int {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_hash(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *VariantType) IsArray() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_array(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsBasic() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_basic(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsContainer() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_container(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsDefinite() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_definite(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsDictEntry() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_dict_entry(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsMaybe() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_maybe(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsSubtypeOf(supertype0 *VariantType) bool {
	var this1 *C.GVariantType
	var supertype1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	supertype1 = (*C.GVariantType)(unsafe.Pointer(supertype0))
	ret1 := C.g_variant_type_is_subtype_of(this1, supertype1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsTuple() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_tuple(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsVariant() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_variant(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) Key() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_key(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) NItems() uint64 {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_n_items(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *VariantType) NewArray() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_new_array(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) NewDictEntry(value0 *VariantType) *VariantType {
	var this1 *C.GVariantType
	var value1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	value1 = (*C.GVariantType)(unsafe.Pointer(value0))
	ret1 := C.g_variant_type_new_dict_entry(this1, value1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) NewMaybe() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_new_maybe(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Next() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_next(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Value() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_value(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func VariantTypeChecked_(unknown0 string) *VariantType {
	var unknown1 *C.char
	unknown1 = _GoStringToGString(unknown0)
	defer C.free(unsafe.Pointer(unknown1))
	ret1 := C.g_variant_type_checked_(unknown1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func VariantTypeStringIsValid(type_string0 string) bool {
	var type_string1 *C.char
	type_string1 = _GoStringToGString(type_string0)
	defer C.free(unsafe.Pointer(type_string1))
	ret1 := C.g_variant_type_string_is_valid(type_string1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func VariantTypeStringScan(string0 string, limit0 string) (string, bool) {
	var string1 *C.char
	var limit1 *C.char
	var endptr1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	limit1 = _GoStringToGString(limit0)
	defer C.free(unsafe.Pointer(limit1))
	ret1 := C.g_variant_type_string_scan(string1, limit1, &endptr1)
	var endptr2 string
	var ret2 bool
	endptr2 = C.GoString(endptr1)
	C.g_free(unsafe.Pointer(endptr1))
	ret2 = ret1 != 0
	return endptr2, ret2
}
// blacklisted: VoidFunc (callback)
const Win32MsgHandle = 19981206
// blacklisted: _StaticAssert_347 (struct)
// blacklisted: access (function)
// blacklisted: array_free (function)
// blacklisted: array_get_element_size (function)
// blacklisted: array_unref (function)
// blacklisted: ascii_digit_value (function)
// blacklisted: ascii_dtostr (function)
// blacklisted: ascii_formatd (function)
// blacklisted: ascii_strcasecmp (function)
// blacklisted: ascii_strdown (function)
// blacklisted: ascii_strncasecmp (function)
// blacklisted: ascii_strtod (function)
// blacklisted: ascii_strtoll (function)
// blacklisted: ascii_strtoull (function)
// blacklisted: ascii_strup (function)
// blacklisted: ascii_tolower (function)
// blacklisted: ascii_toupper (function)
// blacklisted: ascii_xdigit_value (function)
// blacklisted: assert_warning (function)
// blacklisted: assertion_message (function)
// blacklisted: assertion_message_cmpstr (function)
// blacklisted: assertion_message_error (function)
// blacklisted: assertion_message_expr (function)
// blacklisted: atexit (function)
// blacklisted: atomic_int_add (function)
// blacklisted: atomic_int_and (function)
// blacklisted: atomic_int_compare_and_exchange (function)
// blacklisted: atomic_int_dec_and_test (function)
// blacklisted: atomic_int_exchange_and_add (function)
// blacklisted: atomic_int_get (function)
// blacklisted: atomic_int_inc (function)
// blacklisted: atomic_int_or (function)
// blacklisted: atomic_int_set (function)
// blacklisted: atomic_int_xor (function)
// blacklisted: atomic_pointer_add (function)
// blacklisted: atomic_pointer_and (function)
// blacklisted: atomic_pointer_compare_and_exchange (function)
// blacklisted: atomic_pointer_or (function)
// blacklisted: atomic_pointer_set (function)
// blacklisted: atomic_pointer_xor (function)
// blacklisted: base64_decode (function)
// blacklisted: base64_decode_inplace (function)
// blacklisted: base64_decode_step (function)
// blacklisted: base64_encode (function)
// blacklisted: base64_encode_close (function)
// blacklisted: base64_encode_step (function)
// blacklisted: basename (function)
// blacklisted: bit_lock (function)
// blacklisted: bit_nth_lsf (function)
// blacklisted: bit_nth_msf (function)
// blacklisted: bit_storage (function)
// blacklisted: bit_trylock (function)
// blacklisted: bit_unlock (function)
// blacklisted: blow_chunks (function)
// blacklisted: bookmark_file_error_quark (function)
// blacklisted: build_filenamev (function)
// blacklisted: build_pathv (function)
// blacklisted: byte_array_free (function)
// blacklisted: byte_array_unref (function)
// blacklisted: chdir (function)
// blacklisted: check_version (function)
// blacklisted: checksum_type_get_length (function)
// blacklisted: child_watch_add (function)
// blacklisted: child_watch_source_new (function)
// blacklisted: clear_error (function)
// blacklisted: compute_checksum_for_data (function)
// blacklisted: compute_checksum_for_string (function)
// blacklisted: compute_hmac_for_data (function)
// blacklisted: compute_hmac_for_string (function)
// blacklisted: convert (function)
// blacklisted: convert_error_quark (function)
// blacklisted: convert_with_fallback (function)
// blacklisted: convert_with_iconv (function)
// blacklisted: datalist_clear (function)
// blacklisted: datalist_get_flags (function)
// blacklisted: datalist_id_set_data_full (function)
// blacklisted: datalist_init (function)
// blacklisted: datalist_set_flags (function)
// blacklisted: datalist_unset_flags (function)
// blacklisted: dataset_destroy (function)
// blacklisted: dataset_id_set_data_full (function)
// blacklisted: date_get_days_in_month (function)
// blacklisted: date_get_monday_weeks_in_year (function)
// blacklisted: date_get_sunday_weeks_in_year (function)
// blacklisted: date_is_leap_year (function)
// blacklisted: date_strftime (function)
// blacklisted: date_time_compare (function)
// blacklisted: date_time_equal (function)
// blacklisted: date_time_hash (function)
// blacklisted: date_valid_day (function)
// blacklisted: date_valid_dmy (function)
// blacklisted: date_valid_julian (function)
// blacklisted: date_valid_month (function)
// blacklisted: date_valid_weekday (function)
// blacklisted: date_valid_year (function)
// blacklisted: dcgettext (function)
// blacklisted: dgettext (function)
// blacklisted: dir_make_tmp (function)
// blacklisted: direct_equal (function)
// blacklisted: direct_hash (function)
// blacklisted: dngettext (function)
// blacklisted: double_equal (function)
// blacklisted: double_hash (function)
// blacklisted: dpgettext (function)
// blacklisted: dpgettext2 (function)
// blacklisted: dummy_decl (function)
// blacklisted: file_error_from_errno (function)
// blacklisted: file_error_quark (function)
// blacklisted: file_get_contents (function)
// blacklisted: file_open_tmp (function)
// blacklisted: file_read_link (function)
// blacklisted: file_set_contents (function)
// blacklisted: file_test (function)
// blacklisted: filename_display_basename (function)
// blacklisted: filename_display_name (function)
// blacklisted: filename_from_uri (function)
// blacklisted: filename_from_utf8 (function)
// blacklisted: filename_to_uri (function)
// blacklisted: filename_to_utf8 (function)
// blacklisted: find_program_in_path (function)
// blacklisted: format_size (function)
// blacklisted: format_size_for_display (function)
// blacklisted: format_size_full (function)
// blacklisted: free (function)
// blacklisted: get_application_name (function)
// blacklisted: get_charset (function)
// blacklisted: get_current_dir (function)
// blacklisted: get_current_time (function)
// blacklisted: get_environ (function)
// blacklisted: get_filename_charsets (function)
// blacklisted: get_home_dir (function)
// blacklisted: get_host_name (function)
// blacklisted: get_language_names (function)
// blacklisted: get_locale_variants (function)
// blacklisted: get_monotonic_time (function)
// blacklisted: get_prgname (function)
// blacklisted: get_real_name (function)
// blacklisted: get_real_time (function)
// blacklisted: get_system_config_dirs (function)
// blacklisted: get_system_data_dirs (function)
// blacklisted: get_tmp_dir (function)
// blacklisted: get_user_cache_dir (function)
// blacklisted: get_user_config_dir (function)
// blacklisted: get_user_data_dir (function)
// blacklisted: get_user_name (function)
// blacklisted: get_user_runtime_dir (function)
// blacklisted: get_user_special_dir (function)
// blacklisted: getenv (function)
// blacklisted: hash_table_destroy (function)
// blacklisted: hash_table_insert (function)
// blacklisted: hash_table_lookup_extended (function)
// blacklisted: hash_table_remove (function)
// blacklisted: hash_table_remove_all (function)
// blacklisted: hash_table_replace (function)
// blacklisted: hash_table_size (function)
// blacklisted: hash_table_steal (function)
// blacklisted: hash_table_steal_all (function)
// blacklisted: hash_table_unref (function)
// blacklisted: hook_destroy (function)
// blacklisted: hook_destroy_link (function)
// blacklisted: hook_free (function)
// blacklisted: hook_insert_before (function)
// blacklisted: hook_prepend (function)
// blacklisted: hook_unref (function)
// blacklisted: hostname_is_ascii_encoded (function)
// blacklisted: hostname_is_ip_address (function)
// blacklisted: hostname_is_non_ascii (function)
// blacklisted: hostname_to_ascii (function)
// blacklisted: hostname_to_unicode (function)
// blacklisted: idle_add (function)
// blacklisted: idle_remove_by_data (function)
// blacklisted: idle_source_new (function)
// blacklisted: int64_equal (function)
// blacklisted: int64_hash (function)
// blacklisted: int_equal (function)
// blacklisted: int_hash (function)
// blacklisted: intern_static_string (function)
// blacklisted: intern_string (function)
// blacklisted: io_add_watch (function)
// blacklisted: io_channel_error_from_errno (function)
// blacklisted: io_channel_error_quark (function)
// blacklisted: io_create_watch (function)
// blacklisted: key_file_error_quark (function)
// blacklisted: list_pop_allocator (function)
// blacklisted: list_push_allocator (function)
// blacklisted: listenv (function)
// blacklisted: locale_from_utf8 (function)
// blacklisted: locale_to_utf8 (function)
// blacklisted: log_default_handler (function)
// blacklisted: log_remove_handler (function)
// blacklisted: log_set_always_fatal (function)
// blacklisted: log_set_fatal_mask (function)
// blacklisted: main_context_default (function)
// blacklisted: main_context_get_thread_default (function)
// blacklisted: main_current_source (function)
// blacklisted: main_depth (function)
// blacklisted: markup_error_quark (function)
// blacklisted: markup_escape_text (function)
// blacklisted: mem_chunk_info (function)
// blacklisted: mem_is_system_malloc (function)
// blacklisted: mem_profile (function)
// blacklisted: mem_set_vtable (function)
// blacklisted: mkdir_with_parents (function)
// blacklisted: mkdtemp (function)
// blacklisted: mkdtemp_full (function)
// blacklisted: mkstemp (function)
// blacklisted: mkstemp_full (function)
// blacklisted: node_pop_allocator (function)
// blacklisted: node_push_allocator (function)
// blacklisted: nullify_pointer (function)
// blacklisted: on_error_query (function)
// blacklisted: on_error_stack_trace (function)
// blacklisted: once_init_enter (function)
// blacklisted: once_init_enter_impl (function)
// blacklisted: once_init_leave (function)
// blacklisted: option_error_quark (function)
// blacklisted: parse_debug_string (function)
// blacklisted: path_get_basename (function)
// blacklisted: path_get_dirname (function)
// blacklisted: path_is_absolute (function)
// blacklisted: path_skip_root (function)
// blacklisted: pattern_match (function)
// blacklisted: pattern_match_simple (function)
// blacklisted: pattern_match_string (function)
// blacklisted: pointer_bit_lock (function)
// blacklisted: pointer_bit_trylock (function)
// blacklisted: pointer_bit_unlock (function)
// blacklisted: poll (function)
// blacklisted: propagate_error (function)
// blacklisted: ptr_array_add (function)
// blacklisted: ptr_array_remove (function)
// blacklisted: ptr_array_remove_fast (function)
// blacklisted: ptr_array_remove_range (function)
// blacklisted: ptr_array_set_free_func (function)
// blacklisted: ptr_array_set_size (function)
// blacklisted: ptr_array_unref (function)
// blacklisted: quark_from_static_string (function)
// blacklisted: quark_from_string (function)
// blacklisted: quark_to_string (function)
// blacklisted: quark_try_string (function)
// blacklisted: random_double (function)
// blacklisted: random_double_range (function)
// blacklisted: random_int (function)
// blacklisted: random_int_range (function)
// blacklisted: random_set_seed (function)
// blacklisted: regex_check_replacement (function)
// blacklisted: regex_error_quark (function)
// blacklisted: regex_escape_nul (function)
// blacklisted: regex_escape_string (function)
// blacklisted: regex_match_simple (function)
// blacklisted: reload_user_special_dirs_cache (function)
// blacklisted: return_if_fail_warning (function)
// blacklisted: rmdir (function)
// blacklisted: sequence_move (function)
// blacklisted: sequence_move_range (function)
// blacklisted: sequence_remove (function)
// blacklisted: sequence_remove_range (function)
// blacklisted: sequence_set (function)
// blacklisted: sequence_swap (function)
// blacklisted: set_application_name (function)
// blacklisted: set_error_literal (function)
// blacklisted: set_prgname (function)
// blacklisted: setenv (function)
// blacklisted: shell_error_quark (function)
// blacklisted: shell_parse_argv (function)
// blacklisted: shell_quote (function)
// blacklisted: shell_unquote (function)
// blacklisted: slice_free1 (function)
// blacklisted: slice_free_chain_with_offset (function)
// blacklisted: slice_get_config (function)
// blacklisted: slice_get_config_state (function)
// blacklisted: slice_set_config (function)
// blacklisted: slist_pop_allocator (function)
// blacklisted: slist_push_allocator (function)
// blacklisted: source_remove (function)
// blacklisted: source_remove_by_funcs_user_data (function)
// blacklisted: source_remove_by_user_data (function)
// blacklisted: source_set_name_by_id (function)
// blacklisted: spaced_primes_closest (function)
// blacklisted: spawn_async (function)
// blacklisted: spawn_async_with_pipes (function)
// blacklisted: spawn_close_pid (function)
// blacklisted: spawn_command_line_async (function)
// blacklisted: spawn_command_line_sync (function)
// blacklisted: spawn_error_quark (function)
// blacklisted: spawn_sync (function)
// blacklisted: stpcpy (function)
// blacklisted: str_equal (function)
// blacklisted: str_has_prefix (function)
// blacklisted: str_has_suffix (function)
// blacklisted: str_hash (function)
// blacklisted: strcanon (function)
// blacklisted: strcasecmp (function)
// blacklisted: strchomp (function)
// blacklisted: strchug (function)
// blacklisted: strcmp0 (function)
// blacklisted: strcompress (function)
// blacklisted: strdelimit (function)
// blacklisted: strdown (function)
// blacklisted: strdup (function)
// blacklisted: strerror (function)
// blacklisted: strescape (function)
// blacklisted: strfreev (function)
// blacklisted: string_new (function)
// blacklisted: string_new_len (function)
// blacklisted: string_sized_new (function)
// blacklisted: strip_context (function)
// blacklisted: strjoinv (function)
// blacklisted: strlcat (function)
// blacklisted: strlcpy (function)
// blacklisted: strncasecmp (function)
// blacklisted: strndup (function)
// blacklisted: strnfill (function)
// blacklisted: strreverse (function)
// blacklisted: strrstr (function)
// blacklisted: strrstr_len (function)
// blacklisted: strsignal (function)
// blacklisted: strstr_len (function)
// blacklisted: strtod (function)
// blacklisted: strup (function)
// blacklisted: strv_get_type (function)
// blacklisted: strv_length (function)
// blacklisted: test_bug (function)
// blacklisted: test_bug_base (function)
// blacklisted: test_fail (function)
// blacklisted: test_log_type_name (function)
// blacklisted: test_queue_destroy (function)
// blacklisted: test_queue_free (function)
// blacklisted: test_rand_double (function)
// blacklisted: test_rand_double_range (function)
// blacklisted: test_rand_int (function)
// blacklisted: test_rand_int_range (function)
// blacklisted: test_run (function)
// blacklisted: test_run_suite (function)
// blacklisted: test_timer_elapsed (function)
// blacklisted: test_timer_last (function)
// blacklisted: test_timer_start (function)
// blacklisted: test_trap_assertions (function)
// blacklisted: test_trap_fork (function)
// blacklisted: test_trap_has_passed (function)
// blacklisted: test_trap_reached_timeout (function)
// blacklisted: thread_error_quark (function)
// blacklisted: thread_exit (function)
// blacklisted: thread_get_initialized (function)
// blacklisted: thread_init (function)
// blacklisted: thread_init_with_errorcheck_mutexes (function)
// blacklisted: thread_pool_get_max_idle_time (function)
// blacklisted: thread_pool_get_max_unused_threads (function)
// blacklisted: thread_pool_get_num_unused_threads (function)
// blacklisted: thread_pool_set_max_idle_time (function)
// blacklisted: thread_pool_set_max_unused_threads (function)
// blacklisted: thread_pool_stop_unused_threads (function)
// blacklisted: time_val_from_iso8601 (function)
// blacklisted: timeout_add (function)
// blacklisted: timeout_add_seconds (function)
// blacklisted: timeout_source_new (function)
// blacklisted: timeout_source_new_seconds (function)
// blacklisted: trash_stack_height (function)
// blacklisted: trash_stack_push (function)
// blacklisted: ucs4_to_utf16 (function)
// blacklisted: ucs4_to_utf8 (function)
// blacklisted: unichar_break_type (function)
// blacklisted: unichar_combining_class (function)
// blacklisted: unichar_compose (function)
// blacklisted: unichar_decompose (function)
// blacklisted: unichar_digit_value (function)
// blacklisted: unichar_fully_decompose (function)
// blacklisted: unichar_get_mirror_char (function)
// blacklisted: unichar_get_script (function)
// blacklisted: unichar_isalnum (function)
// blacklisted: unichar_isalpha (function)
// blacklisted: unichar_iscntrl (function)
// blacklisted: unichar_isdefined (function)
// blacklisted: unichar_isdigit (function)
// blacklisted: unichar_isgraph (function)
// blacklisted: unichar_islower (function)
// blacklisted: unichar_ismark (function)
// blacklisted: unichar_isprint (function)
// blacklisted: unichar_ispunct (function)
// blacklisted: unichar_isspace (function)
// blacklisted: unichar_istitle (function)
// blacklisted: unichar_isupper (function)
// blacklisted: unichar_iswide (function)
// blacklisted: unichar_iswide_cjk (function)
// blacklisted: unichar_isxdigit (function)
// blacklisted: unichar_iszerowidth (function)
// blacklisted: unichar_to_utf8 (function)
// blacklisted: unichar_tolower (function)
// blacklisted: unichar_totitle (function)
// blacklisted: unichar_toupper (function)
// blacklisted: unichar_type (function)
// blacklisted: unichar_validate (function)
// blacklisted: unichar_xdigit_value (function)
// blacklisted: unicode_canonical_decomposition (function)
// blacklisted: unicode_canonical_ordering (function)
// blacklisted: unicode_script_from_iso15924 (function)
// blacklisted: unicode_script_to_iso15924 (function)
// blacklisted: unlink (function)
// blacklisted: unsetenv (function)
// blacklisted: uri_escape_string (function)
// blacklisted: uri_parse_scheme (function)
// blacklisted: uri_unescape_segment (function)
// blacklisted: uri_unescape_string (function)
// blacklisted: usleep (function)
// blacklisted: utf16_to_ucs4 (function)
// blacklisted: utf16_to_utf8 (function)
// blacklisted: utf8_casefold (function)
// blacklisted: utf8_collate (function)
// blacklisted: utf8_collate_key (function)
// blacklisted: utf8_collate_key_for_filename (function)
// blacklisted: utf8_find_next_char (function)
// blacklisted: utf8_find_prev_char (function)
// blacklisted: utf8_get_char (function)
// blacklisted: utf8_get_char_validated (function)
// blacklisted: utf8_normalize (function)
// blacklisted: utf8_offset_to_pointer (function)
// blacklisted: utf8_pointer_to_offset (function)
// blacklisted: utf8_prev_char (function)
// blacklisted: utf8_strchr (function)
// blacklisted: utf8_strdown (function)
// blacklisted: utf8_strlen (function)
// blacklisted: utf8_strncpy (function)
// blacklisted: utf8_strrchr (function)
// blacklisted: utf8_strreverse (function)
// blacklisted: utf8_strup (function)
// blacklisted: utf8_substring (function)
// blacklisted: utf8_to_ucs4 (function)
// blacklisted: utf8_to_ucs4_fast (function)
// blacklisted: utf8_to_utf16 (function)
// blacklisted: utf8_validate (function)
// blacklisted: variant_get_gtype (function)
// blacklisted: variant_get_type (function)
// blacklisted: variant_is_object_path (function)
// blacklisted: variant_is_signature (function)
// blacklisted: variant_parse (function)
// blacklisted: variant_parser_get_error_quark (function)
// blacklisted: variant_type_checked_ (function)
// blacklisted: variant_type_string_is_valid (function)
// blacklisted: variant_type_string_scan (function)
// blacklisted: warn_message (function)
